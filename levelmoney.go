package ethaddr

// Website: https://app.level.money/
//
// Deployed contracts: https://level-money.gitbook.io/docs/technical-documentation/contract-details
const LevelMoneyProtocol = "levelmoney"

// Level lvlUSD.
//
// map[chainID] = address.
var LvlUSDList = map[int64]string{
	ChainEthereum: "0x7C1156E515aA1A2E851674120074968C905aAF37", // lvlUSD, 0x7c1156e515aa1a2e851674120074968c905aaf37
}

// Level slvlUSD.
//
// map[chainID] = address.
var SlvlUSDList = map[int64]string{
	ChainEthereum: "0x4737D9b4592B40d51e110b94c9C043c6654067Ae", // slvlUSD, 0x4737d9b4592b40d51e110b94c9c043c6654067ae
}
