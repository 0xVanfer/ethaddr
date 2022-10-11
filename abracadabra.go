package ethaddr

import "github.com/0xVanfer/chainId"

const AbracadabraProtocol string = "abracadabra"

// Abracadabra token: SPELL.
//
// map[network] = address.
var SPELLList = map[string]string{
	chainId.EthereumChainName:  "0x090185f2135308bad17527004364ebcc2d37e5f6", // SPELL
	chainId.AvalancheChainName: "0xce1bffbd5374dac86a2893119683f4911a2f7814", // SPELL
}

// Abracadabra magic internet money: MIM.
//
// map[network] = address.
var MIMList = map[string]string{
	chainId.EthereumChainName:  "0x99d8a9c45b2eca8864373a26d1459e3dff1e17f3", // MIM
	chainId.AvalancheChainName: "0x130966628846bfd36ff31a822705796e8cb8c18d", // MIM
}
