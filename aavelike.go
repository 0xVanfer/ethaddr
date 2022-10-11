package ethaddr

// Aave-like protocols.
var AaveLikeProtocols = []string{
	AaveV2Protocol,
	AaveV3Protocol,
}

// Aave-like A token list map.
//
// map[protocol name][network][underlying] = a token.
var AaveLikeATokenListMap = map[string]map[string]map[string]string{
	AaveV2Protocol: AaveATokenV2List,
	AaveV3Protocol: AaveATokenV3List,
}

// Aave-like V token list map.
//
// map[protocol name][network][underlying] = v token.
var AaveLikeVTokenListMap = map[string]map[string]map[string]string{
	AaveV2Protocol: AaveVTokenV2List,
	AaveV3Protocol: AaveVTokenV3List,
}

// Aave-like S token list map.
//
// map[protocol name][network][underlying] = s token.
var AaveLikeSTokenListMap = map[string]map[string]map[string]string{
	AaveV2Protocol: AaveSTokenV2List,
	AaveV3Protocol: AaveSTokenV3List,
}
