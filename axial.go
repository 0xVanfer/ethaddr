package ethaddr

import "github.com/0xVanfer/chainId"

// WARNING: Axial has dead.
//
// Website: https://www.axial.exchange/
//
// X (twitter): https://twitter.com/AxialDeFi
const AxialProtocol string = "axial"

// WARNING: Axial has dead.
//
// Axial token: AXIAL.
//
// map[network] = address.
var AxialTokenList = map[string]string{
	chainId.AvalancheChainName: "0xcF8419A615c57511807236751c0AF38Db4ba3351", // AXIAL, 0xcf8419a615c57511807236751c0af38db4ba3351
}

// WARNING: Axial has dead.
//
// Same as AxialTokenList.
var AXIALList = AxialTokenList

// WARNING: Axial has dead.
//
// Axial chef v3.
//
// map[network] = address.
var AxialChefList = map[string]string{
	chainId.AvalancheChainName: "0x958C0d0baA8F220846d3966742D4Fb5edc5493D3", // 0x958c0d0baa8f220846d3966742d4fb5edc5493d3
}

// WARNING: Axial has dead.
//
// Axial sAXIAL: sAXIAL.
//
// map[network] = address.
var AxialsAXIALList = map[string]string{
	chainId.AvalancheChainName: "0xed7f93C8FD3B96B53c924F601B3948175D2820D8", // sAXIAL, 0xed7f93c8fd3b96b53c924f601b3948175d2820d8
}

// WARNING: Axial has dead.
//
// Axial veAXIAL: veAXIAL.
//
// map[network] = address.
var AxialveAXIALList = map[string]string{
	chainId.AvalancheChainName: "0x3f563F7efc6dC55adFc1B64BC6Bd4bC5F394c4b2", // veAXIAL, 0x3f563f7efc6dc55adfc1b64bc6bd4bc5f394c4b2
}
