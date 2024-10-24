package ethaddr

// Website: https://solv.finance/
const SolvProtocol string = "solv"

// Solv protocol token SOLV.
//
// map[chainID] = address
var SolvList = map[int64]string{
	ChainEthereum: "0x256F2d67e52fE834726D2DDCD8413654F5Eb8b53", // SOLV, 0x256f2d67e52fe834726d2ddcd8413654f5eb8b53
}

// Solv protocol SolvBTC.
//
// map[chainID] = address
var SolvBTCList = map[int64]string{
	ChainEthereum: "0x7A56E1C57C7475CCf742a1832B028F0456652F97", // SolvBTC, 0x7a56e1c57c7475ccf742a1832b028f0456652f97
}

// Solv protocol SolvBTC.b (bridged by `Free`).
//
// map[chainID] = address
var SolvBTCbList = map[int64]string{
	ChainScroll: "0x3Ba89d490AB1C0c9CC2313385b30710e838370a4", // SolvBTC.b, 0x3ba89d490ab1c0c9cc2313385b30710e838370a4
}

// Solv protocol SolvBTC.BBN
//
// map[chainID] = address
var SolvBTCBBNList = map[int64]string{
	ChainEthereum: "0xd9D920AA40f578ab794426F5C90F6C731D159DEf", // SolvBTC.BBN, 0xd9d920aa40f578ab794426f5c90f6c731d159def
}
