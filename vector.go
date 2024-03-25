package ethaddr

// Website: https://vectorfinance.io/
const VectorProtocol string = "vector"

// Vector token: VTX.
//
// map[network] = address.
var VectorTokenList = map[string]string{
	ChainAvalanche: "0x5817D4F0b62A59b17f75207DA1848C2cE75e7AF4", // VTX, 0x5817d4f0b62a59b17f75207da1848c2ce75e7af4
}

// Same as VectorTokenList.
var VTXList = VectorTokenList

// Vector staking contract.
//
// map[network] = address.
var VectorChefList = map[string]string{
	ChainAvalanche: "0x423D0FE33031aA4456a17b150804aA57fc157d97", // 0x423d0fe33031aa4456a17b150804aa57fc157d97
}

// Vector pool names.
// Get pair names: https://api.vectorfinance.io/api/v1/vtx/tvl
//
// map[network][pool type][underlying] = pool name.
var VectorPoolNameList = map[string]map[string]map[string]string{
	ChainAvalanche: {
		// Staking.
		"VectorStaking": {
			TraderzJOEList[ChainAvalanche]:   "ZJOE",
			PlatypusxPTPList[ChainAvalanche]: "XPTP",
			VectorTokenList[ChainAvalanche]:  "VTX",
		},
		// Deprecated: Platypus' dead.
		// Main pools.
		PlatypusMainPoolsName: {
			DAIeList[ChainAvalanche]:  "DAIe",
			USDCList[ChainAvalanche]:  "USDC",
			USDCeList[ChainAvalanche]: "USDCe",
			USDTList[ChainAvalanche]:  "USDT",
			USDTeList[ChainAvalanche]: "USDTe",
			BUSDList[ChainAvalanche]:  "BUSD",
		},
		// Deprecated: Platypus' dead.
		// USDC-FRAX
		Platypus_USDC_FRAX_PairName: {
			USDCList[ChainAvalanche]: "USDC_FRAX",
			FRAXList[ChainAvalanche]: "FRAX",
		},
		// Deprecated: Platypus' dead.
		// USDC-MIM
		Platypus_USDC_MIM_PairName: {
			USDCList[ChainAvalanche]: "USDC_MIM",
			MIMList[ChainAvalanche]:  "MIM",
		},
		// Deprecated: Platypus' dead.
		// AVAX-sAVAX
		Platypus_AVAX_sAVAX_PairName: {
			WAVAXList[ChainAvalanche]: "AVAX",
			SAVAXList[ChainAvalanche]: "SAVAX",
		},
		// Deprecated: Platypus' dead.
		// USDC-TUSD
		Platypus_USDC_TUSD_PairName: {
			USDCList[ChainAvalanche]: "USDC_TUSD",
			TUSDList[ChainAvalanche]: "TUSD",
		},
		// Deprecated: Platypus' dead.
		// USDC-YUSD
		Platypus_USDC_YUSD_PairName: {
			USDCList[ChainAvalanche]: "USDC_YUSD",
			YUSDList[ChainAvalanche]: "YUSD",
		},
		// Deprecated: Platypus' dead.
		// USDC-MONEY
		Platypus_USDC_MONEY_PairName: {
			USDCList[ChainAvalanche]:  "USDC_MONEY",
			MoneyList[ChainAvalanche]: "MONEY",
		},
		// Deprecated: Platypus' dead.
		// AVAX-yyAVAX
		Platypus_AVAX_yyAVAX_PairName: {
			WAVAXList[ChainAvalanche]:  "AVAX_yyAVAX",
			YyAVAXList[ChainAvalanche]: "yyAVAX",
		},
		// Deprecated: Platypus' dead.
		// WBTC-BTC.b
		PlatypusLp_BTCb_WBTCe_BTCb_Address: {
			WBTCList[ChainAvalanche]: "WBTCe",
			BTCbList[ChainAvalanche]: "BTCb",
		},
	},
}

// Vector machine gun pools. 2022.12.15
// Vector helper v4.
//
// map[network][pool name] = address.
var VectorMachinegunPoolList = map[string]map[string]string{
	ChainAvalanche: {
		// Staking.
		VectorPoolNameList[ChainAvalanche]["VectorStaking"][TraderzJOEList[ChainAvalanche]]:   VectorChefList[ChainAvalanche],
		VectorPoolNameList[ChainAvalanche]["VectorStaking"][PlatypusxPTPList[ChainAvalanche]]: VectorChefList[ChainAvalanche],
		VectorPoolNameList[ChainAvalanche]["VectorStaking"][VectorTokenList[ChainAvalanche]]:  VectorChefList[ChainAvalanche],
		// Main pools.
		VectorPoolNameList[ChainAvalanche][PlatypusMainPoolsName][DAIeList[ChainAvalanche]]:  "0x34B6438a9fF4a43FF2c14bE1F745ee03bF35BfF7", // 0x34b6438a9ff4a43ff2c14be1f745ee03bf35bff7
		VectorPoolNameList[ChainAvalanche][PlatypusMainPoolsName][USDCList[ChainAvalanche]]:  "0x7a4a145bb3126fd29fe820c7cafd6a6Ff428621A", // 0x7a4a145bb3126fd29fe820c7cafd6a6ff428621a
		VectorPoolNameList[ChainAvalanche][PlatypusMainPoolsName][USDCeList[ChainAvalanche]]: "0xdac7E96D8a5D3E688faB61F6402a8a7ed5AcFDfC", // 0xdac7e96d8a5d3e688fab61f6402a8a7ed5acfdfc
		VectorPoolNameList[ChainAvalanche][PlatypusMainPoolsName][USDTList[ChainAvalanche]]:  "0xF096380bA5fEE402FEfFaCAadc888942F7Ea757E", // 0xf096380ba5fee402feffacaadc888942f7ea757e
		VectorPoolNameList[ChainAvalanche][PlatypusMainPoolsName][USDTeList[ChainAvalanche]]: "0xf0bfa2c591462971b0a414121A646052b8Cd99Bf", // 0xf0bfa2c591462971b0a414121a646052b8cd99bf
		VectorPoolNameList[ChainAvalanche][PlatypusMainPoolsName][BUSDList[ChainAvalanche]]:  "0x28203152C306e8d78522B5d7317Fd11C1d6C6cb0", // 0x28203152c306e8d78522b5d7317fd11c1d6c6cb0
		// Alt pools.
		VectorPoolNameList[ChainAvalanche][Platypus_USDC_YUSD_PairName][USDCList[ChainAvalanche]]:     "0xE53b42237Bc2fC991202aaE156eA133aF1ee6870", // 0xe53b42237bc2fc991202aae156ea133af1ee6870
		VectorPoolNameList[ChainAvalanche][Platypus_USDC_YUSD_PairName][YUSDList[ChainAvalanche]]:     "0xa583d15A9eC1355766fdC9953F5e2E93f1B4Ae85", // 0xa583d15a9ec1355766fdc9953f5e2e93f1b4ae85
		VectorPoolNameList[ChainAvalanche][Platypus_USDC_MIM_PairName][USDCList[ChainAvalanche]]:      "0x1Ad304A80677A85E27d04ce1ABD94535F6971DaD", // 0x1ad304a80677a85e27d04ce1abd94535f6971dad
		VectorPoolNameList[ChainAvalanche][Platypus_USDC_MIM_PairName][MIMList[ChainAvalanche]]:       "0x3001028C0bBe740Ee5986dE0E3Cd7664b0636926", // 0x3001028c0bbe740ee5986de0e3cd7664b0636926
		VectorPoolNameList[ChainAvalanche][Platypus_USDC_FRAX_PairName][USDCList[ChainAvalanche]]:     "0x88cbEbf7F77c3cc9a964BbB22B7C4ee0C4E967e8", // 0x88cbebf7f77c3cc9a964bbb22b7c4ee0c4e967e8
		VectorPoolNameList[ChainAvalanche][Platypus_USDC_FRAX_PairName][FRAXList[ChainAvalanche]]:     "0x3C0D54752560347Db618F452f13a843cdD130C16", // 0x3c0d54752560347db618f452f13a843cdd130c16
		VectorPoolNameList[ChainAvalanche][Platypus_USDC_MONEY_PairName][USDCList[ChainAvalanche]]:    "0xe9Ea4945a929BfaEab6824ED841a983E0D8d9513", // 0xe9ea4945a929bfaeab6824ed841a983e0d8d9513
		VectorPoolNameList[ChainAvalanche][Platypus_USDC_MONEY_PairName][MoneyList[ChainAvalanche]]:   "0x7Dc0A2854B659a15358659ceae9489E583ad6cbF", // 0x7dc0a2854b659a15358659ceae9489e583ad6cbf
		VectorPoolNameList[ChainAvalanche][Platypus_AVAX_sAVAX_PairName][WAVAXList[ChainAvalanche]]:   "0xab42ed09F43DDa849aa7F62500885A973A38a8Bc", // 0xab42ed09F43DDa849aa7F62500885A973A38a8Bc
		VectorPoolNameList[ChainAvalanche][Platypus_AVAX_sAVAX_PairName][SAVAXList[ChainAvalanche]]:   "0x822C11be60258D6Bf00C5B0907B2015633d11a62", // 0x822c11be60258d6bf00c5b0907b2015633d11a62
		VectorPoolNameList[ChainAvalanche][Platypus_AVAX_yyAVAX_PairName][WAVAXList[ChainAvalanche]]:  "0xAF224FD9cF65e869E150013529694be18bE7b418", // 0xaf224fd9cf65e869e150013529694be18be7b418
		VectorPoolNameList[ChainAvalanche][Platypus_AVAX_yyAVAX_PairName][YyAVAXList[ChainAvalanche]]: "0xf3876553724471EA85E58C0B0D1ec7c4b4147BdF", // 0xf3876553724471ea85e58c0b0d1ec7c4b4147bdf
		VectorPoolNameList[ChainAvalanche][Platypus_BTCb_WBTCe_PairName][BTCbList[ChainAvalanche]]:    "0x5c3666FA2A2950402A3ff53B69C1E009448FBd41", // 0x5c3666fa2a2950402a3ff53b69c1e009448fbd41
		VectorPoolNameList[ChainAvalanche][Platypus_BTCb_WBTCe_PairName][WBTCList[ChainAvalanche]]:    "0xF482b0d19c9bE9a562EF4701D621ce396CAb4C8e", // 0xf482b0d19c9be9a562ef4701d621ce396cab4c8e
	},
}

// Vector helper v4 for AVAX.
// Does not fit the exisiting structure.
// TODO: Find a better solution.
var VectorHelperV4AVAX string = "0xab42ed09F43DDa849aa7F62500885A973A38a8Bc" // 0xab42ed09f43dda849aa7f62500885a973a38a8bc
