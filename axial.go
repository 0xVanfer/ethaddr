package ethaddr

import "github.com/0xVanfer/chainId"

const AxialProtocol string = "axial"

// Axial token: AXIAL.
//
// map[network] = address.
var AxialTokenList = map[string]string{
	chainId.AvalancheChainName: "0xcf8419a615c57511807236751c0af38db4ba3351", // AXIAL
}

// Same as AxialTokenList.
var AXIALList = AxialTokenList

// Axial chef v3.
//
// map[network] = address.
var AxialChefList = map[string]string{
	chainId.AvalancheChainName: "0x958c0d0baa8f220846d3966742d4fb5edc5493d3",
}

// Axial sAXIAL: sAXIAL.
//
// map[network] = address.
var AxialsAXIALList = map[string]string{
	chainId.AvalancheChainName: "0xed7f93c8fd3b96b53c924f601b3948175d2820d8", // sAXIAL
}

// Axial veAXIAL: veAXIAL.
//
// map[network] = address.
var AxialveAXIALList = map[string]string{
	chainId.AvalancheChainName: "0x3f563f7efc6dc55adfc1b64bc6bd4bc5f394c4b2", // veAXIAL
}
