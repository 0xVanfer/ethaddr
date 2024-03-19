package ethaddr

import "github.com/0xVanfer/chainId"

const UniswapProtocolV3 = "uniswapv3"

// Uniswap V3 swap router.
//
// map[network] = address.
var UniswapV3RouterList = map[string]string{
	chainId.EthereumChainName: "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
	chainId.GoerliChainName:   "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
	chainId.PolygonChainName:  "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
	chainId.OptimismChainName: "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
	chainId.ArbitrumChainName: "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
}

// Uniswap v3 factory.
//
// map[network] = address.
var UniswapV3FactoryList = map[string]string{
	chainId.EthereumChainName: "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
	chainId.GoerliChainName:   "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
	chainId.PolygonChainName:  "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
	chainId.OptimismChainName: "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
	chainId.ArbitrumChainName: "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
}
