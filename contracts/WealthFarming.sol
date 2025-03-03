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
        uint256 sessionId;
        address buyer;
        uint256 amount;
        uint256 refund;
        uint256 timestamp;
        bool processed;
    }

    struct SellNFT {
        uint256 id;
        uint256 sessionId;
        address seller;
        address buyer;
        uint256 tokenId;
        uint256 price;
        uint256 timestamp;
        bool processed;
    }

    struct SessionHistory {
        uint256 id;
        uint256 nftPrice;
        uint256 totalAsset;
        uint256 totalBuyNFT;
        uint256 totalSellNFT;
        uint256 totalNFT;
        uint256 winRate;
        uint256 risk;
        uint256 timestamp;
    }

    /**
     * @dev Token USDC use for trade
     */
    IERC20 public usdcToken = IERC20(0xF62B3a6571E0A05E60663D39EB961e0e1814219d);

    IBEQNFT public nftToken = IBEQNFT(0x5CDf2155c97D40c4Fa3BbA4016AB8F84ecBE83e7);

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
    uint256 public sessionCounter = 0;

    /**
     * @dev Fee, this is the fee that will be charged on each transaction.
     */
    uint256 public baseFee = 50;

    /**
     * @dev Factor used for calculating discounts, where 10000 represents 100%.
     */
    uint256 public constant FEE_FACTOR = 10000;

    /**
     * @dev ceil percent, this will be used to calcuate nft price, reduce excessive price increases
     * 70 - 7%
     */
    uint256 public ceil = 0;

    /**
     * @dev ceil percent, this will be used to calcuate nft price, reduce excessive price increases
     * 70 - 7%
     */
    uint256 public constant CEIL_FACTOR = 1000;

    /**
     * @dev floor percent, this will be used to calcuate nft price, reduce excessive price decrease
     * 70 - 7%
     */
    uint256 public floor = 70;

    /**
     * @dev ceil percent, this will be used to calcuate nft price, reduce excessive price increases
     * 70 - 7%
     */
    uint256 public constant FLOOR_FACTOR = 1000;

    /**
     * @dev Address of receive fee trade
     */
    address public receiveFee;

    /**
     * @dev Mapping ID with BuyNFT struct for tracking pending buy nft transaction
     */
    mapping(uint256 => BuyNFT) public buyNFTHistory;

    /**
     * @dev Tracking pending buy nft transaction in a session
     */
    uint256[] public pendingBuyNFTInSession;

    /**
     * @dev Mapping for tracking pending sell
     */
    mapping(uint256 => SellNFT) public sellNFTHistory;

    /**
     * @dev Tracking pending sell nft transaction in a session
     */
    uint256[] public pendingSellNFTInSession;

    /**
    * @dev list of asset
    */
    SessionHistory[] public sessionHistory;

    /**
    * @dev nav price
    */
    uint256[] public lockNFT;

    /**
    * @dev totalBuyInSession
    */
    uint256 public totalBuyInSession;

     /**
    * @dev totalSellInSession
    */
    uint256 public totalSellInSession;

    /**
    * @dev allowContract
    */
    mapping(address => bool) allowContract;


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
    event CancelBuyNFT(uint256 transactionId);
    event CancelSellNFT(uint256 transactionId);
    event NFTMinted(address indexed buyer, uint256 price, uint256 timestamp);
    event PendingBuyNFTProcessed(uint256 transactionId, uint256 timestamp, uint256 refund);
    event PendingSellNFTProcessed(uint256 sellId, uint256 tokenId, uint256 timestamp);
    event PendingSellCreated(uint256 sellId, uint256 tokenId, address seller,uint256 timestamp);

    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        receiveFee = msg.sender;

        sessionHistory.push(SessionHistory({
            id: sessionCounter,
            nftPrice: 1000 * (10 ** usdcDecimal),
            totalAsset: 0,
            totalBuyNFT: 0,
            totalSellNFT: 0,
            totalNFT: 0,
            winRate: 0,
            risk: 0,
            timestamp: block.timestamp
        }));

        sessionCounter++;
    }

    modifier notContract() {
        if (isContract(msg.sender)) {
            if (!allowContract[msg.sender]) {
                revert("Contract not allow to call this function");
            }
        }
        _;
    }

    /**
     * @dev Push nav value to array
     * @param _totalAsset is nav price
     * @param _winRate is win rate
     * @param _risk is risk
     */
    function calculateNFTPrice(uint256 _totalAsset, uint256 _winRate, uint256 _risk) external onlyRole(DEFAULT_ADMIN_ROLE)  {
        require(_totalAsset > 0, "calculateNFTPrice: NAV Price must greater than 0");

        uint256 totalNFT = nftToken.getNextTokenId() - 1;

        SessionHistory memory session = sessionHistory[sessionHistory.length - 1];
        uint256 fee = session.nftPrice * baseFee / FEE_FACTOR;

        // Add more estimate nfts
        uint256 estimateMintNFT = totalBuyInSession / (session.nftPrice + fee);
        if (estimateMintNFT > lockNFT.length) {
            totalNFT = totalNFT + estimateMintNFT - lockNFT.length;
        }

        uint256 balance = usdcToken.balanceOf(address(this));

        // Calculate nft price
        uint256 currentPrice = (_totalAsset + balance - totalSellInSession) / totalNFT;

        // Check limit decrease or increase
        if (sessionHistory.length > 0) {
            uint256 lastPrice = session.nftPrice;

            if (currentPrice > lastPrice && ceil != 0) {
                uint256 delta = currentPrice - lastPrice;
                uint256 maxDelta = lastPrice * ceil / CEIL_FACTOR;

                if (delta > maxDelta) {
                    currentPrice = lastPrice + maxDelta;
                } else {
                    currentPrice = lastPrice + delta;
                }
            } else if (lastPrice > currentPrice) {
                uint256 delta = lastPrice - currentPrice;
                uint256 maxDelta = lastPrice * floor / FLOOR_FACTOR;

                if (delta > maxDelta) {
                    currentPrice = lastPrice - maxDelta;
                } else {
                    currentPrice = lastPrice - delta;
                }
            }
        }

        // update session history
        sessionHistory.push(SessionHistory({
            id: sessionCounter,
            totalNFT: totalNFT,
            nftPrice: currentPrice,
            totalAsset: _totalAsset,
            totalBuyNFT: totalBuyInSession,
            totalSellNFT: totalSellInSession,
            winRate: _winRate,
            risk: _risk,
            timestamp: block.timestamp
        }));

        // reset amount
        totalBuyInSession = 0;
        totalSellInSession = 0;

        // update session counter
        sessionCounter++;
    }

    /**
     * @dev buy nft
     * @param amount to deposit to address
     * user will deposit to smart contract, this transaction will trigger in end of day
     *
     */
    function buyNFT(uint256 amount) external whenNotPaused nonReentrant notContract {
        SessionHistory memory session = sessionHistory[sessionHistory.length - 1];

        require(amount >= session.nftPrice, "BuyNFT: Amount must greater than nft price");
        require(usdcToken.allowance(msg.sender, address(this)) >= amount, "BuyNFT: Insufficient allowance");

        // transfer to this address
        usdcToken.transferFrom(msg.sender, address(this), amount);

        // Save pending mint nft request to map of transaction counter
        buyNFTHistory[buyNFTCounter] = BuyNFT({
            id: buyNFTCounter,
            sessionId: session.id,
            buyer: msg.sender,
            amount: amount,
            timestamp: block.timestamp,
            processed: false,
            refund: 0
        });

        // Emit event buy NFT
        emit BuyNFTRequest(
            msg.sender,
            amount,
            buyNFTCounter,
            block.timestamp
        );

        pendingBuyNFTInSession.push(buyNFTCounter);
        buyNFTCounter++;
        totalBuyInSession += amount;
    }

    /**
     * @dev Finalize mint nft
     * @param _transactionId is transaction id of mint nft request
     */
    function finalizeBuyNFT(uint256 _transactionId) external onlyRole(DEFAULT_ADMIN_ROLE)  {
        BuyNFT storage txn = buyNFTHistory[_transactionId];
        SessionHistory memory session = sessionHistory[sessionHistory.length - 1];

        require(txn.sessionId == session.id - 1, "FinalizeBuyNFT: Transaction not in pool");
        require(!txn.processed, "FinalizeBuyNFT: Transaction already processed");

        // Calculate fee per nav price
        uint256 fee = session.nftPrice * baseFee / FEE_FACTOR;

        // Find number of NFT to mint
        uint256 baseAmount = txn.amount;
        uint256 totalFeeValue = 0;
        uint256 nftPrice = session.nftPrice + fee;

        // Transfer NFT from funding contract to buyer
        while (lockNFT.length > 0 && nftPrice <= baseAmount) {
            totalFeeValue = totalFeeValue + fee;
            uint256 tokenId = lockNFT[lockNFT.length - 1];
            nftToken.transferFrom(address(this), txn.buyer, tokenId);

            lockNFT.pop();
            baseAmount = baseAmount - nftPrice;

            emit NFTMinted(txn.buyer, session.nftPrice, block.timestamp);
        }

        // Mint nft and transfer to buyer
        while(nftPrice <= baseAmount) {
            nftToken.mint(txn.buyer);
            totalFeeValue = totalFeeValue + fee;
            emit NFTMinted(txn.buyer, session.nftPrice, block.timestamp);
            baseAmount = baseAmount - nftPrice;
        }

        // remove pending buy in pool
        _removePendingBuyNFTInSession(_transactionId);

        if (baseAmount > 0) {
             // Refund usdc to buyer
            usdcToken.transfer(txn.buyer, baseAmount);
        }

        // Send fee to admin
        usdcToken.transfer(receiveFee, totalFeeValue);

        // Update txn
        txn.refund = baseAmount;
        txn.processed = true;

        emit PendingBuyNFTProcessed(
            _transactionId,
            block.timestamp,
            baseAmount
        );
    }

    /**
     * @dev cancel mint nft
     * @param _transactionId is transaction id of mint nft request
     */
    function cancelBuyNFT(uint256 _transactionId) external whenNotPaused nonReentrant notContract {
        BuyNFT storage txn = buyNFTHistory[_transactionId];
        require(msg.sender == txn.buyer, "CancelBuy: Buyer not match");
        require(!txn.processed, "CancelBuy: Transaction already processed");

        _removePendingBuyNFTInSession(_transactionId);

        // Refund usdc to buyer
        usdcToken.transfer(txn.buyer, txn.amount);

        txn.processed = true;
        txn.refund = txn.amount;
        totalBuyInSession -= txn.amount;

        emit CancelBuyNFT(_transactionId);
    }


    /**
     * @dev sell nft
     * @param _tokenId is id of nft
     */
    function sellNFT(uint256 _tokenId) external whenNotPaused nonReentrant notContract {

        require(
            nftToken.ownerOf(_tokenId) == msg.sender,
            "Account does not own the specified ERC721 token ID"
        );

        // Get latest Session info
        SessionHistory memory session = sessionHistory[sessionHistory.length - 1];

        // Store pending sell request
        sellNFTHistory[sellNFTCounter] = SellNFT({
            id: sellNFTCounter,
            sessionId: session.id,
            price: session.nftPrice,
            seller: msg.sender,
            buyer: address(this),
            tokenId: _tokenId,
            timestamp: block.timestamp,
            processed: false
        });

        // Transfer nft from owner to marker contract
        nftToken.transferFrom(msg.sender, address(this), _tokenId);

        pendingSellNFTInSession.push(sellNFTCounter);

        // Increase sell counter id
        sellNFTCounter++;

        // Total sell in session
        totalSellInSession += session.nftPrice;

        // Emit event
        emit PendingSellCreated(
            sellNFTCounter,
            _tokenId,
            msg.sender,
            block.timestamp
        );
    }

    /**
     * @dev finalize sell nft
     * @param _sellId is id sell request
     */
    function finalizeSellNFT(uint256 _sellId) external onlyRole(DEFAULT_ADMIN_ROLE)  {
        SellNFT storage txn = sellNFTHistory[_sellId];
        SessionHistory memory session = sessionHistory[sessionHistory.length - 1];

        require(txn.sessionId == session.id - 1, "FinalizeMint: Transaction not in pool");
        require(!txn.processed, "FinalizeMint: Transaction already processed");

        txn.processed = true;

        // calculate fee per nav price
        uint256 fee = txn.price * baseFee / FEE_FACTOR;

        // remove pending sell id
        _removePendingSellNFTInSession(_sellId);

        // refund usdc to buyer
        usdcToken.transfer(txn.seller, txn.price - fee);

        // send fee to admin
        usdcToken.transfer(receiveFee, fee);

        // update lock nft
        lockNFT.push(txn.tokenId);

        emit PendingSellNFTProcessed(_sellId, txn.tokenId, block.timestamp);
    }

    /**
     * @dev cancel sell nft
     * @param _sellId is id sell request
     */
    function cancelSellNFT(uint256 _sellId) external whenNotPaused nonReentrant notContract {
        SellNFT storage txn = sellNFTHistory[_sellId];
        require(msg.sender == txn.seller, "CancelSell: Buyer not match");
        require(!txn.processed, "CancelSell: Transaction already processed");

        // remove pending sell id
        _removePendingSellNFTInSession(_sellId);

        // refund usdc to buyer
        nftToken.transferFrom(address(this), txn.seller, txn.tokenId);

        txn.processed = true;

        emit CancelSellNFT(_sellId);
    }

    function getAllPendingBuyInSession() external view returns (uint256[] memory) {
        return pendingBuyNFTInSession;
    }

    function getAllPendingSellInSession() external view returns (uint256[] memory) {
        return pendingSellNFTInSession;
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

     /**
     * @dev Allow contract for interact with this market
     * @param _contractAddress address of contract
     */
    function addAllowContract(address _contractAddress) external onlyRole(DEFAULT_ADMIN_ROLE) {
        allowContract[_contractAddress] = true;
    }

    /**
     * @dev Remove contract for interact with this market
     * @param _contractAddress address of contract
     */
    function removeAllowContract(address _contractAddress) external onlyRole(DEFAULT_ADMIN_ROLE) {
        allowContract[_contractAddress] = false;
    }

    /**
     * @dev Update ceil percent
     * @param _ceil is ceil percent
     */
    function updateCeil(uint256 _ceil) external onlyRole(DEFAULT_ADMIN_ROLE) {
        ceil = _ceil;
    }

    /**
     * @dev Update floor percent
     * @param _floor is floor percent
     */
    function updateFloor(uint256 _floor) external onlyRole(DEFAULT_ADMIN_ROLE) {
        floor = _floor;
    }

    /**
     * @dev Remove transaction id in pending buy pool
     * @param _transactionId is id of transaction
     */
    function _removePendingBuyNFTInSession(uint256 _transactionId) internal {
        // Find index of transaction in pending pool
        uint index = 0;
        bool found = false;
        for (uint i = 0; i < pendingBuyNFTInSession.length; i++) {
            if (pendingBuyNFTInSession[i] == _transactionId) {
                index = i;
                found = true;
            }
        }

        require(found == true, "Transaction not found in pending pool");

        // Remove transaction id in pending pool
        if (index >= pendingBuyNFTInSession.length) {
            pendingBuyNFTInSession.pop();
        } else {
            for (uint i = index; i < pendingBuyNFTInSession.length - 1; i++) {
                pendingBuyNFTInSession[i] = pendingBuyNFTInSession[i + 1];
            }

            pendingBuyNFTInSession.pop();
        }
    }

    /**
     * @dev Remove transaction id in pending sell pool
     * @param _transactionId is id of transaction
     */
    function _removePendingSellNFTInSession(uint256 _transactionId) internal {
        // Find index of transaction in pending pool
        uint index = 0;
        bool found = false;
        for (uint i = 0; i < pendingSellNFTInSession.length; i++) {
            if (pendingSellNFTInSession[i] == _transactionId) {
                index = i;
                found = true;
            }
        }

        require(found == true, "Transaction not found in pending pool");

        // Remove transaction id in pending pool
        if (index >= pendingSellNFTInSession.length) {
            pendingSellNFTInSession.pop();
        } else {
            for (uint i = index; i < pendingSellNFTInSession.length - 1; i++) {
                pendingSellNFTInSession[i] = pendingSellNFTInSession[i + 1];
            }

            pendingSellNFTInSession.pop();
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
