package ethaddr

import "github.com/0xVanfer/chainId"

// Binance stable coin: BUSD.
//
// map[network] = address.
var BUSDList = map[string]string{
	chainId.EthereumChainName:  "0x4fabb145d64652a948d72533023f6e7a623c7c53", // BUSD
	chainId.AvalancheChainName: "0x9c9e5fd8bbc25984b178fdce6117defa39d2db39", // BUSD
}

// Wrapped binance smart chain token: WBNB.
//
// map[network] = address.
var WBNBList = map[string]string{
	chainId.EthereumChainName:     "0x418d75f65a02b3d53b2418fb8e1fe493759c7605", // WBNB
	chainId.BinanceSmartChainName: "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c", // WBNB
}
