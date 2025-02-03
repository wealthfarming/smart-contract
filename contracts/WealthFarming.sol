// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract WealthFarming is ERC721URIStorage, AccessControl, Pausable, ReentrancyGuard{

    struct MintNFT {
        uint256 id;
        uint256 navId;
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

    struct NavHistory {
        uint256 id;
        uint256 price;
        uint256 winRate;
        uint256 risk;
        uint256 timestamp;
    }

    /**
     * @dev Token USDC use for trade
     */
    IERC20 public usdcToken = IERC20(0xF62B3a6571E0A05E60663D39EB961e0e1814219d);

    /**
     * @dev decimal of token USDC
     */
    uint256 public usdcDecimal = 18;

    /**
    * @dev counter id for nft 
    */
    uint256 public nextTokenId = 1;

    /**
     * @dev counter for mint nft pending transaction
     */
    uint256 public mintNFTCounter = 0; 

    /**
     * @dev counter for mint nft pending transaction
     */
    uint256 public saleCounter = 0;

    /**
    * @dev nav counter
    */
    uint256 public navCounter = 0;

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
    mapping(uint256 => MintNFT) public pendingMintNFTs;

    /**
     * @dev Mapping for tracking pending sell
     */
    mapping(uint256 => PendingSale) public pendingSales; 

    /**
    * @dev Mapper for tracking user NFTs
    */
    mapping(address => uint256[]) public userNFTs;

    /**
    * @dev mapper for tracking wallet address with deposit transaction 
    */
    mapping(address => MintNFT[]) public deposits;

    /**
    * @dev list of asset
    */
    NavHistory[] public navHistory;

    /**
    * @dev nav price
    */
    uint256 navPrice;

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
    event MintNFTRequest(address buyer, uint256 amount, uint256 transactionId, uint256 timestamp);
    event CancelMintNFT(uint256 transactionId);
    event NFTMinted(address indexed buyer, uint256 tokenId, uint256 price, uint256 timestamp);
    event PendingTransactionProcessed(uint256 transactionId, uint256 timestamp, uint256 refund);
    event NFTTransferred(uint256 tokenId, address from, address to, uint256 price);
    event PendingSaleCreated(uint256 saleId, uint256 tokenId, address seller, address buyer, uint256 price, uint256 timestamp);
    event PendingSaleProcessed(uint256 saleId, uint256 timestamp);

    constructor()  ERC721("WealthFarming NFT", "WFNFT") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        receiveFee = msg.sender;
    }

    modifier notContract() {
        require(!isContract(msg.sender), "This function is not allowed in a contract");
        _;
    }
  
    /**
     * @dev Push nav value to array
     * @param _navPrice is nav price
     * @param _winRate is win rate
     * @param _risk is risk 
     *
     */
    function calculateNAV(uint256 _navPrice, uint256 _winRate, uint256 _risk) external onlyRole(DEFAULT_ADMIN_ROLE)  {
        navCounter++;
        uint256 len = navHistory.length;
        if (len == 0) {
            navPrice = _navPrice;

            navHistory.push(NavHistory({
                id: navCounter,
                price: _navPrice,
                winRate: _winRate,
                risk: _risk,
                timestamp: block.timestamp
            }));
        } else {
            NavHistory memory nav = navHistory[len - 1];
            require(block.timestamp >= nav.timestamp + 1 days, "CalculateNAV: T+1 condition not met");

            navPrice = _navPrice;

            navHistory.push(NavHistory({
                id: navCounter,
                price: _navPrice,
                winRate: _winRate,
                risk: _risk,
                timestamp: block.timestamp
            }));
        }
    }

    /**
     * @dev buy nft
     * @param amount to deposit to address
     * user will deposit to smart contract, this transaction will trigger in end of day
     *
     */
    function mintNFT(uint256 amount) external whenNotPaused nonReentrant notContract {
        
        require(amount > 0, "MintNFT: Amount must be greater than 0");
        require(usdcToken.allowance(msg.sender, address(this)) >= amount, "MintNFT: Insufficient allowance");

        // transfer to this address
        usdcToken.transferFrom(msg.sender, address(this), amount);

        NavHistory memory nav = navHistory[navHistory.length - 1];

        MintNFT memory mintNFTRequest = MintNFT({
            id: mintNFTCounter,
            navId: nav.id,
            buyer: msg.sender,
            amount: amount,
            timestamp: block.timestamp,
            processed: false
        });

        // Save pending mint nft request to map of transaction counter
        pendingMintNFTs[mintNFTCounter] = mintNFTRequest;
        
        // Save to map of deposit of user
        deposits[msg.sender].push(mintNFTRequest);

        emit MintNFTRequest(msg.sender, amount, mintNFTCounter, block.timestamp);
        mintNFTCounter++;
    }

    /**
     * @dev finalize mint nft
     * @param _transactionId is transaction id of mint nft request
     */
    function finalizeMint(uint256 _transactionId) external onlyRole(DEFAULT_ADMIN_ROLE)  {
        MintNFT storage txn = pendingMintNFTs[_transactionId];
        NavHistory memory nav = navHistory[navHistory.length - 1];

        require(txn.navId == nav.id - 1, "FinalizeMint: Transaction not in pool");
        require(!txn.processed, "FinalizeMint: Transaction already processed");
        
        txn.processed = true;

        // calculate fee per nav price
        uint256 fee = navPrice * baseFee / DISCOUNT_FACTOR;

        // find number of NFT to mint
        uint256 numNFT = txn.amount / (navPrice + fee);
        
        uint256 totalFeeValue = 0;

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

        // send fee to admin
        usdcToken.transfer(receiveFee, totalFeeValue);

        emit PendingTransactionProcessed(_transactionId, block.timestamp, refundAmount);
    }

    /**
     * @dev cancel mint nft
     * @param _transactionId is transaction id of mint nft request
     */
    function cancelMint(uint256 _transactionId) external whenNotPaused nonReentrant notContract {
        MintNFT storage txn = pendingMintNFTs[_transactionId];
        require(msg.sender == txn.buyer, "CancelMint: Buyer not match");
        require(!txn.processed, "CancelMint: Transaction already processed");

        txn.processed = true;

        // refund usdc to buyer
        usdcToken.transfer(txn.buyer, txn.amount);

        emit CancelMintNFT(_transactionId);
    }

    /**
     * @dev get all deposits of an address
     */
    function getAllDeposits() external view returns(MintNFT[] memory) {
        MintNFT[] memory transactions = deposits[msg.sender];
        return transactions;
    }

    /**
     * @dev Get Pending Transaction 
     *
     */
    function getPendingTransactions() external view returns (MintNFT[] memory) {
        MintNFT[] memory transactions = new MintNFT[](mintNFTCounter);
        for (uint256 i = 0; i < mintNFTCounter; i++) {
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
