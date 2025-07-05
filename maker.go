package ethaddr

// Deprecated: use SkyProtocol instead. (rebranded)
const MakerProtocol string = "maker"

// Maker was rebranded to sky on Sep.2024.
//
// Website: https://sky.money/; previously: https://makerdao.com/
const SkyProtocol string = "sky"

// Maker token: MKR.
//
// User should upgrade the MKR token to SKY token.
//
// map[chainID] = address.
var MakerTokenList = map[int64]string{
	ChainEthereum: "0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2", // MKR, 0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2
}

// Same as MakerTokenList.
var MKRList = MakerTokenList

// DAI
//
// map[chainID] = address.
var DAIList = map[int64]string{
	ChainArbitrum: "0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1", // DAI, 0xda10009cbd5d07dd0cecc66161fc93d7c9000da1
	ChainBase:     "0x50c5725949A6F0c72E6C4a641F24049A917DB0Cb", // DAI, 0x50c5725949a6f0c72e6c4a641f24049a917db0cb
	ChainBSC:      "0x1AF3F329e8BE154074D8769D1FFa4eE058B1DBc3", // DAI, 0x1af3f329e8be154074d8769d1ffa4ee058b1dbc3
	ChainEthereum: "0x6B175474E89094C44Da98b954EedeAC495271d0F", // DAI, 0x6b175474e89094c44da98b954eedeac495271d0f
	ChainOptimism: "0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1", // DAI, 0xda10009cbd5d07dd0cecc66161fc93d7c9000da1
	ChainPolygon:  "0x8f3Cf7ad23Cd3CaDbD9735AFf958023239c6A063", // DAI, 0x8f3cf7ad23cd3cadbd9735aff958023239c6a063
	ChainScroll:   "0xcA77eB3fEFe3725Dc33bccB54eDEFc3D9f764f97", // DAI, 0xca77eb3fefe3725dc33bccb54edefc3d9f764f97
}

// DAI.e. The bridged DAI.
//
// map[chainID] = address.
var DAIeList = map[int64]string{
	ChainAvalanche: "0xd586E7F844cEa2F87f50152665BCbc2C279D8d70", // DAI.e, 0xd586e7f844cea2f87f50152665bcbc2c279d8d70
}

// Upgraded version of DAI. Sky protocol (previously MakerDAO) USDS.
//
// map[chainID] = address.
var USDSList = map[int64]string{
	ChainEthereum: "0xdC035D45d973E3EC169d2276DDab16f1e407384F", // USDS, 0xdc035d45d973e3ec169d2276ddab16f1e407384f
}

// Upgraded version of sDAI. Sky protocol (previously MakerDAO) sUSDS.
//
// map[chainID] = address.
var SUSDSList = map[int64]string{
	ChainEthereum: "0xa3931d71877C0E7a3148CB7Eb4463524FEc27fbD", // sUSDS, 0xa3931d71877c0e7a3148cb7eb4463524fec27fbd
}

// 1:1 Convert DAI to USDS, or USDS to DAI.
var DAIUSDSConverterList = map[int64]string{
	ChainEthereum: "0x3225737a9Bbb6473CB4a45b7244ACa2BeFdB276A", // 0x3225737a9bbb6473cb4a45b7244aca2befdb276a
}
