package ethaddr

import "github.com/0xVanfer/chainId"

// Docs: https://docs.beefy.finance/
//
// Deployed contracts: https://docs.beefy.finance/additional-resources/contract-addresses
const BeefyProtocol string = "beefy"

// Beefy token: BIFI.
//
// map[network] = address.
var BeefyTokenList = map[string]string{
	chainId.AvalancheChainName: "0xd6070ae98b8069de6b494332d1a1a81b6179d960", // BIFI
	chainId.ArbitrumChainName:  "0x99c409e5f62e4bd2ac142f17cafb6810b8f0baae", // BIFI
}

// Same as BeefyTokenList.
var BIFIList = BeefyTokenList
