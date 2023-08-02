package ethaddr

import "github.com/0xVanfer/chainId"

const MaverickProtocol string = "maverick"

// Maverick token: MAV.
//
// map[network] = address.
var MaverickTokenList = map[string]string{
	chainId.EthereumChainName: "0x7448c7456a97769f6cd04f1e83a4a23ccdc46abd",
	chainId.BNBSmartChainName: "0xd691d9a68c887bdf34da8c36f63487333acfd103",
}

// Maverick ve token: veMAV.
//
// map[network] = address.
var VeMAVList = map[string]string{
	chainId.EthereumChainName: "0x4949ac21d5b2a0ccd303c20425eeb29dccba66d8",
	chainId.BNBSmartChainName: "0xe6108f1869d37e5076a56168c66a1607edb10819",
}

// Maverick token: MAV.
//
// map[network] = address.
var MAVList = MaverickTokenList

// map[network] = address.
var MaverickFactoryList = map[string]string{
	chainId.EthereumChainName: "0xeb6625d65a0553c9dbc64449e56abfe519bd9c9b",
	chainId.BNBSmartChainName: "0x76311728ff86054ad4ac52d2e9ca005bc702f589",
}

// map[network] = address.
var MaverickPositionList = map[string]string{
	chainId.EthereumChainName: "0x4a3e49f77a2a5b60682a2d6b8899c7c5211eb646",
	chainId.BNBSmartChainName: "0x23aeaf001e5df9d7410ee6c6916f502b7ac8e9d0",
}

// map[network] = address.
var MaverickPositionInspectorList = map[string]string{
	chainId.EthereumChainName: "0x456a37144162900799f405be34f815de7c3da53c",
	chainId.BNBSmartChainName: "0x70cd6087033e0b99e4e449d3b904fad194d888a0",
}

// map[network] = address.
var MaverickPoolInformationList = map[string]string{
	chainId.EthereumChainName: "0x9980ce3b5570e41324904f46a06ce7b466925e23",
	chainId.BNBSmartChainName: "0x9e10c5ab6dccd5f709d223ee60ca912c2c2dbc56",
}

// map[network] = address.
var MaverickPoolPositionManagerList = map[string]string{
	chainId.EthereumChainName: "0xe7583af5121a8f583efd82767cccfeb71069d93a",
	chainId.BNBSmartChainName: "0x2d11545d36ffa0b8558e83c26e45cfaf14bdbab2",
}

// map[network] = address.
var MaverickRouterList = map[string]string{
	chainId.EthereumChainName: "0xbbf1ee38152e9d8e3470dc47947eaa65dca94913",
	chainId.BNBSmartChainName: "0xd53a9f3fae2bd46d35e9a30ba58112a585542869",
}
