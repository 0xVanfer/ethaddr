package ethaddr

// Website: https://www.swellnetwork.io/
const SwellProtocol string = "swell"

// Swell staked ETH: swETH.
//
// map[chainID] = address
var SwellStakedEtherList = map[int64]string{
	ChainEthereum: "0xf951E335afb289353dc249e82926178EaC7DEd78", // swETH, 0xf951e335afb289353dc249e82926178eac7ded78
}

// Restaked Swell ETH: rswETH.
//
// map[chainID] = address
var SwellRestakedEtherList = map[int64]string{
	ChainEthereum: "0xFAe103DC9cf190eD75350761e95403b7b8aFa6c0", // rswETH, 0xfae103dc9cf190ed75350761e95403b7b8afa6c0
}

// Same as SwellStakedEtherList.
var SWETHList = SwellStakedEtherList

// Same as SwellRestakedEtherList.
var RswETHList = SwellRestakedEtherList
