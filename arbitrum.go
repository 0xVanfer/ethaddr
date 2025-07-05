package ethaddr

// Blockscan: https://arbiscan.io/
const ChainNameArbitrum string = "arbitrum"

// Arbitrum token ARB.
//
// map[chainID] = address.
var ARBList = map[int64]string{
	ChainEthereum: "0xB50721BCf8d664c30412Cfbc6cf7a15145234ad1", // ARB, 0xb50721bcf8d664c30412cfbc6cf7a15145234ad1
	ChainArbitrum: "0x912CE59144191C1204E64559FE8253a0e49E6548", // ARB, 0x912ce59144191c1204e64559fe8253a0e49e6548,
}
