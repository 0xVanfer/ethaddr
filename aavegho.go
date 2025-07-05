package ethaddr

// The GHO token launched by Aave.
//
// map[chainID] = address
//
// Docs: https://docs.gho.xyz/concepts/overview
//
// Github: https://github.com/aave/gho-core
var GHOList = map[int64]string{
	ChainArbitrum:  "0x7dfF72693f6A4149b17e7C6314655f6A9F7c8B33", // GHO, 0x7dff72693f6a4149b17e7c6314655f6a9f7c8b33
	ChainAvalanche: "0xfc421aD3C883Bf9E7C4f42dE845C4e4405799e73", // GHO, 0xfc421ad3c883bf9e7c4f42de845c4e4405799e73
	ChainEthereum:  "0x40D16FC0246aD3160Ccc09B8D0D3A2cD28aE6C2f", // GHO, 0x40d16fc0246ad3160ccc09b8d0d3a2cd28ae6c2f
	ChainBase:      "0x6Bb7a212910682DCFdbd5BCBb3e28FB4E8da10Ee", // GHO, 0x6bb7a212910682dcfdbd5bcbb3e28fb4e8da10ee
}

// GHO is similar to DAI, Aave provide functions to flash mint/burn instead of flash loan/repay.
//
// map[chainID] = address
var GHOFlashMinterAddress = map[int64]string{
	ChainEthereum: "0xb639D208Bcf0589D54FaC24E655C79EC529762B8", // 0xb639d208bcf0589d54fac24e655c79ec529762b8
}

// Data provider of GHO token as a reserve within the Aave Protocol
//
// map[chainID] = address
var GHOUIDataProviderAddress = map[int64]string{
	ChainEthereum: "0x379c1EDD1A41218bdbFf960a9d5AD2818Bf61aE8", // 0x379c1edd1a41218bdbff960a9d5ad2818bf61ae8
}
