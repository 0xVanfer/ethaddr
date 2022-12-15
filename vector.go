package ethaddr

import "github.com/0xVanfer/chainId"

const VectorProtocol string = "vector"

// Vector token: VTX.
//
// map[network] = address.
var VectorTokenList = map[string]string{
	chainId.AvalancheChainName: "0x5817d4f0b62a59b17f75207da1848c2ce75e7af4", // VTX
}

// Same as VectorTokenList.
var VTXList = VectorTokenList

// Vector staking contract.
//
// map[network] = address.
var VectorChefList = map[string]string{
	chainId.AvalancheChainName: "0x423d0fe33031aa4456a17b150804aa57fc157d97",
}

// Vector pool names.
// Get pair names: https://api.vectorfinance.io/api/v1/vtx/tvl
//
// map[network][pool type][underlying] = pool name.
var VectorPoolNameList = map[string]map[string]map[string]string{
	chainId.AvalancheChainName: {
		// Staking.
		"VectorStaking": {
			TraderzJOEList[chainId.AvalancheChainName]:   "ZJOE",
			PlatypusxPTPList[chainId.AvalancheChainName]: "XPTP",
			VectorTokenList[chainId.AvalancheChainName]:  "VTX",
		},
		// Main pools.
		PlatypusMainPoolsName: {
			DAIeList[chainId.AvalancheChainName]:  "DAIe",
			USDCList[chainId.AvalancheChainName]:  "USDC",
			USDCeList[chainId.AvalancheChainName]: "USDCe",
			USDTList[chainId.AvalancheChainName]:  "USDT",
			USDTeList[chainId.AvalancheChainName]: "USDTe",
			BUSDList[chainId.AvalancheChainName]:  "BUSD",
		},
		// USDC-FRAX
		Platypus_USDC_FRAX_PairName: {
			USDCList[chainId.AvalancheChainName]: "USDC_FRAX",
			FRAXList[chainId.AvalancheChainName]: "FRAX",
		},
		// USDC-MIM
		Platypus_USDC_MIM_PairName: {
			USDCList[chainId.AvalancheChainName]: "USDC_MIM",
			MIMList[chainId.AvalancheChainName]:  "MIM",
		},
		// AVAX-sAVAX
		Platypus_AVAX_sAVAX_PairName: {
			WAVAXList[chainId.AvalancheChainName]: "AVAX",
			SAVAXList[chainId.AvalancheChainName]: "SAVAX",
		},
		// Deprecated.
		// USDC-TUSD
		Platypus_USDC_TUSD_PairName: {
			USDCList[chainId.AvalancheChainName]: "USDC_TUSD",
			TUSDList[chainId.AvalancheChainName]: "TUSD",
		},
		// USDC-YUSD
		Platypus_USDC_YUSD_PairName: {
			USDCList[chainId.AvalancheChainName]: "USDC_YUSD",
			YUSDList[chainId.AvalancheChainName]: "YUSD",
		},
		// USDC-MONEY
		Platypus_USDC_MONEY_PairName: {
			USDCList[chainId.AvalancheChainName]:  "USDC_MONEY",
			MoneyList[chainId.AvalancheChainName]: "MONEY",
		},
		// AVAX-yyAVAX
		Platypus_AVAX_yyAVAX_PairName: {
			WAVAXList[chainId.AvalancheChainName]:  "AVAX_yyAVAX",
			YyAVAXList[chainId.AvalancheChainName]: "yyAVAX",
		},
		// WBTC-BTC.b
		PlatypusLp_BTCb_WBTCe_BTCb_Address: {
			WBTCList[chainId.AvalancheChainName]: "WBTCe",
			BTCbList[chainId.AvalancheChainName]: "BTCb",
		},
	},
}

// Vector machine gun pools. 2022.12.15
// Vector helper v4.
//
// map[network][pool name] = address.
var VectorMachinegunPoolList = map[string]map[string]string{
	chainId.AvalancheChainName: {
		// Staking.
		VectorPoolNameList[chainId.AvalancheChainName]["VectorStaking"][TraderzJOEList[chainId.AvalancheChainName]]:   VectorChefList[chainId.AvalancheChainName],
		VectorPoolNameList[chainId.AvalancheChainName]["VectorStaking"][PlatypusxPTPList[chainId.AvalancheChainName]]: VectorChefList[chainId.AvalancheChainName],
		VectorPoolNameList[chainId.AvalancheChainName]["VectorStaking"][VectorTokenList[chainId.AvalancheChainName]]:  VectorChefList[chainId.AvalancheChainName],
		// Main pools.
		VectorPoolNameList[chainId.AvalancheChainName][PlatypusMainPoolsName][DAIeList[chainId.AvalancheChainName]]:  "0x34b6438a9ff4a43ff2c14be1f745ee03bf35bff7",
		VectorPoolNameList[chainId.AvalancheChainName][PlatypusMainPoolsName][USDCList[chainId.AvalancheChainName]]:  "0x7a4a145bb3126fd29fe820c7cafd6a6ff428621a",
		VectorPoolNameList[chainId.AvalancheChainName][PlatypusMainPoolsName][USDCeList[chainId.AvalancheChainName]]: "0xdac7e96d8a5d3e688fab61f6402a8a7ed5acfdfc",
		VectorPoolNameList[chainId.AvalancheChainName][PlatypusMainPoolsName][USDTList[chainId.AvalancheChainName]]:  "0xf096380ba5fee402feffacaadc888942f7ea757e",
		VectorPoolNameList[chainId.AvalancheChainName][PlatypusMainPoolsName][USDTeList[chainId.AvalancheChainName]]: "0xf0bfa2c591462971b0a414121a646052b8cd99bf",
		VectorPoolNameList[chainId.AvalancheChainName][PlatypusMainPoolsName][BUSDList[chainId.AvalancheChainName]]:  "0x28203152c306e8d78522b5d7317fd11c1d6c6cb0",
		// Alt pools.
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_USDC_YUSD_PairName][USDCList[chainId.AvalancheChainName]]:     "0xe53b42237bc2fc991202aae156ea133af1ee6870",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_USDC_YUSD_PairName][YUSDList[chainId.AvalancheChainName]]:     "0xa583d15a9ec1355766fdc9953f5e2e93f1b4ae85",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_USDC_MIM_PairName][USDCList[chainId.AvalancheChainName]]:      "0x1ad304a80677a85e27d04ce1abd94535f6971dad",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_USDC_MIM_PairName][MIMList[chainId.AvalancheChainName]]:       "0x3001028c0bbe740ee5986de0e3cd7664b0636926",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_USDC_FRAX_PairName][USDCList[chainId.AvalancheChainName]]:     "0x88cbebf7f77c3cc9a964bbb22b7c4ee0c4e967e8",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_USDC_FRAX_PairName][FRAXList[chainId.AvalancheChainName]]:     "0x3c0d54752560347db618f452f13a843cdd130c16",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_USDC_MONEY_PairName][USDCList[chainId.AvalancheChainName]]:    "0xe9ea4945a929bfaeab6824ed841a983e0d8d9513",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_USDC_MONEY_PairName][MoneyList[chainId.AvalancheChainName]]:   "0x7dc0a2854b659a15358659ceae9489e583ad6cbf",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_AVAX_sAVAX_PairName][WAVAXList[chainId.AvalancheChainName]]:   "0x4e42d1a0b83fa354882f19e89a316e00bc106a98",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_AVAX_sAVAX_PairName][SAVAXList[chainId.AvalancheChainName]]:   "0x822c11be60258d6bf00c5b0907b2015633d11a62",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_AVAX_yyAVAX_PairName][WAVAXList[chainId.AvalancheChainName]]:  "0xf3876553724471ea85e58c0b0d1ec7c4b4147bdf",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_AVAX_yyAVAX_PairName][YyAVAXList[chainId.AvalancheChainName]]: "0x1ad304a80677a85e27d04ce1abd94535f6971dad",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_BTCb_WBTCe_PairName][BTCbList[chainId.AvalancheChainName]]:    "0x5c3666fa2a2950402a3ff53b69c1e009448fbd41",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_BTCb_WBTCe_PairName][WBTCList[chainId.AvalancheChainName]]:    "0xf482b0d19c9be9a562ef4701d621ce396cab4c8e",
	},
}
