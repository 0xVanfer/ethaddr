package ethaddr

// Website: https://app.symbiotic.fi/restake
var SymbioticProtocol string = "symbiotic"

// The vaults for symbiotic to store user deposits before symbiotic goes online.
// Last updated: Jul.2.2024
//
// map[chainID][underlying] = address
var SymbioticDefaultCollaterals = map[int64]map[string]string{
	ChainEthereum: {
		WSTETHList[ChainEthereum]:  "0xC329400492c6ff2438472D4651Ad17389fCb843a", // DC_wstETH, 0xc329400492c6ff2438472d4651ad17389fcb843a
		SUSDeList[ChainEthereum]:   "0x19d0D8e6294B7a04a2733FE433444704B791939A", // DC_sUSDe, 0x19d0d8e6294b7a04a2733fe433444704b791939a
		CbETHList[ChainEthereum]:   "0xB26ff591F44b04E78de18f43B46f8b70C6676984", // DC_cbETH, 0xb26ff591f44b04e78de18f43b46f8b70c6676984
		ENAList[ChainEthereum]:     "0xe39B5f5638a209c1A6b6cDFfE5d37F7Ac99fCC84", // DC_ENA, 0xe39b5f5638a209c1a6b6cdffe5d37f7ac99fcc84
		WBETHList[ChainEthereum]:   "0x422F5acCC812C396600010f224b320a743695f85", // DC_wBETH, 0x422f5accc812c396600010f224b320a743695f85
		RETHList[ChainEthereum]:    "0x03Bf48b8A1B37FBeAd1EcAbcF15B98B924ffA5AC", // DC_rETH, 0x03bf48b8a1b37fbead1ecabcf15b98b924ffa5ac
		METHList[ChainEthereum]:    "0x475D3Eb031d250070B63Fa145F0fCFC5D97c304a", // DC_mETH, 0x475d3eb031d250070b63fa145f0fcfc5d97c304a
		SWETHList[ChainEthereum]:   "0x38B86004842D3FA4596f0b7A0b53DE90745Ab654", // DC_swETH, 0x38b86004842d3fa4596f0b7a0b53de90745ab654
		SfrxETHList[ChainEthereum]: "0x5198CB44D7B2E993ebDDa9cAd3b9a0eAa32769D2", // DC_sfrxETH, 0x5198cb44d7b2e993ebdda9cad3b9a0eaa32769d2
		ETHxList[ChainEthereum]:    "0xBdea8e677F9f7C294A4556005c640Ee505bE6925", // DC_ETHx, 0xbdea8e677f9f7c294a4556005c640ee505be6925
	},
}
