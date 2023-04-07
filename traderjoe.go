package ethaddr

import "github.com/0xVanfer/chainId"

const TraderJoeProtocol string = "traderjoe"

// Traderjoe token: JOE.
//
// map[network] = address.
var TraderjoeTokenList = map[string]string{
	chainId.AvalancheChainName: "0x6e84a6216ea6dacc71ee8e6b0a5b7322eebc0fdd", // JOE
	chainId.ArbitrumChainName:  "0x371c7ec6d8039ff7933a2aa28eb827ffe1f52f07", // JOE
}

// Same as TraderjoeTokenList.
var JOEList = TraderjoeTokenList

// Traderjoe zJOE.
//
// map[network] = address.
var TraderzJOEList = map[string]string{
	chainId.AvalancheChainName: "0x769bfeb9faacd6eb2746979a8dd0b7e9920ac2a4", // zJOE
}

// Traderjoe stake token: xJOE.(abandoned)
//
// map[network] = address.
var TraderjoeStakeTokenList = map[string]string{
	chainId.AvalancheChainName: "0x57319d41f71e81f3c65f2a47ca4e001ebafd4f33", // xJOE
}

// Same as TraderjoeStakeTokenList.
var XJOEList = TraderjoeStakeTokenList

// Traderjoe router.
//
// map[network] = address.
var TraderjoeRouterList = map[string]string{
	chainId.AvalancheChainName: "0x60ae616a2155ee3d9a68541ba4544862310933d4",
}

// Traderjoe joetroller.
//
// map[network] = address.
var TraderjoeJoetrollerList = map[string]string{
	chainId.AvalancheChainName: "0xdc13687554205e5b89ac783db14bb5bba4a1edac",
}

// Traderjoe factory.
//
// map[network] = address.
var TraderjoeFactoryList = map[string]string{
	chainId.AvalancheChainName: "0x9ad6c38be94206ca50bb0d90783181662f0cfa10",
}

// Traderjoe reward distributor.
//
// map[network] = address.
var TraderjoeRewardDistributorList = map[string]string{
	chainId.AvalancheChainName: "0x45b2c4139d96f44667577c0d7f7a7d170b420324",
}

// Traderjoe chef v3.
//
// map[network] = address.
var TraderjoeChefV3List = map[string]string{
	chainId.AvalancheChainName: "0x188bed1968b795d5c9022f6a0bb5931ac4c18f00",
}

// Traderjoe veJOE boost chef.
//
// map[network] = address.
var TraderjoeVeBoostChef = map[string]string{
	chainId.AvalancheChainName: "0x4483f0b6e2f5486d06958c20f8c39a7abe87bf8f",
}

// Traderjoe liquidity book factory.
//
// map[network] = address.
var TraderjoeLBFactoryList = map[string]string{
	chainId.AvalancheChainName: "0x6e77932a92582f504ff6c4bdbcef7da6c198aeef",
}

// Traderjoe liquidity book router.
//
// map[network] = address.
var TraderjoeLBRouterList = map[string]string{
	chainId.AvalancheChainName: "0xe3ffc583dc176575eea7fd9df2a7c65f7e23f4c3",
}

// Traderjoe liquidity book qouter.
//
// map[network] = address.
var TraderjoeLBQuoterList = map[string]string{
	chainId.AvalancheChainName: "0x51409fd32c15e9857eef9f28275fa0842c3c4545",
}

// Traderjoe c tokens (j tokens).
//
// map[network][underlying] = address.
var TraderjoeCTokenList = map[string]map[string]string{
	chainId.AvalancheChainName: {
		DAIeList[chainId.AvalancheChainName]:  "0xc988c170d0E38197DC634A45bF00169C7Aa7CA19", // jDAI
		USDCeList[chainId.AvalancheChainName]: "0xEd6AaF91a2B084bd594DBd1245be3691F9f637aC", // jUSDC
		USDCList[chainId.AvalancheChainName]:  "0x29472d511808ce925f501d25f9ee9effd2328db2", // jUSDCNative
		USDTeList[chainId.AvalancheChainName]: "0x8b650e26404AC6837539ca96812f0123601E4448", // jUSDT
		WETHList[chainId.AvalancheChainName]:  "0x929f5caB61DFEc79a5431a7734a68D714C4633fa", // jWETH
		WBTCList[chainId.AvalancheChainName]:  "0x3fE38b7b610C0ACD10296fEf69d9b18eB7a9eB1F", // jWBTC
		WAVAXList[chainId.AvalancheChainName]: "0xC22F01ddc8010Ee05574028528614634684EC29e", // jAVAX
		LINKList[chainId.AvalancheChainName]:  "0x585E7bC75089eD111b656faA7aeb1104F5b96c15", // jLINK
		MIMList[chainId.AvalancheChainName]:   "0xcE095A9657A02025081E0607c8D8b081c76A75ea", // jMIM
		XJOEList[chainId.AvalancheChainName]:  "0xC146783a59807154F92084f9243eb139D58Da696", // jXJOE
		JOEList[chainId.AvalancheChainName]:   "0xbfdbe35168953c9d29bdf9a0043f902f233c76e0", // jJOE
		BTCbList[chainId.AvalancheChainName]:  "0x13a7a6c167d75badd316ddeef6c526c8463a090f", // jBTC
		USDTList[chainId.AvalancheChainName]:  "0x50ac14a3ee0a4cb6ef829f7ad65b2da5493e99d2", // jUSDTNative
	},
}
