package ethaddr

import "github.com/0xVanfer/chainId"

// Binance stable coin: BUSD.
//
// map[network] = address.
var BUSDList = map[string]string{
	chainId.EthereumChainName:  "0x4fabb145d64652a948d72533023f6e7a623c7c53", // BUSD
	chainId.AvalancheChainName: "0x9c9e5fd8bbc25984b178fdce6117defa39d2db39", // BUSD
	chainId.BNBSmartChainName:  "0xe9e7cea3dedca5984780bafc599bd69add087d56", // BUSD
}

// Wrapped binance smart chain token: WBNB.
//
// map[network] = address.
var WBNBList = map[string]string{
	chainId.EthereumChainName: "0x418d75f65a02b3d53b2418fb8e1fe493759c7605", // WBNB
	chainId.BNBSmartChainName: "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c", // WBNB
}

// Binance Beacon ETH: BETH.
//
// map[network] = address.
var BETHList = map[string]string{
	chainId.BNBSmartChainName: "0x250632378e573c6be1ac2f97fcdf00515d0aa91b", // BETH
}

// Wrapped Binance Beacon ETH: WBETH.
//
// map[network] = address.
var WBETHList = map[string]string{
	chainId.BNBSmartChainName: "0xa2e3356610840701bdf5611a53974510ae27e2e1", // wBETH
}
