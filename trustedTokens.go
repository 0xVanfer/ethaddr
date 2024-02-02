package ethaddr

import "github.com/0xVanfer/chainId"

// Deprecated: Remove.
type tokenInfo struct {
	Address  string
	Symbol   string
	Decimals int
}

// Deprecated: Remove.
//
// Trusted tokens info.
//
// map[network] = []tokenInfo{Address:"xx", Symbol:"xx", Decimals:xx}
var TrustedTokenList = map[string][]tokenInfo{
	chainId.EthereumChainName: {
		{
			Address:  USDCList[chainId.EthereumChainName],
			Symbol:   "USDC",
			Decimals: 6,
		},
		{
			Address:  USDTList[chainId.EthereumChainName],
			Symbol:   "USDT",
			Decimals: 6,
		},
		{
			Address:  DAIList[chainId.EthereumChainName],
			Symbol:   "DAI",
			Decimals: 18,
		},
		{
			Address:  WETHList[chainId.EthereumChainName],
			Symbol:   "WETH",
			Decimals: 18,
		},
		{
			Address:  WBTCList[chainId.EthereumChainName],
			Symbol:   "WBTC",
			Decimals: 8,
		},
	},
	chainId.AvalancheChainName: {
		{
			Address:  USDCList[chainId.AvalancheChainName],
			Symbol:   "USDC",
			Decimals: 6,
		},
		{
			Address:  USDTList[chainId.AvalancheChainName],
			Symbol:   "USDt",
			Decimals: 6,
		},
		{
			Address:  USDCeList[chainId.AvalancheChainName],
			Symbol:   "USDC.e",
			Decimals: 6,
		},
		{
			Address:  USDTeList[chainId.AvalancheChainName],
			Symbol:   "USDT.e",
			Decimals: 6,
		},
		{
			Address:  DAIeList[chainId.AvalancheChainName],
			Symbol:   "DAI.e",
			Decimals: 18,
		},
		{
			Address:  WETHList[chainId.AvalancheChainName],
			Symbol:   "WETH.e",
			Decimals: 18,
		},
		{
			Address:  WBTCList[chainId.AvalancheChainName],
			Symbol:   "WBTC.e",
			Decimals: 8,
		},
	},
	chainId.PolygonChainName: {
		{
			Address:  USDCList[chainId.PolygonChainName],
			Symbol:   "USDC",
			Decimals: 6,
		},
		{
			Address:  USDTList[chainId.PolygonChainName],
			Symbol:   "USDT",
			Decimals: 6,
		},
		{
			Address:  DAIList[chainId.PolygonChainName],
			Symbol:   "DAI",
			Decimals: 18,
		},
		{
			Address:  WETHList[chainId.PolygonChainName],
			Symbol:   "WETH",
			Decimals: 18,
		},
		{
			Address:  WBTCList[chainId.PolygonChainName],
			Symbol:   "WBTC",
			Decimals: 8,
		},
	},
}
