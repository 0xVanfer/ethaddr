package ethaddr

import "github.com/0xVanfer/chainId"

// Abracadabra magic internet money: MIM.
// map[network] = address.
var MIMList = map[string]string{
	chainId.EthereumChainName:  "0x99d8a9c45b2eca8864373a26d1459e3dff1e17f3", // MIM
	chainId.AvalancheChainName: "0x130966628846bfd36ff31a822705796e8cb8c18d", // MIM
}
