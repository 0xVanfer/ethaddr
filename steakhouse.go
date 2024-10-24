package ethaddr

// https://www.steakhouse.financial/
const SteakHouseProtocol string = "steakhouse"

// Steakhouse symbiotic restaking vault.
//
// map[network] = address
var SteakHouseRestakingVaultList = map[int64]string{
	ChainEthereum: "0xBEEF69Ac7870777598A04B2bd4771c71212E6aBc", // 0xbeef69ac7870777598a04b2bd4771c71212e6abc
}

// Same as SteakHouseRestakingVaultList
var SteakLRTList = SteakHouseRestakingVaultList
