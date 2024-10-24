package ethaddr

// Website: https://www.re7.capital/
const Re7Protocol string = "re7"

// Re7 symbiotic restaked vault: Re7LRT.
//
// map[chainID] = address
var Re7SymbioticLRTList = map[int64]string{
	ChainEthereum: "0x84631c0d0081FDe56DeB72F6DE77abBbF6A9f93a", // Re7LRT, 0x84631c0d0081fde56deb72f6de77abbbf6a9f93a
}

// Same as Re7SymbioticLRTList.
var Re7LRTList = Re7SymbioticLRTList
