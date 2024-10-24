package ethaddr

// Website: https://app.stakestone.io/u/stake
//
// Doc: https://docs.stakestone.io/stakestone
const StakeStoneProtocol string = "stakestone"

// Stakestone lst STONE.
// Is layerzero OFT.
//
// map[chainID] = address.
var STONEList = map[int64]string{
	ChainArbitrum: "0x80137510979822322193FC997d400D5A6C747bf7", // STONE, 0x80137510979822322193fc997d400d5a6c747bf7
	ChainBase:     "0xD2012fc1B913cE50732ebcaa7E601fe37Ac728C6", // STONE, 0xd2012fc1b913ce50732ebcaa7e601fe37ac728c6
	ChainBSC:      "0x80137510979822322193FC997d400D5A6C747bf7", // STONE, 0x80137510979822322193fc997d400d5a6c747bf7
	ChainEthereum: "0x7122985656e38BDC0302Db86685bb972b145bD3C", // STONE, 0x7122985656e38bdc0302db86685bb972b145bd3c
	ChainLinea:    "0x93F4d0ab6a8B4271f4a28Db399b5E30612D21116", // STONE, 0x93f4d0ab6a8b4271f4a28db399b5e30612d21116
	ChainMantle:   "0x2Fde62942759d7C0aaf25952Da4098423bC1264C", // STONE, 0x2fde62942759d7c0aaf25952da4098423bc1264c
	ChainOptimism: "0x80137510979822322193FC997d400D5A6C747bf7", // STONE, 0x80137510979822322193fc997d400d5a6c747bf7
	ChainScroll:   "0x80137510979822322193FC997d400D5A6C747bf7", // STONE, 0x80137510979822322193fc997d400d5a6c747bf7
}

// Stakestone vault where can read current share price.
//
// map[chainID] = address.
var StoneVaultList = map[int64]string{
	ChainEthereum: "0xA62F9C5af106FeEE069F38dE51098D9d81B90572", // 0xa62f9c5af106feee069f38de51098d9d81b90572
}
