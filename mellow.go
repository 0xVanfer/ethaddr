package ethaddr

// Website: https://app.mellow.finance/restake
const MellowProtocol string = "mellow"

// The vaults for Symbiotic wstETH restaking on Mellow frontend.
//
// Last updated: Jul.3.2024
var MellowWstETHVaultsForSymbiotic = map[int64][]string{
	ChainEthereum: {
		AmphrETHList[ChainEthereum], // Amphor
		RstETHList[ChainEthereum],   // P2P
		SteakLRTList[ChainEthereum], // Steakhouse
		Re7LRTList[ChainEthereum],   // Re7
		PzETHList[ChainEthereum],    // Renzo
		IfsETHList[ChainEthereum],   // InfStones
		CoETHList[ChainEthereum],    // Chorus one
		LugaETHList[ChainEthereum],  // Luganodes
	},
}
