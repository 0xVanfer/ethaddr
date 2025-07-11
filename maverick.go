package ethaddr

// Website: https://www.mav.xyz/
//
// Docs: https://docs.mav.xyz/
//
// Deployed contracts: https://docs.mav.xyz/further-information/contract-addresses
const MaverickProtocol string = "maverick"

// Maverick token: MAV.
//
// map[chainID] = address.
var MaverickTokenList = map[int64]string{
	ChainEthereum: "0x7448c7456a97769F6cD04F1E83A4a23cCdC46aBD", // 0x7448c7456a97769f6cd04f1e83a4a23ccdc46abd
	ChainBSC:      "0xd691d9a68C887BDF34DA8c36f63487333ACfD103", // 0xd691d9a68c887bdf34da8c36f63487333acfd103
}

// Maverick ve token: veMAV.
//
// map[chainID] = address.
var VeMAVList = map[int64]string{
	ChainEthereum: "0x4949Ac21d5b2A0cCd303C20425eeb29DCcba66D8", // 0x4949ac21d5b2a0ccd303c20425eeb29dccba66d8
	ChainBSC:      "0xE6108f1869d37E5076a56168C66A1607EdB10819", // 0xe6108f1869d37e5076a56168c66a1607edb10819
}

// Maverick token: MAV.
//
// map[chainID] = address.
var MAVList = MaverickTokenList

// map[chainID] = address.
var MaverickFactoryList = map[int64]string{
	ChainEthereum: "0xEb6625D65a0553c9dBc64449e56abFe519bd9c9B", // 0xeb6625d65a0553c9dbc64449e56abfe519bd9c9b
	ChainBSC:      "0x76311728FF86054Ad4Ac52D2E9Ca005BC702f589", // 0x76311728ff86054ad4ac52d2e9ca005bc702f589
}

// map[chainID] = address.
var MaverickPositionList = map[int64]string{
	ChainEthereum: "0x4A3e49f77a2A5b60682a2D6B8899C7c5211EB646", // 0x4a3e49f77a2a5b60682a2d6b8899c7c5211eb646
	ChainBSC:      "0x23Aeaf001E5DF9d7410EE6C6916f502b7aC8e9D0", // 0x23aeaf001e5df9d7410ee6c6916f502b7ac8e9d0
}

// map[chainID] = address.
var MaverickPositionInspectorList = map[int64]string{
	ChainEthereum: "0x456A37144162900799f405be34f815dE7C3DA53C", // 0x456a37144162900799f405be34f815de7c3da53c
	ChainBSC:      "0x70Cd6087033E0b99e4e449D3B904FaD194D888A0", // 0x70cd6087033e0b99e4e449d3b904fad194d888a0
}

// map[chainID] = address.
var MaverickPoolInformationList = map[int64]string{
	ChainEthereum: "0x9980ce3b5570e41324904f46A06cE7B466925E23", // 0x9980ce3b5570e41324904f46a06ce7b466925e23
	ChainBSC:      "0x9e10c5Ab6DCcD5F709D223EE60Ca912c2C2dbC56", // 0x9e10c5ab6dccd5f709d223ee60ca912c2c2dbc56
}

// map[chainID] = address.
var MaverickPoolPositionManagerList = map[int64]string{
	ChainEthereum: "0xE7583AF5121a8f583EFD82767CcCfEB71069D93A", // 0xe7583af5121a8f583efd82767cccfeb71069d93a
	ChainBSC:      "0x2D11545d36FfA0b8558e83C26e45cFaF14BDBAB2", // 0x2d11545d36ffa0b8558e83c26e45cfaf14bdbab2
}

// map[chainID] = address.
var MaverickRouterList = map[int64]string{
	ChainEthereum: "0xbBF1EE38152E9D8e3470Dc47947eAa65DcA94913", // 0xbbf1ee38152e9d8e3470dc47947eaa65dca94913
	ChainBSC:      "0xD53a9f3FAe2bd46D35E9a30bA58112A585542869", // 0xd53a9f3fae2bd46d35e9a30ba58112a585542869
}
