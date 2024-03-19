package ethaddr

import "github.com/0xVanfer/chainId"

const QuickswapProtocol string = "quickswap"

// Quickswap token: Quick.
//
// map[network] = address.
var QuickswapTokenList = map[string]string{
	chainId.ArbitrumChainName: "0x6c28e052FfDF1D9f46C284c2E2fE60c503246F2f", // QUICK, 0x6c28e052ffdf1d9f46c284c2e2fe60c503246f2f
	chainId.PolygonChainName:  "0xB5C064F955D8e7F38fE0460C556a72987494eE17", // QUICK, 0xb5c064f955d8e7f38fe0460c556a72987494ee17
}

// Same as QuickswapTokenList.
var QUICKList = QuickswapTokenList

// Quickswap factory.
//
// map[network] = address.
var QuickswapFactoryList = map[string]string{
	chainId.PolygonChainName: "0x5757371414417b8C6CAad45bAeF941aBc7d3Ab32", // 0x5757371414417b8c6caad45baef941abc7d3ab32
}

// Quickswap router02.
//
// map[network] = address.
var QuickswapRouter02List = map[string]string{
	chainId.PolygonChainName: "0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff", // 0xa5e0829caced8ffdd4de3c43696c57f7d7a678ff
}

// Quickswap router03.
//
// map[network] = address.
var QuickswapRouter03List = map[string]string{
	chainId.PolygonChainName: "0xf5b509bB0909a69B1c207E495f687a596C168E12", // 0xf5b509bb0909a69b1c207e495f687a596c168e12
}
