package ethaddr

import "github.com/0xVanfer/chainId"

const UniswapProtocolV2 = "uniswapv2"
const UniswapProtocolV3 = "uniswapv3"

// Uniswap token: UNI.
//
// map[network] = address.
var UniswapTokenList = map[string]string{
	chainId.EthereumChainName: "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984", // UNI, 0x1f9840a85d5af5bf1d1762f925bdaddc4201f984
	chainId.ArbitrumChainName: "0xfa7f8980b0f1e64a2062791cc3b0871572f1f7f0", // UNI
}

// Same as UniswapTokenList.
var UNIList = UniswapTokenList

// Uniswap V2 swap router.
//
// map[network] = address.
//
// Uni V2 is only on ethereum.
var UniswapV2RouterList = map[string]string{
	chainId.EthereumChainName: "0x7a250d5630b4cf539739df2c5dacb4c659f2488d",
}

// Uniswap V3 swap router.
//
// map[network] = address.
var UniswapV3RouterList = map[string]string{
	chainId.EthereumChainName: "0xe592427a0aece92de3edee1f18e0157c05861564",
	chainId.GoerliChainName:   "0xe592427a0aece92de3edee1f18e0157c05861564",
	chainId.PolygonChainName:  "0xe592427a0aece92de3edee1f18e0157c05861564",
	chainId.OptimismChainName: "0xe592427a0aece92de3edee1f18e0157c05861564",
	chainId.ArbitrumChainName: "0xe592427a0aece92de3edee1f18e0157c05861564",
}

// Uniswap v3 factory.
//
// map[network] = address.
var UniswapV3FactoryList = map[string]string{
	chainId.EthereumChainName: "0x1f98431c8ad98523631ae4a59f267346ea31f984",
	chainId.GoerliChainName:   "0x1f98431c8ad98523631ae4a59f267346ea31f984",
	chainId.PolygonChainName:  "0x1f98431c8ad98523631ae4a59f267346ea31f984",
	chainId.OptimismChainName: "0x1f98431c8ad98523631ae4a59f267346ea31f984",
	chainId.ArbitrumChainName: "0x1f98431c8ad98523631ae4a59f267346ea31f984",
}
