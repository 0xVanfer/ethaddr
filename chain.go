package ethaddr

import "github.com/0xVanfer/chainId"

// Chain name map to wrapped native tokens.
//
// map[chain name] = address.
var WrappedNativeTokenList = map[string]string{
	// ETH and ETH L2s.

	chainId.EthereumChainName:     WETHList[chainId.EthereumChainName],
	chainId.ArbitrumChainName:     WETHList[chainId.ArbitrumChainName],
	chainId.BaseChainName:         WETHList[chainId.BaseChainName],
	chainId.OptimismChainName:     WETHList[chainId.OptimismChainName],
	chainId.PolygonZkEVMChainName: WETHList[chainId.PolygonZkEVMChainName],
	chainId.ScrollChainName:       WETHList[chainId.ScrollChainName],

	chainId.AvalancheChainName: WAVAXList[chainId.AvalancheChainName],
	chainId.BNBSmartChainName:  WBNBList[chainId.BNBSmartChainName],
	chainId.FantomChainName:    WFTMList[chainId.FantomChainName],
	chainId.PolygonChainName:   WMATICList[chainId.PolygonChainName],
}
