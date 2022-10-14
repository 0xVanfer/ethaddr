package ethaddr

import "github.com/0xVanfer/chainId"

// Reflexer token: RAI.
//
// map[network] = address.
var ReflexerTokenList = map[string]string{
	chainId.EthereumChainName: "0x03ab458634910aad20ef5f1c8ee96f1d6ac54919", // RAI
}

// Same as ReflexerTokenList.
var RAIList = ReflexerTokenList
