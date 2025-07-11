package ethaddr

// Blockscan: https://polygonscan.com/
const ChainNamePolygon string = "polygon"

// Wrapped polygon chain native token: WPOL (previously WMATIC).
//
// map[chainID] = address.
var WPOLList = map[int64]string{
	ChainEthereum: "0x7c9f4C87d911613Fe9ca58b579f737911AAD2D43", // WMATIC, 0x7c9f4c87d911613fe9ca58b579f737911aad2d43
	ChainPolygon:  "0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270", // WPOL, 0x0d500b1d8e8ef31e21c99d1db9a6444d3adf1270
}

var WMATICList = WPOLList

// MATIC.
//
// map[chainID] = address.
var MATICList = map[int64]string{
	ChainBSC: "0xCC42724C6683B7E57334c4E856f4c9965ED682bD", // MATIC, 0xcc42724c6683b7e57334c4e856f4c9965ed682bd
}
