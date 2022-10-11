package ethaddr

import "github.com/0xVanfer/chainId"

var YearnProtocol string = "yearn"

// Yearn token: YFI.
//
// map[network] = address.
var YearnTokenList = map[string]string{
	chainId.EthereumChainName: "0x0bc529c00c6401aef6d220be8c6ea1667f6ad93e", // YFI
	chainId.FantomChainName:   "0x29b0da86e484e1c0029b56e817912d778ac0ec69", // YFI
	chainId.PolygonChainName:  "0xda537104d6a5edd53c6fbba9a898708e465260b6", // YFI
}

// Same as YearnTokenList.
var YFIList = YearnTokenList

// Yearn woofy token: WOOFY.
//
// map[network] = address.
var WoofyList = map[string]string{
	chainId.EthereumChainName: "0xd0660cd418a64a1d44e9214ad8e459324d8157f1", // WOOFY
	chainId.FantomChainName:   "0xd0660cd418a64a1d44e9214ad8e459324d8157f1", // WOOFY
	chainId.PolygonChainName:  "0xd0660cd418a64a1d44e9214ad8e459324d8157f1", // WOOFY
}

// Yearn vault tokens.
//
// map[network][underlying] = address.
var YearnVaultTokenList = map[string]map[string]string{
	chainId.EthereumChainName: {
		YearnTokenList[chainId.EthereumChainName]:                "0xe14d13d8b3b85af791b2aadd661cdbd5e6097db1", // yvYFI
		OneInchTokenList[chainId.EthereumChainName]:              "0xb8c3b7a2a618c552c23b1e4701109a9e756bab67", // yv1INCH
		WETHList[chainId.EthereumChainName]:                      "0xa258c4606ca8206d8aa700ce2143d7db854d168c", // yvWETH
		USDCList[chainId.EthereumChainName]:                      "0x5f18c75abdae578b483e5f43f12a39cf75b973a9", // yvUSDC
		DAIList[chainId.EthereumChainName]:                       "0xda816459f1ab5631232fe5e97a05bbbb94970c95", // yvDAI
		WBTCList[chainId.EthereumChainName]:                      "0xcb550a6d4c8e3517a939bc79d0c7093eb7cf56b5", // yvWBTC
		USDTList[chainId.EthereumChainName]:                      "0x7da96a3891add058ada2e826306d812c638d87a7", // yvUSDT
		CurvePoolsList[chainId.EthereumChainName]["ib3CRV"]:      "0x27b7b1ad7288079a66d12350c828d3c00a6f07d7", // yvCurve-IronBank
		CurvePoolsList[chainId.EthereumChainName]["eCRV"]:        "0xa3d87fffce63b53e0d54faa1cc983b7eb0b74a9c", // yvCurve-sETH
		CurvePoolsList[chainId.EthereumChainName]["steCRV"]:      "0xdcd90c7f6324cfa40d7169ef80b12031770b4325", // yvCurve-stETH
		CurvePoolsList[chainId.EthereumChainName]["crvRenWSBTC"]: "0x8414db07a7f743debafb402070ab01a4e0d2e45e", // yvCurve-sBTC
		CurvePoolsList[chainId.EthereumChainName]["crvRenWBTC"]:  "0x7047f90229a057c13bf847c0744d646cfb6c9e1a", // yvCurve-renBTC
		CurvePoolsList[chainId.EthereumChainName]["oBTC"]:        "0xe9dc63083c464d6edccff23444ff3cfc6886f6fb", // yvCurve-oBTC
		CurvePoolsList[chainId.EthereumChainName]["pBTC"]:        "0x3c5df3077bcf800640b5dae8c91106575a4826e6", // yvCurve-pBTC
		CurvePoolsList[chainId.EthereumChainName]["tBTC"]:        "0x23d3d0f1c697247d5e0a9efb37d8b0ed0c464f7f", // yvCurve-tBTC
		CurvePoolsList[chainId.EthereumChainName]["bBTC"]:        "0x8fa3a9ecd9efb07a8ce90a6eb014cf3c0e3b32ef", // yvCurve-bBTC
		CurvePoolsList[chainId.EthereumChainName]["FRAX3CRV-f"]:  "0xb4ada607b9d6b2c9ee07a275e9616b84ac560139", // yvCurve-FRAX
		CurvePoolsList[chainId.EthereumChainName]["LUSD3CRV-f"]:  "0x5fa5b62c8af877cb37031e0a3b2f34a78e3c56a6", // yvCurve-LUSD
		CurvePoolsList[chainId.EthereumChainName]["saCRV"]:       "0xb4d1be44bff40ad6e506edf43156577a3f8672ec", // yvCurve-sAave
		CurvePoolsList[chainId.EthereumChainName]["yveCRV-DAO"]:  "0x9d409a0a012cfba9b15f6d4b36ac57a46966ab9a", // yvBOOST
	},
}
