package ethaddr

// Aave fork protocols.
var AaveForkProtocols = []string{
	AaveV2Protocol,
	AaveV3Protocol,
	AaveV3ProtocolLidoFork,
	AaveV3ProtocolEtherfiFork,
	SparkProtocol,
}

// Aave fork lending pool list map.
//
// map[protocol name][network] = lending pool.
var AaveForkLendingPoolListMap = map[string]*map[int64]string{
	AaveV2Protocol:            &AaveLendingPoolV2List,
	AaveV3Protocol:            &AaveLendingPoolV3List,
	AaveV3ProtocolLidoFork:    &AaveLendingPoolV3LidoForkList,
	AaveV3ProtocolEtherfiFork: &AaveLendingPoolV3EtherfiForkList,
	SparkProtocol:             &SparkPoolList,
}

// Aave fork A token list map.
//
// map[protocol name][network][underlying] = atoken.
var AaveForkATokenListMap = map[string]*map[int64]map[string]string{
	AaveV2Protocol:            &AaveATokenV2List,
	AaveV3Protocol:            &AaveATokenV3List,
	AaveV3ProtocolLidoFork:    &AaveATokenV3LidoForkList,
	AaveV3ProtocolEtherfiFork: &AaveATokenV3EtherfiForkList,
	SparkProtocol:             &SparkATokenList,
}

// Aave fork V token list map.
//
// map[protocol name][network][underlying] = vtoken.
var AaveForkVTokenListMap = map[string]*map[int64]map[string]string{
	AaveV2Protocol:            &AaveVTokenV2List,
	AaveV3Protocol:            &AaveVTokenV3List,
	AaveV3ProtocolLidoFork:    &AaveVTokenV3LidoForkList,
	AaveV3ProtocolEtherfiFork: &AaveVTokenV3EtherfiForkList,
	SparkProtocol:             &SparkVTokenList,
}

// DEPRECATED: Aave v3 has deprecated sToken.
//
// Aave fork S token list map.
//
// map[protocol name][network][underlying] = stoken.
var AaveForkSTokenListMap = map[string]*map[int64]map[string]string{
	AaveV2Protocol:         &AaveSTokenV2List,
	AaveV3Protocol:         &AaveSTokenV3List,
	AaveV3ProtocolLidoFork: &AaveSTokenV3LidoForkList,
	SparkProtocol:          &SparkSTokenList,
}

var AaveForkPoolAddressProviderListMap = map[string]*map[int64]string{
	AaveV2Protocol:            &AavePoolAddressProviderV2List,
	AaveV3Protocol:            &AavePoolAddressProviderV3List,
	AaveV3ProtocolLidoFork:    &AavePoolAddressProviderV3LidoForkList,
	AaveV3ProtocolEtherfiFork: &AavePoolAddressProviderV3EtherfiForkList,
	SparkProtocol:             &SparkPoolAddressProviderList,
}

var AaveForkUiPoolDataProviderListMap = map[string]*map[int64]string{
	AaveV2Protocol:            &AaveUiPoolDataProviderV2List,
	AaveV3Protocol:            &AaveUiPoolDataProviderV3List,
	AaveV3ProtocolLidoFork:    &AaveUiPoolDataProviderV3LidoForkList,
	AaveV3ProtocolEtherfiFork: &AaveUiPoolDataProviderV3EtherfiForkList,
	SparkProtocol:             &SparkUiPoolDataProviderList,
}

var AaveForkUiIncentiveDataProviderListMap = map[string]*map[int64]string{
	AaveV2Protocol:            &AaveUiIncentiveDataProviderV2List,
	AaveV3Protocol:            &AaveUiIncentiveDataProviderV3List,
	AaveV3ProtocolLidoFork:    &AaveUiIncentiveDataProviderV3LidoForkList,
	AaveV3ProtocolEtherfiFork: &AaveUiIncentiveDataProviderV3EtherfiForkList,
	SparkProtocol:             &SparkUiIncentiveDataProviderList,
}
