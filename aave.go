package ethaddr

import "github.com/0xVanfer/chainId"

const AaveV2Protocol string = "aave"
const AaveV3Protocol string = "aavev3"

// Aave token: AAVE.
//
// map[network] = address.
var AaveTokenList = map[string]string{
	chainId.EthereumChainName:  "0x7fc66500c84a76ad7e9c93437bfc5ac33e2ddae9", // AAVE
	chainId.AvalancheChainName: "0x63a72806098bd3d9520cc43356dd78afe5d386d9", // AAVE.e
	chainId.OptimismChainName:  "0x76fb31fb4af56892a25e32cfc43de717950c9278", // AAVE
}

// Aave incentive controller V2.
//
// map[network] = address.
var AaveIncentiveControllerV2List = map[string]string{
	chainId.EthereumChainName:  "0xd784927ff2f95ba542bfc824c8a8a98f3495f6b5",
	chainId.AvalancheChainName: "0x01d83fe6a10d2f2b7af17034343746188272cac9",
}

// Aave incentive controller V3.
//
// map[network] = address.
var AaveIncentiveControllerV3List = map[string]string{
	chainId.AvalancheChainName: "0x929ec64c34a17401f460460d4b9390518e5b473e",
	chainId.OptimismChainName:  "0x929ec64c34a17401f460460d4b9390518e5b473e",
}

// Aave lending pool v2.
//
// map[network] = address.
var AaveLendingPoolV2List = map[string]string{
	chainId.EthereumChainName:  "0x7d2768de32b0b80b7a3454c06bdac94a69ddc7a9",
	chainId.AvalancheChainName: "0x4f01aed16d97e3ab5ab2b501154dc9bb0f1a5a2c",
}

// Aave lending pool v3.
//
// map[network] = address.
var AaveLendingPoolV3List = map[string]string{
	chainId.AvalancheChainName: "0x794a61358d6845594f94dc1db02a252b5b4814ad",
	chainId.OptimismChainName:  "0x794a61358d6845594f94dc1db02a252b5b4814ad",
}

// Aave pool data provider v3.
//
// map[network] = address.
var AavePoolDataProviderList = map[string]string{
	chainId.AvalancheChainName: "0x69fa688f1dc47d4b5d8029d5a35fb7a548310654",
	chainId.OptimismChainName:  "0x69fa688f1dc47d4b5d8029d5a35fb7a548310654",
}

// Aave ui pool data provider v3.
//
// map[network] = address.
var AaveUiPoolDataProveiderV3List = map[string]string{
	chainId.AvalancheChainName: "0xdbbfafc45983B4659e368a3025b81f69Ab6e5093",
	chainId.OptimismChainName:  "0x64f558d4bfc1c03a8c8b2ff84976ff04c762b51f",
}

// Aave ui incentive data provider v3.
//
// map[network] = address.
var AaveUiIncentiveDataProveiderV3List = map[string]string{
	chainId.AvalancheChainName: "0x270f51cf3f681010b46f5c4ee2ad5120db33026f",
	chainId.OptimismChainName:  "0x6dd4b295b457a26cc2646aaf2519436681afb5d4",
}

// Aave pool address provider v3.
//
// map[network] = address.
var AavePoolAddressProviderV3List = map[string]string{
	chainId.AvalancheChainName: "0xa97684ead0e402dc232d5a977953Df7ecbab3cdb",
	chainId.OptimismChainName:  "0xa97684ead0e402dc232d5a977953Df7ecbab3cdb",
}

// Aave a tokens v2.
//
// map[network][underlying] = address.
var AaveATokenV2List = map[string]map[string]string{
	chainId.AvalancheChainName: {
		DAIeList[chainId.AvalancheChainName]:      "0x47afa96cdc9fab46904a55a6ad4bf6660b53c38a", // avDAI
		USDCeList[chainId.AvalancheChainName]:     "0x46a51127c3ce23fb7ab1de06226147f446e4a857", // avUSDC
		USDTeList[chainId.AvalancheChainName]:     "0x532e6537fea298397212f09a61e03311686f548e", // avUSDT
		WETHList[chainId.AvalancheChainName]:      "0x53f7c5869a859f0aec3d334ee8b4cf01e3492f21", // avWETH
		WBTCList[chainId.AvalancheChainName]:      "0x686bef2417b6dc32c50a3cbfbcc3bb60e1e9a15d", // avWBTC
		WAVAXList[chainId.AvalancheChainName]:     "0xdfe521292ece2a4f44242efbcd66bc594ca9714b", // avWAVAX
		AaveTokenList[chainId.AvalancheChainName]: "0xd45b7c061016102f9fa220502908f2c0f1add1d7", // avAAVE
	},
	chainId.EthereumChainName: {
		USDTList[chainId.EthereumChainName]:              "0x3ed3b47dd13ec9a98b44e6204a523e766b225811", // aUSDT 0
		WBTCList[chainId.EthereumChainName]:              "0x9ff58f4ffb29fa2266ab25e75e2a8b3503311656", // aWBTC 1
		WETHList[chainId.EthereumChainName]:              "0x030ba81f1c18d280636f32af80b9aad02cf0854e", // aWETH 2
		YearnTokenList[chainId.EthereumChainName]:        "0x5165d24277cd063f5ac44efd447b27025e888f37", // aYFI 3
		ZRXList[chainId.EthereumChainName]:               "0xdf7ff54aacacbff42dfe29dd6144a69b629f8c9e", // aZRX 4
		UniswapTokenList[chainId.EthereumChainName]:      "0xb9d7cb55f463405cdfbe4e90a6d2df01c2b92bf1", // aUNI 5
		AaveTokenList[chainId.EthereumChainName]:         "0xffc97d72e13e01096502cb8eb52dee56f74dad7b", // aAAVE 6
		BATList[chainId.EthereumChainName]:               "0x05ec93c0365baaeabf7aeffb0972ea7ecdd39cf1", // aBAT 7
		BUSDList[chainId.EthereumChainName]:              "0xa361718326c15715591c299427c62086f69923d9", // aBUSD 8
		DAIList[chainId.EthereumChainName]:               "0x028171bca77440897b824ca71d1c56cac55b68a3", // aDAI 9
		ENJList[chainId.EthereumChainName]:               "0xac6df26a590f08dcc95d5a4705ae8abbc88509ef", // aENJ 10
		KyberOldTokenList[chainId.EthereumChainName]:     "0x39c6b3e42d6a679d7d776778fe880bc9487c2eda", // aKNC 11
		LINKList[chainId.EthereumChainName]:              "0xa06bc25b5805d5f8d82847d191cb4af5a3e873e0", // aLINK 12
		DecentralandTokenList[chainId.EthereumChainName]: "0xa685a61171bb30d4072b338c80cb7b2c865c873e", // aMANA 13
		MakerTokenList[chainId.EthereumChainName]:        "0xc713e5e149d5d0715dcd1c156a020976e7e56b88", // aMKR 14
		RepublicTokenList[chainId.EthereumChainName]:     "0xcc12abe4ff81c9378d670de1b57f8e0dd228d77a", // aREN 15
		SynthetixTokenList[chainId.EthereumChainName]:    "0x35f6b052c598d933d69a4eec4d04c73a191fe6c2", // aSNX 16
		SUSDList[chainId.EthereumChainName]:              "0x6c5024cd4f8a59110119c56f8933403a539555eb", // aSUSD 17
		TUSDList[chainId.EthereumChainName]:              "0x101cc05f4a51c0319f570d5e146a8c625198e636", // aTUSD 18
		USDCList[chainId.EthereumChainName]:              "0xbcca60bb61934080951369a648fb03df4f96263c", // aUSDC 19
		CurveTokenlist[chainId.EthereumChainName]:        "0x8dae6cb04688c62d939ed9b68d32bc62e49970b1", // aCRV 20
		GeminiTokenList[chainId.EthereumChainName]:       "0xd37ee7e4f452c6638c96536e68090de8cbcdb583", // aGUSD 21
		BalancerTokenList[chainId.EthereumChainName]:     "0x272f97b7a56a387ae942350bbc7df5700f8a4576", // aBAL 22
		XSushiList[chainId.EthereumChainName]:            "0xf256cc7847e919fac9b808cc216cac87ccf2f47a", // aXSUSHI 23
		RenFILList[chainId.EthereumChainName]:            "0x514cd6756ccbe28772d4cb81bc3156ba9d1744aa", // aRENFIL 24
		ReflexerTokenList[chainId.EthereumChainName]:     "0xc9bc48c72154ef3e5425641a3c747242112a46af", // aRAI 25
		AMPLList[chainId.EthereumChainName]:              "0x1e6bb68acec8fefbd87d192be09bb274170a0548", // aAMPL 26
		USDPList[chainId.EthereumChainName]:              "0x2e8f4bdbe3d47d7d7de490437aea9915d930f1a3", // aUSDP 27 (should be aPAX)
		DefiPulseTokenList[chainId.EthereumChainName]:    "0x6f634c6135d2ebd550000ac92f494f9cb8183dae", // aDPI 28
		FRAXList[chainId.EthereumChainName]:              "0xd4937682df3c8aef4fe912a96a74121c0829e664", // aFRAX 29
		FEIList[chainId.EthereumChainName]:               "0x683923db55fead99a79fa01a27eec3cb19679cc3", // aFEI 30
		LidoSTETHList[chainId.EthereumChainName]:         "0x1982b2f5814301d4e9a8b0201555376e62f82428", // aSTETH 31
		ENSList[chainId.EthereumChainName]:               "0x9a14e23a58edf4efdcb360f68cd1b95ce2081a2f", // aENS 32
		USTList[chainId.EthereumChainName]:               "0xc2e2152647f4c26028482efaf64b2aa28779efc4", // aUST 33 deprecated
		ConvexTokenList[chainId.EthereumChainName]:       "0x952749e07d7157bb9644a894dfaf3bad5ef6d918", // aCVX 34
	},
}

// Aave s tokens v2.
//
// map[network][underlying] = address.
var AaveSTokenV2List = map[string]map[string]string{
	chainId.AvalancheChainName: {},
	chainId.EthereumChainName:  {},
}

// Aave v tokens v2.
//
// map[network][underlying] = address.
var AaveVTokenV2List = map[string]map[string]string{
	chainId.AvalancheChainName: {
		DAIeList[chainId.AvalancheChainName]:      "0x1852dc24d1a8956a0b356aa18ede954c7a0ca5ae", // variableDebtvDAI
		USDCeList[chainId.AvalancheChainName]:     "0x848c080d2700cbe1b894a3374ad5e887e5ccb89c", // variableDebtvUSDC
		USDTeList[chainId.AvalancheChainName]:     "0xfc1ada7a288d6fce0d29ccfaaa57bc9114bb2dbe", // variableDebtvUSDT
		WETHList[chainId.AvalancheChainName]:      "0x4e575cacb37bc1b5afec68a0462c4165a5268983", // variableDebtvWETH
		WBTCList[chainId.AvalancheChainName]:      "0x2dc0e35ec3ab070b8a175c829e23650ee604a9eb", // variableDebtvWBTC
		WAVAXList[chainId.AvalancheChainName]:     "0x66a0fe52fb629a6cb4d10b8580afdffe888f5fd4", // variableDebtvWAVAX
		AaveTokenList[chainId.AvalancheChainName]: "0x8352e3fd18b8d84d3c8a1b538d788899073c7a8e", // variableDebtvAAVE
	},
	chainId.EthereumChainName: {
		USDTList[chainId.EthereumChainName]:              "0x531842cebbdd378f8ee36d171d6cc9c4fcf475ec", // variableDebtUSDT 0
		WBTCList[chainId.EthereumChainName]:              "0x9c39809Dec7F95F5e0713634a4D0701329B3b4d2", // variableDebtWBTC 1
		WETHList[chainId.EthereumChainName]:              "0xF63B34710400CAd3e044cFfDcAb00a0f32E33eCf", // variableDebtWETH 2
		YearnTokenList[chainId.EthereumChainName]:        "0x7EbD09022Be45AD993BAA1CEc61166Fcc8644d97", // variableDebtYFI 3
		ZRXList[chainId.EthereumChainName]:               "0x85791D117A392097590bDeD3bD5abB8d5A20491A", // variableDebtZRX 4
		UniswapTokenList[chainId.EthereumChainName]:      "0x5BdB050A92CADcCfCDcCCBFC17204a1C9cC0Ab73", // variableDebtUNI 5
		AaveTokenList[chainId.EthereumChainName]:         "0xF7DBA49d571745D9d7fcb56225B05BEA803EBf3C", // variableDebtAAVE 6
		BATList[chainId.EthereumChainName]:               "0xfc218A6Dfe6901CB34B1a5281FC6f1b8e7E56877", // variableDebtBAT 7
		BUSDList[chainId.EthereumChainName]:              "0xbA429f7011c9fa04cDd46a2Da24dc0FF0aC6099c", // variableDebtBUSD 8
		DAIList[chainId.EthereumChainName]:               "0x6c3c78838c761c6ac7be9f59fe808ea2a6e4379d", // variableDebtDAI 9
		ENJList[chainId.EthereumChainName]:               "0x38995f292a6e31b78203254fe1cdd5ca1010a446", // variableDebtENJ 10
		KyberOldTokenList[chainId.EthereumChainName]:     "0x6B05D1c608015Ccb8e205A690cB86773A96F39f1", // variableDebtKNC 11
		LINKList[chainId.EthereumChainName]:              "0x0b8f12b1788BFdE65Aa1ca52E3e9F3Ba401be16D", // variableDebtLINK 12
		DecentralandTokenList[chainId.EthereumChainName]: "0x0A68976301e46Ca6Ce7410DB28883E309EA0D352", // variableDebtMANA 13
		MakerTokenList[chainId.EthereumChainName]:        "0xba728ead5e496be00dcf66f650b6d7758ecb50f8", // variableDebtMKR 14
		RepublicTokenList[chainId.EthereumChainName]:     "0xcd9d82d33bd737de215cdac57fe2f7f04df77fe0", // variableDebtREN 15
		SynthetixTokenList[chainId.EthereumChainName]:    "0x267eb8cf715455517f9bd5834aeae3cea1ebdbd8", // variableDebtSNX 16
		SUSDList[chainId.EthereumChainName]:              "0xdc6a3ab17299d9c2a412b0e0a4c1f55446ae0817", // variableDebtSUSD 17
		TUSDList[chainId.EthereumChainName]:              "0x01C0eb1f8c6F1C1bF74ae028697ce7AA2a8b0E92", // variableDebtTUSD 18
		USDCList[chainId.EthereumChainName]:              "0x619beb58998eD2278e08620f97007e1116D5D25b", // variableDebtUSDC 19
		CurveTokenlist[chainId.EthereumChainName]:        "0x00ad8eBF64F141f1C81e9f8f792d3d1631c6c684", // variableDebtCRV 20
		GeminiTokenList[chainId.EthereumChainName]:       "0x279af5b99540c1a3a7e3cdd326e19659401ef99e", // variableDebtGUSD 21
		BalancerTokenList[chainId.EthereumChainName]:     "0x13210D4Fe0d5402bd7Ecbc4B5bC5cFcA3b71adB0", // variableDebtBAL 22
		XSushiList[chainId.EthereumChainName]:            "0xfAFEDF95E21184E3d880bd56D4806c4b8d31c69A", // variableDebtXSUSHI 23
		RenFILList[chainId.EthereumChainName]:            "0x348e2ebd5e962854871874e444f4122399c02755", // variableDebtRENFIL 24
		ReflexerTokenList[chainId.EthereumChainName]:     "0xb5385132ee8321977fff44b60cde9fe9ab0b4e6b", // variableDebtRAI 25
		AMPLList[chainId.EthereumChainName]:              "0xf013d90e4e4e3baf420dfea60735e75dbd42f1e1", // variableDebtAMPL 26
		USDPList[chainId.EthereumChainName]:              "0xfdb93b3b10936cf81fa59a02a7523b6e2149b2b7", // variableDebtUSDP 27 (should be aPAX)
		// DefiPulseTokenList[chainId.EthereumChainName]: "",                                           // variableDebtDPI 28
		FRAXList[chainId.EthereumChainName]:      "0xfe8f19b17ffef0fdbfe2671f248903055afaa8ca", // variableDebtFRAX 29
		FEIList[chainId.EthereumChainName]:       "0xC2e10006AccAb7B45D9184FcF5b7EC7763f5BaAe", // variableDebtFEI 30
		LidoSTETHList[chainId.EthereumChainName]: "0xA9DEAc9f00Dc4310c35603FCD9D34d1A750f81Db", // variableDebtSTETH 31
		ENSList[chainId.EthereumChainName]:       "0x176808047cc9b7A2C9AE202c593ED42dDD7C0D13", // variableDebtENS 32
		USTList[chainId.EthereumChainName]:       "0xaf32001cf2E66C4C3af4205F6EA77112AA4160FE", // variableDebtUST 33 deprecated
		// ConvexTokenList[chainId.EthereumChainName]:    "",                                           // variableDebtCVX 34
	},
}

// Aave a tokens v3.
//
// map[network][underlying] = address.
var AaveATokenV3List = map[string]map[string]string{
	chainId.AvalancheChainName: {
		DAIeList[chainId.AvalancheChainName]:       "0x82e64f49ed5ec1bc6e43dad4fc8af9bb3a2312ee", // aAvaDAI
		USDCList[chainId.AvalancheChainName]:       "0x625e7708f30ca75bfd92586e17077590c60eb4cd", // aAvaUSDC
		USDTList[chainId.AvalancheChainName]:       "0x6ab707aca953edaefbc4fd23ba73294241490620", // aAvaUSDT
		WETHList[chainId.AvalancheChainName]:       "0xe50fa9b3c56ffb159cb0fca61f5c9d750e8128c8", // aAvaWETH
		WBTCList[chainId.AvalancheChainName]:       "0x078f358208685046a11c85e8ad32895ded33a249", // aAvaWBTC
		WAVAXList[chainId.AvalancheChainName]:      "0x6d80113e533a2c0fe82eabd35f1875dcea89ea97", // aAvaWAVAX
		LINKList[chainId.AvalancheChainName]:       "0x191c10aa4af7c30e871e70c95db0e4eb77237530", // aAvaLINK
		AaveTokenList[chainId.AvalancheChainName]:  "0xf329e36c7bf6e5e86ce2150875a84ce77f477375", // aAvaAAVE
		BenqiSAVAXList[chainId.AvalancheChainName]: "0x513c7e3a9c69ca3e22550ef58ac1c0088e918fff", // aAvaSAVAX
	},
	chainId.OptimismChainName: {
		DAIList[chainId.OptimismChainName]:       "0x82e64f49ed5ec1bc6e43dad4fc8af9bb3a2312ee", // aOptDAI
		USDCList[chainId.OptimismChainName]:      "0x625e7708f30ca75bfd92586e17077590c60eb4cd", // aOptUSDC
		USDTList[chainId.OptimismChainName]:      "0x6ab707aca953edaefbc4fd23ba73294241490620", // aOptUSDT
		WETHList[chainId.OptimismChainName]:      "0xe50fa9b3c56ffb159cb0fca61f5c9d750e8128c8", // aOptWETH
		WBTCList[chainId.OptimismChainName]:      "0x078f358208685046a11c85e8ad32895ded33a249", // aOptWBTC
		SUSDList[chainId.OptimismChainName]:      "0x6d80113e533a2c0fe82eabd35f1875dcea89ea97", // aOptSUSD
		LINKList[chainId.OptimismChainName]:      "0x191c10aa4af7c30e871e70c95db0e4eb77237530", // aOptLINK
		AaveTokenList[chainId.OptimismChainName]: "0xf329e36C7bF6E5E86ce2150875a84Ce77f477375", // aOptAAVE
	},
}

// Aave s tokens v3.
//
// map[network][underlying] = address.
var AaveSTokenV3List = map[string]map[string]string{
	chainId.AvalancheChainName: {
		DAIeList[chainId.AvalancheChainName]:       "0xd94112b5b62d53c9402e7a60289c6810def1dc9b", // stableDebtAvaDAI
		USDCList[chainId.AvalancheChainName]:       "0x307ffe186f84a3bc2613d1ea417a5737d69a7007", // stableDebtAvaUSDC
		USDTList[chainId.AvalancheChainName]:       "0x70effc565db6eef7b927610155602d31b670e802", // stableDebtAvaUSDT
		WETHList[chainId.AvalancheChainName]:       "0xd8ad37849950903571df17049516a5cd4cbe55f6", // stableDebtAvaWETH
		WBTCList[chainId.AvalancheChainName]:       "0x633b207dd676331c413d4c013a6294b0fe47cd0e", // stableDebtAvaWBTC
		WAVAXList[chainId.AvalancheChainName]:      "0xf15f26710c827dde8acba678682f3ce24f2fb56e", // stableDebtAvaWAVAX
		LINKList[chainId.AvalancheChainName]:       "0x89d976629b7055ff1ca02b927ba3e020f22a44e4", // stableDebtAvaLINK
		AaveTokenList[chainId.AvalancheChainName]:  "0xfaef6a702d15428e588d4c0614aefb4348d83d48", // stableDebtAvaAAVE
		BenqiSAVAXList[chainId.AvalancheChainName]: "0x08cb71192985e936c7cd166a8b268035e400c3c3", // stableDebtAvaSAVAX
	},
}

// Aave v tokens v3.
//
// map[network][underlying] = address.
var AaveVTokenV3List = map[string]map[string]string{
	chainId.AvalancheChainName: {
		DAIeList[chainId.AvalancheChainName]:       "0x8619d80fb0141ba7f184cbf22fd724116d9f7ffc", // variableDebtAvaDAI
		USDCList[chainId.AvalancheChainName]:       "0xfccf3cabbe80101232d343252614b6a3ee81c989", // variableDebtAvaUSDC
		USDTList[chainId.AvalancheChainName]:       "0xfb00ac187a8eb5afae4eace434f493eb62672df7", // variableDebtAvaUSDT
		WETHList[chainId.AvalancheChainName]:       "0x0c84331e39d6658cd6e6b9ba04736cc4c4734351", // variableDebtAvaWETH
		WBTCList[chainId.AvalancheChainName]:       "0x92b42c66840c7ad907b4bf74879ff3ef7c529473", // variableDebtAvaWBTC
		WAVAXList[chainId.AvalancheChainName]:      "0x4a1c3ad6ed28a636ee1751c69071f6be75deb8b8", // variableDebtAvaWAVAX
		LINKList[chainId.AvalancheChainName]:       "0x953a573793604af8d41f306feb8274190db4ae0e", // variableDebtAvaLINK
		AaveTokenList[chainId.AvalancheChainName]:  "0xe80761ea617f66f96274ea5e8c37f03960ecc679", // variableDebtAvaAAVE
		BenqiSAVAXList[chainId.AvalancheChainName]: "0x77ca01483f379e58174739308945f044e1a764dc", // variableDebtAvaSAVAX
	},
	chainId.OptimismChainName: {
		DAIList[chainId.OptimismChainName]:       "0x8619d80fb0141ba7f184cbf22fd724116d9f7ffc", // variableDebtOpt
		USDCList[chainId.OptimismChainName]:      "0xfccf3cabbe80101232d343252614b6a3ee81c989", // variableDebtOpt
		USDTList[chainId.OptimismChainName]:      "0xfb00ac187a8eb5afae4eace434f493eb62672df7", // variableDebtOpt
		WETHList[chainId.OptimismChainName]:      "0x0c84331e39d6658Cd6e6b9ba04736cC4c4734351", // variableDebtOpt
		WBTCList[chainId.OptimismChainName]:      "0x92b42c66840c7ad907b4bf74879ff3ef7c529473", // variableDebtOpt
		SUSDList[chainId.OptimismChainName]:      "0x4a1c3ad6ed28a636ee1751c69071f6be75deb8b8", // variableDebtOpt
		LINKList[chainId.OptimismChainName]:      "0x953a573793604af8d41f306feb8274190db4ae0e", // variableDebtOpt
		AaveTokenList[chainId.OptimismChainName]: "0xe80761ea617f66f96274ea5e8c37f03960ecc679", // variableDebtOpt
	},
}
