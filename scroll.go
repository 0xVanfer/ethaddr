package ethaddr

// Blockscan: https://scrollscan.com/
const ChainNameScroll string = "scroll"

var ScrollL1StandardERC20GateWay = map[int64]string{
	ChainEthereum: "0xD8A791fE2bE73eb6E6cF1eb0cb3F36adC9B3F8f9", // 0xd8a791fe2be73eb6e6cf1eb0cb3f36adc9b3f8f9
}

// Scroll token SCR.
//
// map[chainID] = address
var SCRList = map[int64]string{
	ChainScroll: "0xd29687c813D741E2F938F4aC377128810E217b1b", // 0xd29687c813d741e2f938f4ac377128810e217b1b
}
