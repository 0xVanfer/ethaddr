package ethaddr

// Website: https://app.uniswap.org/swap
//
// GitHub: https://github.com/Uniswap/v3-core
//
// Deployed addresses: https://docs.uniswap.org/contracts/v3/reference/deployments/ethereum-deployments
const UniswapProtocolV3 = "uniswapv3"

// Uniswap V3 swap router.
//
// map[network] = address.
var UniswapV3RouterList = map[int64]string{
	ChainArbitrum: "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
	ChainEthereum: "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
	ChainGoerli:   "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
	ChainOptimism: "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
	ChainPolygon:  "0xE592427A0AEce92De3Edee1F18E0157C05861564", // 0xe592427a0aece92de3edee1f18e0157c05861564
}

// Uniswap V3 swap router 02.
//
// map[network] = address.
var UniswapV3Router02List = map[int64]string{
	ChainAvalanche:  "0xbb00FF08d01D300023C629E8fFfFcb65A5a578cE", // 0xbb00ff08d01d300023c629e8ffffcb65a5a578ce
	ChainBase:       "0x2626664c2603336E57B271c5C0b26F421741e481", // 0x2626664c2603336e57b271c5c0b26f421741e481
	ChainBlast:      "0x549FEB8c9bd4c12Ad2AB27022dA12492aC452B66", // 0x549feb8c9bd4c12ad2ab27022da12492ac452b66
	ChainBSC:        "0xB971eF87ede563556b2ED4b1C0b0019111Dd85d2", // 0xb971ef87ede563556b2ed4b1c0b0019111dd85d2
	ChainEthereum:   "0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45", // 0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45
	ChainWorldChain: "0x091AD9e2e6e5eD44c1c66dB50e49A601F9f36cF6", // 0x091ad9e2e6e5ed44c1c66db50e49a601f9f36cf6
	ChainZkSync:     "0x99c56385daBCE3E81d8499d0b8d0257aBC07E8A3", // 0x99c56385dabce3e81d8499d0b8d0257abc07e8a3
}

// Uniswap v3 factory.
//
// map[network] = address.
var UniswapV3FactoryList = map[int64]string{
	ChainArbitrum:   "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
	ChainAvalanche:  "0x740b1c1de25031C31FF4fC9A62f554A55cdC1baD", // 0x740b1c1de25031c31ff4fc9a62f554a55cdc1bad
	ChainBase:       "0x33128a8fC17869897dcE68Ed026d694621f6FDfD", // 0x33128a8fc17869897dce68ed026d694621f6fdfd
	ChainBlast:      "0x792edAdE80af5fC680d96a2eD80A44247D2Cf6Fd", // 0x792edade80af5fc680d96a2ed80a44247d2cf6fd
	ChainBSC:        "0xdB1d10011AD0Ff90774D0C6Bb92e5C5c8b4461F7", // 0xdb1d10011ad0ff90774d0c6bb92e5c5c8b4461f7
	ChainEthereum:   "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
	ChainGoerli:     "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
	ChainOptimism:   "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
	ChainPolygon:    "0x1F98431c8aD98523631AE4a59f267346ea31F984", // 0x1f98431c8ad98523631ae4a59f267346ea31f984
	ChainWorldChain: "0x7a5028BDa40e7B173C278C5342087826455ea25a", // 0x7a5028bda40e7b173c278c5342087826455ea25a
	ChainZkSync:     "0x8FdA5a7a8dCA67BBcDd10F02Fa0649A937215422", // 0x8fda5a7a8dca67bbcdd10f02fa0649a937215422
}

// Uniswap v3 nonfungible position manager.
//
// map[network] = address.
var UniswapV3NonfungiblePositionManagerList = map[int64]string{
	ChainAvalanche:  "0x655C406EBFa14EE2006250925e54ec43AD184f8B", // 0x655c406ebfa14ee2006250925e54ec43ad184f8b
	ChainBase:       "0x03a520b32C04BF3bEEf7BEb72E919cf822Ed34f1", // 0x03a520b32c04bf3beef7beb72e919cf822ed34f1
	ChainBlast:      "0xB218e4f7cF0533d4696fDfC419A0023D33345F28", // 0xb218e4f7cf0533d4696fdfc419a0023d33345f28
	ChainBSC:        "0x7b8A01B39D58278b5DE7e48c8449c9f4F5170613", // 0x7b8a01b39d58278b5de7e48c8449c9f4f5170613
	ChainEthereum:   "0xC36442b4a4522E871399CD717aBDD847Ab11FE88", // 0xc36442b4a4522e871399cd717abdd847ab11fe88
	ChainWorldChain: "0xec12a9F9a09f50550686363766Cc153D03c27b5e", // 0xec12a9f9a09f50550686363766cc153d03c27b5e
	ChainZkSync:     "0x0616e5762c1E7Dc3723c50663dF10a162D690a86", // 0x0616e5762c1e7dc3723c50663df10a162d690a86
}
