package ethaddr

// Website: https://amphor.io/
const AmphorProtocol string = "amphor"

// Amphor symbiotic restaked ETH: amphrETH.
//
// map[chainID] = address
var AmphorSymbioticRestakedETHList = map[int64]string{
	ChainEthereum: "0x5fD13359Ba15A84B76f7F87568309040176167cd", // amphrETH, 0x5fd13359ba15a84b76f7f87568309040176167cd
}

// Same as AmphorRestakedETHList
var AmphrETHList = AmphorSymbioticRestakedETHList
