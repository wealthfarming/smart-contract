// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/utils/Pausable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "./interfaces/IBEQNFT.sol";

contract WealthFarming is AccessControl, Pausable, ReentrancyGuard{

    struct BuyNFT {
        uint256 id;
        uint256 navId;
        address buyer;
        uint256 amount;
        uint256 refund;
        uint256 timestamp;
        bool processed;
    }

    struct SellNFT {
        uint256 id;
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

    IBEQNFT public nftToken = IBEQNFT(0x742C882b0B82Bc75Dc07c526c0DEE3f92a3Bef18);

    /**
     * @dev decimal of token USDC
     */
    uint256 public usdcDecimal = 18;

    /**
     * @dev counter for mint nft pending transaction
     */
    uint256 public buyNFTCounter = 0; 

    /**
     * @dev counter for mint nft pending transaction
     */
    uint256 public sellNFTCounter = 0;

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
    mapping(uint256 => BuyNFT) public pendingBuyNFT;

    /**
     * @dev Mapping for tracking pending mint transaction
     */
    mapping(address => mapping(uint256 => BuyNFT)) public mapAddressWithBuyNFT;

    /**
     * @dev Mapping of total buy nft of an address
     */
    mapping(address => uint256[]) public buyRequestIdOfAddress;

    /**
     * @dev Mapping for tracking pending sell
     */
    mapping(uint256 => SellNFT) public pendingSellNFT; 
   
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
    event BuyNFTRequest(address buyer, uint256 amount, uint256 transactionId, uint256 timestamp);
    event CancelMintNFT(uint256 transactionId);
    event NFTMinted(address indexed buyer, uint256 price, uint256 timestamp);
    event PendingTransactionProcessed(uint256 transactionId, uint256 timestamp, uint256 refund);
    event NFTTransferred(uint256 tokenId, address from, address to, uint256 price);
    event PendingSellCreated(uint256 sellId, uint256 tokenId, address seller,uint256 timestamp);
    event PendingSaleProcessed(uint256 sellId, uint256 price, uint256 timestamp);

    constructor() {
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
            // NavHistory memory nav = navHistory[len - 1];
            // require(block.timestamp >= nav.timestamp + 1 days, "CalculateNAV: T+1 condition not met");

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
    function buyNFT(uint256 amount) external whenNotPaused nonReentrant notContract {
        
        require(amount > 0, "BuyNFT: Amount must be greater than 0");
        require(usdcToken.allowance(msg.sender, address(this)) >= amount, "BuyNFT: Insufficient allowance");

        // transfer to this address
        usdcToken.transferFrom(msg.sender, address(this), amount);

        NavHistory memory nav = navHistory[navHistory.length - 1];

        BuyNFT memory buyNFTRequest = BuyNFT({
            id: buyNFTCounter,
            navId: nav.id,
            buyer: msg.sender,
            amount: amount,
            timestamp: block.timestamp,
            processed: false,
            refund: 0
        });

        // Save pending mint nft request to map of transaction counter
        pendingBuyNFT[buyNFTCounter] = buyNFTRequest;
        
        // Save to map of deposit of user
        mapAddressWithBuyNFT[msg.sender][buyNFTCounter] = buyNFTRequest;

        emit BuyNFTRequest(msg.sender, amount, buyNFTCounter, block.timestamp);
        buyNFTCounter++;
    }

    /**
     * @dev finalize mint nft
     * @param _transactionId is transaction id of mint nft request
     */
    function finalizeBuyNFT(uint256 _transactionId) external onlyRole(DEFAULT_ADMIN_ROLE)  {
        BuyNFT storage txn = pendingBuyNFT[_transactionId];
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
            nftToken.mint(txn.buyer);
            totalFeeValue = totalFeeValue + fee;
            emit NFTMinted(txn.buyer, navPrice, block.timestamp);
        }

        uint256 refundAmount = txn.amount - ((navPrice + fee) * numNFT);

        // refund usdc to buyer
        usdcToken.transfer(txn.buyer, refundAmount);

        // assign refund amount
        txn.refund = refundAmount;

        // send fee to admin
        usdcToken.transfer(receiveFee, totalFeeValue);
        
        // synchonize transaction
        mapAddressWithBuyNFT[txn.buyer][_transactionId] = txn;

        emit PendingTransactionProcessed(_transactionId, block.timestamp, refundAmount);
    }

    /**
     * @dev cancel mint nft
     * @param _transactionId is transaction id of mint nft request
     */
    function cancelMint(uint256 _transactionId) external whenNotPaused nonReentrant notContract {
        BuyNFT storage txn = pendingBuyNFT[_transactionId];
        require(msg.sender == txn.buyer, "CancelMint: Buyer not match");
        require(!txn.processed, "CancelMint: Transaction already processed");

        // refund usdc to buyer
        usdcToken.transfer(txn.buyer, txn.amount);

        txn.processed = true;
        txn.refund = txn.amount;
        
        // synchonize transaction
        mapAddressWithBuyNFT[txn.buyer][_transactionId] = txn;

        emit CancelMintNFT(_transactionId);
    }

    /**
     * @dev sell nft
     * @param _tokenId is id of nft
     */
    function sell(uint256 _tokenId) external whenNotPaused nonReentrant notContract {
        
        require(
            nftToken.ownerOf(_tokenId) == msg.sender,
            "Account does not own the specified ERC721 token ID"
        );

        nftToken.safeTransferFrom(msg.sender, address(this), _tokenId);

        SellNFT memory sellNFTRequest = SellNFT({
            id: sellNFTCounter,
            price: 0,
            seller: msg.sender,
            buyer: address(this),
            tokenId: _tokenId,
            timestamp: block.timestamp,
            processed: false
        });

        pendingSellNFT[sellNFTCounter] = sellNFTRequest;

        emit PendingSellCreated(
            sellNFTCounter, 
            _tokenId, 
            msg.sender,
            block.timestamp
        );
        
        sellNFTCounter++;
    }

    /**
     * @dev get all buy request of an address
     */
    function getAllBuyRequest() external view returns(BuyNFT[] memory) {
        BuyNFT[] memory transactions = new BuyNFT[](buyRequestIdOfAddress[msg.sender].length);
        uint256[] memory listId = buyRequestIdOfAddress[msg.sender];

        for (uint8 i = 0; i < listId.length;i++) {
            transactions[i] = pendingBuyNFT[listId[i]];
        }
        
        return transactions;
    }

    function pause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _pause();
    }

    function unpause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _unpause();
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
