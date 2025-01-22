// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract WalthFarming is ERC721URIStorage, Ownable, Pausable {
    struct PendingTransaction {
        address buyer;
        uint256 amount;
        uint256 price;
        uint256 timestamp;
        bool processed;
    }

    IERC20 public usdcToken; // Token USDC được sử dụng để giao dịch
    uint256 public nav; // Giá trị NAV hiện tại của quỹ
    uint256 public nextTokenId = 1; // ID NFT tiếp theo sẽ được phát hành
    uint256 public transactionCounter = 0; // Bộ đếm giao dịch pending

    mapping(uint256 => PendingTransaction) public pendingTransactions; // Danh sách giao dịch T+1
    mapping(address => uint256[]) public userNFTs; // Lưu trữ NFT mà một user sở hữu

    event NAVUpdated(uint256 newNAV, uint256 timestamp);
    event NFTMinted(address indexed buyer, uint256 tokenId, uint256 price, uint256 timestamp);
    event PendingTransactionProcessed(uint256 transactionId, uint256 timestamp);
    event NFTTransferred(uint256 tokenId, address from, address to, uint256 price);

    constructor(address usdcAddress) ERC721("FundNFT", "FNFT") {
        usdcToken = IERC20(usdcAddress);
    }

    // ---------------------- Quản lý NAV ----------------------
    function updateNAV(uint256 newNAV) external onlyOwner {
        nav = newNAV;
        emit NAVUpdated(newNAV, block.timestamp);
    }

    function getNAV() external view returns (uint256) {
        return nav;
    }

    // ---------------------- Phát hành và quản lý NFT ----------------------
    function mintNFT(address buyer, uint256 amount) external whenNotPaused {
        require(nav > 0, "NAV must be set before minting");
        require(amount > 0, "Amount must be greater than 0");
        uint256 totalPrice = amount * nav;

        // Chuyển USDC từ buyer vào hợp đồng
        require(usdcToken.transferFrom(buyer, address(this), totalPrice), "USDC transfer failed");

        // Tạo giao dịch pending
        pendingTransactions[transactionCounter] = PendingTransaction({
            buyer: buyer,
            amount: amount,
            price: nav,
            timestamp: block.timestamp,
            processed: false
        });

        transactionCounter++;
    }

    function finalizeMint(uint256 transactionId) external onlyOwner {
        PendingTransaction storage txn = pendingTransactions[transactionId];
        require(!txn.processed, "Transaction already processed");
        require(block.timestamp >= txn.timestamp + 1 days, "T+1 condition not met");

        txn.processed = true;

        for (uint256 i = 0; i < txn.amount; i++) {
            uint256 tokenId = nextTokenId++;
            _mint(txn.buyer, tokenId);
            userNFTs[txn.buyer].push(tokenId);
            emit NFTMinted(txn.buyer, tokenId, txn.price, block.timestamp);
        }

        emit PendingTransactionProcessed(transactionId, block.timestamp);
    }

    function getNFTDetails(uint256 tokenId) external view returns (address owner, uint256 navAtMint) {
        require(_exists(tokenId), "Token does not exist");
        return (ownerOf(tokenId), nav);
    }

    function getPendingTransactions() external view returns (PendingTransaction[] memory) {
        PendingTransaction[] memory transactions = new PendingTransaction[](transactionCounter);
        for (uint256 i = 0; i < transactionCounter; i++) {
            transactions[i] = pendingTransactions[i];
        }
        return transactions;
    }

    // ---------------------- Mua / bán NFT ----------------------
    function sellNFT(uint256 tokenId, address to, uint256 price) external whenNotPaused {
        require(ownerOf(tokenId) == msg.sender, "You do not own this token");

        // Chuyển quyền sở hữu NFT
        _transfer(msg.sender, to, tokenId);

        // Chuyển USDC từ người mua (to) sang người bán (msg.sender)
        require(usdcToken.transferFrom(to, msg.sender, price), "USDC transfer failed");

        emit NFTTransferred(tokenId, msg.sender, to, price);
    }

    function listNFTForSale(uint256 tokenId, uint256 price) external whenNotPaused {
        require(ownerOf(tokenId) == msg.sender, "You do not own this token");
        // Logic để thêm NFT vào marketplace (tuỳ chỉnh theo dApp của bạn)
    }

    function buyNFTFromUser(uint256 tokenId, uint256 price) external whenNotPaused {
        address seller = ownerOf(tokenId);
        require(seller != msg.sender, "You cannot buy your own token");

        // Chuyển USDC từ buyer sang seller
        require(usdcToken.transferFrom(msg.sender, seller, price), "USDC transfer failed");

        // Chuyển quyền sở hữu NFT
        _transfer(seller, msg.sender, tokenId);

        emit NFTTransferred(tokenId, seller, msg.sender, price);
    }

    // ---------------------- Tạm dừng hợp đồng ----------------------
    function pauseContract() external onlyOwner {
        _pause();
    }

    function unpauseContract() external onlyOwner {
        _unpause();
    }

    // ---------------------- Truy vấn phụ trợ ----------------------
    function getTokenOwner(uint256 tokenId) external view returns (address) {
        return ownerOf(tokenId);
    }

    function getUserNFTs(address user) external view returns (uint256[] memory) {
        return userNFTs[user];
    }

    function getMarketplaceListings() external view returns (uint256[] memory) {
        // Trả về danh sách NFT trên marketplace (nếu cần phát triển thêm)
    }
}
