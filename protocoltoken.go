package ethaddr

// Deprecated: meaningless.
//
// Protocol token list map.
//
// map[protocol name][network] = token.
var ProtocolTokenListMap = map[string]map[int64]string{
	AaveV2Protocol:       AAVEList,  // aave v2
	AaveV3Protocol:       AAVEList,  // aave v3
	AbracadabraProtocol:  SPELLList, // abracadabra
	AlphaProtocol:        ALPHAList, // alpha
	AngleProtocol:        ANGLEList, // angle
	AxialProtocol:        AXIALList, // axial
	BalancerProtocol:     BALList,   // balancer
	BeefyProtocol:        BIFIList,  // beefy
	BenqiProtocol:        QIList,    // benqi
	CompoundProtocol:     COMPList,  // compound
	ConvexProtocol:       CVXList,   // convex
	CurveProtocol:        CRVList,   // curve
	DecentralandProtocol: MANAList,  // decentraland
	KyberProtocol:        KNCList,   // kyber
	LidoProtocol:         LDOList,   // lido
	MakerProtocol:        MKRList,   // maker
	PancakeswapProtocol:  CakeList,  // pancakeswap
	PangolinProtocol:     PNGList,   // pangolin
	QuickswapProtocol:    QUICKList, // quickswap
	RepublicProtocol:     RENList,   // republic
	SushiProtocol:        SUSHIList, // sushi
	TraderJoeProtocol:    JOEList,   // traderjoe
	UniswapProtocolV2:    UNIList,   // uniswap v2
	UniswapProtocolV3:    UNIList,   // uniswap v3
	VectorProtocol:       VTXList,   // vector
	YearnProtocol:        YFIList,   // yearn
}
