package ethaddr

const (
	ChainEthereum   int64 = 1
	ChainGoerli     int64 = 5 // Testnet
	ChainOptimism   int64 = 10
	ChainBSC        int64 = 56
	ChainOK         int64 = 66
	ChainGnosis     int64 = 100
	ChainHeco       int64 = 128
	ChainPolygon    int64 = 137
	ChainBlast      int64 = 238
	ChainFantom     int64 = 250
	ChainFraxtal    int64 = 252
	ChainFilecoin   int64 = 314
	ChainZkSync     int64 = 324
	ChainWorldChain int64 = 480
	ChainPolygonZk  int64 = 1101
	ChainMoonbeam   int64 = 1284
	ChainCentrifuge int64 = 2031
	ChainKava       int64 = 2222
	ChainMantle     int64 = 5000
	ChainBase       int64 = 8453
	ChainImmutable  int64 = 13371
	ChainArbitrum   int64 = 42161
	ChainCelo       int64 = 42220
	ChainAvalanche  int64 = 43114
	ChainLinea      int64 = 59144
	ChainScroll     int64 = 534352
	ChainZkLinkNova int64 = 810180
	ChainSepolia    int64 = 11155111 // Testnet
)

// Chain name map to wrapped native tokens.
//
// map[chain name] = address.
var WrappedNativeTokenList = map[int64]string{
	// ETH and ETH L2s.

	ChainEthereum:  WETHList[ChainEthereum],  // 1
	ChainOptimism:  WETHList[ChainOptimism],  // 10
	ChainPolygonZk: WETHList[ChainPolygonZk], // 1101
	ChainBase:      WETHList[ChainBase],      // 8453
	ChainArbitrum:  WETHList[ChainArbitrum],  // 42161
	ChainScroll:    WETHList[ChainScroll],    // 534352

	ChainBSC:       WBNBList[ChainBSC],        // 56
	ChainPolygon:   WMATICList[ChainPolygon],  // 137
	ChainFantom:    WFTMList[ChainFantom],     // 250
	ChainAvalanche: WAVAXList[ChainAvalanche], // 43114
}

// Map chainID to chain name (self defined).
var ChainNameMap = map[int64]string{
	ChainEthereum:   ChainNameEthereum,
	ChainGoerli:     ChainNameGoerli,
	ChainOptimism:   ChainNameOptimism,
	ChainBSC:        ChainNameBSC,
	ChainOK:         ChainNameOkex,
	ChainGnosis:     ChainNameGnosis,
	ChainHeco:       ChainNameHeco,
	ChainPolygon:    ChainNamePolygon,
	ChainBlast:      ChainNameBlast,
	ChainFantom:     ChainNameFantom,
	ChainFraxtal:    ChainNameFraxtal,
	ChainFilecoin:   ChainNameFilecoin,
	ChainZkSync:     ChainNameZkSync,
	ChainWorldChain: ChainNameWorldChain,
	ChainPolygonZk:  ChainNamePolygonZkEVM,
	ChainMoonbeam:   ChainNameMoonbeam,
	ChainCentrifuge: ChainNameCentrifuge,
	ChainKava:       ChainNameKava,
	ChainMantle:     ChainNameMantle,
	ChainBase:       ChainNameBase,
	ChainImmutable:  ChainNameImmutable,
	ChainArbitrum:   ChainNameArbitrum,
	ChainCelo:       ChainNameCelo,
	ChainAvalanche:  ChainNameAvalanche,
	ChainLinea:      ChainNameLinea,
	ChainScroll:     ChainNameScroll,
	ChainZkLinkNova: ChainNameZkLinkNova,
	ChainSepolia:    ChainNameSepolia,
}

// Map chain name (self defined) to chainID.
var ChainNameMapReverse = map[string]int64{
	ChainNameEthereum:     ChainEthereum,
	ChainNameGoerli:       ChainGoerli,
	ChainNameOptimism:     ChainOptimism,
	ChainNameBSC:          ChainBSC,
	ChainNameOkex:         ChainOK,
	ChainNameGnosis:       ChainGnosis,
	ChainNameHeco:         ChainHeco,
	ChainNamePolygon:      ChainPolygon,
	ChainNameBlast:        ChainBlast,
	ChainNameFantom:       ChainFantom,
	ChainNameFraxtal:      ChainFraxtal,
	ChainNameFilecoin:     ChainFilecoin,
	ChainNameZkSync:       ChainZkSync,
	ChainNameWorldChain:   ChainWorldChain,
	ChainNamePolygonZkEVM: ChainPolygonZk,
	ChainNameMoonbeam:     ChainMoonbeam,
	ChainNameCentrifuge:   ChainCentrifuge,
	ChainNameKava:         ChainKava,
	ChainNameMantle:       ChainMantle,
	ChainNameBase:         ChainBase,
	ChainNameImmutable:    ChainImmutable,
	ChainNameArbitrum:     ChainArbitrum,
	ChainNameCelo:         ChainCelo,
	ChainNameAvalanche:    ChainAvalanche,
	ChainNameLinea:        ChainLinea,
	ChainNameScroll:       ChainScroll,
	ChainNameZkLinkNova:   ChainZkLinkNova,
	ChainNameSepolia:      ChainSepolia,
}
