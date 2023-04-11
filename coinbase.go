package ethaddr

import "github.com/0xVanfer/chainId"

// Coinbase cbETH.
//
// map[network] = address.
var CbETHList = map[string]string{
	chainId.EthereumChainName: "0xbe9895146f7af43049ca1c1ae358b0541ea49704",
	chainId.ArbitrumChainName: "0x1debd73e752beaf79865fd6446b0c970eae7732f",
}