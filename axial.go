package ethaddr

import "github.com/0xVanfer/chainId"

// Deprecated: Axial has dead.
//
// Website: https://www.axial.exchange/
const AxialProtocol string = "axial"

// Deprecated: Axial has dead.
//
// Axial token: AXIAL.
//
// map[network] = address.
var AxialTokenList = map[string]string{
	chainId.AvalancheChainName: "0xcf8419a615c57511807236751c0af38db4ba3351", // AXIAL
}

// Deprecated: Axial has dead.
//
// Same as AxialTokenList.
var AXIALList = AxialTokenList

// Deprecated: Axial has dead.
//
// Axial chef v3.
//
// map[network] = address.
var AxialChefList = map[string]string{
	chainId.AvalancheChainName: "0x958c0d0baa8f220846d3966742d4fb5edc5493d3",
}

// Deprecated: Axial has dead.
//
// Axial sAXIAL: sAXIAL.
//
// map[network] = address.
var AxialsAXIALList = map[string]string{
	chainId.AvalancheChainName: "0xed7f93c8fd3b96b53c924f601b3948175d2820d8", // sAXIAL
}

// Deprecated: Axial has dead.
//
// Axial veAXIAL: veAXIAL.
//
// map[network] = address.
var AxialveAXIALList = map[string]string{
	chainId.AvalancheChainName: "0x3f563f7efc6dc55adfc1b64bc6bd4bc5f394c4b2", // veAXIAL
}
