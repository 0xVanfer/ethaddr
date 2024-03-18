package ethaddr

import "github.com/0xVanfer/chainId"

// Frax token: FRAX.
//
// map[network] = address.
var FRAXList = map[string]string{
	chainId.EthereumChainName:  "0x853d955aCEf822Db058eb8505911ED77F175b99e", // FRAX, 0x853d955acef822db058eb8505911ed77f175b99e
	chainId.AvalancheChainName: "0xd24c2ad096400b6fbcd2ad8b24e7acbc21a1da64", // FRAX
	chainId.OptimismChainName:  "0x2e3d870790dc77a83dd1d18184acc7439a53f475", // FRAX
	chainId.ArbitrumChainName:  "0x17fc002b466eec40dae837fc4be5c67993ddbd6f", // FRAX
}

// Frax share token: FXS.
//
// map[network] = address.
var FXSList = map[string]string{
	chainId.EthereumChainName: "0x3432B6A60D23Ca0dFCa7761B7ab56459D9C964D0", // FXS, 0x3432b6a60d23ca0dfca7761b7ab56459d9c964d0
}
