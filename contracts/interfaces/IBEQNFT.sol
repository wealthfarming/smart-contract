// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

/*
 * @author Beq Holding
 * @title IBEQNFT
 * @dev Interface for interacting with the BEQNFT contract.
 */
interface IBEQNFT is IERC721 {

    function mint(address buyer) external;

    function getTokenOwner(uint256 tokenId) external view returns (address);
}
