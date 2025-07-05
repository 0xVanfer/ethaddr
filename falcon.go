package ethaddr

// Website: https://falcon.finance/
//
// Deployed contracts: https://docs.falcon.finance/resources/smart-contracts
const FalconProtocol = "falcon"

// Falcon USD. Aka USDf
//
// map[chainID] = address.
var FalconUSDList = map[int64]string{
	ChainEthereum: "0xFa2B947eEc368f42195f24F36d2aF29f7c24CeC2", // USDf, 0xfa2b947eec368f42195f24f36d2af29f7c24cec2
}

var USDfList = FalconUSDList

// Falcon sUSDf.
//
// map[chainID] = address.
var StakedFalconUSDList = map[int64]string{
	ChainEthereum: "0xc8CF6D7991f15525488b2A83Df53468D682Ba4B0", // sUSDf, 0xc8cf6d7991f15525488b2a83df53468d682ba4b0
}

var SUSDfList = StakedFalconUSDList
