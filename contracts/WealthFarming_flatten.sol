/**
 *Submitted for verification at testnet.bscscan.com on 2025-02-11
*/

// File: @openzeppelin/contracts/utils/Context.sol


// OpenZeppelin Contracts (last updated v5.0.1) (utils/Context.sol)

pragma solidity ^0.8.20;

/**
 * @dev Provides information about the current execution context, including the
 * sender of the transaction and its data. While these are generally available
 * via msg.sender and msg.data, they should not be accessed in such a direct
 * manner, since when dealing with meta-transactions the account sending and
 * paying for execution may not be the actual sender (as far as an application
 * is concerned).
 *
 * This contract is only required for intermediate, library-like contracts.
 */
abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }

    function _contextSuffixLength() internal view virtual returns (uint256) {
        return 0;
    }
}

// File: @openzeppelin/contracts/utils/Pausable.sol


// OpenZeppelin Contracts (last updated v5.0.0) (utils/Pausable.sol)

pragma solidity ^0.8.20;


/**
 * @dev Contract module which allows children to implement an emergency stop
 * mechanism that can be triggered by an authorized account.
 *
 * This module is used through inheritance. It will make available the
 * modifiers `whenNotPaused` and `whenPaused`, which can be applied to
 * the functions of your contract. Note that they will not be pausable by
 * simply including this module, only once the modifiers are put in place.
 */
abstract contract Pausable is Context {
    bool private _paused;

    /**
     * @dev Emitted when the pause is triggered by `account`.
     */
    event Paused(address account);

    /**
     * @dev Emitted when the pause is lifted by `account`.
     */
    event Unpaused(address account);

    /**
     * @dev The operation failed because the contract is paused.
     */
    error EnforcedPause();

    /**
     * @dev The operation failed because the contract is not paused.
     */
    error ExpectedPause();

    /**
     * @dev Initializes the contract in unpaused state.
     */
    constructor() {
        _paused = false;
    }

    /**
     * @dev Modifier to make a function callable only when the contract is not paused.
     *
     * Requirements:
     *
     * - The contract must not be paused.
     */
    modifier whenNotPaused() {
        _requireNotPaused();
        _;
    }

    /**
     * @dev Modifier to make a function callable only when the contract is paused.
     *
     * Requirements:
     *
     * - The contract must be paused.
     */
    modifier whenPaused() {
        _requirePaused();
        _;
    }

    /**
     * @dev Returns true if the contract is paused, and false otherwise.
     */
    function paused() public view virtual returns (bool) {
        return _paused;
    }

    /**
     * @dev Throws if the contract is paused.
     */
    function _requireNotPaused() internal view virtual {
        if (paused()) {
            revert EnforcedPause();
        }
    }

    /**
     * @dev Throws if the contract is not paused.
     */
    function _requirePaused() internal view virtual {
        if (!paused()) {
            revert ExpectedPause();
        }
    }

    /**
     * @dev Triggers stopped state.
     *
     * Requirements:
     *
     * - The contract must not be paused.
     */
    function _pause() internal virtual whenNotPaused {
        _paused = true;
        emit Paused(_msgSender());
    }

    /**
     * @dev Returns to normal state.
     *
     * Requirements:
     *
     * - The contract must be paused.
     */
    function _unpause() internal virtual whenPaused {
        _paused = false;
        emit Unpaused(_msgSender());
    }
}

// File: @openzeppelin/contracts/token/ERC20/IERC20.sol


// OpenZeppelin Contracts (last updated v5.0.0) (token/ERC20/IERC20.sol)

pragma solidity ^0.8.20;

/**
 * @dev Interface of the ERC20 standard as defined in the EIP.
 */
interface IERC20 {
    /**
     * @dev Emitted when `value` tokens are moved from one account (`from`) to
     * another (`to`).
     *
     * Note that `value` may be zero.
     */
    event Transfer(address indexed from, address indexed to, uint256 value);

    /**
     * @dev Emitted when the allowance of a `spender` for an `owner` is set by
     * a call to {approve}. `value` is the new allowance.
     */
    event Approval(address indexed owner, address indexed spender, uint256 value);

    /**
     * @dev Returns the value of tokens in existence.
     */
    function totalSupply() external view returns (uint256);

    /**
     * @dev Returns the value of tokens owned by `account`.
     */
    function balanceOf(address account) external view returns (uint256);

    /**
     * @dev Moves a `value` amount of tokens from the caller's account to `to`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transfer(address to, uint256 value) external returns (bool);

    /**
     * @dev Returns the remaining number of tokens that `spender` will be
     * allowed to spend on behalf of `owner` through {transferFrom}. This is
     * zero by default.
     *
     * This value changes when {approve} or {transferFrom} are called.
     */
    function allowance(address owner, address spender) external view returns (uint256);

    /**
     * @dev Sets a `value` amount of tokens as the allowance of `spender` over the
     * caller's tokens.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * IMPORTANT: Beware that changing an allowance with this method brings the risk
     * that someone may use both the old and the new allowance by unfortunate
     * transaction ordering. One possible solution to mitigate this race
     * condition is to first reduce the spender's allowance to 0 and set the
     * desired value afterwards:
     * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
     *
     * Emits an {Approval} event.
     */
    function approve(address spender, uint256 value) external returns (bool);

    /**
     * @dev Moves a `value` amount of tokens from `from` to `to` using the
     * allowance mechanism. `value` is then deducted from the caller's
     * allowance.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transferFrom(address from, address to, uint256 value) external returns (bool);
}

// File: @openzeppelin/contracts/utils/introspection/IERC165.sol


// OpenZeppelin Contracts (last updated v5.0.0) (utils/introspection/IERC165.sol)

pragma solidity ^0.8.20;

/**
 * @dev Interface of the ERC165 standard, as defined in the
 * https://eips.ethereum.org/EIPS/eip-165[EIP].
 *
 * Implementers can declare support of contract interfaces, which can then be
 * queried by others ({ERC165Checker}).
 *
 * For an implementation, see {ERC165}.
 */
interface IERC165 {
    /**
     * @dev Returns true if this contract implements the interface defined by
     * `interfaceId`. See the corresponding
     * https://eips.ethereum.org/EIPS/eip-165#how-interfaces-are-identified[EIP section]
     * to learn more about how these ids are created.
     *
     * This function call must use less than 30 000 gas.
     */
    function supportsInterface(bytes4 interfaceId) external view returns (bool);
}

// File: @openzeppelin/contracts/token/ERC721/IERC721.sol


// OpenZeppelin Contracts (last updated v5.0.0) (token/ERC721/IERC721.sol)

pragma solidity ^0.8.20;


/**
 * @dev Required interface of an ERC721 compliant contract.
 */
interface IERC721 is IERC165 {
    /**
     * @dev Emitted when `tokenId` token is transferred from `from` to `to`.
     */
    event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);

    /**
     * @dev Emitted when `owner` enables `approved` to manage the `tokenId` token.
     */
    event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId);

    /**
     * @dev Emitted when `owner` enables or disables (`approved`) `operator` to manage all of its assets.
     */
    event ApprovalForAll(address indexed owner, address indexed operator, bool approved);

    /**
     * @dev Returns the number of tokens in ``owner``'s account.
     */
    function balanceOf(address owner) external view returns (uint256 balance);

    /**
     * @dev Returns the owner of the `tokenId` token.
     *
     * Requirements:
     *
     * - `tokenId` must exist.
     */
    function ownerOf(uint256 tokenId) external view returns (address owner);

    /**
     * @dev Safely transfers `tokenId` token from `from` to `to`.
     *
     * Requirements:
     *
     * - `from` cannot be the zero address.
     * - `to` cannot be the zero address.
     * - `tokenId` token must exist and be owned by `from`.
     * - If the caller is not `from`, it must be approved to move this token by either {approve} or {setApprovalForAll}.
     * - If `to` refers to a smart contract, it must implement {IERC721Receiver-onERC721Received}, which is called upon
     *   a safe transfer.
     *
     * Emits a {Transfer} event.
     */
    function safeTransferFrom(address from, address to, uint256 tokenId, bytes calldata data) external;

    /**
     * @dev Safely transfers `tokenId` token from `from` to `to`, checking first that contract recipients
     * are aware of the ERC721 protocol to prevent tokens from being forever locked.
     *
     * Requirements:
     *
     * - `from` cannot be the zero address.
     * - `to` cannot be the zero address.
     * - `tokenId` token must exist and be owned by `from`.
     * - If the caller is not `from`, it must have been allowed to move this token by either {approve} or
     *   {setApprovalForAll}.
     * - If `to` refers to a smart contract, it must implement {IERC721Receiver-onERC721Received}, which is called upon
     *   a safe transfer.
     *
     * Emits a {Transfer} event.
     */
    function safeTransferFrom(address from, address to, uint256 tokenId) external;

    /**
     * @dev Transfers `tokenId` token from `from` to `to`.
     *
     * WARNING: Note that the caller is responsible to confirm that the recipient is capable of receiving ERC721
     * or else they may be permanently lost. Usage of {safeTransferFrom} prevents loss, though the caller must
     * understand this adds an external call which potentially creates a reentrancy vulnerability.
     *
     * Requirements:
     *
     * - `from` cannot be the zero address.
     * - `to` cannot be the zero address.
     * - `tokenId` token must be owned by `from`.
     * - If the caller is not `from`, it must be approved to move this token by either {approve} or {setApprovalForAll}.
     *
     * Emits a {Transfer} event.
     */
    function transferFrom(address from, address to, uint256 tokenId) external;

    /**
     * @dev Gives permission to `to` to transfer `tokenId` token to another account.
     * The approval is cleared when the token is transferred.
     *
     * Only a single account can be approved at a time, so approving the zero address clears previous approvals.
     *
     * Requirements:
     *
     * - The caller must own the token or be an approved operator.
     * - `tokenId` must exist.
     *
     * Emits an {Approval} event.
     */
    function approve(address to, uint256 tokenId) external;

    /**
     * @dev Approve or remove `operator` as an operator for the caller.
     * Operators can call {transferFrom} or {safeTransferFrom} for any token owned by the caller.
     *
     * Requirements:
     *
     * - The `operator` cannot be the address zero.
     *
     * Emits an {ApprovalForAll} event.
     */
    function setApprovalForAll(address operator, bool approved) external;

    /**
     * @dev Returns the account approved for `tokenId` token.
     *
     * Requirements:
     *
     * - `tokenId` must exist.
     */
    function getApproved(uint256 tokenId) external view returns (address operator);

    /**
     * @dev Returns if the `operator` is allowed to manage all of the assets of `owner`.
     *
     * See {setApprovalForAll}
     */
    function isApprovedForAll(address owner, address operator) external view returns (bool);
}

// File: @openzeppelin/contracts/access/IAccessControl.sol


// OpenZeppelin Contracts (last updated v5.0.0) (access/IAccessControl.sol)

pragma solidity ^0.8.20;

/**
 * @dev External interface of AccessControl declared to support ERC165 detection.
 */
interface IAccessControl {
    /**
     * @dev The `account` is missing a role.
     */
    error AccessControlUnauthorizedAccount(address account, bytes32 neededRole);

    /**
     * @dev The caller of a function is not the expected one.
     *
     * NOTE: Don't confuse with {AccessControlUnauthorizedAccount}.
     */
    error AccessControlBadConfirmation();

    /**
     * @dev Emitted when `newAdminRole` is set as ``role``'s admin role, replacing `previousAdminRole`
     *
     * `DEFAULT_ADMIN_ROLE` is the starting admin for all roles, despite
     * {RoleAdminChanged} not being emitted signaling this.
     */
    event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole);

    /**
     * @dev Emitted when `account` is granted `role`.
     *
     * `sender` is the account that originated the contract call, an admin role
     * bearer except when using {AccessControl-_setupRole}.
     */
    event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender);

    /**
     * @dev Emitted when `account` is revoked `role`.
     *
     * `sender` is the account that originated the contract call:
     *   - if using `revokeRole`, it is the admin role bearer
     *   - if using `renounceRole`, it is the role bearer (i.e. `account`)
     */
    event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender);

    /**
     * @dev Returns `true` if `account` has been granted `role`.
     */
    function hasRole(bytes32 role, address account) external view returns (bool);

    /**
     * @dev Returns the admin role that controls `role`. See {grantRole} and
     * {revokeRole}.
     *
     * To change a role's admin, use {AccessControl-_setRoleAdmin}.
     */
    function getRoleAdmin(bytes32 role) external view returns (bytes32);

    /**
     * @dev Grants `role` to `account`.
     *
     * If `account` had not been already granted `role`, emits a {RoleGranted}
     * event.
     *
     * Requirements:
     *
     * - the caller must have ``role``'s admin role.
     */
    function grantRole(bytes32 role, address account) external;

    /**
     * @dev Revokes `role` from `account`.
     *
     * If `account` had been granted `role`, emits a {RoleRevoked} event.
     *
     * Requirements:
     *
     * - the caller must have ``role``'s admin role.
     */
    function revokeRole(bytes32 role, address account) external;

    /**
     * @dev Revokes `role` from the calling account.
     *
     * Roles are often managed via {grantRole} and {revokeRole}: this function's
     * purpose is to provide a mechanism for accounts to lose their privileges
     * if they are compromised (such as when a trusted device is misplaced).
     *
     * If the calling account had been granted `role`, emits a {RoleRevoked}
     * event.
     *
     * Requirements:
     *
     * - the caller must be `callerConfirmation`.
     */
    function renounceRole(bytes32 role, address callerConfirmation) external;
}

// File: @openzeppelin/contracts/utils/introspection/ERC165.sol


// OpenZeppelin Contracts (last updated v5.0.0) (utils/introspection/ERC165.sol)

pragma solidity ^0.8.20;


/**
 * @dev Implementation of the {IERC165} interface.
 *
 * Contracts that want to implement ERC165 should inherit from this contract and override {supportsInterface} to check
 * for the additional interface id that will be supported. For example:
 *
 * ```solidity
 * function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
 *     return interfaceId == type(MyInterface).interfaceId || super.supportsInterface(interfaceId);
 * }
 * ```
 */
abstract contract ERC165 is IERC165 {
    /**
     * @dev See {IERC165-supportsInterface}.
     */
    function supportsInterface(bytes4 interfaceId) public view virtual returns (bool) {
        return interfaceId == type(IERC165).interfaceId;
    }
}

// File: @openzeppelin/contracts/access/AccessControl.sol


// OpenZeppelin Contracts (last updated v5.0.0) (access/AccessControl.sol)

pragma solidity ^0.8.20;




/**
 * @dev Contract module that allows children to implement role-based access
 * control mechanisms. This is a lightweight version that doesn't allow enumerating role
 * members except through off-chain means by accessing the contract event logs. Some
 * applications may benefit from on-chain enumerability, for those cases see
 * {AccessControlEnumerable}.
 *
 * Roles are referred to by their `bytes32` identifier. These should be exposed
 * in the external API and be unique. The best way to achieve this is by
 * using `public constant` hash digests:
 *
 * ```solidity
 * bytes32 public constant MY_ROLE = keccak256("MY_ROLE");
 * ```
 *
 * Roles can be used to represent a set of permissions. To restrict access to a
 * function call, use {hasRole}:
 *
 * ```solidity
 * function foo() public {
 *     require(hasRole(MY_ROLE, msg.sender));
 *     ...
 * }
 * ```
 *
 * Roles can be granted and revoked dynamically via the {grantRole} and
 * {revokeRole} functions. Each role has an associated admin role, and only
 * accounts that have a role's admin role can call {grantRole} and {revokeRole}.
 *
 * By default, the admin role for all roles is `DEFAULT_ADMIN_ROLE`, which means
 * that only accounts with this role will be able to grant or revoke other
 * roles. More complex role relationships can be created by using
 * {_setRoleAdmin}.
 *
 * WARNING: The `DEFAULT_ADMIN_ROLE` is also its own admin: it has permission to
 * grant and revoke this role. Extra precautions should be taken to secure
 * accounts that have been granted it. We recommend using {AccessControlDefaultAdminRules}
 * to enforce additional security measures for this role.
 */
abstract contract AccessControl is Context, IAccessControl, ERC165 {
    struct RoleData {
        mapping(address account => bool) hasRole;
        bytes32 adminRole;
    }

    mapping(bytes32 role => RoleData) private _roles;

    bytes32 public constant DEFAULT_ADMIN_ROLE = 0x00;

    /**
     * @dev Modifier that checks that an account has a specific role. Reverts
     * with an {AccessControlUnauthorizedAccount} error including the required role.
     */
    modifier onlyRole(bytes32 role) {
        _checkRole(role);
        _;
    }

    /**
     * @dev See {IERC165-supportsInterface}.
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return interfaceId == type(IAccessControl).interfaceId || super.supportsInterface(interfaceId);
    }

    /**
     * @dev Returns `true` if `account` has been granted `role`.
     */
    function hasRole(bytes32 role, address account) public view virtual returns (bool) {
        return _roles[role].hasRole[account];
    }

    /**
     * @dev Reverts with an {AccessControlUnauthorizedAccount} error if `_msgSender()`
     * is missing `role`. Overriding this function changes the behavior of the {onlyRole} modifier.
     */
    function _checkRole(bytes32 role) internal view virtual {
        _checkRole(role, _msgSender());
    }

    /**
     * @dev Reverts with an {AccessControlUnauthorizedAccount} error if `account`
     * is missing `role`.
     */
    function _checkRole(bytes32 role, address account) internal view virtual {
        if (!hasRole(role, account)) {
            revert AccessControlUnauthorizedAccount(account, role);
        }
    }

    /**
     * @dev Returns the admin role that controls `role`. See {grantRole} and
     * {revokeRole}.
     *
     * To change a role's admin, use {_setRoleAdmin}.
     */
    function getRoleAdmin(bytes32 role) public view virtual returns (bytes32) {
        return _roles[role].adminRole;
    }

    /**
     * @dev Grants `role` to `account`.
     *
     * If `account` had not been already granted `role`, emits a {RoleGranted}
     * event.
     *
     * Requirements:
     *
     * - the caller must have ``role``'s admin role.
     *
     * May emit a {RoleGranted} event.
     */
    function grantRole(bytes32 role, address account) public virtual onlyRole(getRoleAdmin(role)) {
        _grantRole(role, account);
    }

    /**
     * @dev Revokes `role` from `account`.
     *
     * If `account` had been granted `role`, emits a {RoleRevoked} event.
     *
     * Requirements:
     *
     * - the caller must have ``role``'s admin role.
     *
     * May emit a {RoleRevoked} event.
     */
    function revokeRole(bytes32 role, address account) public virtual onlyRole(getRoleAdmin(role)) {
        _revokeRole(role, account);
    }

    /**
     * @dev Revokes `role` from the calling account.
     *
     * Roles are often managed via {grantRole} and {revokeRole}: this function's
     * purpose is to provide a mechanism for accounts to lose their privileges
     * if they are compromised (such as when a trusted device is misplaced).
     *
     * If the calling account had been revoked `role`, emits a {RoleRevoked}
     * event.
     *
     * Requirements:
     *
     * - the caller must be `callerConfirmation`.
     *
     * May emit a {RoleRevoked} event.
     */
    function renounceRole(bytes32 role, address callerConfirmation) public virtual {
        if (callerConfirmation != _msgSender()) {
            revert AccessControlBadConfirmation();
        }

        _revokeRole(role, callerConfirmation);
    }

    /**
     * @dev Sets `adminRole` as ``role``'s admin role.
     *
     * Emits a {RoleAdminChanged} event.
     */
    function _setRoleAdmin(bytes32 role, bytes32 adminRole) internal virtual {
        bytes32 previousAdminRole = getRoleAdmin(role);
        _roles[role].adminRole = adminRole;
        emit RoleAdminChanged(role, previousAdminRole, adminRole);
    }

    /**
     * @dev Attempts to grant `role` to `account` and returns a boolean indicating if `role` was granted.
     *
     * Internal function without access restriction.
     *
     * May emit a {RoleGranted} event.
     */
    function _grantRole(bytes32 role, address account) internal virtual returns (bool) {
        if (!hasRole(role, account)) {
            _roles[role].hasRole[account] = true;
            emit RoleGranted(role, account, _msgSender());
            return true;
        } else {
            return false;
        }
    }

    /**
     * @dev Attempts to revoke `role` to `account` and returns a boolean indicating if `role` was revoked.
     *
     * Internal function without access restriction.
     *
     * May emit a {RoleRevoked} event.
     */
    function _revokeRole(bytes32 role, address account) internal virtual returns (bool) {
        if (hasRole(role, account)) {
            _roles[role].hasRole[account] = false;
            emit RoleRevoked(role, account, _msgSender());
            return true;
        } else {
            return false;
        }
    }
}

// File: @openzeppelin/contracts/security/ReentrancyGuard.sol


// OpenZeppelin Contracts (last updated v4.9.0) (security/ReentrancyGuard.sol)

pragma solidity ^0.8.0;

/**
 * @dev Contract module that helps prevent reentrant calls to a function.
 *
 * Inheriting from `ReentrancyGuard` will make the {nonReentrant} modifier
 * available, which can be applied to functions to make sure there are no nested
 * (reentrant) calls to them.
 *
 * Note that because there is a single `nonReentrant` guard, functions marked as
 * `nonReentrant` may not call one another. This can be worked around by making
 * those functions `private`, and then adding `external` `nonReentrant` entry
 * points to them.
 *
 * TIP: If you would like to learn more about reentrancy and alternative ways
 * to protect against it, check out our blog post
 * https://blog.openzeppelin.com/reentrancy-after-istanbul/[Reentrancy After Istanbul].
 */
abstract contract ReentrancyGuard {
    // Booleans are more expensive than uint256 or any type that takes up a full
    // word because each write operation emits an extra SLOAD to first read the
    // slot's contents, replace the bits taken up by the boolean, and then write
    // back. This is the compiler's defense against contract upgrades and
    // pointer aliasing, and it cannot be disabled.

    // The values being non-zero value makes deployment a bit more expensive,
    // but in exchange the refund on every call to nonReentrant will be lower in
    // amount. Since refunds are capped to a percentage of the total
    // transaction's gas, it is best to keep them low in cases like this one, to
    // increase the likelihood of the full refund coming into effect.
    uint256 private constant _NOT_ENTERED = 1;
    uint256 private constant _ENTERED = 2;

    uint256 private _status;

    constructor() {
        _status = _NOT_ENTERED;
    }

    /**
     * @dev Prevents a contract from calling itself, directly or indirectly.
     * Calling a `nonReentrant` function from another `nonReentrant`
     * function is not supported. It is possible to prevent this from happening
     * by making the `nonReentrant` function external, and making it call a
     * `private` function that does the actual work.
     */
    modifier nonReentrant() {
        _nonReentrantBefore();
        _;
        _nonReentrantAfter();
    }

    function _nonReentrantBefore() private {
        // On the first call to nonReentrant, _status will be _NOT_ENTERED
        require(_status != _ENTERED, "ReentrancyGuard: reentrant call");

        // Any calls to nonReentrant after this point will fail
        _status = _ENTERED;
    }

    function _nonReentrantAfter() private {
        // By storing the original value once again, a refund is triggered (see
        // https://eips.ethereum.org/EIPS/eip-2200)
        _status = _NOT_ENTERED;
    }

    /**
     * @dev Returns true if the reentrancy guard is currently set to "entered", which indicates there is a
     * `nonReentrant` function in the call stack.
     */
    function _reentrancyGuardEntered() internal view returns (bool) {
        return _status == _ENTERED;
    }
}

// File: contracts/interfaces/IBEQNFT.sol


pragma solidity ^0.8.24;


/*
 * @author Beq Holding
 * @title IBEQNFT
 * @dev Interface for interacting with the BEQNFT contract.
 */
interface IBEQNFT is IERC721 {

    function getNextTokenId() external view returns (uint256);

    function mint(address buyer) external;

    function getTokenOwner(uint256 tokenId) external view returns (address);
}

// File: contracts/WealthFarming.sol


pragma solidity ^0.8.24;








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

    IBEQNFT public nftToken = IBEQNFT(0x22a1FC12a8CEaeCcB39136391cd3D22DD7d7BbBa);

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
    uint256 public ceil = 70;

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
        
        // Calculate nft price
        uint256 currentPrice = (_totalAsset + totalBuyInSession - totalSellInSession) / totalNFT;

        // Check limit decrease or increase
        if (sessionHistory.length > 0) {
            uint256 lastPrice = session.nftPrice;

            if (currentPrice > lastPrice) {
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
        while (lockNFT.length > 0) {
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