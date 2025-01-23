// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract WealthFarming is ERC721URIStorage, AccessControl, Pausable, ReentrancyGuard{

    bytes32 public constant EDITOR_ROLE = keccak256("EDITOR_ROLE");

    struct PendingMintNFT {
        address buyer;
        uint256 amount;
        uint256 timestamp;
        bool processed;
    }

    struct PendingSale {
        address seller;
        address buyer;
        uint256 tokenId;
        uint256 price;
        uint256 timestamp;
        bool processed;
    }

    struct Asset {
        uint256 value;
        string name;
        string code;
    }

    struct Debt {
        uint256 value;
        string name;
        string code;
    }

    /**
     * @dev Token USDC use for trade
     */
    IERC20 public usdcToken;

    /**
     * @dev decimal of token USDC
     */
    uint256 public usdcDecimal;

    /**
    * @dev counter id for nft 
    */
    uint256 public nextTokenId = 1;

    /**
     * @dev counter for mint nft pending transaction
     */
    uint256 public transactionCounter = 0; 

    /**
     * @dev counter for mint nft pending transaction
     */
    uint256 public saleCounter = 0;

    /**
     * @dev total asset value is sum of value of all assets
     */
    uint256 public totalAssetValue = 0;

    /**
     * @dev total debt value is sum of value of all debt
     */
    uint256 public totalDebtValue = 0; 

    /**
     * @dev total fee from trading action, admin can withdraw
     */
    uint256 private totalFeeValue = 0; 

    /**
     * @dev Fee, this is the fee that will be charged on each transaction.
     */
    uint256 public baseFee = 50; 
    
    /**
     * @dev Factor used for calculating discounts, where 10000 represents 100%.
     */
    uint256 public constant DISCOUNT_FACTOR = 10000;

    /**
     * @dev Address of receive fee trade
     */
    address private receiveFee;

    /**
     * @dev Mapping for tracking pending mint transaction
     */
    mapping(uint256 => PendingMintNFT) public pendingMintNFTs;

    /**
     * @dev Mapping for tracking pending sell
     */
    mapping(uint256 => PendingSale) public pendingSales; 

    /**
    * @dev Mapper for tracking user NFTs
    */
    mapping(address => uint256[]) public userNFTs;

    /**
    * @dev list of asset
    */
    Asset[] public assets;

    /** 
    * @dev list of debt
    */
    Debt[] public debts; 

    /**
     * @dev Address representing the native token, e.g., Ether (ETH) on Ethereum.
     */
    address public constant NATIVE_TOKEN_ADDRESS = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    event EditorUpdated(address newEditor);
    event AssetAdded(string name, string code, uint256 value, uint256 index);
    event AssetUpdated(uint256 index, uint256 value);
    event DebtAdded(string name, string code, uint256 value, uint256 index);
    event DebtUpdated(uint256 index, uint256 value);
    event NAVUpdated(uint256 newNAV, uint256 timestamp);
    event NFTMinted(address indexed buyer, uint256 tokenId, uint256 price, uint256 timestamp);
    event PendingTransactionProcessed(uint256 transactionId, uint256 timestamp, uint256 refund);
    event NFTTransferred(uint256 tokenId, address from, address to, uint256 price);
    event PendingSaleCreated(uint256 saleId, uint256 tokenId, address seller, address buyer, uint256 price, uint256 timestamp);
    event PendingSaleProcessed(uint256 saleId, uint256 timestamp);

    constructor(
        address _usdcAddress, 
        uint256 _usdcDecimal,
        address _editor,
        address _receiveFee
    )  ERC721("WealthFarming NFT", "WFNFT") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(EDITOR_ROLE, _editor);
        usdcDecimal = _usdcDecimal;
        usdcToken = IERC20(_usdcAddress);
        receiveFee = _receiveFee;
    }

    /**
     * @dev Receives Ether sent directly to the contract address, typically used to handle accidental transfers.
     */
    receive() external payable {}

    /**
     * @dev Fallback function to handle Ether transfers, useful for managing accidental transfers to the contract.
     */
    fallback() external payable {}


    modifier notContract() {
        require(!isContract(msg.sender), "This function is not allowed in a contract");
        _;
    }
  
    /**
     * @dev Add asset
     *
     * @param _name is address of editor
     * @param _code is code of asset
     * @param _value is value of asset
     *  
     * Requirements:
     *
     * - Only editor
     */
    function addAsset(
        string memory _name, 
        string memory _code, 
        uint256 _value
    ) external onlyRole(EDITOR_ROLE) {
        assets.push(Asset({value: _value, name: _name, code: _code}));
        totalAssetValue += _value;
        uint256 _len = assets.length;
        emit AssetAdded(_name, _code, _value, _len - 1);
    }


    /**
     * @dev update asset
     *
     * @param _index is index of asset on list
     * @param _value is value of asset
     *  
     * Requirements:
     *
     * - Only editor
     */
    function updateAsset(uint256 _index, uint256 _value) external onlyRole(EDITOR_ROLE)  {
        require(_index < assets.length, "Invalid asset index");
        totalAssetValue = totalAssetValue - assets[_index].value + _value;
        assets[_index].value = _value;
        emit AssetUpdated(_index, _value);
    }


    /**
     * @dev get all assets
     *
     */
    function getAssets() external view returns (Asset[] memory) {
        return assets;
    }

    
    /**
     * @dev Add debt
     *
     * @param _name is name of debt
     * @param _code is code of debt
     * @param _value is value of debt
     *  
     * Requirements:
     *
     * - Only editor
     */
    function addDebt(
        string memory _name, 
        string memory _code, 
        uint256 _value
    ) external onlyRole(EDITOR_ROLE)  {
        debts.push(Debt({value: _value, name: _name, code: _code}));
        totalDebtValue += _value;
        uint256 _len = debts.length;
        emit DebtAdded(_name, _code, _value, _len);
    }

    /**
     * @dev update debt
     *
     * @param _index is index of debt 
     * @param _value is value of asset
     *  
     * Requirements:
     *
     * - Only editor
     */
    function updateDebt(uint256 _index, uint256 _value) external onlyRole(EDITOR_ROLE)  {
        require(_index < debts.length, "Invalid debt index");
        totalDebtValue = totalDebtValue - debts[_index].value + _value;
        debts[_index].value = _value;
        emit DebtUpdated(_index, _value);
    }

    /**
     * @dev get all debts
     *
     */
    function getDebts() external view returns (Debt[] memory) {
        return debts;
    }


    /**
     * @dev calculate nav value
     * nav = (total asset value - total debt value) / total nft
     *
     */
    function calculateNAV() public view returns (uint256) {
        uint256 totalNFTs = nextTokenId - 1; // Số lượng NFT đã phát hành
        require(totalNFTs > 0, "No NFTs minted yet");
        return (totalAssetValue - totalDebtValue) * (1 ** usdcDecimal) / totalNFTs;
    }

    /**
     * @dev buy nft
     * @param amount to deposit to address
     * user will deposit to smart contract, this transaction will trigger in end of day
     *
     */
    function mintNFT(uint256 amount) external whenNotPaused nonReentrant notContract {
        
        require(amount > 0, "Amount must be greater than 0");

        require(usdcToken.transferFrom(msg.sender, address(this), amount), "USDC transfer failed");

        pendingMintNFTs[transactionCounter] = PendingMintNFT({
            buyer: msg.sender,
            amount: amount,
            timestamp: block.timestamp,
            processed: false
        });

        transactionCounter++;
    }

    /**
     * @dev finalize mint nft
     * @param _transactionId is transaction id of 
     */
    function finalizeMint(uint256 _transactionId) external onlyRole(DEFAULT_ADMIN_ROLE)  {
        uint256 navPrice = calculateNAV();
        PendingMintNFT storage txn = pendingMintNFTs[_transactionId];
        require(!txn.processed, "Transaction already processed");
        require(block.timestamp >= txn.timestamp + 1 days, "T+1 condition not met");

        txn.processed = true;

        // calculate fee per nav price
        uint256 fee = navPrice * baseFee / DISCOUNT_FACTOR;

        // find number of NFT to mint
        uint256 numNFT = txn.amount / (navPrice + fee);
        
        // mint nft and transfer to buyer
        for (uint256 i = 0; i < numNFT; i++) {
            uint256 tokenId = nextTokenId++;
            _mint(txn.buyer, tokenId);
            _setTokenURI(tokenId, "");
            userNFTs[txn.buyer].push(tokenId);
            totalFeeValue = totalFeeValue + fee;
            emit NFTMinted(txn.buyer, tokenId, navPrice, block.timestamp);
        }

        uint256 refundAmount = txn.amount - ((navPrice + fee) * numNFT);

        // refund usdc to buyer
        usdcToken.transfer(txn.buyer, refundAmount);
        
        emit PendingTransactionProcessed(_transactionId, block.timestamp, refundAmount);
    }

    /**
     * @dev Get Pending Transaction 
     *
     */
    function getPendingTransactions() external view returns (PendingMintNFT[] memory) {
        PendingMintNFT[] memory transactions = new PendingMintNFT[](transactionCounter);
        for (uint256 i = 0; i < transactionCounter; i++) {
            transactions[i] = pendingMintNFTs[i];
        }
        return transactions;
    }

    function sellNFT(uint256 tokenId, address to, uint256 price) external whenNotPaused {
        require(ownerOf(tokenId) == msg.sender, "You do not own this token");

        // Tạo giao dịch bán pending
        pendingSales[saleCounter] = PendingSale({
            seller: msg.sender,
            buyer: to,
            tokenId: tokenId,
            price: price,
            timestamp: block.timestamp,
            processed: false
        });

        saleCounter++;
        emit PendingSaleCreated(saleCounter - 1, tokenId, msg.sender, to, price, block.timestamp);
    }

    function finalizeSale(uint256 saleId) external onlyRole(DEFAULT_ADMIN_ROLE) {
        PendingSale storage sale = pendingSales[saleId];
        require(!sale.processed, "Sale already processed");
        require(block.timestamp >= sale.timestamp + 1 days, "T+1 condition not met");

        sale.processed = true;

        // Chuyển quyền sở hữu NFT
        _transfer(sale.seller, sale.buyer, sale.tokenId);

        // Chuyển USDC từ người mua sang người bán
        require(usdcToken.transferFrom(sale.buyer, sale.seller, sale.price), "USDC transfer failed");

        emit PendingSaleProcessed(saleId, block.timestamp);
        emit NFTTransferred(sale.tokenId, sale.seller, sale.buyer, sale.price);
    }

    function getPendingSales() external view returns (PendingSale[] memory) {
        PendingSale[] memory sales = new PendingSale[](saleCounter);
        for (uint256 i = 0; i < saleCounter; i++) {
            sales[i] = pendingSales[i];
        }
        return sales;
    }

    function pause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _pause();
    }

    function unpause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _unpause();
    }


    function getTokenOwner(uint256 tokenId) external view returns (address) {
        return ownerOf(tokenId);
    }

    function getUserNFTs(address user) external view returns (uint256[] memory) {
        return userNFTs[user];
    }

    function getMarketplaceListings() external view returns (uint256[] memory) {
    
    }

    /**
     * @dev Add Editor
     *
     * @param _editor is address of editor
     *  
     * Requirements:
     *
     * - Only admin
     */
    function grantEditor(address _editor) internal onlyRole(DEFAULT_ADMIN_ROLE) {
        grantRole(EDITOR_ROLE, _editor);
    }

    /**
     * @dev Revoke Editor
     *
     * @param _editor is address of editor
     *  
     * Requirements:
     *
     * - Only admin
     */
    function revokeEditor(address _editor) internal onlyRole(DEFAULT_ADMIN_ROLE) {
        revokeRole(EDITOR_ROLE, _editor);
    }

    /**
     * @dev Sweeps tokens from the contract to the admin's address.
     * Reverts if the caller is not an admin, the token address is zero.
     * Transfers ETH balance if the token is ETH_ADDRESS, otherwise transfers IERC20 token balance.
     * @param _token The address of the token to sweep.
     */
    function sweepToken(address _token) external onlyRole(DEFAULT_ADMIN_ROLE) {

        if (_token == NATIVE_TOKEN_ADDRESS) {
            payable(msg.sender).transfer(address(this).balance);
        } else {
            uint256 balance = IERC20(_token).balanceOf(address(this));
            IERC20(_token).transfer(msg.sender, balance);
        }
    }

     // Override required by Solidity
    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721URIStorage, AccessControl)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }

    function isContract(address account) internal view returns (bool) {
        // This method relies on extcodesize, which returns 0 for contracts in
        // construction, since the code is only stored at the end of the
        // constructor execution.

        uint256 size;
        // solhint-disable-next-line no-inline-assembly
        assembly { size := extcodesize(account) }
        return size > 0;
    }
}
