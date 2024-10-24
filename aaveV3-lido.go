package ethaddr

// Website: https://aave.com/
//
// Docs: https://docs.aave.com/developers/getting-started/readme
//
// Deployed Addresses: https://github.com/bgd-labs/aave-address-book/blob/main/src/ts/AaveV3EthereumLido.ts
//
// Fork of AaveV3 for lido.
const AaveV3ProtocolLidoFork string = "aavev3lido"

// Aave rewards controller V3 for lido fork.
// Previously named as "incentives controller".
//
// map[network] = address.
var AaveIncentivesControllerV3LidoForkList = map[int64]string{
	ChainEthereum: "0x8c885F766E4cF88898e0c1Bb02993048f3E8093f", // 0x8c885f766e4cf88898e0c1bb02993048f3e8093f
}

// Aave lending pool v3 for lido fork.
//
// map[network] = address.
var AaveLendingPoolV3LidoForkList = map[int64]string{
	ChainEthereum: "0x4e033931ad43597d96D6bcc25c280717730B58B1", // 0x4e033931ad43597d96d6bcc25c280717730b58b1
}

// Aave protocol data provider v3 for lido fork.
// Previously named as "pool data provider".
//
// map[network] = address.
//
// Last updated Oct.9.2024. Aave v3.2.0
var AavePoolDataProviderV3LidoForkList = map[int64]string{
	ChainEthereum: "0x08795CFE08C7a81dCDFf482BbAAF474B240f31cD", // 0x08795cfe08c7a81dcdff482bbaaf474b240f31cd
}

// Aave ui pool data provider v3 for lido fork.
//
// map[network] = address.
//
// Last updated Oct.9.2024. Aave v3.2.0
var AaveUiPoolDataProviderV3LidoForkList = map[int64]string{
	ChainEthereum: "0x194324C9Af7f56E22F1614dD82E18621cb9238E7", // 0x194324c9af7f56e22f1614dd82e18621cb9238e7
}

// Aave ui incentive data provider v3 for lido fork.
//
// map[network] = address.
//
// Last updated Oct.9.2024. Aave v3.2.0
var AaveUiIncentiveDataProviderV3LidoForkList = map[int64]string{
	ChainEthereum: "0x5a40cDe2b76Da2beD545efB3ae15708eE56aAF9c", // 0x5a40cde2b76da2bed545efb3ae15708ee56aaf9c
}

// Aave pool address provider v3 for lido fork.
//
// map[network] = address.
var AavePoolAddressProviderV3LidoForkList = map[int64]string{
	ChainEthereum: "0xcfBf336fe147D643B9Cb705648500e101504B16d", // 0xcfbf336fe147d643b9cb705648500e101504b16d
}

// Aave a tokens v3 for lido fork.
//
// map[network][underlying] = address.
var AaveATokenV3LidoForkList = map[int64]map[string]string{
	ChainEthereum: {
		WSTETHList[ChainEthereum]: "0xC035a7cf15375cE2706766804551791aD035E0C2", // 0xc035a7cf15375ce2706766804551791ad035e0c2
		WETHList[ChainEthereum]:   "0xfA1fDbBD71B0aA16162D76914d69cD8CB3Ef92da", // 0xfa1fdbbd71b0aa16162d76914d69cd8cb3ef92da
		USDSList[ChainEthereum]:   "0x09AA30b182488f769a9824F15E6Ce58591Da4781", // 0x09aa30b182488f769a9824f15e6ce58591da4781
		USDCList[ChainEthereum]:   "0x2A1FBcb52Ed4d9b23daD17E1E8Aed4BB0E6079b8", // 0x2a1fbcb52ed4d9b23dad17e1e8aed4bb0e6079b8
	},
}

// DEPRECATED: Aave v3 has deprecated sToken.
//
// Aave s tokens v3 for lido fork.
//
// map[network][underlying] = address.
var AaveSTokenV3LidoForkList = map[int64]map[string]string{
	ChainEthereum: {
		WSTETHList[ChainEthereum]: "0x3d0Fd161363b327C704b013a9E63a8Cc03Bec1c4", // 0x3d0fd161363b327c704b013a9e63a8cc03bec1c4
		WETHList[ChainEthereum]:   "0x97D5Cd1a26243647ddEac87183236Cf215974d30", // 0x97d5cd1a26243647ddeac87183236cf215974d30
		USDSList[ChainEthereum]:   "0x779dB175167C60c2B2193Be6B8d8B3602435e89E", // 0x779db175167c60c2b2193be6b8d8b3602435e89e
		USDCList[ChainEthereum]:   "0x1bd1FaAA75030f63d523450e68F233fd0a25eC2e", // 0x1bd1faaa75030f63d523450e68f233fd0a25ec2e
	},
}

// Aave v tokens v3 for lido fork.
//
// map[network][underlying] = address.
var AaveVTokenV3LidoForkList = map[int64]map[string]string{
	ChainEthereum: {
		WSTETHList[ChainEthereum]: "0xE439edd2625772AA635B437C099C607B6eb7d35f", // 0xe439edd2625772aa635b437c099c607b6eb7d35f
		WETHList[ChainEthereum]:   "0x91b7d78BF92db564221f6B5AeE744D1727d1Dd1e", // 0x91b7d78bf92db564221f6b5aee744d1727d1dd1e
		USDSList[ChainEthereum]:   "0x2D9fe18b6c35FE439cC15D932cc5C943bf2d901E", // 0x2d9fe18b6c35fe439cc15d932cc5c943bf2d901e
		USDCList[ChainEthereum]:   "0xeD90dE2D824Ee766c6Fd22E90b12e598f681dc9F", // 0xed90de2d824ee766c6fd22e90b12e598f681dc9f
	},
}
