package ethaddr

import "github.com/0xVanfer/chainId"

const StaderProtocol string = "stader"

// MaticX by Stader.
//
// map[network] = address.
var MaticXList = map[string]string{
	chainId.EthereumChainName: "0xf03a7eb46d01d9ecaa104558c732cf82f6b6b645", // MaticX
	chainId.PolygonChainName:  "0xfa68fb4628dff1028cfec22b4162fccd0d45efb6", // MaticX
}
