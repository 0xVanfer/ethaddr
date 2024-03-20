package ethaddr

// Chain name map to wrapped native tokens.
//
// map[chain name] = address.
var WrappedNativeTokenList = map[string]string{
	// ETH and ETH L2s.

	chainEthereum:  WETHList[chainEthereum],
	chainArbitrum:  WETHList[chainArbitrum],
	chainBase:      WETHList[chainBase],
	chainOptimism:  WETHList[chainOptimism],
	chainPolygonZk: WETHList[chainPolygonZk],
	chainScroll:    WETHList[chainScroll],

	chainAvalanche: WAVAXList[chainAvalanche],
	chainBSC:       WBNBList[chainBSC],
	chainFantom:    WFTMList[chainFantom],
	chainPolygon:   WMATICList[chainPolygon],
}
