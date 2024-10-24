package ethaddr

// Website: https://enjin.io/
//
// Docs: https://docs.enjin.io/
const EnjinProtocol string = "enjin"

// Enjin token: ENJ.
//
// map[network] = address.
var EnjinTokenList = map[int64]string{
	ChainEthereum: "0xF629cBd94d3791C9250152BD8dfBDF380E2a3B9c", // ENJ, 0xf629cbd94d3791c9250152bd8dfbdf380e2a3b9c
	ChainPolygon:  "0x7eC26842F195c852Fa843bB9f6D8B583a274a157", // ENJ, 0x7ec26842f195c852fa843bb9f6d8b583a274a157
}

// Same as EnjinTokenList.
var ENJList = EnjinTokenList
