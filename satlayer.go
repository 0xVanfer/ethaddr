package ethaddr

// Website: https://satlayer.xyz/
const SatLayerProtocol string = "satlayer"

// SatLayer restaked BTC list.
var SatLayerLSTsList = map[int64]map[string]string{
	ChainEthereum: {
		LBTCList[ChainEthereum]:     "0x067e11Ac5471C853aea205B3C1933a5f6367152F", // satLBTC, 0x067e11ac5471c853aea205b3c1933a5f6367152f
		FBTCList[ChainEthereum]:     "0xe2C6755C10d0B61D8B11Dd2851AE8266Cea912DC", // satFBTC, 0xe2c6755c10d0b61d8b11dd2851ae8266cea912dc
		UniBTCList[ChainEthereum]:   "0xF7De2B7afdb07AA5dD143180Ed758165821E076e", // satUniBTC, 0xf7de2b7afdb07aa5dd143180ed758165821e076e
		WBTCList[ChainEthereum]:     "0x69223B5B36a785Ec08e5f685fd7961399982C566", // satWBTC, 0x69223b5b36a785ec08e5f685fd7961399982c566
		XSolvBTCList[ChainEthereum]: "0x17140b69FfaDfF9e87BF1D86D99119ee10AD24ff", // satSolvBTC, 0x17140b69ffadff9e87bf1d86d99119ee10ad24ff
		PumpBTCList[ChainEthereum]:  "0x0c4dd69705D16d91bC9C0534Cc926966f23430c7", // satPumpBTC, 0x0c4dd69705d16d91bc9c0534cc926966f23430c7
	},
}
