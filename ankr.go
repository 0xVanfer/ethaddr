package ethaddr

// Website: https://www.ankr.com/
const AnkrProtocol string = "ankr"

// Ankr token ANKR.
//
// map[chainID] = address.
var ANKRList = map[int64]string{
	ChainBSC: "0xf307910A4c7bbc79691fD374889b36d8531B08e3", // 0xf307910a4c7bbc79691fd374889b36d8531b08e3
}

// Ankr staked ETH: ankrETH.
//
// map[chainID] = address.
var AnkrETHList = map[int64]string{
	ChainEthereum: "0xE95A203B1a91a908F9B9CE46459d101078c2c3cb", // ankrETH
}

// Ankr staked BNB: ankrBNB.
//
// map[chainID] = address.
var AnkrBNBList = map[int64]string{
	ChainBSC: "0x52F24a5e03aee338Da5fd9Df68D2b6FAe1178827", // ankrBNB
}
