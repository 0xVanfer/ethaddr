package ethaddr

import "github.com/0xVanfer/chainId"

// Ampleforth token: AMPL.
//
// map[network] = address.
var AmpleforthTokenList = map[string]string{
	chainId.EthereumChainName: "0xD46bA6D942050d489DBd938a2C909A5d5039A161", // AMPL, 0xd46ba6d942050d489dbd938a2c909a5d5039a161
}

// Same as AmpleforthTokenList.
var AMPLList = AmpleforthTokenList
