package ethaddr

import "github.com/0xVanfer/chainId"

// Mai stable coin: MAI.
//
// map[network] = address.
var MaiList = map[string]string{
	chainId.AvalancheChainName: "0x5c49b268c9841aff1cc3b0a418ff5c3442ee3f3b", // MAI
}

// Mai MATIC: miMATIC.
//
// map[network] = address.
var MiMATICList = map[string]string{
	chainId.PolygonChainName: "0xa3fa99a148fa48d14ed51d610c367c61876997f1", // miMATIC
}
