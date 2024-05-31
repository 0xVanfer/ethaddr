package ethaddr

// Chain name map to wrapped native tokens.
//
// map[chain name] = address.
var WrappedNativeTokenList = map[int64]string{
	// ETH and ETH L2s.

	ChainEthereum:  WETHList[ChainEthereum],
	ChainArbitrum:  WETHList[ChainArbitrum],
	ChainBase:      WETHList[ChainBase],
	ChainOptimism:  WETHList[ChainOptimism],
	ChainPolygonZk: WETHList[ChainPolygonZk],
	ChainScroll:    WETHList[ChainScroll],

	ChainAvalanche: WAVAXList[ChainAvalanche],
	ChainBSC:       WBNBList[ChainBSC],
	ChainFantom:    WFTMList[ChainFantom],
	ChainPolygon:   WMATICList[ChainPolygon],
}
