// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

contract WealthFarming is ERC721URIStorage, AccessControl, Pausable {

    bytes32 public constant EDITOR_ROLE = keccak256("EDITOR_ROLE");

    struct PendingTransaction {
        address buyer;
        uint256 amount;
        uint256 price;
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

    IERC20 public usdcToken; // Token USDC được sử dụng để giao dịch
    uint256 public nextTokenId = 1; // ID NFT tiếp theo sẽ được phát hành
    uint256 public transactionCounter = 0; // Bộ đếm giao dịch pending
    uint256 public saleCounter = 0; // Bộ đếm giao dịch bán pending

    uint256 public totalAssetValue = 0; // Tổng giá trị tài sản
    uint256 public totalDebtValue = 0; // Tổng giá trị nợ

    address public editor; // Địa chỉ của Editor

    mapping(uint256 => PendingTransaction) public pendingTransactions; // Danh sách giao dịch T+1
    mapping(uint256 => PendingSale) public pendingSales; // Danh sách giao dịch bán T+1
    mapping(address => uint256[]) public userNFTs; // Lưu trữ NFT mà một user sở hữu
    Asset[] public assets; // Danh sách tài sản
    Debt[] public debts; // Danh sách nợ

    event EditorUpdated(address newEditor);
    event AssetAdded(string name, string code, uint256 value);
    event AssetUpdated(uint256 index, uint256 value);
    event DebtAdded(string name, string code, uint256 value);
    event DebtUpdated(uint256 index, uint256 value);
    event NAVUpdated(uint256 newNAV, uint256 timestamp);
    event NFTMinted(address indexed buyer, uint256 tokenId, uint256 price, uint256 timestamp);
    event PendingTransactionProcessed(uint256 transactionId, uint256 timestamp);
    event NFTTransferred(uint256 tokenId, address from, address to, uint256 price);
    event PendingSaleCreated(uint256 saleId, uint256 tokenId, address seller, address buyer, uint256 price, uint256 timestamp);
    event PendingSaleProcessed(uint256 saleId, uint256 timestamp);

    constructor(address usdcAddress) ERC721("WealthFarming NFT", "WFNFT") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        usdcToken = IERC20(usdcAddress);
    }


    // ---------------------- Quản lý tài sản ----------------------
    function addAsset(string memory _name, string memory _code, uint256 _value) external onlyRole(EDITOR_ROLE) {
        assets.push(Asset({value: _value, name: _name, code: _code}));
        totalAssetValue += _value;
        emit AssetAdded(_name, _code, _value);
    }

    function updateAsset(uint256 index, uint256 value) external onlyRole(EDITOR_ROLE)  {
        require(index < assets.length, "Invalid asset index");
        totalAssetValue = totalAssetValue - assets[index].value + value;
        assets[index].value = value;
        emit AssetUpdated(index, value);
    }

    function getAssets() external view returns (Asset[] memory) {
        return assets;
    }

    // ---------------------- Quản lý nợ ----------------------
    function addDebt(string memory _name, string memory _code, uint256 _value) external onlyRole(EDITOR_ROLE)  {
        debts.push(Debt({value: _value, name: _name, code: _code}));
        totalDebtValue += _value;
        emit DebtAdded(_name, _code, _value);
    }

    function updateDebt(uint256 index, uint256 value) external onlyRole(EDITOR_ROLE)  {
        require(index < debts.length, "Invalid debt index");
        totalDebtValue = totalDebtValue - debts[index].value + value;
        debts[index].value = value;
        emit DebtUpdated(index, value);
    }

    function getDebts() external view returns (Debt[] memory) {
        return debts;
    }

    // ---------------------- Tính toán NAV ----------------------
    function calculateNAV() public view returns (uint256) {
        uint256 totalNFTs = nextTokenId - 1; // Số lượng NFT đã phát hành
        require(totalNFTs > 0, "No NFTs minted yet");
        return (totalAssetValue - totalDebtValue) / totalNFTs;
    }

    // ---------------------- Phát hành và quản lý NFT ----------------------
    function mintNFT(address buyer, uint256 amount) external whenNotPaused {
        uint256 nav = calculateNAV();
        require(nav > 0, "NAV must be greater than 0");
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

    function finalizeMint(uint256 transactionId) external onlyRole(DEFAULT_ADMIN_ROLE)  {
        PendingTransaction storage txn = pendingTransactions[transactionId];
        require(!txn.processed, "Transaction already processed");
        require(block.timestamp >= txn.timestamp + 1 days, "T+1 condition not met");

        txn.processed = true;

        for (uint256 i = 0; i < txn.amount; i++) {
            uint256 tokenId = nextTokenId++;
            _mint(txn.buyer, tokenId);
            _setTokenURI(tokenId, "");
            userNFTs[txn.buyer].push(tokenId);
            emit NFTMinted(txn.buyer, tokenId, txn.price, block.timestamp);
        }

        emit PendingTransactionProcessed(transactionId, block.timestamp);
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

    // ---------------------- Tạm dừng hợp đồng ----------------------
    function pause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _pause();
    }

    function unpause() external onlyRole(DEFAULT_ADMIN_ROLE) {
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

     // Override required by Solidity
    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721URIStorage, AccessControl)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}
