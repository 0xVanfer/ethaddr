package ethaddr

// Aave token: AAVE.
//
// NOT STABLE for avalanche. If AAVE is later deployed on avalanche, this address might be changed.
//
// map[network] = address.
//
// More details are seperated into `aaveV2.go` & `aaveV3.go`.
var AaveTokenList = map[int64]string{
	ChainArbitrum:  "0xba5DdD1f9d7F570dc94a51479a000E3BCE967196", // AAVE, 0xba5ddd1f9d7f570dc94a51479a000e3bce967196
	ChainAvalanche: "0x63a72806098Bd3D9520cC43356dD78afe5D386D9", // AAVE.e, 0x63a72806098bd3d9520cc43356dd78afe5d386d9
	ChainBSC:       "0xfb6115445Bff7b52FeB98650C87f44907E58f802", // AAVE, 0xfb6115445bff7b52feb98650c87f44907e58f802
	ChainEthereum:  "0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9", // AAVE, 0x7fc66500c84a76ad7e9c93437bfc5ac33e2ddae9
	ChainOptimism:  "0x76FB31fb4af56892A25e32cFC43De717950c9278", // AAVE, 0x76fb31fb4af56892a25e32cfc43de717950c9278
	ChainPolygon:   "0xD6DF932A45C0f255f85145f286eA0b292B21C90B", // AAVE, 0xd6df932a45c0f255f85145f286ea0b292b21c90b
}

// Aave token on avalanche: AAVE.e.
//
// map[network] = address.
var AAVEeList = map[int64]string{
	ChainAvalanche: "0x63a72806098Bd3D9520cC43356dD78afe5D386D9", // AAVE.e, 0x63a72806098bd3d9520cc43356dd78afe5d386d9
}

// Same as AaveTokenList.
var AAVEList = AaveTokenList
