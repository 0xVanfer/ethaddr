package ethaddr

import "github.com/0xVanfer/chainId"

const AaveV2Protocol string = "aave"

// Aave incentives controller V2.
//
// map[network] = address.
var AaveIncentivesControllerV2List = map[string]string{
	chainId.EthereumChainName:  "0xd784927ff2f95ba542bfc824c8a8a98f3495f6b5",
	chainId.AvalancheChainName: "0x01d83fe6a10d2f2b7af17034343746188272cac9",
}

// Deprecated. Use AaveIncentivesControllerV2List instead.
var AaveIncentiveControllerV2List = AaveIncentivesControllerV2List

// Aave lending pool v2.
//
// map[network] = address.
var AaveLendingPoolV2List = map[string]string{
	chainId.EthereumChainName:  "0x7d2768de32b0b80b7a3454c06bdac94a69ddc7a9",
	chainId.AvalancheChainName: "0x4f01aed16d97e3ab5ab2b501154dc9bb0f1a5a2c",
}

// Aave a tokens v2.
//
// map[network][underlying] = address.
var AaveATokenV2List = map[string]map[string]string{
	chainId.AvalancheChainName: {
		DAIeList[chainId.AvalancheChainName]:  "0x47afa96cdc9fab46904a55a6ad4bf6660b53c38a", // avDAI
		USDCeList[chainId.AvalancheChainName]: "0x46a51127c3ce23fb7ab1de06226147f446e4a857", // avUSDC
		USDTeList[chainId.AvalancheChainName]: "0x532e6537fea298397212f09a61e03311686f548e", // avUSDT
		WETHList[chainId.AvalancheChainName]:  "0x53f7c5869a859f0aec3d334ee8b4cf01e3492f21", // avWETH
		WBTCList[chainId.AvalancheChainName]:  "0x686bef2417b6dc32c50a3cbfbcc3bb60e1e9a15d", // avWBTC
		WAVAXList[chainId.AvalancheChainName]: "0xdfe521292ece2a4f44242efbcd66bc594ca9714b", // avWAVAX
		AAVEList[chainId.AvalancheChainName]:  "0xd45b7c061016102f9fa220502908f2c0f1add1d7", // avAAVE
	},
	chainId.EthereumChainName: {
		USDTList[chainId.EthereumChainName]:          "0x3ed3b47dd13ec9a98b44e6204a523e766b225811", // aUSDT 0
		WBTCList[chainId.EthereumChainName]:          "0x9ff58f4ffb29fa2266ab25e75e2a8b3503311656", // aWBTC 1
		WETHList[chainId.EthereumChainName]:          "0x030ba81f1c18d280636f32af80b9aad02cf0854e", // aWETH 2
		YFIList[chainId.EthereumChainName]:           "0x5165d24277cd063f5ac44efd447b27025e888f37", // aYFI 3
		ZRXList[chainId.EthereumChainName]:           "0xdf7ff54aacacbff42dfe29dd6144a69b629f8c9e", // aZRX 4
		UNIList[chainId.EthereumChainName]:           "0xb9d7cb55f463405cdfbe4e90a6d2df01c2b92bf1", // aUNI 5
		AAVEList[chainId.EthereumChainName]:          "0xffc97d72e13e01096502cb8eb52dee56f74dad7b", // aAAVE 6
		BATList[chainId.EthereumChainName]:           "0x05ec93c0365baaeabf7aeffb0972ea7ecdd39cf1", // aBAT 7
		BUSDList[chainId.EthereumChainName]:          "0xa361718326c15715591c299427c62086f69923d9", // aBUSD 8
		DAIList[chainId.EthereumChainName]:           "0x028171bca77440897b824ca71d1c56cac55b68a3", // aDAI 9
		ENJList[chainId.EthereumChainName]:           "0xac6df26a590f08dcc95d5a4705ae8abbc88509ef", // aENJ 10
		KyberOldTokenList[chainId.EthereumChainName]: "0x39c6b3e42d6a679d7d776778fe880bc9487c2eda", // aKNC 11
		LINKList[chainId.EthereumChainName]:          "0xa06bc25b5805d5f8d82847d191cb4af5a3e873e0", // aLINK 12
		MANAList[chainId.EthereumChainName]:          "0xa685a61171bb30d4072b338c80cb7b2c865c873e", // aMANA 13
		MKRList[chainId.EthereumChainName]:           "0xc713e5e149d5d0715dcd1c156a020976e7e56b88", // aMKR 14
		RENList[chainId.EthereumChainName]:           "0xcc12abe4ff81c9378d670de1b57f8e0dd228d77a", // aREN 15
		SNXList[chainId.EthereumChainName]:           "0x35f6b052c598d933d69a4eec4d04c73a191fe6c2", // aSNX 16
		SUSDList[chainId.EthereumChainName]:          "0x6c5024cd4f8a59110119c56f8933403a539555eb", // aSUSD 17
		TUSDList[chainId.EthereumChainName]:          "0x101cc05f4a51c0319f570d5e146a8c625198e636", // aTUSD 18
		USDCList[chainId.EthereumChainName]:          "0xbcca60bb61934080951369a648fb03df4f96263c", // aUSDC 19
		CRVList[chainId.EthereumChainName]:           "0x8dae6cb04688c62d939ed9b68d32bc62e49970b1", // aCRV 20
		GUSDList[chainId.EthereumChainName]:          "0xd37ee7e4f452c6638c96536e68090de8cbcdb583", // aGUSD 21
		BALList[chainId.EthereumChainName]:           "0x272f97b7a56a387ae942350bbc7df5700f8a4576", // aBAL 22
		XSushiList[chainId.EthereumChainName]:        "0xf256cc7847e919fac9b808cc216cac87ccf2f47a", // aXSUSHI 23
		RenFILList[chainId.EthereumChainName]:        "0x514cd6756ccbe28772d4cb81bc3156ba9d1744aa", // aRENFIL 24
		RAIList[chainId.EthereumChainName]:           "0xc9bc48c72154ef3e5425641a3c747242112a46af", // aRAI 25
		AMPLList[chainId.EthereumChainName]:          "0x1e6bb68acec8fefbd87d192be09bb274170a0548", // aAMPL 26
		USDPList[chainId.EthereumChainName]:          "0x2e8f4bdbe3d47d7d7de490437aea9915d930f1a3", // aUSDP 27 (should be aPAX)
		DPIList[chainId.EthereumChainName]:           "0x6f634c6135d2ebd550000ac92f494f9cb8183dae", // aDPI 28
		FRAXList[chainId.EthereumChainName]:          "0xd4937682df3c8aef4fe912a96a74121c0829e664", // aFRAX 29
		FEIList[chainId.EthereumChainName]:           "0x683923db55fead99a79fa01a27eec3cb19679cc3", // aFEI 30
		STETHList[chainId.EthereumChainName]:         "0x1982b2f5814301d4e9a8b0201555376e62f82428", // aSTETH 31
		ENSList[chainId.EthereumChainName]:           "0x9a14e23a58edf4efdcb360f68cd1b95ce2081a2f", // aENS 32
		USTList[chainId.EthereumChainName]:           "0xc2e2152647f4c26028482efaf64b2aa28779efc4", // aUST 33 deprecated
		CVXList[chainId.EthereumChainName]:           "0x952749e07d7157bb9644a894dfaf3bad5ef6d918", // aCVX 34
		OneInchTokenList[chainId.EthereumChainName]:  "0xb29130cbcc3f791f077eade0266168e808e5151e", // a1INCH
		LUSDList[chainId.EthereumChainName]:          "0xce1871f791548600cb59efbeffc9c38719142079", // aLUSD
	},
}

// Aave s tokens v2.
//
// map[network][underlying] = address.
var AaveSTokenV2List = map[string]map[string]string{
	chainId.AvalancheChainName: {
		DAIeList[chainId.AvalancheChainName]:  "0x3676e4ee689d527ddb89812b63fad0b7501772b3",
		USDCeList[chainId.AvalancheChainName]: "0x5b14679135dbe8b02015ec3ca4924a12e4c6c85a",
		USDTeList[chainId.AvalancheChainName]: "0x9c7b81a867499b7387ed05017a13d4172a0c17bf",
		WETHList[chainId.AvalancheChainName]:  "0x60f6a45006323b97d97cb0a42ac39e2b757ada63",
		WBTCList[chainId.AvalancheChainName]:  "0x3484408989985d68c9700dc1cfdfeae6d2f658cf",
		WAVAXList[chainId.AvalancheChainName]: "0x2920cd5b8a160b2addb00ec5d5f4112255d4ae75",
		AAVEList[chainId.AvalancheChainName]:  "0x66904e4f3f44e3925d22ceca401b6f2da085c98f",
	},
	chainId.EthereumChainName: {
		USDTList[chainId.EthereumChainName]:          "0xe91d55ab2240594855abd11b3faae801fd4c4687",
		WBTCList[chainId.EthereumChainName]:          "0x51b039b9afe64b78758f8ef091211b5387ea717c",
		WETHList[chainId.EthereumChainName]:          "0x4e977830ba4bd783c0bb7f15d3e243f73ff57121",
		YFIList[chainId.EthereumChainName]:           "0xca823f78c2dd38993284bb42ba9b14152082f7bd",
		ZRXList[chainId.EthereumChainName]:           "0x071b4323a24e73a5afeebe34118cd21b8faaf7c3",
		UNIList[chainId.EthereumChainName]:           "0xd939f7430dc8d5a427f156de1012a56c18acb6aa",
		AAVEList[chainId.EthereumChainName]:          "0x079d6a3e844bcecf5720478a718edb6575362c5f",
		BATList[chainId.EthereumChainName]:           "0x277f8676facf4daa5a6ea38ba511b7f65aa02f9f",
		BUSDList[chainId.EthereumChainName]:          "0x4a7a63909a72d268b1d8a93a9395d098688e0e5c",
		DAIList[chainId.EthereumChainName]:           "0x778a13d3eeb110a4f7bb6529f99c000119a08e92",
		ENJList[chainId.EthereumChainName]:           "0x943dcca156b5312aa24c1a08769d67fece4ac14c",
		KyberOldTokenList[chainId.EthereumChainName]: "0x9915dfb872778b2890a117da1f35f335eb06b54f",
		LINKList[chainId.EthereumChainName]:          "0xfb4aec4cc858f2539ebd3d37f2a43eae5b15b98a",
		MANAList[chainId.EthereumChainName]:          "0xd86c74ea2224f4b8591560652b50035e4e5c0a3b",
		MKRList[chainId.EthereumChainName]:           "0xc01c8e4b12a89456a9fd4e4e75b72546bf53f0b5",
		RENList[chainId.EthereumChainName]:           "0x3356ec1efa75d9d150da1ec7d944d9edf73703b7",
		SNXList[chainId.EthereumChainName]:           "0x8575c8ae70bdb71606a53aea1c6789cb0fbf3166",
		SUSDList[chainId.EthereumChainName]:          "0x30b0f7324fedf89d8eff397275f8983397efe4af",
		TUSDList[chainId.EthereumChainName]:          "0x7f38d60d94652072b2c44a18c0e14a481ec3c0dd",
		USDCList[chainId.EthereumChainName]:          "0xe4922afab0bbadd8ab2a88e0c79d884ad337fca6",
		CRVList[chainId.EthereumChainName]:           "0x9288059a74f589c919c7cf1db433251cdfeb874b",
		GUSDList[chainId.EthereumChainName]:          "0xf8ac64ec6ff8e0028b37eb89772d21865321bce0",
		BALList[chainId.EthereumChainName]:           "0xe569d31590307d05da3812964f1edd551d665a0b",
		XSushiList[chainId.EthereumChainName]:        "0x73bfb81d7dba75c904f430ea8bae82db0d41187b",
		RenFILList[chainId.EthereumChainName]:        "0xcaad05c49e14075077915cb5c820eb3245afb950",
		RAIList[chainId.EthereumChainName]:           "0x9c72b8476c33ae214ee3e8c20f0bc28496a62032",
		AMPLList[chainId.EthereumChainName]:          "0x18152c9f77dadc737006e9430db913159645fa87",
		USDPList[chainId.EthereumChainName]:          "0x2387119bc85a74e0bbcbe190d80676cb16f10d4f",
		DPIList[chainId.EthereumChainName]:           "0xa3953f07f389d719f99fc378ebdb9276177d8a6e",
		FRAXList[chainId.EthereumChainName]:          "0x3916e3b6c84b161df1b2733dffc9569a1da710c2",
		FEIList[chainId.EthereumChainName]:           "0xd89cf9e8a858f8b4b31faf793505e112d6c17449",
		STETHList[chainId.EthereumChainName]:         "0x66457616dd8489df5d0afd8678f4a260088aaf55",
		ENSList[chainId.EthereumChainName]:           "0x34441ffd1948e49dc7a607882d0c38efd0083815",
		USTList[chainId.EthereumChainName]:           "0x7fdbfb0412700d94403c42ca3caeeea183f07b26",
		CVXList[chainId.EthereumChainName]:           "0xb01eb1ce1da06179136d561766fc2d609c5f55eb",
		OneInchTokenList[chainId.EthereumChainName]:  "0x1278d6ed804d59d2d18a5aa5638dfd591a79af0a",
		LUSDList[chainId.EthereumChainName]:          "0x39f010127274b2dbdb770b45e1de54d974974526",
	},
}

// Aave v tokens v2.
//
// map[network][underlying] = address.
var AaveVTokenV2List = map[string]map[string]string{
	chainId.AvalancheChainName: {
		DAIeList[chainId.AvalancheChainName]:  "0x1852dc24d1a8956a0b356aa18ede954c7a0ca5ae", // variableDebtvDAI
		USDCeList[chainId.AvalancheChainName]: "0x848c080d2700cbe1b894a3374ad5e887e5ccb89c", // variableDebtvUSDC
		USDTeList[chainId.AvalancheChainName]: "0xfc1ada7a288d6fce0d29ccfaaa57bc9114bb2dbe", // variableDebtvUSDT
		WETHList[chainId.AvalancheChainName]:  "0x4e575cacb37bc1b5afec68a0462c4165a5268983", // variableDebtvWETH
		WBTCList[chainId.AvalancheChainName]:  "0x2dc0e35ec3ab070b8a175c829e23650ee604a9eb", // variableDebtvWBTC
		WAVAXList[chainId.AvalancheChainName]: "0x66a0fe52fb629a6cb4d10b8580afdffe888f5fd4", // variableDebtvWAVAX
		AAVEList[chainId.AvalancheChainName]:  "0x8352e3fd18b8d84d3c8a1b538d788899073c7a8e", // variableDebtvAAVE
	},
	chainId.EthereumChainName: {
		USDTList[chainId.EthereumChainName]:          "0x531842cebbdd378f8ee36d171d6cc9c4fcf475ec", // variableDebtUSDT 0
		WBTCList[chainId.EthereumChainName]:          "0x9c39809Dec7F95F5e0713634a4D0701329B3b4d2", // variableDebtWBTC 1
		WETHList[chainId.EthereumChainName]:          "0xF63B34710400CAd3e044cFfDcAb00a0f32E33eCf", // variableDebtWETH 2
		YFIList[chainId.EthereumChainName]:           "0x7EbD09022Be45AD993BAA1CEc61166Fcc8644d97", // variableDebtYFI 3 deprecated
		ZRXList[chainId.EthereumChainName]:           "0x85791D117A392097590bDeD3bD5abB8d5A20491A", // variableDebtZRX 4 deprecated
		UNIList[chainId.EthereumChainName]:           "0x5BdB050A92CADcCfCDcCCBFC17204a1C9cC0Ab73", // variableDebtUNI 5
		AAVEList[chainId.EthereumChainName]:          "0xF7DBA49d571745D9d7fcb56225B05BEA803EBf3C", // variableDebtAAVE 6
		BATList[chainId.EthereumChainName]:           "0xfc218A6Dfe6901CB34B1a5281FC6f1b8e7E56877", // variableDebtBAT 7 deprecated
		BUSDList[chainId.EthereumChainName]:          "0xbA429f7011c9fa04cDd46a2Da24dc0FF0aC6099c", // variableDebtBUSD 8
		DAIList[chainId.EthereumChainName]:           "0x6c3c78838c761c6ac7be9f59fe808ea2a6e4379d", // variableDebtDAI 9
		ENJList[chainId.EthereumChainName]:           "0x38995f292a6e31b78203254fe1cdd5ca1010a446", // variableDebtENJ 10 deprecated
		KyberOldTokenList[chainId.EthereumChainName]: "0x6B05D1c608015Ccb8e205A690cB86773A96F39f1", // variableDebtKNC 11 deprecated
		LINKList[chainId.EthereumChainName]:          "0x0b8f12b1788BFdE65Aa1ca52E3e9F3Ba401be16D", // variableDebtLINK 12
		MANAList[chainId.EthereumChainName]:          "0x0A68976301e46Ca6Ce7410DB28883E309EA0D352", // variableDebtMANA 13 deprecated
		MKRList[chainId.EthereumChainName]:           "0xba728ead5e496be00dcf66f650b6d7758ecb50f8", // variableDebtMKR 14
		RENList[chainId.EthereumChainName]:           "0xcd9d82d33bd737de215cdac57fe2f7f04df77fe0", // variableDebtREN 15 deprecated
		SNXList[chainId.EthereumChainName]:           "0x267eb8cf715455517f9bd5834aeae3cea1ebdbd8", // variableDebtSNX 16
		SUSDList[chainId.EthereumChainName]:          "0xdc6a3ab17299d9c2a412b0e0a4c1f55446ae0817", // variableDebtSUSD 17
		TUSDList[chainId.EthereumChainName]:          "0x01C0eb1f8c6F1C1bF74ae028697ce7AA2a8b0E92", // variableDebtTUSD 18
		USDCList[chainId.EthereumChainName]:          "0x619beb58998eD2278e08620f97007e1116D5D25b", // variableDebtUSDC 19
		CRVList[chainId.EthereumChainName]:           "0x00ad8eBF64F141f1C81e9f8f792d3d1631c6c684", // variableDebtCRV 20
		GUSDList[chainId.EthereumChainName]:          "0x279af5b99540c1a3a7e3cdd326e19659401ef99e", // variableDebtGUSD 21
		BALList[chainId.EthereumChainName]:           "0x13210D4Fe0d5402bd7Ecbc4B5bC5cFcA3b71adB0", // variableDebtBAL 22 deprecated
		XSushiList[chainId.EthereumChainName]:        "0xfAFEDF95E21184E3d880bd56D4806c4b8d31c69A", // variableDebtXSUSHI 23 deprecated
		RenFILList[chainId.EthereumChainName]:        "0x348e2ebd5e962854871874e444f4122399c02755", // variableDebtRENFIL 24 deprecated
		RAIList[chainId.EthereumChainName]:           "0xb5385132ee8321977fff44b60cde9fe9ab0b4e6b", // variableDebtRAI 25 deprecated
		AMPLList[chainId.EthereumChainName]:          "0xf013d90e4e4e3baf420dfea60735e75dbd42f1e1", // variableDebtAMPL 26 deprecated
		USDPList[chainId.EthereumChainName]:          "0xfdb93b3b10936cf81fa59a02a7523b6e2149b2b7", // variableDebtUSDP 27 (should be aPAX)
		DPIList[chainId.EthereumChainName]:           "0x4ddff5885a67e4effec55875a3977d7e60f82ae0", // variableDebtDPI 28
		FRAXList[chainId.EthereumChainName]:          "0xfe8f19b17ffef0fdbfe2671f248903055afaa8ca", // variableDebtFRAX 29
		FEIList[chainId.EthereumChainName]:           "0xC2e10006AccAb7B45D9184FcF5b7EC7763f5BaAe", // variableDebtFEI 30 deprecated
		STETHList[chainId.EthereumChainName]:         "0xA9DEAc9f00Dc4310c35603FCD9D34d1A750f81Db", // variableDebtSTETH 31
		ENSList[chainId.EthereumChainName]:           "0x176808047cc9b7A2C9AE202c593ED42dDD7C0D13", // variableDebtENS 32
		USTList[chainId.EthereumChainName]:           "0xaf32001cf2E66C4C3af4205F6EA77112AA4160FE", // variableDebtUST 33 deprecated
		CVXList[chainId.EthereumChainName]:           "0x4ae5e4409c6dbc84a00f9f89e4ba096603fb7d50", // variableDebtCVX 34 deprecated
		OneInchTokenList[chainId.EthereumChainName]:  "0xd7896c1b9b4455aff31473908eb15796ad2295da",
		LUSDList[chainId.EthereumChainName]:          "0x411066489ab40442d6fc215ad7c64224120d33f2",
	},
}
