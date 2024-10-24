package ethaddr

// Website: https://p2p.org/
const P2PProtocol string = "p2p"

// P2P symbiotic restaking vault ETH: rstETH.
//
// map[chainID] = address
var P2PRestakingVaultETHList = map[int64]string{
	ChainEthereum: "0x7a4EffD87C2f3C55CA251080b1343b605f327E3a", // rstETH, 0x7a4effd87c2f3c55ca251080b1343b605f327e3a
}

// Same as P2PRestakingVaultETHList.
var RstETHList = P2PRestakingVaultETHList
