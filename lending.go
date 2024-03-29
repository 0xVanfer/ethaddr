package ethaddr

// Lending tokens info.
type LendingTokenLists struct {
	ATokenList *map[string]map[string]string // map[network][underlying]=address
	VTokenList *map[string]map[string]string // map[network][underlying]=address
	STokenList *map[string]map[string]string // map[network][underlying]=address
	CTokenList *map[string]map[string]string // map[network][underlying]=address
}

// If the list not pass regular check, the protocol is not supported.
func (list LendingTokenLists) RegularCheck() bool {
	return list.ATokenList != nil || list.CTokenList != nil
}

// map[protocol name]=LendingTokenLists
var LendingTokenListsMap = map[string]LendingTokenLists{
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
	BenqiProtocol: {
		CTokenList: &BenqiCTokenList,
	},
	TraderJoeProtocol: {
		CTokenList: &TraderjoeCTokenList,
	},
	CompoundProtocol: {
		CTokenList: &CompoundCTokenList,
	},
	SparkProtocol: {
		ATokenList: &SparkATokenList,
		VTokenList: &SparkVTokenList,
		STokenList: &SparkSTokenList,
	},
}
