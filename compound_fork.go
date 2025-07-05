package ethaddr

// Compound V2 fork protocols.
var CompoundV2ForkProtocols = []string{
	BenqiProtocol,
	CompoundV2Protocol,
	TraderJoeProtocol,
	VenusProtocol,
	RhoProtocol,
}

// Compound V2 fork token list map.
//
// map[protocol name][network][underlying] = c token.
var CompoundV2ForkCTokenListMap = map[string]*map[int64]map[string]string{
	CompoundV2Protocol: &CompoundV2CTokenList,
	BenqiProtocol:      &BenqiCTokenList,
	TraderJoeProtocol:  &TraderjoeCTokenList,
	VenusProtocol:      &VenusCTokenList,
	RhoProtocol:        &RhoCTokenList,
}

// Compound V3 fork protocols.
var CompoundV3ForkProtocols = []string{
	CompoundV3Protocol,
}

// Compound V3 fork token list map.
//
// map[protocol name][network][underlying] = c token.
var CompoundV3ForkCTokenListMap = map[string]*map[int64]map[string]string{
	CompoundV3Protocol: &CompoundV3CTokenList,
}
