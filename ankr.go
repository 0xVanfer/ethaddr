package ethaddr

const AnkrProtocol string = "ankr"

// Ankr staked ETH: ankrETH.
//
// map[chainID] = address.
var AnkrETHList = map[int64]string{
	ChainEthereum: "0xE95A203B1a91a908F9B9CE46459d101078c2c3cb", // ankrETH
}
