package ethaddr

import "github.com/0xVanfer/chainId"

const DecentralandProtocol string = "decentraland"

// Decentraland token: MANA.
// map[network] = address.
var DecentralandTokenList = map[string]string{
	chainId.EthereumChainName: "0x0f5d2fb29fb7d3cfee444a200298f468908cc942", // MANA
}
