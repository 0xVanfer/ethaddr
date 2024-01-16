package ethaddr

import "github.com/0xVanfer/chainId"

const PendleProtocol string = "pendle"

// Pendle token: PENDLE.
//
// map[network] = address.
var PENDLEList = map[string]string{
	chainId.ArbitrumChainName: "0x0c880f6761f1af8d9aa9c466984b80dab9a8c9e8", // PENDLE
}
