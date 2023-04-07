package ethaddr

import "github.com/0xVanfer/chainId"

// Frax token: FRAX.
//
// map[network] = address.
var FRAXList = map[string]string{
	chainId.EthereumChainName:  "0x853d955acef822db058eb8505911ed77f175b99e", // FRAX
	chainId.AvalancheChainName: "0xd24c2ad096400b6fbcd2ad8b24e7acbc21a1da64", // FRAX
	chainId.OptimismChainName:  "0x2e3d870790dc77a83dd1d18184acc7439a53f475", // FRAX
	chainId.ArbitrumChainName:  "0x17fc002b466eec40dae837fc4be5c67993ddbd6f", // FRAX
}
