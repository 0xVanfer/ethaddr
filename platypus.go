package ethaddr

import "github.com/0xVanfer/chainId"

const PlatypusProtocol string = "platypus"

// Platypus token: PTP.
//
// map[network] = address.
var PlatypusTokenList = map[string]string{
	chainId.AvalancheChainName: "0x22d4002028f537599be9f666d1c4fa138522f9c8", // PTP
}

// Same as PlatypusTokenList.
var PTPList = PlatypusTokenList

// Platypus ve token: vePTP.
//
// map[network] = address.
var PlatypusvePTPList = map[string]string{
	chainId.AvalancheChainName: "0x5857019c749147eee22b1fe63500f237f3c1b692", // vePTP
}

// Platypus xPTP.
//
// map[network] = address.
var PlatypusxPTPList = map[string]string{
	chainId.AvalancheChainName: "0x060556209e507d30f2167a101bfc6d256ed2f3e1", // xPTP
}

// Platypus pair names.
const (
	// Platypus main pools, including USDC, USDT, USDC.e, USDT.e, DAI.e, BUSD, MIM(deprecated).
	PlatypusMainPoolsName string = "PlatypusMainPools"
	// Platypus USDC - FRAX alt pool.
	Platypus_USDC_FRAX_PairName string = "Platypus_USDC_FRAX_pair"
	// Platypus USDC - MIM alt pool.
	Platypus_USDC_MIM_PairName string = "Platypus_USDC_MIM_pair"
	// Platypus USDC - YUSD alt pool.
	Platypus_USDC_YUSD_PairName string = "Platypus_USDC_YUSD_pair"
	// Platypus USDC - TUSD alt pool.(deprecated)
	Platypus_USDC_TUSD_PairName string = "Platypus_USDC_TUSD_pair"
	// Platypus USDC - UST alt pool.(deprecated)
	Platypus_USDC_UST_PairName string = "Platypus_USDC_UST_pair"
	// Platypus USDC - MONEY alt pool.
	Platypus_USDC_MONEY_PairName string = "Platypus_USDC_MONEY_pair"
	// Platypus AVAX - sAVAX alt pool.
	Platypus_AVAX_sAVAX_PairName string = "Platypus_AVAX_sAVAX_pair"
	// Platypus AVAX - yyAVAX alt pool.
	Platypus_AVAX_yyAVAX_PairName string = "Platypus_AVAX_yyAVAX_pair"
	// Platypus BTC.b - WBTC.e alt pool.
	Platypus_BTCb_WBTCe_PairName string = "Platypus_BTCb_WBTCe_pair"
)

// Platypus router.
//
// map[network][pool name] = address.
var PlatypusRouterList = map[string]map[string]string{
	chainId.AvalancheChainName: {
		// main
		PlatypusMainPoolsName: "0x66357dcace80431aee0a7507e2e361b7e2402370",

		// alt
		Platypus_USDC_FRAX_PairName:  "0xb8e567fc23c39c94a1f6359509d7b43d1fbed824", // USDC - FRAX
		Platypus_USDC_MIM_PairName:   "0x30c30d826be87cd0a4b90855c2f38f7fcfe4eaa7", // USDC - MIM
		Platypus_USDC_TUSD_PairName:  "0x13329c7905f1ee55c3c7d7bfc26c1197c512c207", // USDC - TUSD(deprecated)
		Platypus_USDC_YUSD_PairName:  "0xc828d995c686aaba78a4ac89dfc8ec0ff4c5be83", // USDC - YUSD
		Platypus_USDC_MONEY_PairName: "0x27912ae6ba9a54219d8287c3540a8969ff35500b", // USDC - MONEY

		Platypus_AVAX_sAVAX_PairName:  "0x4658ea7e9960d6158a261104aaa160cc953bb6ba", // AVAX - sAVAX
		Platypus_AVAX_yyAVAX_PairName: "0x8b4a45da5b0705ae4f47ebefc180c099345cf57e", // AVAX - yyAVAX

		Platypus_BTCb_WBTCe_PairName: "0x39de4e02f76dbd4352ec2c926d8d64db8abdf5b2", // BTC.b - WBTC.e
	},
}

// Platypus master platypus v2.
//
// map[network] = address.
var PlatypusMasterPlatypusV2List = map[string]string{
	chainId.AvalancheChainName: "0x68c5f4374228beedfa078e77b5ed93c28a2f713e",
}

// Platypus master platypus v4.
//
// map[network] = address.
var PlatypusMasterPlatypusV4List = map[string]string{
	chainId.AvalancheChainName: "0xff6934aac9c94e1c39358d4fdcf70aeca77d0ab0",
}

// Platypus LP Tokens.
const (
	// Platypus main pool - USDC.e LP.
	PlatypusLpUSDCeAddress string = "0x909b0ce4fac1a0dca78f8ca7430bbafeeca12871"
	// Platypus main pool - USDT.e LP.
	PlatypusLpUSDTeAddress string = "0x0d26d103c91f63052fbca88aaf01d5304ae40015"
	// Platypus main pool - USDC LP.
	PlatypusLpUSDCAddress string = "0xaef735b1e7ecfaf8209ea46610585817dc0a2e16"
	// Platypus main pool - USDT LP.
	PlatypusLpUSDTAddress string = "0x776628a5c37335608dd2a9538807b9bba3869e14"
	// Platypus main pool - DAI.e LP.
	PlatypusLpDAIeAddress string = "0xc1daa16e6979c2d1229cb1fd0823491ea44555be"
	// Platypus main pool - BUSD LP.
	PlatypusLpBUSDAddress string = "0xe23f8ccdeb4e8ce5d9fe76782718cd85d12689c8"
	// Platypus main pool - MIM LP.(deprecated)
	PlatypusLpMIMAddress string = "0x6220baad9d08dee465befae4f82ee251cf7c8b82"

	// Platypus USDC - FRAX pool - FRAX LP.
	PlatypusLp_USDC_FRAX_FRAX_Address string = "0x6fd4b4c38ed80727ecd0d58505565f9e422c965f"
	// Platypus USDC - FRAX pool - USDC LP.
	PlatypusLp_USDC_FRAX_USDC_Address string = "0x035d7d7f209b5d18e2ab5c2072e85b32e1d43760"

	// Platypus USDC - UST pool - UST LP.(deprecated)
	PlatypusLp_USDC_UST_UST_Address string = "0xc7388d98fa86b6639d71a0a6d410d5cdfc63a1d0"
	// Platypus USDC - UST pool - USDC LP.(deprecated)
	PlatypusLp_USDC_UST_USDC_Address string = "0xfc95481f79ec965a535ed8cef4630e1dd308d319"

	// Platypus USDC - MIM pool - MIM LP.
	PlatypusLp_USDC_MIM_MIM_Address string = "0xf01cea00598d87cb9792a01b040d04b0bd8ca781"
	// Platypus USDC - MIM pool - USDC LP.
	PlatypusLp_USDC_MIM_USDC_Address string = "0x4e5704991b43c1d33b9ccd1bc33b211bf068385a"

	// Platypus USDC - TUSD pool - TUSD LP.(deprecated)
	PlatypusLp_USDC_TUSD_TUSD_Address string = "0xc75b2b90079492922af96ba53988d7b384158335"
	// Platypus USDC - TUSD pool - USDC LP.(deprecated)
	PlatypusLp_USDC_TUSD_USDC_Address string = "0xa551480dc5399921f8a73e02ec327f2fd7e5ddc0"

	// Platypus USDC - YUSD pool - YUSD LP.
	PlatypusLp_USDC_YUSD_YUSD_Address string = "0x7716307350c0819ed05c3e7f6c478b27caed5361"
	// Platypus USDC - YUSD pool - USDC LP.
	PlatypusLp_USDC_YUSD_USDC_Address string = "0x4b851118a4a4948799f24d0cbe17fa3dad09e2d5"

	// Platypus USDC - MONEY pool - MONEY LP.
	PlatypusLp_USDC_MONEY_MONEY_Address string = "0xe08947ee864af325d9f98743b3b905875ae0ec99"
	// Platypus USDC - MONEY pool - USDC LP.
	PlatypusLp_USDC_MONEY_USDC_Address string = "0x551c259bf4d88edfdabb04179342a73daa759583"

	// Platypus AVAX - sAVAX pool - AVAX LP.
	PlatypusLp_AVAX_sAVAX_AVAX_Address string = "0xc73eed4494382093c6a7c284426a9a00f6c79939"
	// Platypus AVAX - sAVAX pool - sAVAX LP.
	PlatypusLp_AVAX_sAVAX_sAVAX_Address string = "0xa2a7ee49750ff12bb60b407da2531db3c50a1789"

	// Platypus AVAX - yyAVAX pool - AVAX LP.
	PlatypusLp_AVAX_yyAVAX_AVAX_Address string = "0x12141b8fd20b4bbdd5f4e911bf91575258a3eabd"
	// Platypus AVAX - yyAVAX pool - yyAVAX LP.
	PlatypusLp_AVAX_yyAVAX_yyAVAX_Address string = "0x3beb0d3db537b79d377131ce81950b683d382ec9"

	// Platypus BTC.b - WBTC.e pool - BTC.b LP.
	PlatypusLp_BTCb_WBTCe_BTCb_Address string = "0x209a0399a2905900c0d1a9a382fe23e37024dc84"
	// Platypus BTC.b - WBTC.e pool - WBTC.e LP.
	PlatypusLp_BTCb_WBTCe_WBTCe_Address string = "0xc09c12093b037866bf68c9474ecdb5113160fbce"
)

// Platypus lp list.
//
// map[network][pool name][underlying] = address.
var PlatypusLpList = map[string]map[string]map[string]string{
	chainId.AvalancheChainName: {
		// main pools
		PlatypusMainPoolsName: {
			DAIeList[chainId.AvalancheChainName]:  PlatypusLpDAIeAddress,
			USDCList[chainId.AvalancheChainName]:  PlatypusLpUSDCAddress,
			USDCeList[chainId.AvalancheChainName]: PlatypusLpUSDCeAddress,
			USDTList[chainId.AvalancheChainName]:  PlatypusLpUSDTAddress,
			USDTeList[chainId.AvalancheChainName]: PlatypusLpUSDTeAddress,
			BUSDList[chainId.AvalancheChainName]:  PlatypusLpBUSDAddress,
		},

		// USDC - FRAX
		Platypus_USDC_FRAX_PairName: {
			USDCList[chainId.AvalancheChainName]: PlatypusLp_USDC_FRAX_USDC_Address,
			FRAXList[chainId.AvalancheChainName]: PlatypusLp_USDC_FRAX_FRAX_Address,
		},
		// USDC - MIM
		Platypus_USDC_MIM_PairName: {
			USDCList[chainId.AvalancheChainName]: PlatypusLp_USDC_MIM_USDC_Address,
			MIMList[chainId.AvalancheChainName]:  PlatypusLp_USDC_MIM_MIM_Address,
		},
		// USDC - YUSD
		Platypus_USDC_YUSD_PairName: {
			USDCList[chainId.AvalancheChainName]: PlatypusLp_USDC_YUSD_USDC_Address,
			YUSDList[chainId.AvalancheChainName]: PlatypusLp_USDC_YUSD_YUSD_Address,
		},
		// USDC - MONEY
		Platypus_USDC_MONEY_PairName: {
			USDCList[chainId.AvalancheChainName]:  PlatypusLp_USDC_MONEY_USDC_Address,
			MoneyList[chainId.AvalancheChainName]: PlatypusLp_USDC_MONEY_MONEY_Address,
		},

		// BTC.b - WBTC.e
		Platypus_BTCb_WBTCe_PairName: {
			WBTCList[chainId.AvalancheChainName]: PlatypusLp_BTCb_WBTCe_WBTCe_Address,
			BTCbList[chainId.AvalancheChainName]: PlatypusLp_BTCb_WBTCe_BTCb_Address,
		},

		// AVAX - sAVAX
		Platypus_AVAX_sAVAX_PairName: {
			WAVAXList[chainId.AvalancheChainName]: PlatypusLp_AVAX_sAVAX_AVAX_Address,
			SAVAXList[chainId.AvalancheChainName]: PlatypusLp_AVAX_sAVAX_sAVAX_Address,
		},
		// AVAX - yyAVAX
		Platypus_AVAX_yyAVAX_PairName: {
			WAVAXList[chainId.AvalancheChainName]:  PlatypusLp_AVAX_yyAVAX_AVAX_Address,
			YyAVAXList[chainId.AvalancheChainName]: PlatypusLp_AVAX_yyAVAX_yyAVAX_Address,
		},
		"PlatypusDeprecated": {
			"mim": PlatypusLpMIMAddress,

			"usdc_tusd_tusd": PlatypusLp_USDC_TUSD_TUSD_Address,
			"usdc_tusd_usdc": PlatypusLp_USDC_TUSD_USDC_Address,

			"usdc_ust_usdc": PlatypusLp_USDC_UST_USDC_Address,
			"usdc_ust_ust":  PlatypusLp_USDC_UST_UST_Address,
		},
	},
}
