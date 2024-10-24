package ethaddr

// Website: https://aave.com/
//
// Docs: https://docs.aave.com/developers/getting-started/readme
//
// Deployed Addresses: https://github.com/bgd-labs/aave-address-book/blob/main/src/ts/AaveV3EthereumEtherFi.ts
//
// Fork of AaveV3 for etherfi.
const AaveV3ProtocolEtherfiFork string = "aavev3etherfi"

// Aave rewards controller V3 for etherfi fork.
// Previously named as "incentives controller".
//
// map[network] = address.
var AaveIncentivesControllerV3EtherfiForkList = map[int64]string{
	ChainEthereum: "0x8164Cc65827dcFe994AB23944CBC90e0aa80bFcb", // 0x8164cc65827dcfe994ab23944cbc90e0aa80bfcb
}

// Aave lending pool v3 for etherfi fork.
//
// map[network] = address.
var AaveLendingPoolV3EtherfiForkList = map[int64]string{
	ChainEthereum: "0x0AA97c284e98396202b6A04024F5E2c65026F3c0", // 0x0aa97c284e98396202b6a04024f5e2c65026f3c0
}

// Aave protocol data provider v3 for etherfi fork.
// Previously named as "pool data provider".
//
// map[network] = address.
//
// Last updated Oct.24.2024. Aave v3.2.0
var AavePoolDataProviderV3EtherfiForkList = map[int64]string{
	ChainEthereum: "0xE7d490885A68f00d9886508DF281D67263ed5758", // 0xe7d490885a68f00d9886508df281d67263ed5758
}

// Aave ui pool data provider v3 for etherfi fork.
//
// map[network] = address.
//
// Last updated Oct.24.2024. Aave v3.2.0
var AaveUiPoolDataProviderV3EtherfiForkList = map[int64]string{
	ChainEthereum: "0x3F78BBD206e4D3c504Eb854232EdA7e47E9Fd8FC", // 0x3f78bbd206e4d3c504eb854232eda7e47e9fd8fc
}

// Aave ui incentive data provider v3 for etherfi fork.
//
// map[network] = address.
//
// Last updated Oct.24.2024. Aave v3.2.0
var AaveUiIncentiveDataProviderV3EtherfiForkList = map[int64]string{
	ChainEthereum: "0xe3dFf4052F0bF6134ACb73bEaE8fe2317d71F047", // 0xe3dff4052f0bf6134acb73beae8fe2317d71f047
}

// Aave pool address provider v3 for etherfi fork.
//
// map[network] = address.
var AavePoolAddressProviderV3EtherfiForkList = map[int64]string{
	ChainEthereum: "0xeBa440B438Ad808101d1c451C1C5322c90BEFCdA", // 0xeba440b438ad808101d1c451c1c5322c90befcda
}

// Aave a tokens v3 for etherfi fork.
//
// map[network][underlying] = address.
var AaveATokenV3EtherfiForkList = map[int64]map[string]string{
	ChainEthereum: {
		WEETHList[ChainEthereum]: "0xbe1F842e7e0afd2c2322aae5d34bA899544b29db", // 0xbe1f842e7e0afd2c2322aae5d34ba899544b29db
		USDCList[ChainEthereum]:  "0x7380c583cDe4409eFF5DD3320D93a45D96B80E2e", // 0x7380c583cde4409eff5dd3320d93a45d96b80e2e
		PyUSDList[ChainEthereum]: "0xdF7f48892244C6106EA784609f7de10AB36F9c7e", // 0xdf7f48892244c6106ea784609f7de10ab36f9c7e
		FRAXList[ChainEthereum]:  "0x6914ECCf50837dC61b43ee478a9BD9B439648956", // 0x6914eccf50837dc61b43ee478a9bd9b439648956
	},
}

// Aave v tokens v3 for etherfi fork.
//
// map[network][underlying] = address.
var AaveVTokenV3EtherfiForkList = map[int64]map[string]string{
	ChainEthereum: {
		WEETHList[ChainEthereum]: "0x16264412CB72F0d16A446f7D928Dd0D822810048", // 0x16264412cb72f0d16a446f7d928dd0d822810048
		USDCList[ChainEthereum]:  "0x9355032d747f1e08F8720CD01950E652eE15cdB7", // 0x9355032d747f1e08f8720cd01950e652ee15cdb7
		PyUSDList[ChainEthereum]: "0xD2cf07dEE40d3D530D15b88d689f5cd97A31FC3D", // 0xd2cf07dee40d3d530d15b88d689f5cd97a31fc3d
		FRAXList[ChainEthereum]:  "0xfd3aDA5AAbdc6531C7C2AC46c00eBf870f5a0E6B", // 0xfd3ada5aabdc6531c7c2ac46c00ebf870f5a0e6b
	},
}
