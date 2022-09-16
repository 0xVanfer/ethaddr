package ethaddr

import "github.com/0xVanfer/chainId"

const UniswapProtocolV2 = "uniswapv2"
const UniswapProtocolV3 = "uniswapv3"

// Uniswap token: UNI.
// map[network] = address.
var UniswapTokenList = map[string]string{
	chainId.EthereumChainName: "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984", // UNI
}
