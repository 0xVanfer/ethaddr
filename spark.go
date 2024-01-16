package ethaddr

import "github.com/0xVanfer/chainId"

// Website: https://app.spark.fi/markets/
const SparkProtocol string = "spark"

// Spark Saving DAI address.
//
// map[network] = address.
var SDAIList = map[string]string{
	chainId.EthereumChainName: "0x83f20f44975d03b1b09e64809b757c47f942beea",
}

var SparkPoolList = map[string]string{
	chainId.EthereumChainName: "0xc13e21b648a5ee794902342038ff3adab66be987",
}

var SparkPoolAddressProviderList = map[string]string{
	chainId.EthereumChainName: "0x02c3ea4e34c0cbd694d2adfa2c690eecbc1793ee",
}

var SparkDataProviderList = map[string]string{
	chainId.EthereumChainName: "0xfc21d6d146e6086b8359705c8b28512a983db0cb",
}

var SparkTreasuryList = map[string]string{
	chainId.EthereumChainName: "0xb137e7d16564c81ae2b0c8ee6b55de81dd46ece5",
}

var SparkTreasuryControllerList = map[string]string{
	chainId.EthereumChainName: "0x92ef091c5a1e01b3ce1ba0d0150c84412d818f7a",
}

var SparkIncentivesList = map[string]string{
	chainId.EthereumChainName: "0x4370d3b6c9588e02ce9d22e684387859c7ff5b34",
}

// Spark a tokens.
//
// map[network][underlying] = address.
var SparkATokenList = map[string]map[string]string{
	chainId.EthereumChainName: {
		DAIList[chainId.EthereumChainName]:    "0x4dedf26112b3ec8ec46e7e31ea5e123490b05b8b", // spDAI
		USDCList[chainId.EthereumChainName]:   "0x377c3bd93f2a2984e1e7be6a5c22c525ed4a4815",
		USDTList[chainId.EthereumChainName]:   "0xe7df13b8e3d6740fe17cbe928c7334243d86c92f",
		RETHList[chainId.EthereumChainName]:   "0x9985df20d7e9103ecbceb16a84956434b6f06ae8",
		SDAIList[chainId.EthereumChainName]:   "0x78f897f0fe2d3b5690ebae7f19862deacedf10a7",
		WBTCList[chainId.EthereumChainName]:   "0x4197ba364ae6698015ae5c1468f54087602715b2",
		WETHList[chainId.EthereumChainName]:   "0x59cd1c87501baa753d0b5b5ab5d8416a45cd71db",
		WSTETHList[chainId.EthereumChainName]: "0x12b54025c112aa61face2cdb7118740875a566e9",
	},
}

// Spark v tokens.
//
// map[network][underlying] = address.
var SparkVTokenList = map[string]map[string]string{
	chainId.EthereumChainName: {
		DAIList[chainId.EthereumChainName]:    "0xf705d2b7e92b3f38e6ae7afadaa2fee110fe5914", // variableDebtDAI
		USDCList[chainId.EthereumChainName]:   "0x7b70d04099cb9cfb1db7b6820badafb4c5c70a67",
		USDTList[chainId.EthereumChainName]:   "0x529b6158d1d2992e3129f7c69e81a7c677dc3b12",
		RETHList[chainId.EthereumChainName]:   "0xba2c8f2ea5b56690bfb8b709438f049e5dd76b96",
		SDAIList[chainId.EthereumChainName]:   "0xabc57081c04d921388240393ec4088aa47c6832b",
		WBTCList[chainId.EthereumChainName]:   "0xf6fee3a8ac8040c3d6d81d9a4a168516ec9b51d2",
		WETHList[chainId.EthereumChainName]:   "0x2e7576042566f8d6990e07a1b61ad1efd86ae70d",
		WSTETHList[chainId.EthereumChainName]: "0xd5c3e3b566a42a6110513ac7670c1a86d76e13e6",
	},
}

// Spark s tokens.
var SparkSTokenList = map[string]map[string]string{}
