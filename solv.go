package ethaddr

// Website: https://solv.finance/
const SolvProtocol string = "solv"

// Solv protocol token SOLV.
//
// map[chainID] = address
var SolvList = map[int64]string{
	ChainEthereum: "0x256F2d67e52fE834726D2DDCD8413654F5Eb8b53", // SOLV, 0x256f2d67e52fe834726d2ddcd8413654f5eb8b53
}

// ==================== SolvBTC ====================

// Solv protocol SolvBTC.
//
// map[chainID] = address
var SolvBTCList = map[int64]string{
	ChainArbitrum:  "0x3647c54c4c2C65bC7a2D63c0Da2809B399DBBDC0", // SolvBTC, 0x3647c54c4c2c65bc7a2d63c0da2809b399dbbdc0
	ChainAvalanche: "0xbc78D84Ba0c46dFe32cf2895a19939c86b81a777", // SolvBTC, 0xbc78d84ba0c46dfe32cf2895a19939c86b81a777
	ChainBase:      "0x3B86Ad95859b6AB773f55f8d94B4b9d443EE931f", // SolvBTC, 0x3b86ad95859b6ab773f55f8d94b4b9d443ee931f
	ChainBSC:       "0x4aae823a6a0b376De6A78e74eCC5b079d38cBCf7", // SolvBTC, 0x4aae823a6a0b376de6a78e74ecc5b079d38cbcf7
	ChainEthereum:  "0x7A56E1C57C7475CCf742a1832B028F0456652F97", // SolvBTC, 0x7a56e1c57c7475ccf742a1832b028f0456652f97
	ChainLinea:     "0x541FD749419CA806a8bc7da8ac23D346f2dF8B77", // SolvBTC, 0x541fd749419ca806a8bc7da8ac23d346f2df8b77
	ChainMantle:    "0xa68d25fC2AF7278db4BcdcAabce31814252642a9", // SolvBTC, 0xa68d25fc2af7278db4bcdcaabce31814252642a9
	ChainMode:      "0x541FD749419CA806a8bc7da8ac23D346f2dF8B77", // SolvBTC, 0x541fd749419ca806a8bc7da8ac23d346f2df8b77
	ChainSei:       "0x541FD749419CA806a8bc7da8ac23D346f2dF8B77", // SolvBTC, 0x541fd749419ca806a8bc7da8ac23d346f2df8b77
	ChainSonic:     "0x541FD749419CA806a8bc7da8ac23D346f2dF8B77", // SolvBTC, 0x541fd749419ca806a8bc7da8ac23d346f2df8b77
	ChainTaiko:     "0x541FD749419CA806a8bc7da8ac23D346f2dF8B77", // SolvBTC, 0x541fd749419ca806a8bc7da8ac23d346f2df8b77
}

// Solv protocol SolvBTC.b (bridged by `Free`).
//
// map[chainID] = address
var SolvBTCbList = map[int64]string{
	ChainScroll: "0x3Ba89d490AB1C0c9CC2313385b30710e838370a4", // SolvBTC.b, 0x3ba89d490ab1c0c9cc2313385b30710e838370a4
}

// ==================== SolvBTC LSTs ====================

// Solv protocol xSolvBTC (Previously SolvBTC.BBN)
//
// map[chainID] = address
var XSolvBTCList = map[int64]string{
	ChainArbitrum:  "0x346c574C56e1A4aAa8dc88Cda8F7EB12b39947aB", // xSolvBTC, 0x346c574c56e1a4aaa8dc88cda8f7eb12b39947ab
	ChainAvalanche: "0xCC0966D8418d412c599A6421b760a847eB169A8c", // xSolvBTC, 0xcc0966d8418d412c599a6421b760a847eb169a8c
	ChainBase:      "0xC26C9099BD3789107888c35bb41178079B282561", // xSolvBTC, 0xc26c9099bd3789107888c35bb41178079b282561
	ChainBSC:       "0x1346b618dC92810EC74163e4c27004c921D446a5", // xSolvBTC, 0x1346b618dc92810ec74163e4c27004c921d446a5
	ChainEthereum:  "0xd9D920AA40f578ab794426F5C90F6C731D159DEf", // xSolvBTC, 0xd9d920aa40f578ab794426f5c90f6c731d159def
	ChainLinea:     "0xCC0966D8418d412c599A6421b760a847eB169A8c", // xSolvBTC, 0xcc0966d8418d412c599a6421b760a847eb169a8c
	ChainMantle:    "0x1d40baFC49c37CdA49F2a5427E2FB95E1e3FCf20", // xSolvBTC, 0x1d40bafc49c37cda49f2a5427e2fb95e1e3fcf20
	ChainSonic:     "0xCC0966D8418d412c599A6421b760a847eB169A8c", // xSolvBTC, 0xcc0966d8418d412c599a6421b760a847eb169a8c
	ChainTaiko:     "0xCC0966D8418d412c599A6421b760a847eB169A8c", // xSolvBTC, 0xcc0966d8418d412c599a6421b760a847eb169a8c
}

// Solv protocol xSolvBTC (Previously SolvBTC.BBN)
//
// map[chainID] = address
var SolvBTCBBNList = XSolvBTCList

// Solv protocol SolvBTC.CORE
//
// map[chainID] = address
var SolvBTCCOREList = map[int64]string{
	ChainBSC: "0xb9f59cAB0d6AA9D711acE5c3640003Bc09C15Faf", // SolvBTC.CORE, 0xb9f59cab0d6aa9d711ace5c3640003bc09c15faf
}

// Solv protocol SolvBTC.JUP
//
// map[chainID] = address
var SolvBTCJUPList = map[int64]string{
	ChainBSC: "0x38a001e57430f781404ffF7a81DE4Bd67d1f6117", // SolvBTC.JUP, 0x38a001e57430f781404fff7a81de4bd67d1f6117
}

// Solv protocol SolvBTC.ENA
//
// map[chainID] = address
var SolvBTCENAList = map[int64]string{
	ChainArbitrum: "0xaFAfd68AFe3fe65d376eEC9Eab1802616cFacCb8", // SolvBTC.ENA, 0xafafd68afe3fe65d376eec9eab1802616cfaccb8
	ChainBSC:      "0x53E63a31fD1077f949204b94F431bCaB98F72BCE", // SolvBTC.ENA, 0x53e63a31fd1077f949204b94f431bcab98f72bce
	ChainEthereum: "0x325DC9EBceC31940C658aCACa45f8293418d811E", // SolvBTC.ENA, 0x325dc9ebcec31940c658acaca45f8293418d811e
}

// ==================== Routers ====================

// Router for SolvBTC.
//
// map[chainID] = address
var SolvBTCRouterList = map[int64]string{
	ChainArbitrum:  "0xe9eD7530427Cb41A56C9e004e00e074cCc168C44", // 0xe9ed7530427cb41a56c9e004e00e074ccc168c44
	ChainAvalanche: "0x5b60F7e24Ac48C1146d1aedb6a72B62c83378730", // 0x5b60f7e24ac48c1146d1aedb6a72b62c83378730
	ChainBase:      "0x65EFfDA5e69dF470d4dBd31a805e15855Cae65c7", // 0x65effda5e69df470d4dbd31a805e15855cae65c7
	ChainBSC:       "0x5c1215712F174dF2Cbc653eDce8B53FA4CAF2201", // 0x5c1215712f174df2cbc653edce8b53fa4caf2201
	ChainEthereum:  "0x1fF7d7C0A7D8E94046708C611DeC5056A9d2B823", // 0x1ff7d7c0a7d8e94046708c611dec5056a9d2b823
	ChainLinea:     "0x516dcd9F731c61591c849D7FAAf49AE146C58D12", // 0x516dcd9f731c61591c849d7faaf49ae146c58d12
	ChainMantle:    "0x900637B3258E6B86FE2E713fBcA4510aD934eE7e", // 0x900637b3258e6b86fe2e713fbca4510ad934ee7e
}

// Router for SolvBTC LSTs.
//
// map[chainID] = address
var SolvBTCLSTRouterList = map[int64]string{
	ChainArbitrum:  "0x836bE4332347440995Bc06103aA740AdaFC0068a", // 0x836be4332347440995bc06103aa740adafc0068a
	ChainAvalanche: "0x2181e83ba95df14fa339f0f854A706E7D289F78a", // 0x2181e83ba95df14fa339f0f854a706e7d289f78a
	ChainBase:      "0x814F3ae67dF0da9fe2399a29516FD14b9085263a", // 0x814f3ae67df0da9fe2399a29516fd14b9085263a
	ChainBSC:       "0x8EC6Ef69a423045cEa97d2Bd0D768D042A130aA7", // 0x8ec6ef69a423045cea97d2bd0d768d042a130aa7
	ChainEthereum:  "0x01024AaeD5561fa6237C0ad4073417576C591261", // 0x01024aaed5561fa6237c0ad4073417576c591261
	ChainLinea:     "0xa0db6ab82ea2F44dd15ECDB228811656B446181b", // 0xa0db6ab82ea2f44dd15ecdb228811656b446181b
	ChainMantle:    "0xa6F5cf9259E4ba72b195b5F3caBe2577ce205dF2", // 0xa6f5cf9259e4ba72b195b5f3cabe2577ce205df2
}
