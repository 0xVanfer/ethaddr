package ethaddr

import "github.com/0xVanfer/chainId"

// Uniswap token: UNI.
//
// map[network] = address.
var UniswapTokenList = map[string]string{
	chainId.EthereumChainName: "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984", // UNI, 0x1f9840a85d5af5bf1d1762f925bdaddc4201f984
	chainId.ArbitrumChainName: "0xFa7F8980b0f1E64A2062791cc3b0871572f1F7f0", // UNI, 0xfa7f8980b0f1e64a2062791cc3b0871572f1f7f0
}

// Same as UniswapTokenList.
var UNIList = UniswapTokenList
