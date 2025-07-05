package ethaddr

// Website: https://app.stakestone.io/u/stake
//
// Doc: https://docs.stakestone.io/stakestone
const StakeStoneProtocol string = "stakestone"

// Stakestone ETH lst STONE.
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

// Stakestone vault.
//
// map[chainID] = address.
var StoneVaultList = map[int64]string{
	ChainEthereum: "0xA62F9C5af106FeEE069F38dE51098D9d81B90572", // 0xa62f9c5af106feee069f38de51098d9d81b90572
}

// Stakestone yield bearing BTC.
//
// map[chainID] = address.
var STONEBTCList = map[int64]string{
	ChainEthereum: "0x2D83F5BfC83cF0B09B8884101C015FA9c74c32F3", // STONEBTC, 0x2d83f5bfc83cf0b09b8884101c015fa9c74c32f3
}

// Stakestone yield bearing BTC vault.
//
// map[chainID] = address.
var StoneBTCVaultList = map[int64]string{
	ChainEthereum: "0x1fC603779DC6b4866769A58067777D2C52628226", // 0x1fc603779dc6b4866769a58067777d2c52628226
}

// Stakestone SBTC.
//
// map[chainID] = address.
var SBTCList = map[int64]string{
	ChainBSC:      "0x15469528C11E8Ace863F3F9e5a8329216e33dD7d", // SBTC, 0x15469528c11e8ace863f3f9e5a8329216e33dd7d
	ChainEthereum: "0x094c0e36210634c3CfA25DC11B96b562E0b07624", // SBTC, 0x094c0e36210634c3cfa25dc11b96b562e0b07624
	ChainScroll:   "0xE630Abc6a480AC27270Fa9Ce615bFA5917e85525", // SBTC, 0xe630abc6a480ac27270fa9ce615bfa5917e85525
	ChainSei:      "0x095957CEB9f317ac1328f0aB3123622401766D71", // SBTC, 0x095957ceb9f317ac1328f0ab3123622401766d71
}

// Stakestone SBTC vault.
//
// map[chainID] = address.
var SBTCVaultList = map[int64]string{
	ChainBSC:      "0x3aa0670E24Cb122e1d5307Ed74b0c44d619aFF9b", // 0x3aa0670e24cb122e1d5307ed74b0c44d619aff9b
	ChainEthereum: "0x7dBAC0aA440A25D7FB43951f7b178FF7A809108D", // 0x7dbac0aa440a25d7fb43951f7b178ff7a809108d
}
