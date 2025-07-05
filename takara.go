package ethaddr

// Webiste: https://app.takaralend.com/
const TakaraProtocol string = "takara"

// Takara unitroller, proxy for comptroller.
//
// map[chainID] = address
var TakaraUnitrollerList = map[int64]string{
	ChainSei: "0x71034bf5eC0FAd7aEE81a213403c8892F3d8CAeE", // 0x71034bf5ec0fad7aee81a213403c8892f3d8caee
}

// Takara cTokens (tTokens)
//
// map[chainID][underlying] = ctoken address
var TakaraCTokenList = map[int64]map[string]string{
	ChainSei: {
		USDTList[ChainSei]:         "0xc68351B9B3638A6f4A3Ae100Bd251e227BbD7479", // tUSDT, 0xc68351b9b3638a6f4a3ae100bd251e227bbd7479
		USDCList[ChainSei]:         "0xC3c9e322F4aAe352ace79D0E62ADe3563fB86e87", // tUSDC, 0xc3c9e322f4aae352ace79d0e62ade3563fb86e87
		WSEIList[ChainSei]:         "0xA26b9BFe606d29F16B5Aecf30F9233934452c4E2", // tSEI, 0xa26b9bfe606d29f16b5aecf30f9233934452c4e2
		FastUSDList[ChainSei]:      "0x92e51466482146E71b692ced2265284968E8B3d6", // tfastUSD, 0x92e51466482146e71b692ced2265284968e8b3d6
		ISEIList[ChainSei]:         "0xda642A7821E91eD285262fead162E5fd17200429", // tiSEI, 0xda642a7821e91ed285262fead162e5fd17200429
		UBTCList_BSuared[ChainSei]: "0xabFb7A392a6DaaC50f99c5D14B5f27EFfd08Fe03", // tuBTC, 0xabfb7a392a6daac50f99c5d14b5f27efdf08fe03
		USDT0List[ChainSei]:        "0xA82a40324DBf7B57E87bD07C9e1D722E9754be9B", // tUSDT0, 0xa82a40324dbf7b57e87bd07c9e1d722e9754be9b
		WFBTCList[ChainSei]:        "0x51Be1bAde695f7aDb37ecf597EC2F97ebD99FbFE", // tFBTC, 0x51be1bade695f7adb37ecf597ec2f97ebd99fbfe
		SBTCList[ChainSei]:         "0x413498E8eD8608053DE3f407DBEDd29F19f80c04", // tSBTC, 0x413498e8ed8608053de3f407dbedd29f19f80c04
		SolvBTCList[ChainSei]:      "0xA54a39D8d2126C2aaE1622443B30F19414C74f3B", // tSolvBTC, 0xa54a39d8d2126c2aae1622443b30f19414c74f3b
		MerlinBTCList[ChainSei]:    "0x963Db326b734FD58a9396C020BBb52C14acaFb02", // tMBTC, 0x963db326b734fd58a9396c020bbb52c14acafb02
	},
}
