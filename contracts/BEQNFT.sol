// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

contract BEQNFT is ERC721URIStorage, AccessControl, Pausable{

    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    /**
    * @dev counter id for nft 
    */
    uint256 public nextTokenId = 1;

    /**
     * @dev Address representing the native token, e.g., Ether (ETH) on Ethereum.
     */
    address public constant NATIVE_TOKEN_ADDRESS = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    constructor()  ERC721("BEQ NFT", "BNFT") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    /**
     * @dev mint nft
     * @param buyer is address of buyer
     */
    function mint(address buyer) external onlyRole(MINTER_ROLE)  {
        uint256 tokenId = nextTokenId++;
        _mint(buyer, tokenId);
        _setTokenURI(tokenId, "");
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
}
