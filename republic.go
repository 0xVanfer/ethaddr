package ethaddr

import "github.com/0xVanfer/chainId"

const RepublicProtocol string = "republic"

// Republic token: REN.
//
// map[network] = address.
var RepublicTokenList = map[string]string{
	chainId.EthereumChainName: "0x408e41876cCCDC0F92210600ef50372656052a38", // REN, 0x408e41876cccdc0f92210600ef50372656052a38
}

// Same as RepublicTokenList.
var RENList = RepublicTokenList

// Republic btc: renBTC.
//
// map[network] = address.
var RenBTCList = map[string]string{
	chainId.AvalancheChainName: "0xDBf31dF14B66535aF65AaC99C32e9eA844e14501", // renBTC, 0xdbf31df14b66535af65aac99c32e9ea844e14501
	chainId.EthereumChainName:  "0xEB4C2781e4ebA804CE9a9803C67d0893436bB27D", // renBTC, 0xeb4c2781e4eba804ce9a9803c67d0893436bb27d
}

// Republic fil: renFIL.
//
// map[network] = address.
var RenFILList = map[string]string{
	chainId.EthereumChainName: "0xD5147bc8e386d91Cc5DBE72099DAC6C9b99276F5", // renFIL, 0xd5147bc8e386d91cc5dbe72099dac6c9b99276f5
}
