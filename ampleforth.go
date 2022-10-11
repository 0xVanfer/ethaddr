package ethaddr

import "github.com/0xVanfer/chainId"

// Ampleforth token: AMPL.
//
// map[network] = address.
var AmpleforthTokenList = map[string]string{
	chainId.EthereumChainName: "0xd46ba6d942050d489dbd938a2c909a5d5039a161", // AMPL
}

// Same as AmpleforthTokenList.
var AMPLList = AmpleforthTokenList
