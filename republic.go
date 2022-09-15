package ethaddr

import "github.com/0xVanfer/chainId"

const RepublicProtocol string = "republic"

// Republic token: REN.
var RepublicTokenList = map[string]string{
	chainId.EthereumChainName: "0x408e41876cccdc0f92210600ef50372656052a38", // REN
}

// Republic btc: renBTC.
var RenBTCList = map[string]string{
	chainId.EthereumChainName:  "0xeb4c2781e4eba804ce9a9803c67d0893436bb27d", // renBTC
	chainId.AvalancheChainName: "0xdbf31df14b66535af65aac99c32e9ea844e14501", // renBTC
}

// Republic fil: renFIL.
var RenFILList = map[string]string{
	chainId.EthereumChainName: "0xd5147bc8e386d91cc5dbe72099dac6c9b99276f5", // renFIL
}
