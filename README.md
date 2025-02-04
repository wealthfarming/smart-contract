# Smart Contract Documentation


### 1. Build contract to Go package

Step 1: Build the contract
```
solc --abi contracts/WealthFarming_flatten.sol -o build
```

Step 2: Generate to Golang package
```
abigen --abi=./build/WealthFarming.abi --pkg=farming --out=farming.go
```