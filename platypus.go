package ethaddr

// WARNING: Platypus was dead and relaunched on Jan.2024.
// Last checked Feb.2025, Platypus has shut down, and deleted twitter account.

// Website: https://platypus.finance/
//
// X (twitter): https://twitter.com/Platypusdefi
const PlatypusProtocol string = "platypus"

// WARNING: Platypus has dead and relaunched. They might have deployed new contracts and the addresses here might be out dated.
//
// Platypus token: PTP.
//
// map[chainID] = address.
var PlatypusTokenList = map[int64]string{
	ChainAvalanche: "0x22d4002028f537599bE9f666d1c4Fa138522f9c8", // PTP, 0x22d4002028f537599be9f666d1c4fa138522f9c8
}

// Same as PlatypusTokenList.
var PTPList = PlatypusTokenList

// WARNING: Platypus has dead and relaunched. They might have deployed new contracts and the addresses here might be out dated.
//
// Platypus ve token: vePTP.
//
// map[chainID] = address.
var PlatypusvePTPList = map[int64]string{
	ChainAvalanche: "0x5857019c749147EEE22b1Fe63500F237F3c1B692", // vePTP, 0x5857019c749147eee22b1fe63500f237f3c1b692
}

// WARNING: Platypus has dead and relaunched. They might have deployed new contracts and the addresses here might be out dated.
//
// Platypus xPTP.
//
// map[chainID] = address.
var PlatypusxPTPList = map[int64]string{
	ChainAvalanche: "0x060556209E507d30f2167a101bFC6D256Ed2f3e1", // xPTP, 0x060556209e507d30f2167a101bfc6d256ed2f3e1
}

// WARNING: Platypus has dead and relaunched. They might have deployed new contracts and the addresses here might be out dated.
//
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

// WARNING: Platypus has dead and relaunched. They might have deployed new contracts and the addresses here might be out dated.
//
// Platypus router.
//
// map[network][pool name] = address.
var PlatypusRouterList = map[int64]map[string]string{
	ChainAvalanche: {
		// main
		PlatypusMainPoolsName: "0x66357dCaCe80431aee0A7507e2E361B7e2402370", // 0x66357dcace80431aee0a7507e2e361b7e2402370

		// alt
		Platypus_USDC_FRAX_PairName:  "0xB8E567fc23c39C94a1f6359509D7b43D1Fbed824", // USDC - FRAX, 0xb8e567fc23c39c94a1f6359509d7b43d1fbed824
		Platypus_USDC_MIM_PairName:   "0x30C30d826be87Cd0A4b90855C2F38f7FcfE4eaA7", // USDC - MIM, 0x30c30d826be87cd0a4b90855c2f38f7fcfe4eaa7
		Platypus_USDC_TUSD_PairName:  "0x13329C7905F1EE55c3C7D7Bfc26c1197c512c207", // USDC - TUSD, 0x13329c7905f1ee55c3c7d7bfc26c1197c512c207
		Platypus_USDC_YUSD_PairName:  "0xC828D995C686AaBA78A4aC89dfc8eC0Ff4C5be83", // USDC - YUSD, 0xc828d995c686aaba78a4ac89dfc8ec0ff4c5be83
		Platypus_USDC_MONEY_PairName: "0x27912AE6Ba9a54219d8287C3540A8969FF35500B", // USDC - MONEY, 0x27912ae6ba9a54219d8287c3540a8969ff35500b

		Platypus_AVAX_sAVAX_PairName:  "0x4658EA7e9960D6158a261104aAA160cC953bb6ba", // AVAX - sAVAX, 0x4658ea7e9960d6158a261104aaa160cc953bb6ba
		Platypus_AVAX_yyAVAX_PairName: "0x8B4A45da5B0705Ae4f47EBeFC180C099345cF57e", // AVAX - yyAVAX, 0x8b4a45da5b0705ae4f47ebefc180c099345cf57e

		Platypus_BTCb_WBTCe_PairName: "0x39dE4e02F76Dbd4352Ec2c926D8d64Db8aBdf5b2", // BTC.b - WBTC.e, 0x39de4e02f76dbd4352ec2c926d8d64db8abdf5b2
	},
}

// WARNING: Platypus has dead and relaunched. They might have deployed new contracts and the addresses here might be out dated.
//
// Platypus master platypus v2.
//
// map[chainID] = address.
var PlatypusMasterPlatypusV2List = map[int64]string{
	ChainAvalanche: "0x68c5f4374228BEEdFa078e77b5ed93C28a2f713E", // 0x68c5f4374228beedfa078e77b5ed93c28a2f713e
}

// WARNING: Platypus has dead and relaunched. They might have deployed new contracts and the addresses here might be out dated.
//
// Platypus master platypus v4.
//
// map[chainID] = address.
var PlatypusMasterPlatypusV4List = map[int64]string{
	ChainAvalanche: "0xfF6934aAC9C94E1C39358D4fDCF70aeca77D0AB0", // 0xff6934aac9c94e1c39358d4fdcf70aeca77d0ab0
}

// WARNING: Platypus has dead and relaunched. They might have deployed new contracts and the addresses here might be out dated.
//
// Platypus LP Tokens.
const (
	// Platypus main pool - USDC.e LP.
	PlatypusLpUSDCeAddress string = "0x909B0ce4FaC1A0dCa78F8Ca7430bBAfeEcA12871" // 0x909b0ce4fac1a0dca78f8ca7430bbafeeca12871
	// Platypus main pool - USDT.e LP.
	PlatypusLpUSDTeAddress string = "0x0D26D103c91F63052Fbca88aAF01d5304Ae40015" // 0x0d26d103c91f63052fbca88aaf01d5304ae40015
	// Platypus main pool - USDC LP.
	PlatypusLpUSDCAddress string = "0xAEf735B1E7EcfAf8209ea46610585817Dc0a2E16" // 0xaef735b1e7ecfaf8209ea46610585817dc0a2e16
	// Platypus main pool - USDT LP.
	PlatypusLpUSDTAddress string = "0x776628A5C37335608DD2a9538807b9bba3869E14" // 0x776628a5c37335608dd2a9538807b9bba3869e14
	// Platypus main pool - DAI.e LP.
	PlatypusLpDAIeAddress string = "0xc1Daa16E6979C2D1229cB1fd0823491eA44555Be" // 0xc1daa16e6979c2d1229cb1fd0823491ea44555be
	// Platypus main pool - BUSD LP.
	PlatypusLpBUSDAddress string = "0xe23F8CCDeB4e8Ce5d9fE76782718cD85D12689C8" // 0xe23f8ccdeb4e8ce5d9fe76782718cd85d12689c8
	// Platypus main pool - MIM LP.(deprecated)
	PlatypusLpMIMAddress string = "0x6220BaAd9D08Dee465BefAE4f82ee251cF7c8b82" // 0x6220baad9d08dee465befae4f82ee251cf7c8b82

	// Platypus USDC - FRAX pool - FRAX LP.
	PlatypusLp_USDC_FRAX_FRAX_Address string = "0x6FD4b4c38ED80727EcD0d58505565F9e422c965f" // 0x6fd4b4c38ed80727ecd0d58505565f9e422c965f
	// Platypus USDC - FRAX pool - USDC LP.
	PlatypusLp_USDC_FRAX_USDC_Address string = "0x035D7D7F209B5d18e2AB5C2072E85B32e1D43760" // 0x035d7d7f209b5d18e2ab5c2072e85b32e1d43760

	// Platypus USDC - UST pool - UST LP.(deprecated)
	PlatypusLp_USDC_UST_UST_Address string = "0xc7388D98Fa86B6639d71A0A6d410D5cDfc63A1d0" // 0xc7388d98fa86b6639d71a0a6d410d5cdfc63a1d0
	// Platypus USDC - UST pool - USDC LP.(deprecated)
	PlatypusLp_USDC_UST_USDC_Address string = "0xFC95481F79eC965A535Ed8cef4630e1dd308d319" // 0xfc95481f79ec965a535ed8cef4630e1dd308d319

	// Platypus USDC - MIM pool - MIM LP.
	PlatypusLp_USDC_MIM_MIM_Address string = "0xF01cEA00598d87Cb9792a01B040d04b0bd8Ca781" // 0xf01cea00598d87cb9792a01b040d04b0bd8ca781
	// Platypus USDC - MIM pool - USDC LP.
	PlatypusLp_USDC_MIM_USDC_Address string = "0x4E5704991b43C1D33b9Ccd1BC33B211bf068385A" // 0x4e5704991b43c1d33b9ccd1bc33b211bf068385a

	// Platypus USDC - TUSD pool - TUSD LP.(deprecated)
	PlatypusLp_USDC_TUSD_TUSD_Address string = "0xc75b2b90079492922aF96bA53988D7b384158335" // 0xc75b2b90079492922af96ba53988d7b384158335
	// Platypus USDC - TUSD pool - USDC LP.(deprecated)
	PlatypusLp_USDC_TUSD_USDC_Address string = "0xA551480Dc5399921F8a73e02EC327f2Fd7E5dDc0" // 0xa551480dc5399921f8a73e02ec327f2fd7e5ddc0

	// Platypus USDC - YUSD pool - YUSD LP.
	PlatypusLp_USDC_YUSD_YUSD_Address string = "0x7716307350c0819eD05C3e7f6c478b27CAED5361" // 0x7716307350c0819ed05c3e7f6c478b27caed5361
	// Platypus USDC - YUSD pool - USDC LP.
	PlatypusLp_USDC_YUSD_USDC_Address string = "0x4b851118a4A4948799f24d0CBE17FA3dad09e2D5" // 0x4b851118a4a4948799f24d0cbe17fa3dad09e2d5

	// Platypus USDC - MONEY pool - MONEY LP.
	PlatypusLp_USDC_MONEY_MONEY_Address string = "0xE08947eE864Af325D9F98743B3b905875Ae0Ec99" // 0xe08947ee864af325d9f98743b3b905875ae0ec99
	// Platypus USDC - MONEY pool - USDC LP.
	PlatypusLp_USDC_MONEY_USDC_Address string = "0x551C259Bf4D88edFdAbb04179342a73dAa759583" // 0x551c259bf4d88edfdabb04179342a73daa759583

	// Platypus AVAX - sAVAX pool - AVAX LP.
	PlatypusLp_AVAX_sAVAX_AVAX_Address string = "0xC73eeD4494382093C6a7C284426A9a00f6C79939" // 0xc73eed4494382093c6a7c284426a9a00f6c79939
	// Platypus AVAX - sAVAX pool - sAVAX LP.
	PlatypusLp_AVAX_sAVAX_sAVAX_Address string = "0xA2A7EE49750Ff12bb60b407da2531dB3c50A1789" // 0xa2a7ee49750ff12bb60b407da2531db3c50a1789

	// Platypus AVAX - yyAVAX pool - AVAX LP.
	PlatypusLp_AVAX_yyAVAX_AVAX_Address string = "0x12141b8FD20b4bBdd5F4e911bF91575258A3eABD" // 0x12141b8fd20b4bbdd5f4e911bf91575258a3eabd
	// Platypus AVAX - yyAVAX pool - yyAVAX LP.
	PlatypusLp_AVAX_yyAVAX_yyAVAX_Address string = "0x3BEB0D3DB537b79D377131Ce81950B683d382Ec9" // 0x3beb0d3db537b79d377131ce81950b683d382ec9

	// Platypus BTC.b - WBTC.e pool - BTC.b LP.
	PlatypusLp_BTCb_WBTCe_BTCb_Address string = "0x209a0399A2905900C0d1a9a382fe23e37024dC84" // 0x209a0399a2905900c0d1a9a382fe23e37024dc84
	// Platypus BTC.b - WBTC.e pool - WBTC.e LP.
	PlatypusLp_BTCb_WBTCe_WBTCe_Address string = "0xc09c12093b037866Bf68C9474EcDb5113160fBcE" // 0xc09c12093b037866bf68c9474ecdb5113160fbce
)

// WARNING: Platypus has dead and relaunched. They might have deployed new contracts and the addresses here might be out dated.
//
// Platypus lp list.
//
// map[network][pool name][underlying] = address.
var PlatypusLpList = map[int64]map[string]map[string]string{
	ChainAvalanche: {
		// main pools
		PlatypusMainPoolsName: {
			DAIeList[ChainAvalanche]:  PlatypusLpDAIeAddress,
			USDCList[ChainAvalanche]:  PlatypusLpUSDCAddress,
			USDCeList[ChainAvalanche]: PlatypusLpUSDCeAddress,
			USDTList[ChainAvalanche]:  PlatypusLpUSDTAddress,
			USDTeList[ChainAvalanche]: PlatypusLpUSDTeAddress,
			BUSDList[ChainAvalanche]:  PlatypusLpBUSDAddress,
		},

		// USDC - FRAX
		Platypus_USDC_FRAX_PairName: {
			USDCList[ChainAvalanche]: PlatypusLp_USDC_FRAX_USDC_Address,
			FRAXList[ChainAvalanche]: PlatypusLp_USDC_FRAX_FRAX_Address,
		},
		// USDC - MIM
		Platypus_USDC_MIM_PairName: {
			USDCList[ChainAvalanche]: PlatypusLp_USDC_MIM_USDC_Address,
			MIMList[ChainAvalanche]:  PlatypusLp_USDC_MIM_MIM_Address,
		},
		// USDC - YUSD
		Platypus_USDC_YUSD_PairName: {
			USDCList[ChainAvalanche]: PlatypusLp_USDC_YUSD_USDC_Address,
			YUSDList[ChainAvalanche]: PlatypusLp_USDC_YUSD_YUSD_Address,
		},
		// USDC - MONEY
		Platypus_USDC_MONEY_PairName: {
			USDCList[ChainAvalanche]:  PlatypusLp_USDC_MONEY_USDC_Address,
			MoneyList[ChainAvalanche]: PlatypusLp_USDC_MONEY_MONEY_Address,
		},

		// BTC.b - WBTC.e
		Platypus_BTCb_WBTCe_PairName: {
			WBTCeList[ChainAvalanche]: PlatypusLp_BTCb_WBTCe_WBTCe_Address,
			BTCbList[ChainAvalanche]:  PlatypusLp_BTCb_WBTCe_BTCb_Address,
		},

		// AVAX - sAVAX
		Platypus_AVAX_sAVAX_PairName: {
			WAVAXList[ChainAvalanche]: PlatypusLp_AVAX_sAVAX_AVAX_Address,
			SAVAXList[ChainAvalanche]: PlatypusLp_AVAX_sAVAX_sAVAX_Address,
		},
		// AVAX - yyAVAX
		Platypus_AVAX_yyAVAX_PairName: {
			WAVAXList[ChainAvalanche]:  PlatypusLp_AVAX_yyAVAX_AVAX_Address,
			YyAVAXList[ChainAvalanche]: PlatypusLp_AVAX_yyAVAX_yyAVAX_Address,
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
