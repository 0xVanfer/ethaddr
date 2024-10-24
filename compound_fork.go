package ethaddr

// Compound V2 fork protocols.
var CompoundV2ForkProtocols = []string{
	BenqiProtocol,
	CompoundV2Protocol,
	TraderJoeProtocol,
}

// Compound V2 fork token list map.
//
// map[protocol name][network][underlying] = c token.
var CompoundV2ForkCTokenListMap = map[string]*map[int64]map[string]string{
	CompoundV2Protocol: &CompoundCTokenList,
	BenqiProtocol:      &BenqiCTokenList,
	TraderJoeProtocol:  &TraderjoeCTokenList,
}
