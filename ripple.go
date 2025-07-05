package ethaddr

// Website: https://ripple.com/
const RippleProtocol string = "ripple"

// Ripple XRP.
//
// map[chainID] = address.
var XRPList = map[int64]string{
	ChainBSC: "0x1D2F0da169ceB9fC7B3144628dB156f3F6c60dBE", // XRP, 0x1d2f0da169ceb9fc7b3144628db156f3f6c60dbe
}

// Ripple RLUSD.
//
// map[chainID] = address.
var RLUSDList = map[int64]string{
	ChainEthereum: "0x8292Bb45bf1Ee4d140127049757C2E0fF06317eD", // RLUSD, 0x8292bb45bf1ee4d140127049757c2e0ff06317ed
}
