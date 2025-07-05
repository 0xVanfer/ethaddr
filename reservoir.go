package ethaddr

// Website: https://www.reservoir.xyz/
//
// Deployed contracts: https://docs.reservoir.xyz/security-and-compliance/smart-contract-addresses
const ReservoirProtocol string = "reservoir"

// Reservoir USD: rUSD.
//
// map[chainID] = address.
var ReservoirUSDList = map[int64]string{
	ChainEthereum: "0x09D4214C03D01F49544C0448DBE3A27f768F2b34", // rUSD, 0x09d4214c03d01f49544c0448dbe3a27f768f2b34
}

// Saving Reservoir USD: srUSD.
//
// map[chainID] = address.
var SavingReservoirUSDList = map[int64]string{
	ChainEthereum: "0x738d1115B90efa71AE468F1287fc864775e23a31", // srUSD, 0x738d1115b90efa71ae468f1287fc864775e23a31
}
