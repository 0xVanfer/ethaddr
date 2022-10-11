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
//
// map[network][pool type][underlying] = pool name.
var VectorPoolNameList = map[string]map[string]map[string]string{
	chainId.AvalancheChainName: {
		"VectorStaking": {
			TraderzJOEList[chainId.AvalancheChainName]:   "ZJOE",
			PlatypusxPTPList[chainId.AvalancheChainName]: "XPTP",
			VectorTokenList[chainId.AvalancheChainName]:  "VTX",
		},
		PlatypusMainPoolsName: {
			DAIeList[chainId.AvalancheChainName]:  "DAIe",
			USDCList[chainId.AvalancheChainName]:  "USDC",
			USDCeList[chainId.AvalancheChainName]: "USDCe",
			USDTList[chainId.AvalancheChainName]:  "USDT",
			USDTeList[chainId.AvalancheChainName]: "USDTe",
		},
		Platypus_USDC_FRAX_PairName: {
			USDCList[chainId.AvalancheChainName]: "USDC_FRAX",
			FRAXList[chainId.AvalancheChainName]: "FRAX",
		},
		Platypus_USDC_MIM_PairName: {
			USDCList[chainId.AvalancheChainName]: "USDC_MIM",
			MIMList[chainId.AvalancheChainName]:  "MIM",
		},
		Platypus_AVAX_sAVAX_PairName: {
			WAVAXList[chainId.AvalancheChainName]: "AVAX",
			SAVAXList[chainId.AvalancheChainName]: "SAVAX",
		},
		Platypus_USDC_TUSD_PairName: {
			USDCList[chainId.AvalancheChainName]: "USDC_TUSD",
			TUSDList[chainId.AvalancheChainName]: "TUSD",
		},
		Platypus_USDC_YUSD_PairName: {
			USDCList[chainId.AvalancheChainName]: "USDC_YUSD",
			YUSDList[chainId.AvalancheChainName]: "YUSD",
		},
	},
}

// Vector machine gun pools. 2022.06.13
//
// map[network][pool name] = address.
var VectorMachinegunPoolList = map[string]map[string]string{
	chainId.AvalancheChainName: {
		VectorPoolNameList[chainId.AvalancheChainName]["VectorStaking"][TraderzJOEList[chainId.AvalancheChainName]]:         VectorChefList[chainId.AvalancheChainName],
		VectorPoolNameList[chainId.AvalancheChainName]["VectorStaking"][PlatypusxPTPList[chainId.AvalancheChainName]]:       VectorChefList[chainId.AvalancheChainName],
		VectorPoolNameList[chainId.AvalancheChainName]["VectorStaking"][VectorTokenList[chainId.AvalancheChainName]]:        VectorChefList[chainId.AvalancheChainName],
		VectorPoolNameList[chainId.AvalancheChainName][PlatypusMainPoolsName][DAIeList[chainId.AvalancheChainName]]:         "0xc1ac7d1405b87259b8d380e0041d0573fb0acb8c",
		VectorPoolNameList[chainId.AvalancheChainName][PlatypusMainPoolsName][USDCList[chainId.AvalancheChainName]]:         "0x994f0e36ceb953105d05897537bf55d201245156",
		VectorPoolNameList[chainId.AvalancheChainName][PlatypusMainPoolsName][USDCeList[chainId.AvalancheChainName]]:        "0x257d69aa678e0a8da6dfda6a16cdf2052a460b45",
		VectorPoolNameList[chainId.AvalancheChainName][PlatypusMainPoolsName][USDTList[chainId.AvalancheChainName]]:         "0x5c6c7bc771f9a4231af8d5a463e6d495833011f0",
		VectorPoolNameList[chainId.AvalancheChainName][PlatypusMainPoolsName][USDTeList[chainId.AvalancheChainName]]:        "0x834eed8b99463a5d58b4d4b3a16b5904c37d7a2e",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_USDC_YUSD_PairName][USDCList[chainId.AvalancheChainName]]:   "0x7550b2d6a1f039dd6a3d54a857fefcbf77213d80",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_USDC_YUSD_PairName][YUSDList[chainId.AvalancheChainName]]:   "0xbd6a6a5bdb0ca1d5142519fb21b021109414bd1c",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_USDC_MIM_PairName][USDCList[chainId.AvalancheChainName]]:    "0x711a2aaf789ae342f3bf0ac290994f45408e694d",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_USDC_MIM_PairName][MIMList[chainId.AvalancheChainName]]:     "0x222ea06ff0a6d11fc1ef30b5d294e5e71e837caa",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_USDC_FRAX_PairName][USDCList[chainId.AvalancheChainName]]:   "0x1338b4065e25ad681c511644aa319181fc3d64cc",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_USDC_FRAX_PairName][FRAXList[chainId.AvalancheChainName]]:   "0xccce9e28fe4ceff031fd3b8837cb569cfb7a6843",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_AVAX_sAVAX_PairName][WAVAXList[chainId.AvalancheChainName]]: "0xff5386af93cf4bd8d5aecad6df7f4f4be381fd69",
		VectorPoolNameList[chainId.AvalancheChainName][Platypus_AVAX_sAVAX_PairName][SAVAXList[chainId.AvalancheChainName]]: "0x812b7c3b5a9164270dd8a0b3bc47550877aecdb1",
	},
}
