package ethaddr

import "github.com/0xVanfer/chainId"

// Chain name map to wrapped chain tokens.
var WrappedChainTokenList = map[string]string{
	chainId.AvalancheChainName:     WAVAXList[chainId.AvalancheChainName],
	chainId.BinanaceSmartChainName: WBNBList[chainId.BinanaceSmartChainName],
	chainId.EthereumChainName:      WETHList[chainId.EthereumChainName],
	chainId.PolygonChainName:       WMATICList[chainId.PolygonChainName],
}
