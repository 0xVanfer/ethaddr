package ethaddr

// Aave-like protocols.
var AaveLikeProtocols = []string{
	AaveV2Protocol,
	AaveV3Protocol,
	SparkProtocol,
}

// Aave-like A token list map.
//
// map[protocol name][network][underlying] = a token.
var AaveLikeATokenListMap = map[string]*map[int64]map[string]string{
	AaveV2Protocol: &AaveATokenV2List,
	AaveV3Protocol: &AaveATokenV3List,
	SparkProtocol:  &SparkATokenList,
}

// Aave-like V token list map.
//
// map[protocol name][network][underlying] = v token.
var AaveLikeVTokenListMap = map[string]*map[int64]map[string]string{
	AaveV2Protocol: &AaveVTokenV2List,
	AaveV3Protocol: &AaveVTokenV3List,
	SparkProtocol:  &SparkVTokenList,
}

// Aave-like S token list map.
//
// map[protocol name][network][underlying] = s token.
var AaveLikeSTokenListMap = map[string]*map[int64]map[string]string{
	AaveV2Protocol: &AaveSTokenV2List,
	AaveV3Protocol: &AaveSTokenV3List,
	SparkProtocol:  &SparkSTokenList,
}
