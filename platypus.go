package ethaddr

import "github.com/0xVanfer/chainId"

const PlatypusProtocol string = "platypus"

// Platypus token: PTP.
//
// map[network] = address.
var PlatypusTokenList = map[string]string{
	chainId.AvalancheChainName: "0x22d4002028f537599be9f666d1c4fa138522f9c8", // PTP
}

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

// platypus pair names
const PlatypusMainPoolsName string = "PlatypusMainPools"
const PlatypusUSDCFRAXPairName string = "PlatypusUSDCFRAXpair"
const PlatypusUSDCMIMPairName string = "PlatypusUSDCMIMpair"
const PlatypusAVAXsAVAXPairName string = "PlatypusAVAXsAVAXpair"
const PlatypusUSDCTUSDPairName string = "PlatypusUSDCTUSDpair"
const PlatypusUSDCYUSDPairName string = "PlatypusUSDCYUSDpair"
const PlatypusBTCbWBTCePairName string = "PlatypusBTCbWBTCepair"

// Platypus router.
//
// map[network][pool name] = address.
var PlatypusRouterList = map[string]map[string]string{
	chainId.AvalancheChainName: {
		PlatypusMainPoolsName:     "0x66357dcace80431aee0a7507e2e361b7e2402370",
		PlatypusUSDCFRAXPairName:  "0xb8e567fc23c39c94a1f6359509d7b43d1fbed824",
		PlatypusUSDCMIMPairName:   "0x30c30d826be87cd0a4b90855c2f38f7fcfe4eaa7",
		PlatypusAVAXsAVAXPairName: "0x4658ea7e9960d6158a261104aaa160cc953bb6ba",
		PlatypusUSDCTUSDPairName:  "0x13329c7905f1ee55c3c7d7bfc26c1197c512c207",
		PlatypusUSDCYUSDPairName:  "0xc828d995c686aaba78a4ac89dfc8ec0ff4c5be83",
		PlatypusBTCbWBTCePairName: "0x39de4e02f76dbd4352ec2c926d8d64db8abdf5b2",
	},
}

// Platypus master platypus v2.
//
// map[network] = address.
var PlatypusMasterPlatypusV2List = map[string]string{
	chainId.AvalancheChainName: "0x68c5f4374228beedfa078e77b5ed93c28a2f713e",
}

const PlatypusLpUSDCeAddress string = "0x909b0ce4fac1a0dca78f8ca7430bbafeeca12871"
const PlatypusLpUSDTeAddress string = "0x0d26d103c91f63052fbca88aaf01d5304ae40015"
const PlatypusLpUSDCAddress string = "0xaef735b1e7ecfaf8209ea46610585817dc0a2e16"
const PlatypusLpUSDTAddress string = "0x776628a5c37335608dd2a9538807b9bba3869e14"
const PlatypusLpDAIeAddress string = "0xc1daa16e6979c2d1229cb1fd0823491ea44555be"
const PlatypusLpBUSDAddress string = "0xe23f8ccdeb4e8ce5d9fe76782718cd85d12689c8"
const PlatypusLpMIMDeprecatedAddress string = "0x6220baad9d08dee465befae4f82ee251cf7c8b82"

const PlatypusLpFRAXAddress string = "0x6fd4b4c38ed80727ecd0d58505565f9e422c965f"
const PlatypusLpUSDC_FRAXAddress string = "0x035d7d7f209b5d18e2ab5c2072e85b32e1d43760"

const PlatypusLpUSTAddress string = "0xc7388d98fa86b6639d71a0a6d410d5cdfc63a1d0"
const PlatypusLpUSDC_USTAddress string = "0xfc95481f79ec965a535ed8cef4630e1dd308d319"

const PlatypusLpMIMAddress string = "0xf01cea00598d87cb9792a01b040d04b0bd8ca781"
const PlatypusLpUSDC_MIMAddress string = "0x4e5704991b43c1d33b9ccd1bc33b211bf068385a"

const PlatypusLpAVAXAddress string = "0xc73eed4494382093c6a7c284426a9a00f6c79939"
const PlatypusLpsAVAXAddress string = "0xa2a7ee49750ff12bb60b407da2531db3c50a1789"

const PlatypusLpTUSDAddress string = "0xc75b2b90079492922af96ba53988d7b384158335"
const PlatypusLpUSDC_TUSDAddress string = "0xa551480dc5399921f8a73e02ec327f2fd7e5ddc0"

const PlatypusLpYUSDAddress string = "0x7716307350c0819ed05c3e7f6c478b27caed5361"
const PlatypusLpUSDC_YUSDAddress string = "0x4b851118a4a4948799f24d0cbe17fa3dad09e2d5"

const PlatypusLpWBTCeAddress string = "0xc09c12093b037866bf68c9474ecdb5113160fbce"
const PlatypusLpBTCbAddress string = "0x209a0399a2905900c0d1a9a382fe23e37024dc84"

// Platypus lp list.
//
// map[network][pool name][underlying] = address.
var PlatypusLpList = map[string]map[string]map[string]string{
	chainId.AvalancheChainName: {
		PlatypusMainPoolsName: {
			DAIeList[chainId.AvalancheChainName]:  PlatypusLpDAIeAddress,
			USDCList[chainId.AvalancheChainName]:  PlatypusLpUSDCAddress,
			USDCeList[chainId.AvalancheChainName]: PlatypusLpUSDCeAddress,
			USDTList[chainId.AvalancheChainName]:  PlatypusLpUSDTAddress,
			USDTeList[chainId.AvalancheChainName]: PlatypusLpUSDTeAddress,
			BUSDList[chainId.AvalancheChainName]:  PlatypusLpBUSDAddress,
		},
		PlatypusUSDCFRAXPairName: {
			USDCList[chainId.AvalancheChainName]: PlatypusLpUSDC_FRAXAddress,
			FRAXList[chainId.AvalancheChainName]: PlatypusLpFRAXAddress,
		},
		PlatypusUSDCMIMPairName: {
			USDCList[chainId.AvalancheChainName]: PlatypusLpUSDC_MIMAddress,
			MIMList[chainId.AvalancheChainName]:  PlatypusLpMIMAddress,
		},
		PlatypusAVAXsAVAXPairName: {
			WAVAXList[chainId.AvalancheChainName]:      PlatypusLpAVAXAddress,
			BenqiSAVAXList[chainId.AvalancheChainName]: PlatypusLpsAVAXAddress,
		},
		PlatypusUSDCYUSDPairName: {
			USDCList[chainId.AvalancheChainName]: PlatypusLpUSDC_YUSDAddress,
			YUSDList[chainId.AvalancheChainName]: PlatypusLpYUSDAddress,
		},
		PlatypusBTCbWBTCePairName: {
			WBTCList[chainId.AvalancheChainName]: PlatypusLpWBTCeAddress,
			BTCbList[chainId.AvalancheChainName]: PlatypusLpBTCbAddress,
		},
		"PlatypusDeprecated": {
			"mim":            PlatypusLpMIMDeprecatedAddress,
			"usdc_tusd_tusd": PlatypusLpTUSDAddress,
			"usdc_tusd_usdc": PlatypusLpUSDC_TUSDAddress,
			"usdc_ust_usdc":  PlatypusLpUSDC_USTAddress,
			"usdc_ust_ust":   PlatypusLpUSTAddress,
		},
	},
}
