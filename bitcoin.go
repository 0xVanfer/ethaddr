package ethaddr

// Wrapped BTC: WBTC.
//
// For chain avalanche (43114), use WBTCeList instead.
// If the protocol later deploy the token on avalanche, this address might be changed.
//
// map[chainID] = address.
var WBTCList = map[int64]string{
	ChainArbitrum:  "0x2f2a2543B76A4166549F7aaB2e75Bef0aefC5B0f", // WBTC, 0x2f2a2543b76a4166549f7aab2e75bef0aefc5b0f
	ChainAvalanche: "0x50b7545627a5162F82A992c33b87aDc75187B218", // WBTC.e, 0x50b7545627a5162f82a992c33b87adc75187b218
	ChainBera:      "0x0555E30da8f98308EdB960aa94C0Db47230d2B9c", // WBTC, 0x0555e30da8f98308edb960aa94c0db47230d2b9c
	ChainEthereum:  "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599", // WBTC, 0x2260fac5e5542a773aa44fbcfedf7c193bc2c599
	ChainMerlin:    "0xF6D226f9Dc15d9bB51182815b320D3fBE324e1bA", // WBTC, 0xf6d226f9dc15d9bb51182815b320d3fbe324e1ba
	ChainOptimism:  "0x68f180fcCe6836688e9084f035309E29Bf0A2095", // WBTC, 0x68f180fcce6836688e9084f035309e29bf0a2095
	ChainPolygon:   "0x1BFD67037B42Cf73acF2047067bd4F2C47D9BfD6", // WBTC, 0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6
	ChainScroll:    "0x3C1BCa5a656e69edCD0D4E36BEbb3FcDAcA60Cf1", // WBTC, 0x3c1bca5a656e69edcd0d4e36bebb3fcdaca60cf1
	ChainSei:       "0x0555E30da8f98308EdB960aa94C0Db47230d2B9c", // WBTC, 0x0555e30da8f98308edb960aa94c0db47230d2b9c
}

// Wrapped btc on avalanche: WBTC.e.
//
// map[chainID] = address.
var WBTCeList = map[int64]string{
	ChainAvalanche: "0x50b7545627a5162F82A992c33b87aDc75187B218", // WBTC.e, 0x50b7545627a5162f82a992c33b87adc75187b218
}

// BTC.b. Bridged BTC.
// Different from BTCB, which is the BTC withdrawn from binance on BSC.
//
// map[chainID] = address.
var BTCbList = map[int64]string{
	ChainArbitrum:  "0x2297aEbD383787A160DD0d9F71508148769342E3", // BTC.b, 0x2297aebd383787a160dd0d9f71508148769342e3
	ChainAvalanche: "0x152b9d0FdC40C096757F570A51E494bd4b943E50", // BTC.b, 0x152b9d0fdc40c096757f570a51e494bd4b943e50
}

// Bitcoin cash token.
//
// map[chainID] = address.
var BCHList = map[int64]string{
	ChainBSC: "0x8fF795a6F4D97E7887C79beA79aba5cc76444aDf", // BCH, 0x8ff795a6f4d97e7887c79bea79aba5cc76444adf
}
