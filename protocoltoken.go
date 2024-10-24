package ethaddr

// Protocol token list map.
//
// map[protocol name][network] = token.
//
// Not suggested to use this map.
var ProtocolTokenListMap = map[string]map[int64]string{
	AaveV2Protocol:            AAVEList,   // aave v2
	AaveV3Protocol:            AAVEList,   // aave v3
	AaveV3ProtocolLidoFork:    AAVEList,   // aave v3 lido
	AaveV3ProtocolEtherfiFork: AAVEList,   // aave v3 etherfi
	AbracadabraProtocol:       SPELLList,  // abracadabra
	AlphaProtocol:             ALPHAList,  // alpha
	AngleProtocol:             ANGLEList,  // angle
	AxialProtocol:             AXIALList,  // axial
	BalancerProtocol:          BALList,    // balancer
	BeefyProtocol:             BIFIList,   // beefy
	BenqiProtocol:             QIList,     // benqi
	ChainlinkProtocol:         LINKList,   // chainlink
	CompoundV2Protocol:        COMPList,   // compound
	ConvexProtocol:            CVXList,    // convex
	CurveProtocol:             CRVList,    // curve
	DecentralandProtocol:      MANAList,   // decentraland
	EnjinProtocol:             ENJList,    // enjin
	EthenaProtocol:            ENAList,    // ethena
	EtherfiProtocol:           ETHFIList,  // etherfi
	KyberProtocol:             KNCList,    // kyber
	LidoProtocol:              LDOList,    // lido
	MakerProtocol:             MKRList,    // maker
	MaverickProtocol:          MAVList,    // maverick
	MorphoOptimizersProtocol:  MORPHOList, // morphooptimizers
	MorphoBlueProtocol:        MORPHOList, // morphoblue
	OneInchProtocol:           OINCHLIST,  // 1inch
	PancakeswapProtocol:       CakeList,   // pancakeswap
	PangolinProtocol:          PNGList,    // pangolin
	PendleProtocol:            PENDLEList, // pendle
	PlatypusProtocol:          PTPList,    // platypus
	QuickswapProtocol:         QUICKList,  // quickswap
	RepublicProtocol:          RENList,    // republic
	RocketPoolProtocol:        RPLList,    // rocketpool
	StargateProtocol:          STGList,    // stargate
	SushiProtocol:             SUSHIList,  // sushi
	SynthetixProtocol:         SNXList,    // synthetix
	TraderJoeProtocol:         JOEList,    // traderjoe
	UniswapProtocolV2:         UNIList,    // uniswap v2
	UniswapProtocolV3:         UNIList,    // uniswap v3
	UniswapProtocolV4:         UNIList,    // uniswap v4
	VectorProtocol:            VTXList,    // vector
	VenusProtocol:             XVSList,    // venus
	YearnProtocol:             YFIList,    // yearn
}
