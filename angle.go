package ethaddr

import "github.com/0xVanfer/chainId"

const AngleProtocol string = "angle"

// Angle token: ANGLE.
//
// map[network] = address.
var AngleTokenList = map[string]string{
	chainId.PolygonChainName: "0x900f717ea076e1e7a484ad9dd2db81ceec60ebf1", // ANGLE
}

// Same as AngleTokenList.
var ANGLEList = AngleTokenList

// Angle EUR token: agEUR.
//
// map[network] = address.
var AgEURList = map[string]string{
	chainId.PolygonChainName:  "0xe0b52e49357fd4daf2c15e02058dce6bc0057db4", // agEUR
	chainId.ArbitrumChainName: "0xfa5ed56a203466cbbc2430a43c66b9d8723528e7", // agEUR
}
