package ethaddr

// Website: https://makerdao.com/
const MakerProtocol string = "maker"

// Maker token: MKR.
//
// map[network] = address.
var MakerTokenList = map[int64]string{
	ChainEthereum: "0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2", // MKR, 0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2
}

// Same as MakerTokenList.
var MKRList = MakerTokenList

// DAI
//
// map[network] = address.
var DAIList = map[int64]string{
	ChainArbitrum: "0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1", // DAI, 0xda10009cbd5d07dd0cecc66161fc93d7c9000da1
	ChainBSC:      "0x1AF3F329e8BE154074D8769D1FFa4eE058B1DBc3", // DAI, 0x1af3f329e8be154074d8769d1ffa4ee058b1dbc3
	ChainEthereum: "0x6B175474E89094C44Da98b954EedeAC495271d0F", // DAI, 0x6b175474e89094c44da98b954eedeac495271d0f
	ChainOptimism: "0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1", // DAI, 0xda10009cbd5d07dd0cecc66161fc93d7c9000da1
	ChainPolygon:  "0x8f3Cf7ad23Cd3CaDbD9735AFf958023239c6A063", // DAI, 0x8f3cf7ad23cd3cadbd9735aff958023239c6a063
}

// DAI.e. The bridged DAI.
//
// map[network] = address.
var DAIeList = map[int64]string{
	ChainAvalanche: "0xd586E7F844cEa2F87f50152665BCbc2C279D8d70", // DAI.e, 0xd586e7f844cea2f87f50152665bcbc2c279d8d70
}
