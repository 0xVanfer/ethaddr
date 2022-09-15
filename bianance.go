package ethaddr

import "github.com/0xVanfer/chainId"

// Binance stable coin: BUSD.
var BUSDList = map[string]string{
	chainId.EthereumChainName: "0x4fabb145d64652a948d72533023f6e7a623c7c53", // BUSD
}

// Wrapped binance smart chain token: WBNB.
var WBNBList = map[string]string{
	chainId.EthereumChainName:      "0x418d75f65a02b3d53b2418fb8e1fe493759c7605", // WBNB
	chainId.BinanaceSmartChainName: "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c", // WBNB
}
