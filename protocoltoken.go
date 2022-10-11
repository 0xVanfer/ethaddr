package ethaddr

// Protocol token list map.
//
// map[protocol name][network] = token.
var ProtocolTokenListMap = map[string]map[string]string{
	AaveV2Protocol:       AaveTokenList,         // aave v2
	AaveV3Protocol:       AaveTokenList,         // aave v3
	AbracadabraProtocol:  SPELLList,             // abracadabra
	AlphaProtocol:        AlphaTokenList,        // alpha
	AngleProtocol:        ANGLEList,             // angle
	AxialProtocol:        AxialTokenList,        // axial
	BalancerProtocol:     BalancerTokenList,     // balancer
	BeefyProtocol:        BeefyTokenList,        // beefy
	BenqiProtocol:        BenqiTokenList,        // benqi
	CompoundProtocol:     CompoundTokenList,     // compound
	ConvexProtocol:       ConvexTokenList,       // convex
	CurveProtocol:        CurveTokenlist,        // curve
	DecentralandProtocol: DecentralandTokenList, // decentraland
	KyberProtocol:        KyberTokenList,        // kyber
	LidoProtocol:         LidoTokenList,         // lido
	MakerProtocol:        MakerTokenList,        // maker
	PancakeswapProtocol:  PancakeswapTokenList,  // pancakeswap
	PangolinProtocol:     PangolinTokenList,     // pangolin
	QuickswapProtocol:    QuickswapTokenList,    // quickswap
	RepublicProtocol:     RepublicTokenList,     // republic
	SushiProtocol:        SushiTokenList,        // sushi
	TraderJoeProtocol:    TraderjoeTokenList,    // traderjoe
	UniswapProtocolV2:    UniswapTokenList,      // uniswap v2
	UniswapProtocolV3:    UniswapTokenList,      // uniswap v3
	VectorProtocol:       VectorTokenList,       // vector
	YearnProtocol:        YearnTokenList,        // yearn
}
