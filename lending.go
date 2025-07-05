package ethaddr

// Lending tokens info.
type LendingTokenLists struct {
	ATokenList *map[int64]map[string]string // map[network][underlying]=address
	VTokenList *map[int64]map[string]string // map[network][underlying]=address
	STokenList *map[int64]map[string]string // map[network][underlying]=address
	CTokenList *map[int64]map[string]string // map[network][underlying]=address
}

// map[protocol name]=LendingTokenLists
var LendingTokenListsMap = map[string]LendingTokenLists{
	// Aave forks

	AaveV2Protocol: {
		ATokenList: &AaveATokenV2List,
		VTokenList: &AaveVTokenV2List,
		STokenList: &AaveSTokenV2List,
	},
	AaveV3Protocol: {
		ATokenList: &AaveATokenV3List,
		VTokenList: &AaveVTokenV3List,
		STokenList: &AaveSTokenV3List,
	},
	AaveV3ProtocolLidoFork: {
		ATokenList: &AaveATokenV3LidoForkList,
		VTokenList: &AaveVTokenV3LidoForkList,
		STokenList: &AaveSTokenV3LidoForkList,
	},
	AaveV3ProtocolEtherfiFork: {
		ATokenList: &AaveATokenV3EtherfiForkList,
		VTokenList: &AaveVTokenV3EtherfiForkList,
	},
	SparkProtocol: {
		ATokenList: &SparkATokenList,
		VTokenList: &SparkVTokenList,
		STokenList: &SparkSTokenList,
	},

	// CompoundV2 forks

	BenqiProtocol: {
		CTokenList: &BenqiCTokenList,
	},
	TraderJoeProtocol: {
		CTokenList: &TraderjoeCTokenList,
	},
	CompoundV2Protocol: {
		CTokenList: &CompoundV2CTokenList,
	},
	RhoProtocol: {
		CTokenList: &RhoCTokenList,
	},
	VenusProtocol: {
		CTokenList: &VenusCTokenList,
	},

	// CompoundV3 forks

	CompoundV3Protocol: {
		CTokenList: &CompoundV3CTokenList,
	},
}
