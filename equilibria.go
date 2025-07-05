package ethaddr

// Website: https://equilibria.fi/home
const EquilibriaProtocol string = "equilibria"

// Equilibria token, EQB.
//
// map[chainID] = address
var EQBList = map[int64]string{
	ChainArbitrum: "0xBfbCFe8873fE28Dfa25f1099282b088D52bbAD9C", // EQB, 0xbfbcfe8873fe28dfa25f1099282b088d52bbad9c
	ChainBSC:      "0x374Ca32fd7934c5d43240E1e73fa9B2283468609", // EQB, 0x374ca32fd7934c5d43240e1e73fa9b2283468609
	ChainEthereum: "0xfE80D611c6403f70e5B1b9B722D2B3510B740B2B", // EQB, 0xfe80d611c6403f70e5b1b9b722d2b3510b740b2b
	ChainMantle:   "0x3e7eF8f50246f725885102E8238CBba33F276747", // EQB, 0x3e7ef8f50246f725885102e8238cbba33f276747
	ChainOptimism: "0xaf3A6f67Af1624d3878A8d30b09FAe7915DcA2a0", // EQB, 0xaf3a6f67af1624d3878a8d30b09fae7915dca2a0
}

// Equilibria token, xEQB.
//
// map[chainID] = address
var XEQBList = map[int64]string{
	ChainArbitrum: "0x96C4A48Abdf781e9c931cfA92EC0167Ba219ad8E", // xEQB, 0x96c4a48abdf781e9c931cfa92ec0167ba219ad8e
	ChainBSC:      "0xfE80D611c6403f70e5B1b9B722D2B3510B740B2B", // xEQB, 0xfe80d611c6403f70e5b1b9b722d2b3510b740b2b
	ChainEthereum: "0xd6eCfD0d5f1Dfd3ad30f267a3a29b3E1bC4fd54f", // xEQB, 0xd6ecfd0d5f1dfd3ad30f267a3a29b3e1bc4fd54f
	ChainMantle:   "0x96C4A48Abdf781e9c931cfA92EC0167Ba219ad8E", // xEQB, 0x96c4a48abdf781e9c931cfa92ec0167ba219ad8e
	ChainOptimism: "0xd6eCfD0d5f1Dfd3ad30f267a3a29b3E1bC4fd54f", // xEQB, 0xd6ecfd0d5f1dfd3ad30f267a3a29b3e1bc4fd54f
}

// Equilibria contract.
//
// map[chainID] = address
var EquilibriaPendleBooster = map[int64]string{
	ChainArbitrum: "0x4D32C8Ff2fACC771eC7Efc70d6A8468bC30C26bF", // 0x4d32c8ff2facc771ec7efc70d6a8468bc30c26bf
	ChainBSC:      "0x4D32C8Ff2fACC771eC7Efc70d6A8468bC30C26bF", // 0x4d32c8ff2facc771ec7efc70d6a8468bc30c26bf
	ChainEthereum: "0x4D32C8Ff2fACC771eC7Efc70d6A8468bC30C26bF", // 0x4d32c8ff2facc771ec7efc70d6a8468bc30c26bf
	ChainMantle:   "0x920873E5b302A619C54c908aDFB77a1C4256A3B8", // 0x920873e5b302a619c54c908adfb77a1c4256a3b8
	ChainOptimism: "0x18C61629E6CBAdB85c29ba7993f251b3EbE2B356", // 0x18c61629e6cbadb85c29ba7993f251b3ebe2b356
}
