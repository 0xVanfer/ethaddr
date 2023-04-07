package ethaddr

import "github.com/0xVanfer/chainId"

const QuickswapProtocol string = "quickswap"

// Quickswap token: Quick.
//
// map[network] = address.
var QuickswapTokenList = map[string]string{
	chainId.PolygonChainName:  "0xb5c064f955d8e7f38fe0460c556a72987494ee17", // QUICK
	chainId.ArbitrumChainName: "0x6c28e052ffdf1d9f46c284c2e2fe60c503246f2f", // QUICK
}

// Same as QuickswapTokenList.
var QUICKList = QuickswapTokenList

// Quickswap factory.
//
// map[network] = address.
var QuickswapFactoryList = map[string]string{
	chainId.PolygonChainName: "0x5757371414417b8c6caad45baef941abc7d3ab32",
}

// Quickswap router02.
//
// map[network] = address.
var QuickswapRouter02List = map[string]string{
	chainId.PolygonChainName: "0xa5e0829caced8ffdd4de3c43696c57f7d7a678ff",
}

// Quickswap router03.
//
// map[network] = address.
var QuickswapRouter03List = map[string]string{
	chainId.PolygonChainName: "0xf5b509bb0909a69b1c207e495f687a596c168e12",
}
