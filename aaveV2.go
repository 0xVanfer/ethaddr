package ethaddr

// Last edited: March.19.2024.
// The listed or delisted tokens might be not up to date.
// For now, only ETH, POL, AVAX are still working.

// Website: https://aave.com/
//
// Docs: https://docs.aave.com/developers/getting-started/readme
//
// Deployed contracts: https://docs.aave.com/developers/deployed-contracts/deployed-contracts
//
// X (twitter): https://twitter.com/aave
const AaveV2Protocol string = "aave"

// Aave incentives controller V2.
//
// map[network] = address.
var AaveIncentivesControllerV2List = map[string]string{
	ChainAvalanche: "0x01D83Fe6A10D2f2B7AF17034343746188272cAc9", // 0x01d83fe6a10d2f2b7af17034343746188272cac9
	ChainEthereum:  "0xd784927Ff2f95ba542BfC824c8a8a98F3495f6b5", // 0xd784927ff2f95ba542bfc824c8a8a98f3495f6b5
	ChainPolygon:   "0x357D51124f59836DeD84c8a1730D72B749d8BC23", // 0x357d51124f59836ded84c8a1730d72b749d8bc23
}

// Aave lending pool v2.
//
// map[network] = address.
var AaveLendingPoolV2List = map[string]string{
	ChainAvalanche: "0x4F01AeD16D97E3aB5ab2B501154DC9bb0F1A5A2C", // 0x4f01aed16d97e3ab5ab2b501154dc9bb0f1a5a2c
	ChainEthereum:  "0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9", // 0x7d2768de32b0b80b7a3454c06bdac94a69ddc7a9
	ChainPolygon:   "0x8dFf5E27EA6b7AC08EbFdf9eB090F32ee9a30fcf", // 0x8dff5e27ea6b7ac08ebfdf9eb090f32ee9a30fcf
}

// Aave oracle v2.
//
// map[network] = address.
var AaveOracleV2List = map[string]string{
	ChainAvalanche: "0xdC336Cd4769f4cC7E9d726DA53e6d3fC710cEB89", // 0xdc336cd4769f4cc7e9d726da53e6d3fc710ceb89
	ChainEthereum:  "0xA50ba011c48153De246E5192C8f9258A2ba79Ca9", // 0xa50ba011c48153de246e5192c8f9258a2ba79ca9
	ChainPolygon:   "0x0229F777B0fAb107F9591a41d5F02E4e98dB6f2d", // 0x0229f777b0fab107f9591a41d5f02e4e98db6f2d
}

// Aave lending pool address provider v2.
//
// map[network] = address.
var AavePoolAddressProviderV2List = map[string]string{
	ChainAvalanche: "0xb6A86025F0FE1862B372cb0ca18CE3EDe02A318f", // 0xb6a86025f0fe1862b372cb0ca18ce3ede02a318f
	ChainEthereum:  "0xB53C1a33016B2DC2fF3653530bfF1848a515c8c5", // 0xb53c1a33016b2dc2ff3653530bff1848a515c8c5
	ChainPolygon:   "0xd05e3E715d945B59290df0ae8eF85c1BdB684744", // 0xd05e3e715d945b59290df0ae8ef85c1bdb684744
}

// Aave ui pool data provider v2.
//
// map[network] = address.
var AaveUiPoolDataProviderV2List = map[string]string{
	ChainAvalanche: "0x00e50FAB64eBB37b87df06Aa46b8B35d5f1A4e1A", // 0x00e50fab64ebb37b87df06aa46b8b35d5f1a4e1a
	ChainEthereum:  "0x00e50FAB64eBB37b87df06Aa46b8B35d5f1A4e1A", // 0x00e50fab64ebb37b87df06aa46b8b35d5f1a4e1a
	ChainPolygon:   "0x204f2Eb81D996729829debC819f7992DCEEfE7b1", // 0x204f2eb81d996729829debc819f7992dceefe7b1
}

// Aave ui incentives data provider v2.
//
// map[network] = address.
var AaveUiIncentiveDataProviderV2List = map[string]string{
	ChainAvalanche: "0x11979886A6dBAE27D7a72c49fCF3F23240D647bF", // 0x11979886a6dbae27d7a72c49fcf3f23240d647bf
	ChainEthereum:  "0xD01ab9a6577E1D84F142e44D49380e23A340387d", // 0xd01ab9a6577e1d84f142e44d49380e23a340387d
	ChainPolygon:   "0x645654D59A5226CBab969b1f5431aA47CBf64ab8", // 0x645654d59a5226cbab969b1f5431aa47cbf64ab8
}

// Aave a tokens v2.
//
// map[network][underlying] = address.
var AaveATokenV2List = map[string]map[string]string{
	ChainAvalanche: {
		DAIeList[ChainAvalanche]:  "0x47AFa96Cdc9fAb46904A55a6ad4bf6660B53c38a", // avDAI, 0x47afa96cdc9fab46904a55a6ad4bf6660b53c38a
		USDCeList[ChainAvalanche]: "0x46A51127C3ce23fb7AB1DE06226147F446e4a857", // avUSDC, 0x46a51127c3ce23fb7ab1de06226147f446e4a857
		USDTeList[ChainAvalanche]: "0x532E6537FEA298397212F09A61e03311686f548e", // avUSDT, 0x532e6537fea298397212f09a61e03311686f548e
		WETHList[ChainAvalanche]:  "0x53f7c5869a859F0AeC3D334ee8B4Cf01E3492f21", // avWETH, 0x53f7c5869a859f0aec3d334ee8b4cf01e3492f21
		WBTCList[ChainAvalanche]:  "0x686bEF2417b6Dc32C50a3cBfbCC3bb60E1e9a15D", // avWBTC, 0x686bef2417b6dc32c50a3cbfbcc3bb60e1e9a15d, delisted
		WAVAXList[ChainAvalanche]: "0xDFE521292EcE2A4f44242efBcD66Bc594CA9714B", // avWAVAX, 0xdfe521292ece2a4f44242efbcd66bc594ca9714b
		AAVEList[ChainAvalanche]:  "0xD45B7c061016102f9FA220502908f2c0f1add1D7", // avAAVE, 0xd45b7c061016102f9fa220502908f2c0f1add1d7
	},
	ChainEthereum: {
		USDTList[ChainEthereum]:          "0x3Ed3B47Dd13EC9a98b44e6204A523E766B225811", // aUSDT 0, 0x3ed3b47dd13ec9a98b44e6204a523e766b225811
		WBTCList[ChainEthereum]:          "0x9ff58f4fFB29fA2266Ab25e75e2A8b3503311656", // aWBTC 1, 0x9ff58f4ffb29fa2266ab25e75e2a8b3503311656
		WETHList[ChainEthereum]:          "0x030bA81f1c18d280636F32af80b9AAd02Cf0854e", // aWETH 2, 0x030ba81f1c18d280636f32af80b9aad02cf0854e
		YFIList[ChainEthereum]:           "0x5165d24277cD063F5ac44Efd447B27025e888f37", // aYFI 3, 0x5165d24277cd063f5ac44efd447b27025e888f37, delisted
		ZRXList[ChainEthereum]:           "0xDf7FF54aAcAcbFf42dfe29DD6144A69b629f8C9e", // aZRX 4, 0xdf7ff54aacacbff42dfe29dd6144a69b629f8c9e, delisted
		UNIList[ChainEthereum]:           "0xB9D7CB55f463405CDfBe4E90a6D2Df01C2B92BF1", // aUNI 5, 0xb9d7cb55f463405cdfbe4e90a6d2df01c2b92bf1, delisted
		AAVEList[ChainEthereum]:          "0xFFC97d72E13E01096502Cb8Eb52dEe56f74DAD7B", // aAAVE 6, 0xffc97d72e13e01096502cb8eb52dee56f74dad7b
		BATList[ChainEthereum]:           "0x05Ec93c0365baAeAbF7AefFb0972ea7ECdD39CF1", // aBAT 7, 0x05ec93c0365baaeabf7aeffb0972ea7ecdd39cf1, delisted
		BUSDList[ChainEthereum]:          "0xA361718326c15715591c299427c62086F69923D9", // aBUSD 8, 0xa361718326c15715591c299427c62086f69923d9, delisted
		DAIList[ChainEthereum]:           "0x028171bCA77440897B824Ca71D1c56caC55b68A3", // aDAI 9, 0x028171bca77440897b824ca71d1c56cac55b68a3
		ENJList[ChainEthereum]:           "0xaC6Df26a590F08dcC95D5a4705ae8abbc88509Ef", // aENJ 10, 0xac6df26a590f08dcc95d5a4705ae8abbc88509ef, delisted
		KyberOldTokenList[ChainEthereum]: "0x39C6b3e42d6A679d7D776778Fe880BC9487C2EDA", // aKNC 11, 0x39c6b3e42d6a679d7d776778fe880bc9487c2eda, delisted
		LINKList[ChainEthereum]:          "0xa06bC25B5805d5F8d82847D191Cb4Af5A3e873E0", // aLINK 12, 0xa06bc25b5805d5f8d82847d191cb4af5a3e873e0, delisted
		MANAList[ChainEthereum]:          "0xa685a61171bb30d4072B338c80Cb7b2c865c873E", // aMANA 13, 0xa685a61171bb30d4072b338c80cb7b2c865c873e, delisted
		MKRList[ChainEthereum]:           "0xc713e5E149D5D0715DcD1c156a020976e7E56B88", // aMKR 14, 0xc713e5e149d5d0715dcd1c156a020976e7e56b88, delisted
		RENList[ChainEthereum]:           "0xCC12AbE4ff81c9378D670De1b57F8e0Dd228D77a", // aREN 15, 0xcc12abe4ff81c9378d670de1b57f8e0dd228d77a, delisted
		SNXList[ChainEthereum]:           "0x35f6B052C598d933D69A4EEC4D04c73A191fE6c2", // aSNX 16, 0x35f6b052c598d933d69a4eec4d04c73a191fe6c2, delisted
		SUSDList[ChainEthereum]:          "0x6C5024Cd4F8A59110119C56f8933403A539555EB", // aSUSD 17, 0x6c5024cd4f8a59110119c56f8933403a539555eb
		TUSDList[ChainEthereum]:          "0x101cc05f4A51C0319f570d5E146a8C625198e636", // aTUSD 18, 0x101cc05f4a51c0319f570d5e146a8c625198e636, delisted
		USDCList[ChainEthereum]:          "0xBcca60bB61934080951369a648Fb03DF4F96263C", // aUSDC 19, 0xbcca60bb61934080951369a648fb03df4f96263c
		CRVList[ChainEthereum]:           "0x8dAE6Cb04688C62d939ed9B68d32Bc62e49970b1", // aCRV 20, 0x8dae6cb04688c62d939ed9b68d32bc62e49970b1, delisted
		GUSDList[ChainEthereum]:          "0xD37EE7e4f452C6638c96536e68090De8cBcdb583", // aGUSD 21, 0xd37ee7e4f452c6638c96536e68090de8cbcdb583
		BALList[ChainEthereum]:           "0x272F97b7a56a387aE942350bBC7Df5700f8a4576", // aBAL 22, 0x272f97b7a56a387ae942350bbc7df5700f8a4576, delisted
		XSushiList[ChainEthereum]:        "0xF256CC7847E919FAc9B808cC216cAc87CCF2f47a", // aXSUSHI 23, 0xf256cc7847e919fac9b808cc216cac87ccf2f47a, delisted
		RenFILList[ChainEthereum]:        "0x514cd6756CCBe28772d4Cb81bC3156BA9d1744aa", // aRENFIL 24, 0x514cd6756ccbe28772d4cb81bc3156ba9d1744aa, delisted
		RAIList[ChainEthereum]:           "0xc9BC48c72154ef3e5425641a3c747242112a46AF", // aRAI 25, 0xc9bc48c72154ef3e5425641a3c747242112a46af, delisted
		AMPLList[ChainEthereum]:          "0x1E6bb68Acec8fefBD87D192bE09bb274170a0548", // aAMPL 26, 0x1e6bb68acec8fefbd87d192be09bb274170a0548, delisted
		USDPList[ChainEthereum]:          "0x2e8F4bdbE3d47d7d7DE490437AeA9915D930F1A3", // aUSDP 27, 0x2e8f4bdbe3d47d7d7de490437aea9915d930f1a3
		DPIList[ChainEthereum]:           "0x6F634c6135D2EBD550000ac92F494F9CB8183dAe", // aDPI 28, 0x6f634c6135d2ebd550000ac92f494f9cb8183dae, delisted
		FRAXList[ChainEthereum]:          "0xd4937682df3C8aEF4FE912A96A74121C0829E664", // aFRAX 29, 0xd4937682df3c8aef4fe912a96a74121c0829e664
		FEIList[ChainEthereum]:           "0x683923dB55Fead99A79Fa01A27EeC3cB19679cC3", // aFEI 30, 0x683923db55fead99a79fa01a27eec3cb19679cc3, delisted
		STETHList[ChainEthereum]:         "0x1982b2F5814301d4e9a8b0201555376e62F82428", // aSTETH 31, 0x1982b2f5814301d4e9a8b0201555376e62f82428
		ENSList[ChainEthereum]:           "0x9a14e23A58edf4EFDcB360f68cd1b95ce2081a2F", // aENS 32, 0x9a14e23a58edf4efdcb360f68cd1b95ce2081a2f, delisted
		USTList[ChainEthereum]:           "0xc2e2152647F4C26028482Efaf64b2Aa28779EFC4", // aUST 33, 0xc2e2152647f4c26028482efaf64b2aa28779efc4, delisted
		CVXList[ChainEthereum]:           "0x952749E07d7157bb9644A894dFAF3Bad5eF6D918", // aCVX 34, 0x952749e07d7157bb9644a894dfaf3bad5ef6d918, delisted
		OINCHLIST[ChainEthereum]:         "0xB29130CBcC3F791f077eAdE0266168E808E5151e", // a1INCH 35, 0xb29130cbcc3f791f077eade0266168e808e5151e, delisted
		LUSDList[ChainEthereum]:          "0xce1871f791548600cb59efbefFC9c38719142079", // aLUSD 36, 0xce1871f791548600cb59efbeffc9c38719142079
	},
	ChainPolygon: {
		AAVEList[ChainPolygon]:   "0x1d2a0E5EC8E5bBDCA5CB219e649B565d8e5c3360", // amAAVE, 0x1d2a0e5ec8e5bbdca5cb219e649b565d8e5c3360
		DAIList[ChainPolygon]:    "0x27F8D03b3a2196956ED754baDc28D73be8830A6e", // amDAI, 0x27f8d03b3a2196956ed754badc28d73be8830a6e
		USDCeList[ChainPolygon]:  "0x1a13F4Ca1d028320A707D99520AbFefca3998b7F", // amUSDC, 0x1a13f4ca1d028320a707d99520abfefca3998b7f
		USDTList[ChainPolygon]:   "0x60D55F02A771d515e077c9C2403a1ef324885CeC", // amUSDT, 0x60d55f02a771d515e077c9c2403a1ef324885cec
		WBTCList[ChainPolygon]:   "0x5c2ed810328349100A66B82b78a1791B101C9D61", // amWBTC, 0x5c2ed810328349100a66b82b78a1791b101c9d61
		WETHList[ChainPolygon]:   "0x28424507fefb6f7f8E9D3860F56504E4e5f5f390", // amWETH, 0x28424507fefb6f7f8e9d3860f56504e4e5f5f390
		WMATICList[ChainPolygon]: "0x8dF3aad3a84da6b69A4DA8aeC3eA40d9091B2Ac4", // amWMATIC, 0x8df3aad3a84da6b69a4da8aec3ea40d9091b2ac4
	},
}

// Aave s tokens v2.
//
// map[network][underlying] = address.
var AaveSTokenV2List = map[string]map[string]string{
	ChainAvalanche: {
		DAIeList[ChainAvalanche]:  "0x3676E4EE689D527dDb89812B63fAD0B7501772B3", // stableDebtvDAI, 0x3676e4ee689d527ddb89812b63fad0b7501772b3
		USDCeList[ChainAvalanche]: "0x5B14679135dbE8B02015ec3Ca4924a12E4C6C85a", // stableDebtvUSDC, 0x5b14679135dbe8b02015ec3ca4924a12e4c6c85a
		USDTeList[ChainAvalanche]: "0x9c7B81A867499B7387ed05017a13d4172a0c17bF", // stableDebtvUSDT, 0x9c7b81a867499b7387ed05017a13d4172a0c17bf
		WETHList[ChainAvalanche]:  "0x60F6A45006323B97d97cB0a42ac39e2b757ADA63", // stableDebtvWETH, 0x60f6a45006323b97d97cb0a42ac39e2b757ada63
		WBTCList[ChainAvalanche]:  "0x3484408989985d68C9700dc1CFDFeAe6d2f658CF", // stableDebtvWBTC, 0x3484408989985d68c9700dc1cfdfeae6d2f658cf, delisted
		WAVAXList[ChainAvalanche]: "0x2920CD5b8A160b2Addb00Ec5d5f4112255d4ae75", // stableDebtvWAVAX, 0x2920cd5b8a160b2addb00ec5d5f4112255d4ae75
		AAVEList[ChainAvalanche]:  "0x66904E4F3f44e3925D22ceca401b6F2DA085c98f", // stableDebtvAAVE, 0x66904e4f3f44e3925d22ceca401b6f2da085c98f
	},
	ChainEthereum: {
		USDTList[ChainEthereum]:          "0xe91D55AB2240594855aBd11b3faAE801Fd4c4687", // stableDebtUSDT 0, 0xe91d55ab2240594855abd11b3faae801fd4c4687
		WBTCList[ChainEthereum]:          "0x51B039b9AFE64B78758f8Ef091211b5387eA717c", // stableDebtWBTC 1, 0x51b039b9afe64b78758f8ef091211b5387ea717c
		WETHList[ChainEthereum]:          "0x4e977830ba4bd783C0BB7F15d3e243f73FF57121", // stableDebtWETH 2, 0x4e977830ba4bd783c0bb7f15d3e243f73ff57121
		YFIList[ChainEthereum]:           "0xca823F78C2Dd38993284bb42Ba9b14152082F7BD", // stableDebtYFI 3, 0xca823f78c2dd38993284bb42ba9b14152082f7bd, delisted
		ZRXList[ChainEthereum]:           "0x071B4323a24E73A5afeEbe34118Cd21B8FAAF7C3", // stableDebtZRX 4, 0x071b4323a24e73a5afeebe34118cd21b8faaf7c3, delisted
		UNIList[ChainEthereum]:           "0xD939F7430dC8D5a427f156dE1012A56C18AcB6Aa", // stableDebtUNI 5, 0xd939f7430dc8d5a427f156de1012a56c18acb6aa, delisted
		AAVEList[ChainEthereum]:          "0x079D6a3E844BcECf5720478A718Edb6575362C5f", // stableDebtAAVE 6, 0x079d6a3e844bcecf5720478a718edb6575362c5f
		BATList[ChainEthereum]:           "0x277f8676FAcf4dAA5a6EA38ba511B7F65AA02f9F", // stableDebtBAT 7, 0x277f8676facf4daa5a6ea38ba511b7f65aa02f9f, delisted
		BUSDList[ChainEthereum]:          "0x4A7A63909A72D268b1D8a93a9395d098688e0e5C", // stableDebtBUSD 8, 0x4a7a63909a72d268b1d8a93a9395d098688e0e5c, delisted
		DAIList[ChainEthereum]:           "0x778A13D3eeb110A4f7bb6529F99c000119a08E92", // stableDebtDAI 9, 0x778a13d3eeb110a4f7bb6529f99c000119a08e92
		ENJList[ChainEthereum]:           "0x943DcCA156b5312Aa24c1a08769D67FEce4ac14C", // stableDebtENJ 10, 0x943dcca156b5312aa24c1a08769d67fece4ac14c, delisted
		KyberOldTokenList[ChainEthereum]: "0x9915dfb872778B2890a117DA1F35F335eb06B54f", // stableDebtKNC 11, 0x9915dfb872778b2890a117da1f35f335eb06b54f, delisted
		LINKList[ChainEthereum]:          "0xFB4AEc4Cc858F2539EBd3D37f2a43eAe5b15b98a", // stableDebtLINK 12, 0xfb4aec4cc858f2539ebd3d37f2a43eae5b15b98a, delisted
		MANAList[ChainEthereum]:          "0xD86C74eA2224f4B8591560652b50035E4e5c0a3b", // stableDebtMANA 13, 0xd86c74ea2224f4b8591560652b50035e4e5c0a3b, delisted
		MKRList[ChainEthereum]:           "0xC01C8E4b12a89456a9fD4e4e75B72546Bf53f0B5", // stableDebtMKR 14, 0xc01c8e4b12a89456a9fd4e4e75b72546bf53f0b5, delisted
		RENList[ChainEthereum]:           "0x3356Ec1eFA75d9D150Da1EC7d944D9EDf73703B7", // stableDebtREN 15, 0x3356ec1efa75d9d150da1ec7d944d9edf73703b7, delisted
		SNXList[ChainEthereum]:           "0x8575c8ae70bDB71606A53AeA1c6789cB0fBF3166", // stableDebtSNX 16, 0x8575c8ae70bdb71606a53aea1c6789cb0fbf3166, delisted
		SUSDList[ChainEthereum]:          "0x30B0f7324feDF89d8eff397275F8983397eFe4af", // stableDebtSUSD 17, 0x30b0f7324fedf89d8eff397275f8983397efe4af
		TUSDList[ChainEthereum]:          "0x7f38d60D94652072b2C44a18c0e14A481EC3C0dd", // stableDebtTUSD 18, 0x7f38d60d94652072b2c44a18c0e14a481ec3c0dd, delisted
		USDCList[ChainEthereum]:          "0xE4922afAB0BbaDd8ab2a88E0C79d884Ad337fcA6", // stableDebtUSDC 19, 0xe4922afab0bbadd8ab2a88e0c79d884ad337fca6
		CRVList[ChainEthereum]:           "0x9288059a74f589C919c7Cf1Db433251CdFEB874B", // stableDebtCRV 20, 0x9288059a74f589c919c7cf1db433251cdfeb874b, delisted
		GUSDList[ChainEthereum]:          "0xf8aC64ec6Ff8E0028b37EB89772d21865321bCe0", // stableDebtGUSD 21, 0xf8ac64ec6ff8e0028b37eb89772d21865321bce0
		BALList[ChainEthereum]:           "0xe569d31590307d05DA3812964F1eDd551D665a0b", // stableDebtBAL 22, 0xe569d31590307d05da3812964f1edd551d665a0b, delisted
		XSushiList[ChainEthereum]:        "0x73Bfb81D7dbA75C904f430eA8BAe82DB0D41187B", // stableDebtXSUSHI 23, 0x73bfb81d7dba75c904f430ea8bae82db0d41187b, delisted
		RenFILList[ChainEthereum]:        "0xcAad05C49E14075077915cB5C820EB3245aFb950", // stableDebtRENFIL 24, 0xcaad05c49e14075077915cb5c820eb3245afb950, delisted
		RAIList[ChainEthereum]:           "0x9C72B8476C33AE214ee3e8C20F0bc28496a62032", // stableDebtRAI 25, 0x9c72b8476c33ae214ee3e8c20f0bc28496a62032, delisted
		AMPLList[ChainEthereum]:          "0x18152C9f77DAdc737006e9430dB913159645fa87", // stableDebtAMPL 26, 0x18152c9f77dadc737006e9430db913159645fa87, delisted
		USDPList[ChainEthereum]:          "0x2387119bc85A74e0BBcbe190d80676CB16F10D4F", // stableDebtUSDP 27, 0x2387119bc85a74e0bbcbe190d80676cb16f10d4f
		DPIList[ChainEthereum]:           "0xa3953F07f389d719F99FC378ebDb9276177d8A6e", // stableDebtDPI 28, 0xa3953f07f389d719f99fc378ebdb9276177d8a6e, delisted
		FRAXList[ChainEthereum]:          "0x3916e3B6c84b161df1b2733dFfc9569a1dA710c2", // stableDebtFRAX 29, 0x3916e3b6c84b161df1b2733dffc9569a1da710c2
		FEIList[ChainEthereum]:           "0xd89cF9E8A858F8B4b31Faf793505e112d6c17449", // stableDebtFEI 30, 0xd89cf9e8a858f8b4b31faf793505e112d6c17449, delisted
		STETHList[ChainEthereum]:         "0x66457616Dd8489dF5D0AFD8678F4A260088aAF55", // stableDebtSTETH 31, 0x66457616dd8489df5d0afd8678f4a260088aaf55
		ENSList[ChainEthereum]:           "0x34441FFD1948E49dC7a607882D0c38Efd0083815", // stableDebtENS 32, 0x34441ffd1948e49dc7a607882d0c38efd0083815, delisted
		USTList[ChainEthereum]:           "0x7FDbfB0412700D94403c42cA3CAEeeA183F07B26", // stableDebtUST 33, 0x7fdbfb0412700d94403c42ca3caeeea183f07b26, delisted
		CVXList[ChainEthereum]:           "0xB01Eb1cE1Da06179136D561766fc2d609C5F55Eb", // stableDebtCVX 34, 0xb01eb1ce1da06179136d561766fc2d609c5f55eb, delisted
		OINCHLIST[ChainEthereum]:         "0x1278d6ED804d59d2d18a5Aa5638DfD591A79aF0a", // stableDebt1INCH 35, 0x1278d6ed804d59d2d18a5aa5638dfd591a79af0a, delisted
		LUSDList[ChainEthereum]:          "0x39f010127274b2dBdB770B45e1de54d974974526", // stableDebtLUSD 36, 0x39f010127274b2dbdb770b45e1de54d974974526
	},
	ChainPolygon: {
		AAVEList[ChainPolygon]:   "0x17912140e780B29Ba01381F088f21E8d75F954F9", // stableDebtmAAVE, 0x17912140e780b29ba01381f088f21e8d75f954f9
		DAIList[ChainPolygon]:    "0x2238101B7014C279aaF6b408A284E49cDBd5DB55", // stableDebtmDAI, 0x2238101b7014c279aaf6b408a284e49cdbd5db55
		USDCeList[ChainPolygon]:  "0xdeb05676dB0DB85cecafE8933c903466Bf20C572", // stableDebtmUSDC, 0xdeb05676db0db85cecafe8933c903466bf20c572
		USDTList[ChainPolygon]:   "0xe590cfca10e81FeD9B0e4496381f02256f5d2f61", // stableDebtmUSDT, 0xe590cfca10e81fed9b0e4496381f02256f5d2f61
		WBTCList[ChainPolygon]:   "0x2551B15dB740dB8348bFaDFe06830210eC2c2F13", // stableDebtmWBTC, 0x2551b15db740db8348bfadfe06830210ec2c2f13
		WETHList[ChainPolygon]:   "0xc478cBbeB590C76b01ce658f8C4dda04f30e2C6f", // stableDebtmWETH, 0xc478cbbeb590c76b01ce658f8c4dda04f30e2c6f
		WMATICList[ChainPolygon]: "0xb9A6E29fB540C5F1243ef643EB39b0AcbC2e68E3", // stableDebtmWMATIC, 0xb9a6e29fb540c5f1243ef643eb39b0acbc2e68e3
	},
}

// Aave v tokens v2.
//
// map[network][underlying] = address.
var AaveVTokenV2List = map[string]map[string]string{
	ChainAvalanche: {
		DAIeList[ChainAvalanche]:  "0x1852DC24d1a8956a0B356AA18eDe954c7a0Ca5ae", // variableDebtvDAI, 0x1852dc24d1a8956a0b356aa18ede954c7a0ca5ae
		USDCeList[ChainAvalanche]: "0x848c080d2700CBE1B894a3374AD5E887E5cCb89c", // variableDebtvUSDC, 0x848c080d2700cbe1b894a3374ad5e887e5ccb89c
		USDTeList[ChainAvalanche]: "0xfc1AdA7A288d6fCe0d29CcfAAa57Bc9114bb2DbE", // variableDebtvUSDT, 0xfc1ada7a288d6fce0d29ccfaaa57bc9114bb2dbe
		WETHList[ChainAvalanche]:  "0x4e575CacB37bc1b5afEc68a0462c4165A5268983", // variableDebtvWETH, 0x4e575cacb37bc1b5afec68a0462c4165a5268983
		WBTCList[ChainAvalanche]:  "0x2dc0E35eC3Ab070B8a175C829e23650Ee604a9eB", // variableDebtvWBTC, 0x2dc0e35ec3ab070b8a175c829e23650ee604a9eb, delisted
		WAVAXList[ChainAvalanche]: "0x66A0FE52Fb629a6cB4D10B8580AFDffE888F5Fd4", // variableDebtvWAVAX, 0x66a0fe52fb629a6cb4d10b8580afdffe888f5fd4
		AAVEList[ChainAvalanche]:  "0x8352E3fd18B8d84D3c8a1b538d788899073c7A8E", // variableDebtvAAVE, 0x8352e3fd18b8d84d3c8a1b538d788899073c7a8e
	},
	ChainEthereum: {
		USDTList[ChainEthereum]:          "0x531842cEbbdD378f8ee36D171d6cC9C4fcf475Ec", // variableDebtUSDT 0, 0x531842cebbdd378f8ee36d171d6cc9c4fcf475ec
		WBTCList[ChainEthereum]:          "0x9c39809Dec7F95F5e0713634a4D0701329B3b4d2", // variableDebtWBTC 1, 0x9c39809dec7f95f5e0713634a4d0701329b3b4d2
		WETHList[ChainEthereum]:          "0xF63B34710400CAd3e044cFfDcAb00a0f32E33eCf", // variableDebtWETH 2, 0xf63b34710400cad3e044cffdcab00a0f32e33ecf
		YFIList[ChainEthereum]:           "0x7EbD09022Be45AD993BAA1CEc61166Fcc8644d97", // variableDebtYFI 3, 0x7ebd09022be45ad993baa1cec61166fcc8644d97, delisted
		ZRXList[ChainEthereum]:           "0x85791D117A392097590bDeD3bD5abB8d5A20491A", // variableDebtZRX 4, 0x85791d117a392097590bded3bd5abb8d5a20491a, delisted
		UNIList[ChainEthereum]:           "0x5BdB050A92CADcCfCDcCCBFC17204a1C9cC0Ab73", // variableDebtUNI 5, 0x5bdb050a92cadccfcdcccbfc17204a1c9cc0ab73, delisted
		AAVEList[ChainEthereum]:          "0xF7DBA49d571745D9d7fcb56225B05BEA803EBf3C", // variableDebtAAVE 6, 0xf7dba49d571745d9d7fcb56225b05bea803ebf3c
		BATList[ChainEthereum]:           "0xfc218A6Dfe6901CB34B1a5281FC6f1b8e7E56877", // variableDebtBAT 7, 0xfc218a6dfe6901cb34b1a5281fc6f1b8e7e56877, delisted
		BUSDList[ChainEthereum]:          "0xbA429f7011c9fa04cDd46a2Da24dc0FF0aC6099c", // variableDebtBUSD 8, 0xba429f7011c9fa04cdd46a2da24dc0ff0ac6099c, delisted
		DAIList[ChainEthereum]:           "0x6C3c78838c761c6Ac7bE9F59fe808ea2A6E4379d", // variableDebtDAI 9, 0x6c3c78838c761c6ac7be9f59fe808ea2a6e4379d
		ENJList[ChainEthereum]:           "0x38995F292a6E31b78203254fE1cdd5Ca1010A446", // variableDebtENJ 10, 0x38995f292a6e31b78203254fe1cdd5ca1010a446, delisted
		KyberOldTokenList[ChainEthereum]: "0x6B05D1c608015Ccb8e205A690cB86773A96F39f1", // variableDebtKNC 11, 0x6b05d1c608015ccb8e205a690cb86773a96f39f1, delisted
		LINKList[ChainEthereum]:          "0x0b8f12b1788BFdE65Aa1ca52E3e9F3Ba401be16D", // variableDebtLINK 12, 0x0b8f12b1788bfde65aa1ca52e3e9f3ba401be16d, delisted
		MANAList[ChainEthereum]:          "0x0A68976301e46Ca6Ce7410DB28883E309EA0D352", // variableDebtMANA 13, 0x0a68976301e46ca6ce7410db28883e309ea0d352, delisted
		MKRList[ChainEthereum]:           "0xba728eAd5e496BE00DCF66F650b6d7758eCB50f8", // variableDebtMKR 14, 0xba728ead5e496be00dcf66f650b6d7758ecb50f8, delisted
		RENList[ChainEthereum]:           "0xcd9D82d33bd737De215cDac57FE2F7f04DF77FE0", // variableDebtREN 15, 0xcd9d82d33bd737de215cdac57fe2f7f04df77fe0, delisted
		SNXList[ChainEthereum]:           "0x267EB8Cf715455517F9BD5834AeAE3CeA1EBdbD8", // variableDebtSNX 16, 0x267eb8cf715455517f9bd5834aeae3cea1ebdbd8, delisted
		SUSDList[ChainEthereum]:          "0xdC6a3Ab17299D9C2A412B0e0a4C1f55446AE0817", // variableDebtSUSD 17, 0xdc6a3ab17299d9c2a412b0e0a4c1f55446ae0817
		TUSDList[ChainEthereum]:          "0x01C0eb1f8c6F1C1bF74ae028697ce7AA2a8b0E92", // variableDebtTUSD 18, 0x01c0eb1f8c6f1c1bf74ae028697ce7aa2a8b0e92, delisted
		USDCList[ChainEthereum]:          "0x619beb58998eD2278e08620f97007e1116D5D25b", // variableDebtUSDC 19, 0x619beb58998ed2278e08620f97007e1116d5d25b
		CRVList[ChainEthereum]:           "0x00ad8eBF64F141f1C81e9f8f792d3d1631c6c684", // variableDebtCRV 20, 0x00ad8ebf64f141f1c81e9f8f792d3d1631c6c684, delisted
		GUSDList[ChainEthereum]:          "0x279AF5b99540c1A3A7E3CDd326e19659401eF99e", // variableDebtGUSD 21, 0x279af5b99540c1a3a7e3cdd326e19659401ef99e
		BALList[ChainEthereum]:           "0x13210D4Fe0d5402bd7Ecbc4B5bC5cFcA3b71adB0", // variableDebtBAL 22, 0x13210d4fe0d5402bd7ecbc4b5bc5cfca3b71adb0, delisted
		XSushiList[ChainEthereum]:        "0xfAFEDF95E21184E3d880bd56D4806c4b8d31c69A", // variableDebtXSUSHI 23, 0xfafedf95e21184e3d880bd56d4806c4b8d31c69a, delisted
		RenFILList[ChainEthereum]:        "0x348e2eBD5E962854871874E444F4122399c02755", // variableDebtRENFIL 24, 0x348e2ebd5e962854871874e444f4122399c02755, delisted
		RAIList[ChainEthereum]:           "0xB5385132EE8321977FfF44b60cDE9fE9AB0B4e6b", // variableDebtRAI 25, 0xb5385132ee8321977fff44b60cde9fe9ab0b4e6b, delisted
		AMPLList[ChainEthereum]:          "0xf013D90E4e4E3Baf420dFea60735e75dbd42f1e1", // variableDebtAMPL 26, 0xf013d90e4e4e3baf420dfea60735e75dbd42f1e1, delisted
		USDPList[ChainEthereum]:          "0xFDb93B3b10936cf81FA59A02A7523B6e2149b2B7", // variableDebtUSDP 27, 0xfdb93b3b10936cf81fa59a02a7523b6e2149b2b7
		DPIList[ChainEthereum]:           "0x4dDff5885a67E4EffeC55875a3977D7E60F82ae0", // variableDebtDPI 28, 0x4ddff5885a67e4effec55875a3977d7e60f82ae0, delisted
		FRAXList[ChainEthereum]:          "0xfE8F19B17fFeF0fDbfe2671F248903055AFAA8Ca", // variableDebtFRAX 29, 0xfe8f19b17ffef0fdbfe2671f248903055afaa8ca
		FEIList[ChainEthereum]:           "0xC2e10006AccAb7B45D9184FcF5b7EC7763f5BaAe", // variableDebtFEI 30, 0xc2e10006accab7b45d9184fcf5b7ec7763f5baae, delisted
		STETHList[ChainEthereum]:         "0xA9DEAc9f00Dc4310c35603FCD9D34d1A750f81Db", // variableDebtSTETH 31, 0xa9deac9f00dc4310c35603fcd9d34d1a750f81db
		ENSList[ChainEthereum]:           "0x176808047cc9b7A2C9AE202c593ED42dDD7C0D13", // variableDebtENS 32, 0x176808047cc9b7a2c9ae202c593ed42ddd7c0d13, delisted
		USTList[ChainEthereum]:           "0xaf32001cf2E66C4C3af4205F6EA77112AA4160FE", // variableDebtUST 33, 0xaf32001cf2e66c4c3af4205f6ea77112aa4160fe, delisted
		CVXList[ChainEthereum]:           "0x4Ae5E4409C6Dbc84A00f9f89e4ba096603fb7d50", // variableDebtCVX 34, 0x4ae5e4409c6dbc84a00f9f89e4ba096603fb7d50, delisted
		OINCHLIST[ChainEthereum]:         "0xD7896C1B9b4455aFf31473908eB15796ad2295DA", // variableDebt1INCH 35, 0xd7896c1b9b4455aff31473908eb15796ad2295da, delisted
		LUSDList[ChainEthereum]:          "0x411066489AB40442d6Fc215aD7c64224120D33F2", // variableDebtLUSD 36, 0x411066489ab40442d6fc215ad7c64224120d33f2
	},
	ChainPolygon: {
		AAVEList[ChainPolygon]:   "0x1c313e9d0d826662F5CE692134D938656F681350", // variableDebtmAAVE, 0x1c313e9d0d826662f5ce692134d938656f681350
		DAIList[ChainPolygon]:    "0x75c4d1Fb84429023170086f06E682DcbBF537b7d", // variableDebtmDAI, 0x75c4d1fb84429023170086f06e682dcbbf537b7d
		USDCeList[ChainPolygon]:  "0x248960A9d75EdFa3de94F7193eae3161Eb349a12", // variableDebtmUSDC, 0x248960a9d75edfa3de94f7193eae3161eb349a12
		USDTList[ChainPolygon]:   "0x8038857FD47108A07d1f6Bf652ef1cBeC279A2f3", // variableDebtmUSDT, 0x8038857fd47108a07d1f6bf652ef1cbec279a2f3
		WBTCList[ChainPolygon]:   "0xF664F50631A6f0D72ecdaa0e49b0c019Fa72a8dC", // variableDebtmWBTC, 0xf664f50631a6f0d72ecdaa0e49b0c019fa72a8dc
		WETHList[ChainPolygon]:   "0xeDe17e9d79fc6f9fF9250D9EEfbdB88Cc18038b5", // variableDebtmWETH, 0xede17e9d79fc6f9ff9250d9eefbdb88cc18038b5
		WMATICList[ChainPolygon]: "0x59e8E9100cbfCBCBAdf86b9279fa61526bBB8765", // variableDebtmWMATIC, 0x59e8e9100cbfcbcbadf86b9279fa61526bbb8765
	},
}
