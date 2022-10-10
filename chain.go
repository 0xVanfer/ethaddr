package ethaddr

import "github.com/0xVanfer/chainId"

// Chain name map to wrapped chain tokens.
//
// map[chain name] = address.
var WrappedChainTokenList = map[string]string{
	chainId.AvalancheChainName:    WAVAXList[chainId.AvalancheChainName],
	chainId.BinanceSmartChainName: WBNBList[chainId.BinanceSmartChainName],
	chainId.EthereumChainName:     WETHList[chainId.EthereumChainName],
	chainId.PolygonChainName:      WMATICList[chainId.PolygonChainName],
}
