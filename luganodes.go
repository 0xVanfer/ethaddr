package ethaddr

// Website: https://www.luganodes.com/
const LuganodesProtocol string = "luganodes"

// Luganodes symbiotic restaked ETH: LugaETH.
//
// map[chainID] = address
var LuganodesRestakedETHList = map[int64]string{
	ChainEthereum: "0x82dc3260f599f4fC4307209A1122B6eAa007163b", // LugaETH, 0x82dc3260f599f4fc4307209a1122b6eaa007163b
}

// Same as LuganodesRestakedETHList.
var LugaETHList = LuganodesRestakedETHList
