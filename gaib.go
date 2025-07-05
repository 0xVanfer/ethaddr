package ethaddr

// Website: https://gaib.ai/
const GaibProtocol string = "gaib"

// Gaib AI token.
//
// map[chainID][underlying] = address.
var GaibAIDTokenList = map[int64]map[string]string{
	ChainSei: {
		USDT0List[ChainSei]: "0xDc45e7027A0489FE6C2E4A0735097d8E6952A340", // AIDaUSDT, 0xdc45e7027a0489fe6c2e4a0735097d8e6952a340
	},
}
