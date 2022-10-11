package ethaddr

import "github.com/0xVanfer/chainId"

const PangolinProtocol string = "pangolin"

// Pangolin token: PNG.
//
// map[network] = address.
var PangolinTokenList = map[string]string{
	chainId.AvalancheChainName: "0x60781c2586d68229fde47564546784ab3faca982", // PNG
}

// Same as PangolinTokenList.
var PNGList = PangolinTokenList

// Pangolin chef v2.
//
// map[network] = address.
var PangolinChefV2List = map[string]string{
	chainId.AvalancheChainName: "0x1f806f7c8ded893fd3cae279191ad7aa3798e928",
}

// Pangolin factory.
//
// map[network] = address.
var PangolinFactoryList = map[string]string{
	chainId.AvalancheChainName: "0xefa94de7a4656d787667c749f7e1223d71e9fd88",
}

// Pangolin pool manager v2.
//
// map[network] = address.
var PangolinPoolManagerV2List = map[string]string{
	chainId.AvalancheChainName: "0x912b5d41656048ef681efa9d32488a3ffe397994",
}
