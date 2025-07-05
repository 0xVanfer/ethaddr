package ethaddr

// Website: https://www.bitfi.one/
const BitfiProtocol = "bitfi"

// Bitfi BTC token.
//
// map[chainID] = address.
var BfBTCList = map[int64]string{
	ChainBSC:      "0x623F2774d9f27B59bc6b954544487532CE79d9DF", // bfBTC, 0x623f2774d9f27b59bc6b954544487532ce79d9df
	ChainEthereum: "0xCdFb58c8C859Cb3F62ebe9Cf2767F9e036C7fb15", // bfBTC, 0xcdfb58c8c859cb3f62ebe9cf2767f9e036c7fb15
}
