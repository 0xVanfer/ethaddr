package ethaddr

// Website: https://app.spark.fi/markets/
const SparkProtocol string = "spark"

// Spark Saving DAI address.
//
// map[network] = address.
var SDAIList = map[int64]string{
	ChainEthereum: "0x83F20F44975D03b1b09e64809B757c47f942BEeA", // sDAI, 0x83f20f44975d03b1b09e64809b757c47f942beea
}

var SparkPoolList = map[int64]string{
	ChainEthereum: "0xC13e21B648A5Ee794902342038FF3aDAB66BE987", // 0xc13e21b648a5ee794902342038ff3adab66be987
}

var SparkPoolAddressProviderList = map[int64]string{
	ChainEthereum: "0x02C3eA4e34C0cBd694D2adFa2c690EECbC1793eE", // 0x02c3ea4e34c0cbd694d2adfa2c690eecbc1793ee
}

var SparkDataProviderList = map[int64]string{
	ChainEthereum: "0xFc21d6d146E6086B8359705C8b28512a983db0cb", // 0xfc21d6d146e6086b8359705c8b28512a983db0cb
}

var SparkTreasuryList = map[int64]string{
	ChainEthereum: "0xb137E7d16564c81ae2b0C8ee6B55De81dd46ECe5", // 0xb137e7d16564c81ae2b0c8ee6b55de81dd46ece5
}

var SparkTreasuryControllerList = map[int64]string{
	ChainEthereum: "0x92eF091C5a1E01b3CE1ba0D0150C84412d818F7a", // 0x92ef091c5a1e01b3ce1ba0d0150c84412d818f7a
}

var SparkIncentivesList = map[int64]string{
	ChainEthereum: "0x4370D3b6C9588E02ce9D22e684387859c7Ff5b34", // 0x4370d3b6c9588e02ce9d22e684387859c7ff5b34
}

var SparkUiPoolDataProviderList = map[int64]string{
	ChainEthereum: "0xF028c2F4b19898718fD0F77b9b881CbfdAa5e8Bb", // 0xf028c2f4b19898718fd0f77b9b881cbfdaa5e8bb
}

var SparkUiIncentiveDataProviderList = map[int64]string{
	ChainEthereum: "0xA7F8A757C4f7696c015B595F51B2901AC0121B18", // 0xa7f8a757c4f7696c015b595f51b2901ac0121b18
}

// Spark a tokens.
//
// map[network][underlying] = address.
var SparkATokenList = map[int64]map[string]string{
	ChainEthereum: {
		DAIList[ChainEthereum]:    "0x4DEDf26112B3Ec8eC46e7E31EA5e123490B05B8B", // spDAI, 0x4dedf26112b3ec8ec46e7e31ea5e123490b05b8b
		USDCList[ChainEthereum]:   "0x377C3bd93f2a2984E1E7bE6A5C22c525eD4A4815", // spUSDC, 0x377c3bd93f2a2984e1e7be6a5c22c525ed4a4815
		USDTList[ChainEthereum]:   "0xe7dF13b8e3d6740fe17CBE928C7334243d86c92f", // spUSDT, 0xe7df13b8e3d6740fe17cbe928c7334243d86c92f
		RETHList[ChainEthereum]:   "0x9985dF20D7e9103ECBCeb16a84956434B6f06ae8", // sprETH, 0x9985df20d7e9103ecbceb16a84956434b6f06ae8
		SDAIList[ChainEthereum]:   "0x78f897F0fE2d3B5690EbAe7f19862DEacedF10a7", // spsDAI, 0x78f897f0fe2d3b5690ebae7f19862deacedf10a7
		WBTCList[ChainEthereum]:   "0x4197ba364AE6698015AE5c1468f54087602715b2", // spWBTC, 0x4197ba364ae6698015ae5c1468f54087602715b2
		WETHList[ChainEthereum]:   "0x59cD1C87501baa753d0B5B5Ab5D8416A45cD71DB", // spWETH, 0x59cd1c87501baa753d0b5b5ab5d8416a45cd71db
		WSTETHList[ChainEthereum]: "0x12B54025C112Aa61fAce2CDB7118740875A566E9", // spwstETH, 0x12b54025c112aa61face2cdb7118740875a566e9
		WEETHList[ChainEthereum]:  "0x3CFd5C0D4acAA8Faee335842e4f31159fc76B008", // spweETH, 0x3cfd5c0d4acaa8faee335842e4f31159fc76b008
		GNOList[ChainEthereum]:    "0x7b481aCC9fDADDc9af2cBEA1Ff2342CB1733E50F", // spGNO, 0x7b481acc9fdaddc9af2cbea1ff2342cb1733e50f
		CbBTCList[ChainEthereum]:  "0xb3973D459df38ae57797811F2A1fd061DA1BC123", // spcbBTC, 0xb3973d459df38ae57797811f2a1fd061da1bc123
		SUSDSList[ChainEthereum]:  "0x6715bc100A183cc65502F05845b589c1919ca3d3", // spsUSDS, 0x6715bc100a183cc65502f05845b589c1919ca3d3
	},
}

// Spark v tokens.
//
// map[network][underlying] = address.
var SparkVTokenList = map[int64]map[string]string{
	ChainEthereum: {
		DAIList[ChainEthereum]:    "0xf705d2B7e92B3F38e6ae7afaDAA2fEE110fE5914", // variableDebtDAI, 0xf705d2b7e92b3f38e6ae7afadaa2fee110fe5914
		USDCList[ChainEthereum]:   "0x7B70D04099CB9cfb1Db7B6820baDAfB4C5C70A67", // variableDebtUSDC, 0x7b70d04099cb9cfb1db7b6820badafb4c5c70a67
		USDTList[ChainEthereum]:   "0x529b6158d1D2992E3129F7C69E81a7c677dc3B12", // variableDebtUSDT, 0x529b6158d1d2992e3129f7c69e81a7c677dc3b12
		RETHList[ChainEthereum]:   "0xBa2C8F2eA5B56690bFb8b709438F049e5Dd76B96", // variableDebtrETH, 0xba2c8f2ea5b56690bfb8b709438f049e5dd76b96
		SDAIList[ChainEthereum]:   "0xaBc57081C04D921388240393ec4088Aa47c6832B", // variableDebtsDAI, 0xabc57081c04d921388240393ec4088aa47c6832b
		WBTCList[ChainEthereum]:   "0xf6fEe3A8aC8040C3d6d81d9A4a168516Ec9B51D2", // variableDebtWBTC, 0xf6fee3a8ac8040c3d6d81d9a4a168516ec9b51d2
		WETHList[ChainEthereum]:   "0x2e7576042566f8D6990e07A1B61Ad1efd86Ae70d", // variableDebtWETH, 0x2e7576042566f8d6990e07a1b61ad1efd86ae70d
		WSTETHList[ChainEthereum]: "0xd5c3E3B566a42A6110513Ac7670C1a86D76E13E6", // variableDebtwstETH, 0xd5c3e3b566a42a6110513ac7670c1a86d76e13e6
		WEETHList[ChainEthereum]:  "0xc2bD6d2fEe70A0A73a33795BdbeE0368AeF5c766", // variableDebtweETH, 0xc2bd6d2fee70a0a73a33795bdbee0368aef5c766
		GNOList[ChainEthereum]:    "0x57a2957651DA467fCD4104D749f2F3684784c25a", // variableDebtGNO, 0x57a2957651da467fcd4104d749f2f3684784c25a
		CbBTCList[ChainEthereum]:  "0x661fE667D2103eb52d3632a3eB2cAbd123F27938", // variableDebtcbBTC, 0x661fe667d2103eb52d3632a3eb2cabd123f27938
		SUSDSList[ChainEthereum]:  "0x4e89b83f426fED3f2EF7Bb2d7eb5b53e288e1A13", // variableDebtsUSDS, 0x4e89b83f426fed3f2ef7bb2d7eb5b53e288e1a13
	},
}

// Spark s tokens.
var SparkSTokenList = map[int64]map[string]string{
	ChainEthereum: {
		DAIList[ChainEthereum]:    "0xfe2B7a7F4cC0Fb76f7Fc1C6518D586F1e4559176", // stableDebtDAI, 0xfe2b7a7f4cc0fb76f7fc1c6518d586f1e4559176
		USDCList[ChainEthereum]:   "0x887Ac022983Ff083AEb623923789052A955C6798", // stableDebtUSDC, 0x887ac022983ff083aeb623923789052a955c6798
		USDTList[ChainEthereum]:   "0x0Dae62F953Ceb2E969fB4dE85f3F9074fa920776", // stableDebtUSDT, 0x0dae62f953ceb2e969fb4de85f3f9074fa920776
		RETHList[ChainEthereum]:   "0xa9a4037295Ea3a168DC3F65fE69FdA524d52b3e1", // stableDebtrETH, 0xa9a4037295ea3a168dc3f65fe69fda524d52b3e1
		SDAIList[ChainEthereum]:   "0xEc6C6aBEd4DC03299EFf82Ac8A0A83643d3cB335", // stableDebtsDAI, 0xec6c6abed4dc03299eff82ac8a0a83643d3cb335
		WBTCList[ChainEthereum]:   "0x4b29e6cBeE62935CfC92efcB3839eD2c2F35C1d9", // stableDebtWBTC, 0x4b29e6cbee62935cfc92efcb3839ed2c2f35c1d9
		WETHList[ChainEthereum]:   "0x3c6b93D38ffA15ea995D1BC950d5D0Fa6b22bD05", // stableDebtWETH, 0x3c6b93d38ffa15ea995d1bc950d5d0fa6b22bd05
		WSTETHList[ChainEthereum]: "0x9832D969a0c8662D98fFf334A4ba7FeE62b109C2", // stableDebtwstETH, 0x9832d969a0c8662d98fff334a4ba7fee62b109c2
		WEETHList[ChainEthereum]:  "0x5B1F8aF3E6C0BF4d20e8e5220a4e4A3A8fA6Dc0A", // stableDebtweETH, 0x5b1f8af3e6c0bf4d20e8e5220a4e4a3a8fa6dc0a
		GNOList[ChainEthereum]:    "0xbf13910620722D4D4F8A03962894EB3335Bf4FaE", // stableDebtGNO, 0xbf13910620722d4d4f8a03962894eb3335bf4fae
		CbBTCList[ChainEthereum]:  "0x26a76E2fa1EaDbe7C30f0c333059Bcc3642c28d2", // stableDebtcbBTC, 0x26a76e2fa1eadbe7c30f0c333059bcc3642c28d2
		SUSDSList[ChainEthereum]:  "0x55580770e14E008082aB2E8d08a16Cc1dC192741", // stableDebtsUSDS, 0x55580770e14e008082ab2e8d08a16cc1dc192741
	},
}
