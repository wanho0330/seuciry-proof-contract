// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract ProofNFT is ERC721, Ownable {
    struct Proof {
        uint256 idx;
        string[] firstImageHashes;
        string[] secondImageHashes;
    }

    uint256 private _tokenIdCounter;
    mapping(uint256 => Proof) private _proofs;

    event ProofConfirmed(uint256 indexed tokenId, uint256 idx, string firstImageHash, string secondImageHash);
    event ProofUpdated(uint256 indexed tokenId, string updateFirstImageHash, string updateSecondImageHash);

    constructor() ERC721("ProofNFT", "PNFT") Ownable(msg.sender) {
        _tokenIdCounter = 1; // 초기 토큰 ID 설정
    }

    function confirmProof(uint256 idx, string memory firstImageHash, string memory secondImageHash) external onlyOwner {
        uint256 tokenId = _tokenIdCounter;
        _tokenIdCounter++;

        // Proof 초기화 및 저장
        Proof storage proof = _proofs[tokenId];
        proof.idx = idx;
        proof.firstImageHashes.push(firstImageHash);
        proof.secondImageHashes.push(secondImageHash);

        // NFT 발행
        _safeMint(msg.sender, tokenId);

        emit ProofConfirmed(tokenId, idx, firstImageHash, secondImageHash);
    }

    function ConfirmUpdateProof(uint256 tokenId, string memory firstImageHash, string memory secondImageHash) external onlyOwner {

        // 새로운 해시 추가
        _proofs[tokenId].firstImageHashes.push(firstImageHash);
        _proofs[tokenId].secondImageHashes.push(secondImageHash);

        emit ProofUpdated(tokenId, firstImageHash, secondImageHash);
    }

    function ReadImageHashes(uint256 tokenId) external view returns (string[] memory, string[] memory) {
        return (_proofs[tokenId].firstImageHashes, _proofs[tokenId].secondImageHashes);
    }

    function ReadLatestImageHash(uint256 tokenId) external view returns (string memory, string memory) {
        uint256 lastIdx = _proofs[tokenId].firstImageHashes.length - 1;
        return (_proofs[tokenId].firstImageHashes[lastIdx], _proofs[tokenId].secondImageHashes[lastIdx]);
    }
}
