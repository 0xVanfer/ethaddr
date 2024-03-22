package ethaddr

// Chainlink token: LINK.
//
// map[network] = address.
var ChainlinkTokenList = map[string]string{
	ChainArbitrum:  "0xf97f4df75117a78c1A5a0DBb814Af92458539FB4", // LINK, 0xf97f4df75117a78c1a5a0dbb814af92458539fb4
	ChainAvalanche: "0x5947BB275c521040051D82396192181b413227A3", // LINK.e, 0x5947bb275c521040051d82396192181b413227a3
	ChainBSC:       "0xF8A0BF9cF54Bb92F17374d9e9A321E6a111a51bD", // LINK, 0xf8a0bf9cf54bb92f17374d9e9a321e6a111a51bd
	ChainEthereum:  "0x514910771AF9Ca656af840dff83E8264EcF986CA", // LINK, 0x514910771af9ca656af840dff83e8264ecf986ca
	ChainOptimism:  "0x350a791Bfc2C21F9Ed5d10980Dad2e2638ffa7f6", // LINK, 0x350a791bfc2c21f9ed5d10980dad2e2638ffa7f6
	ChainPolygon:   "0x53E0bca35eC356BD5ddDFebbD1Fc0fD03FaBad39", // LINK, 0x53e0bca35ec356bd5dddfebbd1fc0fd03fabad39
}

// Same as ChainlinkTokenList.
var LINKList = ChainlinkTokenList
