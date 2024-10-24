package ethaddr

// Website: https://traderjoexyz.com/ethereum
const TraderJoeProtocol string = "traderjoe"

// Traderjoe token: JOE.
//
// map[network] = address.
var TraderjoeTokenList = map[int64]string{
	ChainArbitrum:  "0x371c7ec6D8039ff7933a2AA28EB827Ffe1F52f07", // JOE, 0x371c7ec6d8039ff7933a2aa28eb827ffe1f52f07
	ChainAvalanche: "0x6e84a6216eA6dACC71eE8E6b0a5B7322EEbC0fDd", // JOE, 0x6e84a6216ea6dacc71ee8e6b0a5b7322eebc0fdd
}

// Same as TraderjoeTokenList.
var JOEList = TraderjoeTokenList

// Traderjoe zJOE.
//
// map[network] = address.
var TraderzJOEList = map[int64]string{
	ChainAvalanche: "0x769bfeb9fAacD6Eb2746979a8dD0b7e9920aC2A4", // zJOE, 0x769bfeb9faacd6eb2746979a8dd0b7e9920ac2a4
}

// Traderjoe stake token: xJOE.(abandoned)
//
// map[network] = address.
var TraderjoeStakeTokenList = map[int64]string{
	ChainAvalanche: "0x57319d41F71E81F3c65F2a47CA4e001EbAFd4F33", // xJOE, 0x57319d41f71e81f3c65f2a47ca4e001ebafd4f33
}

// Same as TraderjoeStakeTokenList.
var XJOEList = TraderjoeStakeTokenList

// Traderjoe router.
//
// map[network] = address.
var TraderjoeRouterList = map[int64]string{
	ChainAvalanche: "0x60aE616a2155Ee3d9A68541Ba4544862310933d4", // 0x60ae616a2155ee3d9a68541ba4544862310933d4
}

// Traderjoe joetroller.
//
// map[network] = address.
var TraderjoeJoetrollerList = map[int64]string{
	ChainAvalanche: "0xdc13687554205E5b89Ac783db14bb5bba4A1eDaC", // 0xdc13687554205e5b89ac783db14bb5bba4a1edac
}

// Traderjoe factory.
//
// map[network] = address.
var TraderjoeFactoryList = map[int64]string{
	ChainAvalanche: "0x9Ad6C38BE94206cA50bb0d90783181662f0Cfa10", // 0x9ad6c38be94206ca50bb0d90783181662f0cfa10
}

// Traderjoe reward distributor.
//
// map[network] = address.
var TraderjoeRewardDistributorList = map[int64]string{
	ChainAvalanche: "0x45B2C4139d96F44667577C0D7F7a7D170B420324", // 0x45b2c4139d96f44667577c0d7f7a7d170b420324
}

// Traderjoe chef v3.
//
// map[network] = address.
var TraderjoeChefV3List = map[int64]string{
	ChainAvalanche: "0x188bED1968b795d5c9022F6a0bb5931Ac4c18F00", // 0x188bed1968b795d5c9022f6a0bb5931ac4c18f00
}

// Traderjoe veJOE boost chef.
//
// map[network] = address.
var TraderjoeVeBoostChef = map[int64]string{
	ChainAvalanche: "0x4483f0b6e2F5486D06958C20f8C39A7aBe87bf8F", // 0x4483f0b6e2f5486d06958c20f8c39a7abe87bf8f
}

// Traderjoe liquidity book factory.
//
// map[network] = address.
var TraderjoeLBFactoryList = map[int64]string{
	ChainAvalanche: "0x6E77932A92582f504FF6c4BdbCef7Da6c198aEEf", // 0x6e77932a92582f504ff6c4bdbcef7da6c198aeef
}

// Traderjoe liquidity book router.
//
// map[network] = address.
var TraderjoeLBRouterList = map[int64]string{
	ChainAvalanche: "0xE3Ffc583dC176575eEA7FD9dF2A7c65F7E23f4C3", // 0xe3ffc583dc176575eea7fd9df2a7c65f7e23f4c3
}

// Traderjoe liquidity book qouter.
//
// map[network] = address.
var TraderjoeLBQuoterList = map[int64]string{
	ChainAvalanche: "0x51409fD32C15E9857eeF9F28275FA0842C3C4545", // 0x51409fd32c15e9857eef9f28275fa0842c3c4545
}

// Traderjoe c tokens (j tokens).
//
// map[network][underlying] = address.
var TraderjoeCTokenList = map[int64]map[string]string{
	ChainAvalanche: {
		DAIeList[ChainAvalanche]:  "0xc988c170d0E38197DC634A45bF00169C7Aa7CA19", // jDAI, 0xc988c170d0e38197dc634a45bf00169c7aa7ca19
		USDCeList[ChainAvalanche]: "0xEd6AaF91a2B084bd594DBd1245be3691F9f637aC", // jUSDC, 0xed6aaf91a2b084bd594dbd1245be3691f9f637ac
		USDCList[ChainAvalanche]:  "0x29472D511808Ce925F501D25F9Ee9efFd2328db2", // jUSDCNative, 0x29472d511808ce925f501d25f9ee9effd2328db2
		USDTeList[ChainAvalanche]: "0x8b650e26404AC6837539ca96812f0123601E4448", // jUSDT, 0x8b650e26404ac6837539ca96812f0123601e4448
		WETHeList[ChainAvalanche]: "0x929f5caB61DFEc79a5431a7734a68D714C4633fa", // jWETH, 0x929f5cab61dfec79a5431a7734a68d714c4633fa
		WBTCeList[ChainAvalanche]: "0x3fE38b7b610C0ACD10296fEf69d9b18eB7a9eB1F", // jWBTC, 0x3fe38b7b610c0acd10296fef69d9b18eb7a9eb1f
		WAVAXList[ChainAvalanche]: "0xC22F01ddc8010Ee05574028528614634684EC29e", // jAVAX, 0xc22f01ddc8010ee05574028528614634684ec29e
		LINKeList[ChainAvalanche]: "0x585E7bC75089eD111b656faA7aeb1104F5b96c15", // jLINK, 0x585e7bc75089ed111b656faa7aeb1104f5b96c15
		MIMList[ChainAvalanche]:   "0xcE095A9657A02025081E0607c8D8b081c76A75ea", // jMIM, 0xce095a9657a02025081e0607c8d8b081c76a75ea
		XJOEList[ChainAvalanche]:  "0xC146783a59807154F92084f9243eb139D58Da696", // jXJOE, 0xc146783a59807154f92084f9243eb139d58da696
		JOEList[ChainAvalanche]:   "0xBFdbE35168953c9d29bDF9a0043F902F233c76e0", // jJOE, 0xbfdbe35168953c9d29bdf9a0043f902f233c76e0
		BTCbList[ChainAvalanche]:  "0x13A7a6C167d75baDd316DDeef6C526c8463A090F", // jBTC, 0x13a7a6c167d75badd316ddeef6c526c8463a090f
		USDTList[ChainAvalanche]:  "0x50ac14A3Ee0a4cb6Ef829F7ad65B2dA5493e99d2", // jUSDTNative, 0x50ac14a3ee0a4cb6ef829f7ad65b2da5493e99d2
	},
}
