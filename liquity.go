package ethaddr

import "github.com/0xVanfer/chainId"

// Liquity stable coin: LUSD.
//
// map[network] = address.
var LUSDList = map[string]string{
	chainId.EthereumChainName: "0x5f98805A4E8be255a32880FDeC7F6728C6568bA0", // LUSD, 0x5f98805a4e8be255a32880fdec7f6728c6568ba0
	chainId.ArbitrumChainName: "0x93b346b6bc2548da6a1e7d98e9a421b42541425b", // LUSD
}
