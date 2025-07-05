package ethaddr

// Website: https://tether.to/
const TetherProtocol string = "tether"

// Website: https://docs.usdt0.to/
const USDT0Protocol string = "usdt0"

// USDT
//
// map[chainID] = address.
var USDTList = map[int64]string{
	ChainArbitrum:  "0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9", // USD₮0, 0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9
	ChainAvalanche: "0x9702230A8Ea53601f5cD2dc00fDBc13d4dF4A8c7", // USDt, 0x9702230a8ea53601f5cd2dc00fdbc13d4df4a8c7
	ChainBase:      "0xfde4C96c8593536E31F229EA8f37b2ADa2699bb2", // USDT, 0xfde4c96c8593536e31f229ea8f37b2ada2699bb2
	ChainBSC:       "0x55d398326f99059fF775485246999027B3197955", // USDT, 0x55d398326f99059ff775485246999027b3197955
	ChainEthereum:  "0xdAC17F958D2ee523a2206206994597C13D831ec7", // USDT, 0xdac17f958d2ee523a2206206994597c13d831ec7
	ChainOptimism:  "0x94b008aA00579c1307B0EF2c499aD98a8ce58e58", // USDT, 0x94b008aa00579c1307b0ef2c499ad98a8ce58e58
	ChainPolygon:   "0xc2132D05D31c914a87C6611C10748AEb04B58e8F", // USDT, 0xc2132d05d31c914a87c6611c10748aeb04b58e8f
	ChainScroll:    "0xf55BEC9cafDbE8730f096Aa55dad6D22d44099Df", // USDT, 0xf55bec9cafdbe8730f096aa55dad6d22d44099df
	ChainSei:       "0xB75D0B03c06A926e488e2659DF1A861F860bD3d1", // USDT, 0xb75d0b03c06a926e488e2659df1a861f860bd3d1; name: USDT.kava, bridged by IBC
	ChainTaiko:     "0x2DEF195713CF4a606B49D07E520e22C17899a736", // USDT, 0x2def195713cf4a606b49d07e520e22c17899a736
}

// USDT.e for avalanche
//
// map[chainID] = address.
var USDTeList = map[int64]string{
	ChainAvalanche: "0xc7198437980c041c805A1EDcbA50c1Ce5db95118", // USDT.e, 0xc7198437980c041c805a1edcba50c1ce5db95118
}

var USDT0List = map[int64]string{
	ChainArbitrum: "0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9", // USD₮0, 0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9
	ChainOptimism: "0x01bFF41798a0BcF287b996046Ca68b395DbC1071", // USD₮0, 0x01bff41798a0bcf287b996046ca68b395dbc1071
	ChainSei:      "0x9151434b16b9763660705744891fA906F660EcC5", // USD₮0, 0x9151434b16b9763660705744891fa906f660ecc5
}
