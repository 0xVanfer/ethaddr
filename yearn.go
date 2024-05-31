package ethaddr

const YearnProtocol string = "yearn"

// Yearn token: YFI.
//
// map[network] = address.
var YearnTokenList = map[int64]string{
	ChainEthereum: "0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e", // YFI, 0x0bc529c00c6401aef6d220be8c6ea1667f6ad93e
	ChainFantom:   "0x29b0Da86e484E1C0029B56e817912d778aC0EC69", // YFI, 0x29b0da86e484e1c0029b56e817912d778ac0ec69
	ChainPolygon:  "0xDA537104D6A5edd53c6fBba9A898708E465260b6", // YFI, 0xda537104d6a5edd53c6fbba9a898708e465260b6
	ChainArbitrum: "0x82e3A8F066a6989666b031d916c43672085b1582", // YFI, 0x82e3a8f066a6989666b031d916c43672085b1582
}

// Same as YearnTokenList.
var YFIList = YearnTokenList

// Yearn woofy token: WOOFY.
//
// map[network] = address.
var WoofyList = map[int64]string{
	ChainEthereum: "0xD0660cD418a64a1d44E9214ad8e459324D8157f1", // WOOFY, 0xd0660cd418a64a1d44e9214ad8e459324d8157f1
	ChainFantom:   "0xD0660cD418a64a1d44E9214ad8e459324D8157f1", // WOOFY, 0xd0660cd418a64a1d44e9214ad8e459324d8157f1
	ChainPolygon:  "0xD0660cD418a64a1d44E9214ad8e459324D8157f1", // WOOFY, 0xd0660cd418a64a1d44e9214ad8e459324d8157f1
}

// Yearn vault tokens.
//
// map[network][underlying] = address.
var YearnVaultTokenList = map[int64]map[string]string{
	ChainEthereum: {
		YearnTokenList[ChainEthereum]:                "0xE14d13d8B3b85aF791b2AADD661cDBd5E6097Db1", // yvYFI, 0xe14d13d8b3b85af791b2aadd661cdbd5e6097db1
		OINCHLIST[ChainEthereum]:                     "0xB8C3B7A2A618C552C23B1E4701109a9E756Bab67", // yv1INCH, 0xb8c3b7a2a618c552c23b1e4701109a9e756bab67
		WETHList[ChainEthereum]:                      "0xa258C4606Ca8206D8aA700cE2143D7db854D168c", // yvWETH, 0xa258c4606ca8206d8aa700ce2143d7db854d168c
		USDCList[ChainEthereum]:                      "0x5f18C75AbDAe578b483E5F43f12a39cF75b973a9", // yvUSDC, 0x5f18c75abdae578b483e5f43f12a39cf75b973a9
		DAIList[ChainEthereum]:                       "0xdA816459F1AB5631232FE5e97a05BBBb94970c95", // yvDAI, 0xda816459f1ab5631232fe5e97a05bbbb94970c95
		WBTCList[ChainEthereum]:                      "0xcB550A6D4C8e3517A939BC79d0c7093eb7cF56B5", // yvWBTC, 0xcb550a6d4c8e3517a939bc79d0c7093eb7cf56b5
		USDTList[ChainEthereum]:                      "0x7Da96a3891Add058AdA2E826306D812C638D87a7", // yvUSDT, 0x7da96a3891add058ada2e826306d812c638d87a7
		CurvePoolsList[ChainEthereum]["ib3CRV"]:      "0x27b7b1ad7288079A66d12350c828D3C00A6F07d7", // yvCurve-IronBank, 0x27b7b1ad7288079a66d12350c828d3c00a6f07d7
		CurvePoolsList[ChainEthereum]["eCRV"]:        "0xA3D87FffcE63B53E0d54fAa1cc983B7eB0b74A9c", // yvCurve-sETH, 0xa3d87fffce63b53e0d54faa1cc983b7eb0b74a9c
		CurvePoolsList[ChainEthereum]["steCRV"]:      "0xdCD90C7f6324cfa40d7169ef80b12031770B4325", // yvCurve-stETH, 0xdcd90c7f6324cfa40d7169ef80b12031770b4325
		CurvePoolsList[ChainEthereum]["crvRenWSBTC"]: "0x8414Db07a7F743dEbaFb402070AB01a4E0d2E45e", // yvCurve-sBTC, 0x8414db07a7f743debafb402070ab01a4e0d2e45e
		CurvePoolsList[ChainEthereum]["crvRenWBTC"]:  "0x7047F90229a057C13BF847C0744D646CFb6c9E1A", // yvCurve-renBTC, 0x7047f90229a057c13bf847c0744d646cfb6c9e1a
		CurvePoolsList[ChainEthereum]["oBTC"]:        "0xe9Dc63083c464d6EDcCFf23444fF3CFc6886f6FB", // yvCurve-oBTC, 0xe9dc63083c464d6edccff23444ff3cfc6886f6fb
		CurvePoolsList[ChainEthereum]["pBTC"]:        "0x3c5DF3077BcF800640B5DAE8c91106575a4826E6", // yvCurve-pBTC, 0x3c5df3077bcf800640b5dae8c91106575a4826e6
		CurvePoolsList[ChainEthereum]["tBTC"]:        "0x23D3D0f1c697247d5e0a9efB37d8b0ED0C464f7f", // yvCurve-tBTC, 0x23d3d0f1c697247d5e0a9efb37d8b0ed0c464f7f
		CurvePoolsList[ChainEthereum]["bBTC"]:        "0x8fA3A9ecd9EFb07A8CE90A6eb014CF3c0E3B32Ef", // yvCurve-bBTC, 0x8fa3a9ecd9efb07a8ce90a6eb014cf3c0e3b32ef
		CurvePoolsList[ChainEthereum]["FRAX3CRV-f"]:  "0xB4AdA607B9d6b2c9Ee07A275e9616B84AC560139", // yvCurve-FRAX, 0xb4ada607b9d6b2c9ee07a275e9616b84ac560139
		CurvePoolsList[ChainEthereum]["LUSD3CRV-f"]:  "0x5fA5B62c8AF877CB37031e0a3B2f34A78e3C56A6", // yvCurve-LUSD, 0x5fa5b62c8af877cb37031e0a3b2f34a78e3c56a6
		CurvePoolsList[ChainEthereum]["saCRV"]:       "0xb4D1Be44BfF40ad6e506edf43156577a3f8672eC", // yvCurve-sAave, 0xb4d1be44bff40ad6e506edf43156577a3f8672ec
		CurvePoolsList[ChainEthereum]["yveCRV-DAO"]:  "0x9d409a0A012CFbA9B15F6D4B36Ac57A46966Ab9a", // yvBOOST, 0x9d409a0a012cfba9b15f6d4b36ac57a46966ab9a
	},
}
