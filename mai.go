package ethaddr

import "github.com/0xVanfer/chainId"

// Mai stable coin: MAI.
//
// map[network] = address.
var MAIList = map[string]string{
	chainId.ArbitrumChainName:  "0x3F56e0c36d275367b8C502090EDF38289b3dEa0d", // MAI, 0x3f56e0c36d275367b8c502090edf38289b3dea0d
	chainId.AvalancheChainName: "0x5c49b268c9841AFF1Cc3B0a418ff5c3442eE3F3b", // MAI, 0x5c49b268c9841aff1cc3b0a418ff5c3442ee3f3b
}

// Mai MATIC: miMATIC.
//
// map[network] = address.
var MiMATICList = map[string]string{
	chainId.PolygonChainName: "0xa3Fa99A148fA48D14Ed51d610c367C61876997F1", // miMATIC, 0xa3fa99a148fa48d14ed51d610c367c61876997f1
}
