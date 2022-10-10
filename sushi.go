package ethaddr

import "github.com/0xVanfer/chainId"

const SushiProtocol string = "sushi"

// Sushi token: SUSHI.
//
// map[network] = address.
var SushiTokenList = map[string]string{
	chainId.EthereumChainName:  "0x6b3595068778dd592e39a122f4f5a5cf09c90fe2", // SUSHI
	chainId.AvalancheChainName: "0x37b608519f91f70f2eeb0e5ed9af4061722e4f76", // SUSHI.e
	chainId.PolygonChainName:   "0x0b3f868e0be5597d5db7feb59e1cadbb0fdda50a", // SUSHI
}

// XSUSHI.
//
// map[network] = address.
var XSushiList = map[string]string{
	chainId.EthereumChainName: "0x8798249c2e607446efb7ad49ec89dd1865ff4272", // XSUSHI
}
