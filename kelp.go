package ethaddr

// https://stage.kelpdao.xyz/
const KelpProtocol string = "kelp"

// Kelp restaking ETH: rsETH.
//
// Other chains except Ethereum use layerzero to bridge the original rsETH on Ethereum.
// Therefore, the addresses use different contracts and abis.
//
// map[chainID] = address
var RsETHList = map[int64]string{
	ChainArbitrum: "0x4186BFC76E2E237523CBC30FD220FE055156b41F", // rsETH, 0x4186bfc76e2e237523cbc30fd220fe055156b41f
	ChainBase:     "0x1Bc71130A0e39942a7658878169764Bbd8A45993", // rsETH, 0x1bc71130a0e39942a7658878169764bbd8a45993
	ChainBlast:    "0x4186BFC76E2E237523CBC30FD220FE055156b41F", // rsETH, 0x4186bfc76e2e237523cbc30fd220fe055156b41f
	ChainEthereum: "0xA1290d69c65A6Fe4DF752f95823fae25cB99e5A7", // rsETH, 0xa1290d69c65a6fe4df752f95823fae25cb99e5a7
	ChainLinea:    "0x4186BFC76E2E237523CBC30FD220FE055156b41F", // rsETH, 0x4186bfc76e2e237523cbc30fd220fe055156b41f
	ChainOptimism: "0x4186BFC76E2E237523CBC30FD220FE055156b41F", // rsETH, 0x4186bfc76e2e237523cbc30fd220fe055156b41f
	ChainScroll:   "0x65421ba909200b81640d98B979d07487C9781B66", // rsETH, 0x65421ba909200b81640d98b979d07487c9781b66
}

// Kelp wrapped rsETH for L2s: wrsETH.
//
// wrsETH can be converted 1:1 to rsETH.
//
// map[chainID] = address
var WrsETHList = map[int64]string{
	ChainBlast:    "0xe7903B1F75C534Dd8159b313d92cDCfbC62cB3Cd", // wrsETH, 0xe7903b1f75c534dd8159b313d92cdcfbc62cb3cd
	ChainBase:     "0xEDfa23602D0EC14714057867A78d01e94176BEA0", // wrsETH, 0xedfa23602d0ec14714057867a78d01e94176bea0
	ChainLinea:    "0xD2671165570f41BBB3B0097893300b6EB6101E6C", // wrsETH, 0xd2671165570f41bbb3b0097893300b6eb6101e6c
	ChainScroll:   "0xa25b25548B4C98B0c7d3d27dcA5D5ca743d68b7F", // wrsETH, 0xa25b25548b4c98b0c7d3d27dca5d5ca743d68b7f
	ChainOptimism: "0x87eEE96D50Fb761AD85B1c982d28A042169d61b1", // wrsETH, 0x87eee96d50fb761ad85b1c982d28a042169d61b1
}

// Kelp Deposit Pools.
//
// For different chains, the contracts are not the same.
//
// map[chainID] = address
var KelpDepositPools = map[int64]string{
	ChainArbitrum: "0x376A7564AF88242D6B8598A5cfdD2E9759711B61", // 0x376a7564af88242d6b8598a5cfdd2e9759711b61
	ChainBase:     "0x291088312150482826b3A37d5A69a4c54DAa9118", // 0x291088312150482826b3a37d5a69a4c54daa9118
	ChainBlast:    "0x1558959f1a032F83f24A14Ff539944A926C51bdf", // 0x1558959f1a032f83f24a14ff539944a926c51bdf
	ChainEthereum: "0x036676389e48133B63a802f8635AD39E752D375D", // 0x036676389e48133b63a802f8635ad39e752d375d
	ChainLinea:    "0x057297e44A3364139EDCF3e1594d6917eD7688c2", // 0x057297e44a3364139edcf3e1594d6917ed7688c2
	ChainOptimism: "0xaAA687e218F9B53183A6AA9639FBD9D6e69EcB73", // 0xaaa687e218f9b53183a6aa9639fbd9d6e69ecb73
	ChainScroll:   "0xb80deaecd7F4Bca934DE201B11a8711644156a0a", // 0xb80deaecd7f4bca934de201b11a8711644156a0a
}

// Kelp gain ETH. agETH.
//
// https://stage.kelpdao.xyz/gain/
//
// map[chainID] = address
var KelpGainETHList = map[int64]string{
	ChainEthereum: "0xe1B4d34E8754600962Cd944B535180Bd758E6c2e", // agETH, 0xe1b4d34e8754600962cd944b535180bd758e6c2e
}

var AgETHList = KelpGainETHList
