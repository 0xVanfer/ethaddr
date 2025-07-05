package ethaddr

// Silo staking on SEI Chain.
// Website: https://www.silostaking.io/
const SiloStakingProtocol string = "silostaking"

// SiloStaking staked SEI, iSEI.
//
// map[chainID] = address.
var ISEIList = map[int64]string{
	ChainSei: "0x5Cf6826140C1C56Ff49C808A1A75407Cd1DF9423", // iSEI, 0x5cf6826140c1c56ff49c808a1a75407cd1df9423
}
