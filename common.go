package ethaddr

import (
	"strings"

	"github.com/0xVanfer/chainId"
)

// Get the symbol from the address.
//
// map[network][address] = symbol
func GetCommonTokensAddressSymbolMap() map[string]map[string]string {
	mapp := make(map[string]map[string]string)
	// Commonly used networks.
	networks := []string{
		chainId.EthereumChainName,
		chainId.PolygonChainName,
		chainId.ArbitrumChainName,
		chainId.AvalancheChainName,
		chainId.BinanceSmartChainName,
		chainId.OptimismChainName,
	}
	// Commonly used tokens.
	tokenMaps := map[string]map[string]string{
		"USDT":   USDTList,
		"USDT.e": USDTeList,
		"USDC":   USDCList,
		"USDC.e": USDCeList,
		"DAI":    DAIList,

		"FRAX": FRAXList,
		"TUSD": TUSDList,
		"MIM":  MIMList,
		"LUSD": LUSDList,
		"BUSD": BUSDList,

		"WETH":   WETHList,
		"stETH":  STETHList,
		"wstETH": WSTETHList,
		"cbETH":  CbETHList,
		"rETH":   RETHList,

		"WBTC":  WBTCList,
		"BTC.b": BTCbList,

		"WAVAX": WAVAXList,
		"sAVAX": SAVAXList,

		"WMATIC":  WMATICList,
		"MATICX":  MaticXList,
		"stMATIC": STMATICList,

		"WBNB": WBNBList,

		"1INCH": OneInchTokenList,
		"AAVE":  AAVEList,
		"BAL":   BALList,
		"CRV":   CRVList,
		"LINK":  LINKList,
		"MKR":   MKRList,
		"SNX":   SNXList,
		"UNI":   UNIList,
	}

	for _, network := range networks {
		chainMapp := make(map[string]string)
		for symbol, tokenMap := range tokenMaps {
			// The token map did not record this chain.
			if tokenMap[network] == "" {
				continue
			}
			chainMapp[strings.ToLower(tokenMap[network])] = symbol
		}
		chainMapp[EtherAddress] = chainId.ChainTokenSymbolList[network]
		mapp[network] = chainMapp
	}
	return mapp
}
