package ethaddr

import "github.com/0xVanfer/chainId"

const AaveV3Protocol string = "aavev3"

// Aave incentive controller V3.
//
// map[network] = address.
var AaveIncentiveControllerV3List = map[string]string{
	chainId.AvalancheChainName: "0x929ec64c34a17401f460460d4b9390518e5b473e",
	chainId.OptimismChainName:  "0x929ec64c34a17401f460460d4b9390518e5b473e",
	chainId.PolygonChainName:   "0x929ec64c34a17401f460460d4b9390518e5b473e",
}

// Aave lending pool v3.
//
// map[network] = address.
var AaveLendingPoolV3List = map[string]string{
	chainId.AvalancheChainName: "0x794a61358d6845594f94dc1db02a252b5b4814ad",
	chainId.OptimismChainName:  "0x794a61358d6845594f94dc1db02a252b5b4814ad",
	chainId.PolygonChainName:   "0x794a61358d6845594f94dc1db02a252b5b4814ad",
}

// Aave pool data provider v3.
//
// map[network] = address.
var AavePoolDataProviderList = map[string]string{
	chainId.AvalancheChainName: "0x69fa688f1dc47d4b5d8029d5a35fb7a548310654",
	chainId.OptimismChainName:  "0x69fa688f1dc47d4b5d8029d5a35fb7a548310654",
	chainId.PolygonChainName:   "0x69fa688f1dc47d4b5d8029d5a35fb7a548310654",
}

// Aave ui pool data provider v3.
//
// map[network] = address.
var AaveUiPoolDataProveiderV3List = map[string]string{
	chainId.AvalancheChainName: "0xdbbfafc45983B4659e368a3025b81f69Ab6e5093",
	chainId.OptimismChainName:  "0x64f558d4bfc1c03a8c8b2ff84976ff04c762b51f",
	chainId.PolygonChainName:   "0x7006e5a16e449123a3f26920746d03337ff37340",
}

// Aave ui incentive data provider v3.
//
// map[network] = address.
var AaveUiIncentiveDataProveiderV3List = map[string]string{
	chainId.AvalancheChainName: "0x270f51cf3f681010b46f5c4ee2ad5120db33026f",
	chainId.OptimismChainName:  "0x6dd4b295b457a26cc2646aaf2519436681afb5d4",
	chainId.PolygonChainName:   "0xf43efc9789736baf550dc016c7389210c43e7997",
}

// Aave pool address provider v3.
//
// map[network] = address.
var AavePoolAddressProviderV3List = map[string]string{
	chainId.AvalancheChainName: "0xa97684ead0e402dc232d5a977953Df7ecbab3cdb",
	chainId.OptimismChainName:  "0xa97684ead0e402dc232d5a977953Df7ecbab3cdb",
	chainId.PolygonChainName:   "0xa97684ead0e402dc232d5a977953Df7ecbab3cdb",
}

// Aave a tokens v3.
//
// map[network][underlying] = address.
var AaveATokenV3List = map[string]map[string]string{
	chainId.AvalancheChainName: {
		DAIeList[chainId.AvalancheChainName]:  "0x82e64f49ed5ec1bc6e43dad4fc8af9bb3a2312ee", // aAvaDAI
		USDCList[chainId.AvalancheChainName]:  "0x625e7708f30ca75bfd92586e17077590c60eb4cd", // aAvaUSDC
		USDTList[chainId.AvalancheChainName]:  "0x6ab707aca953edaefbc4fd23ba73294241490620", // aAvaUSDT
		WETHList[chainId.AvalancheChainName]:  "0xe50fa9b3c56ffb159cb0fca61f5c9d750e8128c8", // aAvaWETH
		WBTCList[chainId.AvalancheChainName]:  "0x078f358208685046a11c85e8ad32895ded33a249", // aAvaWBTC
		WAVAXList[chainId.AvalancheChainName]: "0x6d80113e533a2c0fe82eabd35f1875dcea89ea97", // aAvaWAVAX
		LINKList[chainId.AvalancheChainName]:  "0x191c10aa4af7c30e871e70c95db0e4eb77237530", // aAvaLINK
		AAVEList[chainId.AvalancheChainName]:  "0xf329e36c7bf6e5e86ce2150875a84ce77f477375", // aAvaAAVE
		SAVAXList[chainId.AvalancheChainName]: "0x513c7e3a9c69ca3e22550ef58ac1c0088e918fff", // aAvaSAVAX
		BTCbList[chainId.AvalancheChainName]:  "0x8ffdf2de812095b1d19cb146e4c004587c0a0692", // aAvaBTC.b
		FRAXList[chainId.AvalancheChainName]:  "0xc45a479877e1e9dfe9fcd4056c699575a1045daa",
		MAIList[chainId.AvalancheChainName]:   "0x8eb270e296023e9d92081fdf967ddd7878724424",
	},
	chainId.OptimismChainName: {
		DAIList[chainId.OptimismChainName]:  "0x82e64f49ed5ec1bc6e43dad4fc8af9bb3a2312ee", // aOptDAI
		USDCList[chainId.OptimismChainName]: "0x625e7708f30ca75bfd92586e17077590c60eb4cd", // aOptUSDC
		USDTList[chainId.OptimismChainName]: "0x6ab707aca953edaefbc4fd23ba73294241490620", // aOptUSDT
		WETHList[chainId.OptimismChainName]: "0xe50fa9b3c56ffb159cb0fca61f5c9d750e8128c8", // aOptWETH
		WBTCList[chainId.OptimismChainName]: "0x078f358208685046a11c85e8ad32895ded33a249", // aOptWBTC
		SUSDList[chainId.OptimismChainName]: "0x6d80113e533a2c0fe82eabd35f1875dcea89ea97", // aOptSUSD
		LINKList[chainId.OptimismChainName]: "0x191c10aa4af7c30e871e70c95db0e4eb77237530", // aOptLINK
		AAVEList[chainId.OptimismChainName]: "0xf329e36C7bF6E5E86ce2150875a84Ce77f477375", // aOptAAVE
	},
	chainId.PolygonChainName: {
		AgEURList[chainId.PolygonChainName]:   "0x8437d7c167dfb82ed4cb79cd44b7a32a1dd95c77", // aPolAGEUR
		DAIList[chainId.PolygonChainName]:     "0x82e64f49ed5ec1bc6e43dad4fc8af9bb3a2312ee", // aPolDAI
		EURSList[chainId.PolygonChainName]:    "0x38d693ce1df5aadf7bc62595a37d667ad57922e5",
		JEURList[chainId.PolygonChainName]:    "0x6533afac2e7bccb20dca161449a13a32d391fb00",
		MiMATICList[chainId.PolygonChainName]: "0xebe517846d0f36eced99c735cbf6131e1feb775d",
		USDCList[chainId.PolygonChainName]:    "0x625e7708f30ca75bfd92586e17077590c60eb4cd",
		USDTList[chainId.PolygonChainName]:    "0x6ab707aca953edaefbc4fd23ba73294241490620",
		AAVEList[chainId.PolygonChainName]:    "0xf329e36c7bf6e5e86ce2150875a84ce77f477375",
		BALList[chainId.PolygonChainName]:     "0x8ffdf2de812095b1d19cb146e4c004587c0a0692",
		CRVList[chainId.PolygonChainName]:     "0x513c7e3a9c69ca3e22550ef58ac1c0088e918fff",
		DPIList[chainId.PolygonChainName]:     "0x724dc807b04555b71ed48a6896b6f41593b8c637",
		GHSTList[chainId.PolygonChainName]:    "0x8eb270e296023e9d92081fdf967ddd7878724424",
		LINKList[chainId.PolygonChainName]:    "0x191c10aa4af7c30e871e70c95db0e4eb77237530",
		MaticXList[chainId.PolygonChainName]:  "0x80ca0d8c38d2e2bcbab66aa1648bd1c7160500fe",
		STMATICList[chainId.PolygonChainName]: "0xea1132120ddcdda2f119e99fa7a27a0d036f7ac9",
		SUSHIList[chainId.PolygonChainName]:   "0xc45a479877e1e9dfe9fcd4056c699575a1045daa",
		WBTCList[chainId.PolygonChainName]:    "0x078f358208685046a11c85e8ad32895ded33a249",
		WETHList[chainId.PolygonChainName]:    "0xe50fa9b3c56ffb159cb0fca61f5c9d750e8128c8",
		WMATICList[chainId.PolygonChainName]:  "0x6d80113e533a2c0fe82eabd35f1875dcea89ea97",
	},
}

// Aave s tokens v3.
//
// map[network][underlying] = address.
var AaveSTokenV3List = map[string]map[string]string{
	chainId.AvalancheChainName: {
		DAIeList[chainId.AvalancheChainName]:  "0xd94112b5b62d53c9402e7a60289c6810def1dc9b", // stableDebtAvaDAI
		USDCList[chainId.AvalancheChainName]:  "0x307ffe186f84a3bc2613d1ea417a5737d69a7007", // stableDebtAvaUSDC
		USDTList[chainId.AvalancheChainName]:  "0x70effc565db6eef7b927610155602d31b670e802", // stableDebtAvaUSDT
		WETHList[chainId.AvalancheChainName]:  "0xd8ad37849950903571df17049516a5cd4cbe55f6", // stableDebtAvaWETH
		WBTCList[chainId.AvalancheChainName]:  "0x633b207dd676331c413d4c013a6294b0fe47cd0e", // stableDebtAvaWBTC
		WAVAXList[chainId.AvalancheChainName]: "0xf15f26710c827dde8acba678682f3ce24f2fb56e", // stableDebtAvaWAVAX
		LINKList[chainId.AvalancheChainName]:  "0x89d976629b7055ff1ca02b927ba3e020f22a44e4", // stableDebtAvaLINK
		AAVEList[chainId.AvalancheChainName]:  "0xfaef6a702d15428e588d4c0614aefb4348d83d48", // stableDebtAvaAAVE
		SAVAXList[chainId.AvalancheChainName]: "0x08cb71192985e936c7cd166a8b268035e400c3c3", // stableDebtAvaSAVAX
		BTCbList[chainId.AvalancheChainName]:  "0xa5e408678469d23efdb7694b1b0a85bb0669e8bd", //
		FRAXList[chainId.AvalancheChainName]:  "0x78246294a4c6fbf614ed73ccc9f8b875ca8ee841", //
		MAIList[chainId.AvalancheChainName]:   "0x3ef10dff4928279c004308ebadc4db8b7620d6fc", //
	},
	chainId.PolygonChainName: {
		AgEURList[chainId.PolygonChainName]:   "0x40b4baecc69b882e8804f9286b12228c27f8c9bf", // stableDebtPolAGEUR
		DAIList[chainId.PolygonChainName]:     "0xd94112b5b62d53c9402e7a60289c6810def1dc9b", //
		EURSList[chainId.PolygonChainName]:    "0x8a9fde6925a839f6b1932d16b36ac026f8d3fbdb",
		JEURList[chainId.PolygonChainName]:    "0x6b4b37618d85db2a7b469983c888040f7f05ea3d",
		MiMATICList[chainId.PolygonChainName]: "0x687871030477bf974725232f764aa04318a8b9c8",
		USDCList[chainId.PolygonChainName]:    "0x307ffe186f84a3bc2613d1ea417a5737d69a7007",
		USDTList[chainId.PolygonChainName]:    "0x70effc565db6eef7b927610155602d31b670e802",
		AAVEList[chainId.PolygonChainName]:    "0xfaef6a702d15428e588d4c0614aefb4348d83d48",
		BALList[chainId.PolygonChainName]:     "0xa5e408678469d23efdb7694b1b0a85bb0669e8bd",
		CRVList[chainId.PolygonChainName]:     "0x08cb71192985e936c7cd166a8b268035e400c3c3",
		DPIList[chainId.PolygonChainName]:     "0xdc1fad70953bb3918592b6fcc374fe05f5811b6a",
		GHSTList[chainId.PolygonChainName]:    "0x3ef10dff4928279c004308ebadc4db8b7620d6fc",
		LINKList[chainId.PolygonChainName]:    "0x89d976629b7055ff1ca02b927ba3e020f22a44e4",
		MaticXList[chainId.PolygonChainName]:  "0x62fc96b27a510cf4977b59ff952dc32378cc221d",
		STMATICList[chainId.PolygonChainName]: "0x1ffd28689da7d0148ff0fcb669e9f9f0fc13a219",
		SUSHIList[chainId.PolygonChainName]:   "0x78246294a4c6fbf614ed73ccc9f8b875ca8ee841",
		WBTCList[chainId.PolygonChainName]:    "0x633b207dd676331c413d4c013a6294b0fe47cd0e",
		WETHList[chainId.PolygonChainName]:    "0xd8ad37849950903571df17049516a5cd4cbe55f6",
		WMATICList[chainId.PolygonChainName]:  "0xf15f26710c827dde8acba678682f3ce24f2fb56e",
	},
}

// Aave v tokens v3.
//
// map[network][underlying] = address.
var AaveVTokenV3List = map[string]map[string]string{
	chainId.AvalancheChainName: {
		DAIeList[chainId.AvalancheChainName]:  "0x8619d80fb0141ba7f184cbf22fd724116d9f7ffc", // variableDebtAvaDAI
		USDCList[chainId.AvalancheChainName]:  "0xfccf3cabbe80101232d343252614b6a3ee81c989", // variableDebtAvaUSDC
		USDTList[chainId.AvalancheChainName]:  "0xfb00ac187a8eb5afae4eace434f493eb62672df7", // variableDebtAvaUSDT
		WETHList[chainId.AvalancheChainName]:  "0x0c84331e39d6658cd6e6b9ba04736cc4c4734351", // variableDebtAvaWETH
		WBTCList[chainId.AvalancheChainName]:  "0x92b42c66840c7ad907b4bf74879ff3ef7c529473", // variableDebtAvaWBTC
		WAVAXList[chainId.AvalancheChainName]: "0x4a1c3ad6ed28a636ee1751c69071f6be75deb8b8", // variableDebtAvaWAVAX
		LINKList[chainId.AvalancheChainName]:  "0x953a573793604af8d41f306feb8274190db4ae0e", // variableDebtAvaLINK
		AAVEList[chainId.AvalancheChainName]:  "0xe80761ea617f66f96274ea5e8c37f03960ecc679", // variableDebtAvaAAVE
		SAVAXList[chainId.AvalancheChainName]: "0x77ca01483f379e58174739308945f044e1a764dc", // variableDebtAvaSAVAX
		BTCbList[chainId.AvalancheChainName]:  "0xa8669021776bc142dfca87c21b4a52595bcbb40a", // variableDebtAvaBTC.b
		FRAXList[chainId.AvalancheChainName]:  "0x34e2ed44ef7466d5f9e0b782b5c08b57475e7907", //
		MAIList[chainId.AvalancheChainName]:   "0xce186f6cccb0c955445bb9d10c59cae488fea559", //
	},
	chainId.OptimismChainName: {
		DAIList[chainId.OptimismChainName]:  "0x8619d80fb0141ba7f184cbf22fd724116d9f7ffc", // variableDebtOpt
		USDCList[chainId.OptimismChainName]: "0xfccf3cabbe80101232d343252614b6a3ee81c989", // variableDebtOpt
		USDTList[chainId.OptimismChainName]: "0xfb00ac187a8eb5afae4eace434f493eb62672df7", // variableDebtOpt
		WETHList[chainId.OptimismChainName]: "0x0c84331e39d6658Cd6e6b9ba04736cC4c4734351", // variableDebtOpt
		WBTCList[chainId.OptimismChainName]: "0x92b42c66840c7ad907b4bf74879ff3ef7c529473", // variableDebtOpt
		SUSDList[chainId.OptimismChainName]: "0x4a1c3ad6ed28a636ee1751c69071f6be75deb8b8", // variableDebtOpt
		LINKList[chainId.OptimismChainName]: "0x953a573793604af8d41f306feb8274190db4ae0e", // variableDebtOpt
		AAVEList[chainId.OptimismChainName]: "0xe80761ea617f66f96274ea5e8c37f03960ecc679", // variableDebtOpt
	},
	chainId.PolygonChainName: {
		AgEURList[chainId.PolygonChainName]:   "0x3ca5fa07689f266e907439afd1fbb59c44fe12f6", // variableDebtPolAGEUR
		DAIList[chainId.PolygonChainName]:     "0x8619d80fb0141ba7f184cbf22fd724116d9f7ffc", //
		EURSList[chainId.PolygonChainName]:    "0x5d557b07776d12967914379c71a1310e917c7555",
		JEURList[chainId.PolygonChainName]:    "0x44705f578135cc5d703b4c9c122528c73eb87145",
		MiMATICList[chainId.PolygonChainName]: "0x18248226c16bf76c032817854e7c83a2113b4f06",
		USDCList[chainId.PolygonChainName]:    "0xfccf3cabbe80101232d343252614b6a3ee81c989",
		USDTList[chainId.PolygonChainName]:    "0xfb00ac187a8eb5afae4eace434f493eb62672df7",
		AAVEList[chainId.PolygonChainName]:    "0xe80761ea617f66f96274ea5e8c37f03960ecc679",
		BALList[chainId.PolygonChainName]:     "0xa8669021776bc142dfca87c21b4a52595bcbb40a",
		CRVList[chainId.PolygonChainName]:     "0x77ca01483f379e58174739308945f044e1a764dc",
		DPIList[chainId.PolygonChainName]:     "0xf611aeb5013fd2c0511c9cd55c7dc5c1140741a6",
		GHSTList[chainId.PolygonChainName]:    "0xce186f6cccb0c955445bb9d10c59cae488fea559",
		LINKList[chainId.PolygonChainName]:    "0x953a573793604af8d41f306feb8274190db4ae0e",
		MaticXList[chainId.PolygonChainName]:  "0xb5b46f918c2923fc7f26db76e8a6a6e9c4347cf9",
		STMATICList[chainId.PolygonChainName]: "0x6b030ff3fb9956b1b69f475b77ae0d3cf2cc5afa",
		SUSHIList[chainId.PolygonChainName]:   "0x34e2ed44ef7466d5f9e0b782b5c08b57475e7907",
		WBTCList[chainId.PolygonChainName]:    "0x92b42c66840c7ad907b4bf74879ff3ef7c529473",
		WETHList[chainId.PolygonChainName]:    "0x0c84331e39d6658cd6e6b9ba04736cc4c4734351",
		WMATICList[chainId.PolygonChainName]:  "0x4a1c3ad6ed28a636ee1751c69071f6be75deb8b8",
	},
}

// Aave v3 isolated tokens.
//
// map[network]= []underlyings.
var AaveV3IsolatedTokens = map[string][]string{
	chainId.AvalancheChainName: {
		USDTList[chainId.AvalancheChainName],
		FRAXList[chainId.AvalancheChainName],
		MAIList[chainId.AvalancheChainName],
	},
	chainId.PolygonChainName: {
		USDTList[chainId.PolygonChainName],
		EURSList[chainId.PolygonChainName],
		MiMATICList[chainId.PolygonChainName],
	},
}

// Aave v3 tokens that are not allowed to be set as collateral.
//
// map[network]= []underlyings.
var AaveV3NotAllowCollateralTokens = map[string][]string{
	chainId.AvalancheChainName: {},
	chainId.PolygonChainName: {
		JEURList[chainId.PolygonChainName],
		AgEURList[chainId.PolygonChainName],
	},
}
