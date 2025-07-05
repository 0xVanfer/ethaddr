package ethaddr

const RepublicProtocol string = "republic"

// Republic token: REN.
//
// map[chainID] = address.
var RepublicTokenList = map[int64]string{
	ChainEthereum: "0x408e41876cCCDC0F92210600ef50372656052a38", // REN, 0x408e41876cccdc0f92210600ef50372656052a38
}

// Same as RepublicTokenList.
var RENList = RepublicTokenList

// Republic btc: renBTC.
//
// map[chainID] = address.
var RenBTCList = map[int64]string{
	ChainAvalanche: "0xDBf31dF14B66535aF65AaC99C32e9eA844e14501", // renBTC, 0xdbf31df14b66535af65aac99c32e9ea844e14501
	ChainEthereum:  "0xEB4C2781e4ebA804CE9a9803C67d0893436bB27D", // renBTC, 0xeb4c2781e4eba804ce9a9803c67d0893436bb27d
}

// Republic fil: renFIL.
//
// map[chainID] = address.
var RenFILList = map[int64]string{
	ChainEthereum: "0xD5147bc8e386d91Cc5DBE72099DAC6C9b99276F5", // renFIL, 0xd5147bc8e386d91cc5dbe72099dac6c9b99276f5
}
