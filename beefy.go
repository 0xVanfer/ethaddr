package ethaddr

// Website: https://app.beefy.finance/
//
// Docs: https://docs.beefy.finance/
//
// Deployed contracts: https://docs.beefy.finance/additional-resources/contract-addresses
const BeefyProtocol string = "beefy"

// Beefy token: BIFI.
//
// map[chainID] = address.
var BeefyTokenList = map[int64]string{
	ChainArbitrum:  "0x99C409E5f62E4bd2AC142f17caFb6810B8F0BAAE", // BIFI, 0x99c409e5f62e4bd2ac142f17cafb6810b8f0baae
	ChainAvalanche: "0xd6070ae98b8069de6B494332d1A1a81B6179D960", // BIFI, 0xd6070ae98b8069de6b494332d1a1a81b6179d960
	ChainEthereum:  "0xB1F1ee126e9c96231Cc3d3fAD7C08b4cf873b1f1", // BIFI, 0xb1f1ee126e9c96231cc3d3fad7c08b4cf873b1f1
}

// Same as BeefyTokenList.
var BIFIList = BeefyTokenList
