package ethaddr

// Website: https://app.uniswap.org/swap
const UniswapProtocolV2 = "uniswapv2"

// Uniswap V2 swap router.
//
// map[chainID] = address.
//
// Uni V2 is only on ethereum.
var UniswapV2RouterList = map[int64]string{
	ChainEthereum: "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D", // 0x7a250d5630b4cf539739df2c5dacb4c659f2488d
}
