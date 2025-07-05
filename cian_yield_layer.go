package ethaddr

// Website: https://cian.app/
var CianProtocol string = "cian"

var CianYieldLayerstETHList = map[int64]string{
	ChainEthereum: "0xB13aa2d0345b0439b064f26B82D8dCf3f508775d", // ylstETH, 0xb13aa2d0345b0439b064f26b82d8dcf3f508775d
}

var CianYieldLayerstETHBridgedFromEthereumList = map[int64]string{
	ChainScroll: "0xBAC6DD1b1F186EF7cf4d64737235a9C53878cB27", // ylstETH, 0xbac6dd1b1f186ef7cf4d64737235a9c53878cb27
}

var CianStakedYieldLayerstETHList = map[int64]string{
	ChainEthereum: "0x3498fDed9C88Ae83b3BC6a302108F2da408e613b", // sylstETH, 0x3498fded9c88ae83b3bc6a302108f2da408e613b
}

var CianYieldLayerrsETHList = map[int64]string{
	ChainEthereum: "0xd87a19fF681AE98BF10d2220D1AE3Fbd374ADE4e", // ylrsETH, 0xd87a19ff681ae98bf10d2220d1ae3fbd374ade4e
	ChainArbitrum: "0x15cbFF12d53e7BdE3f1618844CaaEf99b2836d2A", // ylrsETH, 0x15cbff12d53e7bde3f1618844caaef99b2836d2a
}

var CianYieldLayerezETHList = map[int64]string{
	ChainEthereum: "0x3D086B688D7c0362BE4f9600d626f622792c4a20", // ylezETH, 0x3d086b688d7c0362be4f9600d626f622792c4a20
}

var CianYieldLayeruniBTCList = map[int64]string{
	ChainEthereum: "0xcc7E6dE27DdF225E24E8652F62101Dab4656E20A", // yluniBTC, 0xcc7e6de27ddf225e24e8652f62101dab4656e20a
}

var CianYieldLayerpumpBTCList = map[int64]string{
	ChainEthereum: "0xd4Cc9b31e9eF33E392FF2f81AD52BE8523e0993b", // ylpumpBTC, 0xd4cc9b31e9ef33e392ff2f81ad52be8523e0993b
}

var CianYieldLayercbBTCList = map[int64]string{
	ChainEthereum: "0x7BA7c46e9F44d93AEF0Ddd37b80134438f60e15e", // ylcbBTC, 0x7ba7c46e9f44d93aef0ddd37b80134438f60e15e
}

var CianYieldLayerFBTCAntalphaList = map[int64]string{
	ChainEthereum: "0x6c77bdE03952BbcB923815d90A73a7eD7EC895D1", // ylBTCLST, 0x6c77bde03952bbcb923815d90a73a7ed7ec895d1
}

// Cian yield layer stETH, ylstETH.
//
// map[chainID] = address.
var YLstETHList = CianYieldLayerstETHList

// Cian staked yield layer stETH, sylstETH.
//
// map[chainID] = address.
var SYLstETHList = CianStakedYieldLayerstETHList

// Cian yield layer rsETH, ylrsETH.
//
// map[chainID] = address.
var YLrsETHList = CianYieldLayerrsETHList

// Cian yield layer stETH `Bridged From Ethereum`, ylstETH.
//
// map[chainID] = address.
var YLstETHBFEList = CianYieldLayerstETHBridgedFromEthereumList

// Cian yield layer ezETH, ylezETH.
//
// map[chainID] = address.
var YLezETHList = CianYieldLayerezETHList

// Cian yield layer uniBTC, yluniBTC.
//
// map[chainID] = address.
var YLuniBTCList = CianYieldLayeruniBTCList

// Cian yield layer pumpBTC, ylpumpBTC.
//
// map[chainID] = address.
var YLpumpBTCList = CianYieldLayerpumpBTCList

// Cian yield layer cbBTC, ylcbBTC.
//
// map[chainID] = address.
var YLcbBTCList = CianYieldLayercbBTCList

// Cian yield layer FBTC for AntAlpha, ylBTCLST.
//
// map[chainID] = address.
var YLBTCLSTList = CianYieldLayerFBTCAntalphaList
