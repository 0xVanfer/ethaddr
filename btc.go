package ethaddr

// BTC.b or BTCB.
//
// map[network] = address.
var BTCbList = map[string]string{
	ChainAvalanche: "0x152b9d0FdC40C096757F570A51E494bd4b943E50", // BTC.b, 0x152b9d0fdc40c096757f570a51e494bd4b943e50
	ChainArbitrum:  "0x2297aEbD383787A160DD0d9F71508148769342E3", // BTC.b, 0x2297aebd383787a160dd0d9f71508148769342e3
	ChainBSC:       "0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c", // BTCB, 0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c
}

// Bitcoin cash token.
//
// map[network] = address.
var BCHList = map[string]string{
	ChainBSC: "0x8fF795a6F4D97E7887C79beA79aba5cc76444aDf", // BCH, 0x8ff795a6f4d97e7887c79bea79aba5cc76444adf
}
