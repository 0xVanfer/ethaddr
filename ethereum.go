package ethaddr

// Wrapped eth: WETH.
//
// NOT STABLE for avalanche. If WETH is later deployed on avalanche, this address might be changed.
//
// map[network] = address.
var WETHList = map[string]string{
	ChainArbitrum:  "0x82aF49447D8a07e3bd95BD0d56f35241523fBab1", // WETH, 0x82af49447d8a07e3bd95bd0d56f35241523fbab1
	ChainAvalanche: "0x49D5c2BdFfac6CE2BFdB6640F4F80f226bc10bAB", // WETH.e, 0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab
	ChainBase:      "0x4200000000000000000000000000000000000006", // WETH, 0x4200000000000000000000000000000000000006
	ChainEthereum:  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", // WETH, 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2
	ChainOptimism:  "0x4200000000000000000000000000000000000006", // WETH, 0x4200000000000000000000000000000000000006
	ChainPolygon:   "0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619", // WETH, 0x7ceb23fd6bc0add59e62ac25578270cff1b9f619
	ChainPolygonZk: "0x4F9A0e7FD2Bf6067db6994CF12E4495Df938E6e9", // WETH, 0x4f9a0e7fd2bf6067db6994cf12e4495df938e6e9
	ChainScroll:    "0x5300000000000000000000000000000000000004", // WETH, 0x5300000000000000000000000000000000000004
}

// Wrapped eth on avalanche: WETH.e.
//
// map[network] = address.
var WETHeList = map[string]string{
	ChainAvalanche: "0x49D5c2BdFfac6CE2BFdB6640F4F80f226bc10bAB", // WETH.e, 0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab
}

// BNB chain ETH is BEP20 token.
//
// map[network] = address.
var ETHList = map[string]string{
	ChainBSC: "0x2170Ed0880ac9A755fd29B2688956BD959F933F8", // ETH, 0x2170ed0880ac9a755fd29b2688956bd959f933f8
}
