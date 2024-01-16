package ethaddr

import "github.com/0xVanfer/chainId"

// Docs: https://docs.sushi.com/docs/intro
//
// Deployed contracts: https://docs.sushi.com/docs/Products/Classic%20AMM/Deployment%20Addresses
const SushiProtocol string = "sushi"

// Sushi token: SUSHI.
//
// map[network] = address.
var SushiTokenList = map[string]string{
	chainId.EthereumChainName:  "0x6b3595068778dd592e39a122f4f5a5cf09c90fe2", // SUSHI
	chainId.AvalancheChainName: "0x37b608519f91f70f2eeb0e5ed9af4061722e4f76", // SUSHI.e
	chainId.PolygonChainName:   "0x0b3f868e0be5597d5db7feb59e1cadbb0fdda50a", // SUSHI
	chainId.ArbitrumChainName:  "0xd4d42f0b6def4ce0383636770ef773390d85c61a", // SUSHI
}

// Same as SushiTokenList.
var SUSHIList = SushiTokenList

// XSUSHI.
//
// map[network] = address.
var XSushiList = map[string]string{
	chainId.EthereumChainName: "0x8798249c2e607446efb7ad49ec89dd1865ff4272", // XSUSHI
}

// v2 Router02
//
// https://docs.sushi.com/docs/Products/Classic%20AMM/Deployment%20Addresses
//
// map[network] = address
var SushiV2RouterList = map[string]string{
	chainId.EthereumChainName:  "0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F",
	chainId.ArbitrumChainName:  "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506",
	chainId.AvalancheChainName: "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506",
	chainId.BNBSmartChainName:  "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506",
	chainId.FantomChainName:    "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506",
	chainId.GoerliChainName:    "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506",
	chainId.OkChainName:        "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506",
	chainId.PolygonChainName:   "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506",
}

// v2 factory
//
// https://docs.sushi.com/docs/Products/Classic%20AMM/Deployment%20Addresses
//
// map[network] = address
var SushiV2FactoryList = map[string]string{
	chainId.EthereumChainName:  "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac",
	chainId.ArbitrumChainName:  "0xc35DADB65012eC5796536bD9864eD8773aBc74C4",
	chainId.AvalancheChainName: "0xc35DADB65012eC5796536bD9864eD8773aBc74C4",
	chainId.BNBSmartChainName:  "0xc35DADB65012eC5796536bD9864eD8773aBc74C4",
	chainId.FantomChainName:    "0xc35DADB65012eC5796536bD9864eD8773aBc74C4",
	chainId.GoerliChainName:    "0xc35DADB65012eC5796536bD9864eD8773aBc74C4",
	chainId.OkChainName:        "0xc35DADB65012eC5796536bD9864eD8773aBc74C4",
	chainId.PolygonChainName:   "0xc35DADB65012eC5796536bD9864eD8773aBc74C4",
}

// v3 factory
//
// https://docs.sushi.com/docs/Products/V3%20AMM/Core/Deployment%20Addresses
//
// map[network] = address
var SushiV3FactoryList = map[string]string{
	chainId.EthereumChainName:     "0xbACEB8eC6b9355Dfc0269C18bac9d6E2Bdc29C4F",
	chainId.ArbitrumChainName:     "0x1af415a1EbA07a4986a52B6f2e7dE7003D82231e",
	chainId.AvalancheChainName:    "0x3e603C14aF37EBdaD31709C4f848Fc6aD5BEc715",
	chainId.BNBSmartChainName:     "0x126555dd55a39328F69400d6aE4F782Bd4C34ABb",
	chainId.FantomChainName:       "0x7770978eED668a3ba661d51a773d3a992Fc9DDCB",
	chainId.OptimismChainName:     "0x9c6522117e2ed1fE5bdb72bb0eD5E3f2bdE7DBe0",
	chainId.PolygonChainName:      "0x917933899c6a5F8E37F31E19f92CdBFF7e8FF0e2",
	chainId.PolygonZkEVMChainName: "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506",
	chainId.ScrollChainName:       "0x46B3fDF7b5CDe91Ac049936bF0bDb12c5d22202e",
	chainId.FilecoinChainName:     "0xc35DADB65012eC5796536bD9864eD8773aBc74C4",
}
