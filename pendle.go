package ethaddr

// Website: https://www.pendle.finance/
//
// Api: https://api-v2.pendle.finance/api/core/docs
//
// Docs: https://docs.pendle.finance/Developers/Overview
//
// Deployed contracts: https://github.com/pendle-finance/pendle-core-v2-public/blob/main/deployments/1-core.json
const PendleProtocol string = "pendle"

// Pendle token: PENDLE.
//
// map[network] = address.
var PENDLEList = map[int64]string{
	ChainArbitrum: "0x0c880f6761F1af8d9Aa9C466984b80DAb9a8c9e8", // PENDLE, 0x0c880f6761f1af8d9aa9c466984b80dab9a8c9e8
	ChainEthereum: "0x808507121B80c02388fAd14726482e061B8da827", // PENDLE, 0x808507121b80c02388fad14726482e061b8da827
}

// Pendle Router V4.
//
// map[network] = address.
var PendleRouterV4List = map[int64]string{
	ChainEthereum: "0x888888888889758F76e7103c6CbF23ABbF58F946", // 0x888888888889758f76e7103c6cbf23abbf58f946
}

// Pendle Router Static, similar to aave ui data provider.
//
// map[network] = address.
var PendleRouterStaticList = map[int64]string{
	ChainEthereum: "0x263833d47eA3fA4a30f269323aba6a107f9eB14C", // 0x263833d47ea3fa4a30f269323aba6a107f9eb14c
}

// Pendle vePendle. Not a ERC20.
//
// map[network] = address.
var VePENDLE = map[int64]string{
	ChainEthereum: "0x4f30A9D41B80ecC5B94306AB4364951AE3170210", // 0x4f30a9d41b80ecc5b94306ab4364951ae3170210
}
