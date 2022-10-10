package ethaddr

import "github.com/0xVanfer/chainId"

// Angel EUR token: agEUR.
//
// map[network] = address.
//
// Deprecated: Use AgEURList instead.
var AngelEURList = map[string]string{
	chainId.PolygonChainName: "0xe0b52e49357fd4daf2c15e02058dce6bc0057db4", // agEUR
}

// Angel EUR token: agEUR.
//
// map[network] = address.
var AgEURList = map[string]string{
	chainId.PolygonChainName: "0xe0b52e49357fd4daf2c15e02058dce6bc0057db4", // agEUR
}
