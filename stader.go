package ethaddr

import "github.com/0xVanfer/chainId"

// Docs: https://www.staderlabs.com/docs/
//
// Deployed contracts: https://www.staderlabs.com/docs-v1/Ethereum/smart-contracts
const StaderProtocol string = "stader"

// MaticX by Stader.
//
// map[network] = address.
var MaticXList = map[string]string{
	chainId.EthereumChainName: "0xf03a7eb46d01d9ecaa104558c732cf82f6b6b645", // MaticX
	chainId.PolygonChainName:  "0xfa68fb4628dff1028cfec22b4162fccd0d45efb6", // MaticX
}

// ETHx by stader.
//
// map[network] = address.
var ETHxList = map[string]string{
	chainId.EthereumChainName: "0xA35b1B31Ce002FBF2058D22F30f95D405200A15b", // ETHx
}

// Stader PoS staking contract.
//
// map[network] = address.
var StaderStakeManagerList = map[string]string{
	chainId.EthereumChainName: "0x5e3ef299fddf15eaa0432e6e66473ace8c13d908",
}

var StaderChildPoolList = map[string]string{
	chainId.PolygonChainName: "0xfd225c9e6601c9d38d8f98d8731bf59efcf8c0e3",
}
