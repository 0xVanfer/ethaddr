package ethaddr

const UniswapProtocolV3 = "uniswapv3"

// Uniswap V3 swap router.
//
// map[network] = address.
var UniswapV3RouterList = map[string]string{
	chainEthereum: "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
	chainGoerli:   "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
	chainPolygon:  "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
	chainOptimism: "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
	chainArbitrum: "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
}

// Uniswap v3 factory.
//
// map[network] = address.
var UniswapV3FactoryList = map[string]string{
	chainEthereum: "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
	chainGoerli:   "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
	chainPolygon:  "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
	chainOptimism: "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
	chainArbitrum: "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
}
