package ethaddr

import "github.com/0xVanfer/chainId"

// Tron TRX.
//
// map[network] = address.
var TRXList = map[string]string{
	chainId.BNBSmartChainName: "0xce7de646e7208a4ef112cb6ed5038fa6cc6b12e3", // TRX
}

// Old Tron TRX.
//
// map[network] = address.
var TRXOldList = map[string]string{
	chainId.BNBSmartChainName: "0x85eac5ac2f758618dfa09bdbe0cf174e7d574d5b", // TRXOLD
}
