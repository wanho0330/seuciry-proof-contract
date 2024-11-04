#!/bin/bash

# Compile Solidity contracts using solc
solc --abi --bin solidity/*.sol --base-path /Users/wanho/GolandProjects/security-proof-contract --include-path node_modules --output-dir solidity/output/ --overwrite
echo "Solidity contracts compiled to ABI and BIN files in solidity/output/."

# Generate Go bindings for ProofNFT contract using abigen
abigen --abi /Users/wanho/GolandProjects/security-proof-contract/solidity/output/ProofNFT.abi \
       --bin /Users/wanho/GolandProjects/security-proof-contract/solidity/output/ProofNFT.bin \
       --pkg proof \
       --out /Users/wanho/GolandProjects/security-proof-contract/pkg/contract/proof/proofNFT.go

echo "Go bindings generated for ProofNFT contract at pkg/contract/proof/proofNFT.go."
