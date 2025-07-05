package ethaddr

// Website: https://www.pangolin.exchange/
const PangolinProtocol string = "pangolin"

// Pangolin token: PNG.
//
// map[chainID] = address.
var PangolinTokenList = map[int64]string{
	ChainAvalanche: "0x60781C2586D68229fde47564546784ab3fACA982", // PNG, 0x60781c2586d68229fde47564546784ab3faca982
}

// Same as PangolinTokenList.
var PNGList = PangolinTokenList

// Pangolin chef v2.
//
// map[chainID] = address.
var PangolinChefV2List = map[int64]string{
	ChainAvalanche: "0x1f806f7C8dED893fd3caE279191ad7Aa3798E928", // 0x1f806f7c8ded893fd3cae279191ad7aa3798e928
}

// Pangolin factory.
//
// map[chainID] = address.
var PangolinFactoryList = map[int64]string{
	ChainAvalanche: "0xefa94DE7a4656D787667C749f7E1223D71E9FD88", // 0xefa94de7a4656d787667c749f7e1223d71e9fd88
}

// Pangolin pool manager v2.
//
// map[chainID] = address.
var PangolinPoolManagerV2List = map[int64]string{
	ChainAvalanche: "0x912b5D41656048Ef681eFa9D32488a3fFE397994", // 0x912b5d41656048ef681efa9d32488a3ffe397994
}
