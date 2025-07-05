package ethaddr

// EXPERIMENTAL. DO NOT USE.
type BridgeInfo struct {
	Protocol      string `json:"protocol"`
	IsOfficial    bool   `json:"is_official"`
	TargetChainID int64  `json:"target_chain_id"`
}

// EXPERIMENTAL. DO NOT USE.
//
// map[from chainID][bridge address] = target chainID
var Bridges = map[int64]map[string]BridgeInfo{
	ChainEthereum: {
		// Scroll 官方桥
		"0xD8A791fE2bE73eb6E6cF1eb0cb3F36adC9B3F8f9": {
			IsOfficial:    true,
			TargetChainID: ChainScroll,
		},
	},
}
