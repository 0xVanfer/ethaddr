package ethaddr

// Website: https://lista.org/
const ListaProtocol string = "lista"

// Lista lending. Aka moolah.
const MoolahProtocol string = "moolah"

// Lista dao staked BNB, slisBNB.
//
// Minted on BSC and LZ OFT to ethereum.
//
// map[chainID] = address.
var SlisBNBList = map[int64]string{
	ChainBSC:      "0xB0b84D294e0C75A6abe60171b70edEb2EFd14A1B", // slisBNB, 0xb0b84d294e0c75a6abe60171b70edeb2efd14a1b
	ChainEthereum: "0x44388Ef3bc730BDE8670a3b4831281dd7E89C584", // slisBNB, 0x44388ef3bc730bde8670a3b4831281dd7e89c584
}

// Lista dao collateral BNB, clisBNB.
//
// map[chainID] = address.
var ClisBNBList = map[int64]string{
	ChainBSC: "0x4b30fcAA7945fE9fDEFD2895aae539ba102Ed6F6", // clisBNB, 0x4b30fcaa7945fe9fdefd2895aae539ba102ed6f6
}

// Lista dao collateral ETH, clisETH.
//
// map[chainID] = address.
var ClisETHList = map[int64]string{
	ChainBSC: "0x620e897d529EfECa41E57d0344eBB24f75864D1E", // clisETH, 0x620e897d529efeca41e57d0344ebb24f75864d1e
}

// Lista dao staked BNB stake manager.
//
// map[chainID] = address.
var SlisBNBStakeManagerList = map[int64]string{
	ChainBSC: "0x1adB950d8bB3dA4bE104211D5AB038628e477fE6", // 0x1adb950d8bb3da4be104211d5ab038628e477fe6
}

// Lista USD, lisUSD
//
// map[chainID] = address.
var LisUSDList = map[int64]string{
	ChainBSC: "0x0782b6d8c4551B9760e74c0545a9bCD90bdc41E5", // lisUSD, 0x0782b6d8c4551b9760e74c0545a9bcd90bdc41e5
}

// Lista dao lending pool. Aka Moolah.
//
// map[chainID] = address.
var ListaLendingPoolList = map[int64]string{
	ChainBSC: "0x8F73b65B4caAf64FBA2aF91cC5D4a2A1318E5D8C", // 0x8f73b65b4caaf64fba2af91cc5d4a2a1318e5d8c
}
