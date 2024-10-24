package ethaddr

// ETH LST
var AllETHLSTs = []map[int64]string{
	AnkrETHList,         // ankr
	BETHList, WBETHList, // binance
	CbETHList,               // coinbase
	FrxETHList, SfrxETHList, // frax
	STETHList, WSTETHList, // lido
	METHList,  // mantle
	RETHList,  // rocketpool
	ETHxList,  // stader
	OSETHList, // stakewise
	SWETHList, // swell
}

// ETH LRT, restaked by eigenlayer
var AllETHEigenLayerLRTs = []map[int64]string{
	EETHList, WEETHList, // etherfi
	RsETHList,  // Kelp
	EzETHList,  // renzo
	RswETHList, // swell
}

// ETH LRT, restaked symbiotic
var AllETHSymbioticLRTs = []map[int64]string{
	AmphrETHList, // amphor
	CoETHList,    // chorusone
	IfsETHList,   // infStones
	LugaETHList,  // luganodes
	RstETHList,   // p2p
	Re7LRTList,   // re7
	PzETHList,    // renzo
	SteakLRTList, // steakhouse
}

// ETH LRT
var AllETHLRTs = append(AllETHEigenLayerLRTs, AllETHSymbioticLRTs...)

// ETH LST & LRT
var AllETHRelated = append(AllETHLSTs, AllETHLRTs...)
