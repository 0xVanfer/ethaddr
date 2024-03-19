package ethaddr

import "github.com/0xVanfer/chainId"

// Now: https://stellaxyz.io/ Previous: https://alphaventuredao.io/
//
// Docs: https://docs.stellaxyz.io/stella-doc/overview/getting-started
const AlphaProtocol string = "alpha"

// Alpha token: ALPHA.
//
// map[network] = address.
var AlphaTokenList = map[string]string{
	chainId.AvalancheChainName: "0x2147EFFF675e4A4eE1C2f918d181cDBd7a8E208f", // ALPHA.e, 0x2147efff675e4a4ee1c2f918d181cdbd7a8e208f
}

// Same as AlphaTokenList.
var ALPHAList = AlphaTokenList
