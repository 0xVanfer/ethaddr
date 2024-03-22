package ethaddr

// Wrapped btc: WBTC.
//
// NOT STABLE for avalanche. If WBTC is later deployed on avalanche, this address might be changed.
//
// map[network] = address.
var WBTCList = map[string]string{
	ChainArbitrum:  "0x2f2a2543B76A4166549F7aaB2e75Bef0aefC5B0f", // WBTC, 0x2f2a2543b76a4166549f7aab2e75bef0aefc5b0f
	ChainAvalanche: "0x50b7545627a5162F82A992c33b87aDc75187B218", // WBTC.e, 0x50b7545627a5162f82a992c33b87adc75187b218
	ChainEthereum:  "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599", // WBTC, 0x2260fac5e5542a773aa44fbcfedf7c193bc2c599
	ChainOptimism:  "0x68f180fcCe6836688e9084f035309E29Bf0A2095", // WBTC, 0x68f180fcce6836688e9084f035309e29bf0a2095
	ChainPolygon:   "0x1BFD67037B42Cf73acF2047067bd4F2C47D9BfD6", // WBTC, 0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6
}

// Wrapped btc on avalanche: WBTC.e.
//
// map[network] = address.
var WBTCeList = map[string]string{
	ChainAvalanche: "0x50b7545627a5162F82A992c33b87aDc75187B218", // WBTC.e, 0x50b7545627a5162f82a992c33b87adc75187b218
}
