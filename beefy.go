package ethaddr

import "github.com/0xVanfer/chainId"

const BeefyProtocol string = "beefy"

// Beefy token: BIFI.
//
// map[network] = address.
var BeefyTokenList = map[string]string{
	chainId.AvalancheChainName: "0xd6070ae98b8069de6b494332d1a1a81b6179d960", // BIFI
}
