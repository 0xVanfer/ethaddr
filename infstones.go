package ethaddr

// https://infstones.com/
const InfStonesProtocol string = "infstones"

// InfStones symbiotic restaked ETH: ifsETH.
//
// map[chainID] = address
var InfStonesRestakedETHList = map[int64]string{
	ChainEthereum: "0x49cd586dd9BA227Be9654C735A659a1dB08232a9", // ifsETH, 0x49cd586dd9ba227be9654c735a659a1db08232a9
}

// Same as InfStonesRestakedETHList.
var IfsETHList = InfStonesRestakedETHList
