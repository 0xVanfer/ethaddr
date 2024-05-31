package ethaddr

// Aave-like protocols.
var CompoundLikeProtocols = []string{
	BenqiProtocol,
	CompoundProtocol,
	TraderJoeProtocol,
}

// Compound-like C token list map.
//
// map[protocol name][network][underlying] = c token.
var CompoundLikeCTokenListMap = map[string]*map[int64]map[string]string{
	CompoundProtocol:  &CompoundCTokenList,
	BenqiProtocol:     &BenqiCTokenList,
	TraderJoeProtocol: &TraderjoeCTokenList,
}
