package ethaddr

// Website: https://app.uniswap.org/swap
const UniswapProtocolV3 = "uniswapv3"

// Uniswap V3 swap router.
//
// map[network] = address.
var UniswapV3RouterList = map[int64]string{
	ChainEthereum: "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
	ChainGoerli:   "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
	ChainPolygon:  "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
	ChainOptimism: "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
	ChainArbitrum: "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
}

// Uniswap v3 factory.
//
// map[network] = address.
var UniswapV3FactoryList = map[int64]string{
	ChainEthereum: "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
	ChainGoerli:   "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
	ChainPolygon:  "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
	ChainOptimism: "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
	ChainArbitrum: "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
}
