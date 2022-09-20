package ethaddr

import "github.com/0xVanfer/chainId"

// map[network] = oracle address.
var AVAXOracleMap = map[string]string{
	chainId.AvalancheChainName: "0x0a77230d17318075983913bc2145db16c7366156",
}

// map[network] = oracle address.
var SAVAXOracleMap = map[string]string{
	chainId.AvalancheChainName: "0xc9245871D69BF4c36c6F2D15E0D68Ffa883FE1A7",
}
