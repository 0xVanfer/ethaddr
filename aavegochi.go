package ethaddr

import "github.com/0xVanfer/chainId"

// Aavegochi token: GHST.
//
// map[network] = address.
var AavegochiTokenList = map[string]string{
	chainId.PolygonChainName: "0x385eeac5cb85a38a9a07a70c73e0a3271cfb54a7", // GHST
}

// Same as AavegochiTokenList.
var GHSTList = AavegochiTokenList
