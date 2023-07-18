package ethaddr

import "github.com/0xVanfer/chainId"

// Chain name map to wrapped native tokens.
//
// map[chain name] = address.
var WrappedNativeTokenList = map[string]string{
	chainId.AvalancheChainName: WAVAXList[chainId.AvalancheChainName],
	chainId.BNBSmartChainName:  WBNBList[chainId.BNBSmartChainName],
	chainId.EthereumChainName:  WETHList[chainId.EthereumChainName],
	chainId.PolygonChainName:   WMATICList[chainId.PolygonChainName],
}

// Deprecated: Use WrappedNativeTokenList instead.
var WrappedChainTokenList = WrappedNativeTokenList
