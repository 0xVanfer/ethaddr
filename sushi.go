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
	chainId.ArbitrumChainName:  "0xd4d42F0b6DEF4CE0383636770eF773390d85c61A", // SUSHI, 0xd4d42f0b6def4ce0383636770ef773390d85c61a
	chainId.AvalancheChainName: "0x37B608519F91f70F2EeB0e5Ed9AF4061722e4F76", // SUSHI.e, 0x37b608519f91f70f2eeb0e5ed9af4061722e4f76
	chainId.EthereumChainName:  "0x6B3595068778DD592e39A122f4f5a5cF09C90fE2", // SUSHI, 0x6b3595068778dd592e39a122f4f5a5cf09c90fe2
	chainId.PolygonChainName:   "0x0b3F868E0BE5597D5DB7fEB59E1CADBb0fdDa50a", // SUSHI, 0x0b3f868e0be5597d5db7feb59e1cadbb0fdda50a
}

// Same as SushiTokenList.
var SUSHIList = SushiTokenList

// XSUSHI.
//
// map[network] = address.
var XSushiList = map[string]string{
	chainId.EthereumChainName: "0x8798249c2E607446EfB7Ad49eC89dD1865Ff4272", // XSUSHI, 0x8798249c2e607446efb7ad49ec89dd1865ff4272
}

// v2 Router02
//
// https://docs.sushi.com/docs/Products/Classic%20AMM/Deployment%20Addresses
//
// map[network] = address
var SushiV2RouterList = map[string]string{
	chainId.ArbitrumChainName:  "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", // 0x1b02da8cb0d097eb8d57a175b88c7d8b47997506
	chainId.AvalancheChainName: "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", // 0x1b02da8cb0d097eb8d57a175b88c7d8b47997506
	chainId.BNBSmartChainName:  "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", // 0x1b02da8cb0d097eb8d57a175b88c7d8b47997506
	chainId.EthereumChainName:  "0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F", // 0xd9e1ce17f2641f24ae83637ab66a2cca9c378b9f
	chainId.FantomChainName:    "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", // 0x1b02da8cb0d097eb8d57a175b88c7d8b47997506
	chainId.GoerliChainName:    "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", // 0x1b02da8cb0d097eb8d57a175b88c7d8b47997506
	chainId.OkChainName:        "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", // 0x1b02da8cb0d097eb8d57a175b88c7d8b47997506
	chainId.PolygonChainName:   "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", // 0x1b02da8cb0d097eb8d57a175b88c7d8b47997506
}

// v2 factory
//
// https://docs.sushi.com/docs/Products/Classic%20AMM/Deployment%20Addresses
//
// map[network] = address
var SushiV2FactoryList = map[string]string{
	chainId.ArbitrumChainName:  "0xc35DADB65012eC5796536bD9864eD8773aBc74C4", // 0xc35dadb65012ec5796536bd9864ed8773abc74c4
	chainId.AvalancheChainName: "0xc35DADB65012eC5796536bD9864eD8773aBc74C4", // 0xc35dadb65012ec5796536bd9864ed8773abc74c4
	chainId.BNBSmartChainName:  "0xc35DADB65012eC5796536bD9864eD8773aBc74C4", // 0xc35dadb65012ec5796536bd9864ed8773abc74c4
	chainId.EthereumChainName:  "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac", // 0xc0aee478e3658e2610c5f7a4a2e1777ce9e4f2ac
	chainId.FantomChainName:    "0xc35DADB65012eC5796536bD9864eD8773aBc74C4", // 0xc35dadb65012ec5796536bd9864ed8773abc74c4
	chainId.GoerliChainName:    "0xc35DADB65012eC5796536bD9864eD8773aBc74C4", // 0xc35dadb65012ec5796536bd9864ed8773abc74c4
	chainId.OkChainName:        "0xc35DADB65012eC5796536bD9864eD8773aBc74C4", // 0xc35dadb65012ec5796536bd9864ed8773abc74c4
	chainId.PolygonChainName:   "0xc35DADB65012eC5796536bD9864eD8773aBc74C4", // 0xc35dadb65012ec5796536bd9864ed8773abc74c4
}

// v3 factory
//
// https://docs.sushi.com/docs/Products/V3%20AMM/Core/Deployment%20Addresses
//
// map[network] = address
var SushiV3FactoryList = map[string]string{
	chainId.ArbitrumChainName:     "0x1af415a1EbA07a4986a52B6f2e7dE7003D82231e", // 0x1af415a1eba07a4986a52b6f2e7de7003d82231e
	chainId.AvalancheChainName:    "0x3e603C14aF37EBdaD31709C4f848Fc6aD5BEc715", // 0x3e603c14af37ebdad31709c4f848fc6ad5bec715
	chainId.BNBSmartChainName:     "0x126555dd55a39328F69400d6aE4F782Bd4C34ABb", // 0x126555dd55a39328f69400d6ae4f782bd4c34abb
	chainId.EthereumChainName:     "0xbACEB8eC6b9355Dfc0269C18bac9d6E2Bdc29C4F", // 0xbaceb8ec6b9355dfc0269c18bac9d6e2bdc29c4f
	chainId.FantomChainName:       "0x7770978eED668a3ba661d51a773d3a992Fc9DDCB", // 0x7770978eed668a3ba661d51a773d3a992fc9ddcb
	chainId.FilecoinChainName:     "0xc35DADB65012eC5796536bD9864eD8773aBc74C4", // 0xc35dadb65012ec5796536bd9864ed8773abc74c4
	chainId.OptimismChainName:     "0x9c6522117e2ed1fE5bdb72bb0eD5E3f2bdE7DBe0", // 0x9c6522117e2ed1fe5bdb72bb0ed5e3f2bde7dbe0
	chainId.PolygonChainName:      "0x917933899c6a5F8E37F31E19f92CdBFF7e8FF0e2", // 0x917933899c6a5f8e37f31e19f92cdbff7e8ff0e2
	chainId.PolygonZkEVMChainName: "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", // 0x1b02da8cb0d097eb8d57a175b88c7d8b47997506
	chainId.ScrollChainName:       "0x46B3fDF7b5CDe91Ac049936bF0bDb12c5d22202e", // 0x46b3fdf7b5cde91ac049936bf0bdb12c5d22202e
}
