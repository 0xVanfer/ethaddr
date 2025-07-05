package ethaddr

const (
	ChainEthereum     int64 = 1
	ChainGoerli       int64 = 5 // Testnet
	ChainOptimism     int64 = 10
	ChainBSC          int64 = 56
	ChainOK           int64 = 66
	ChainGnosis       int64 = 100
	ChainHeco         int64 = 128
	ChainPolygon      int64 = 137
	ChainMonad        int64 = 143
	ChainSonic        int64 = 146
	ChainTac          int64 = 239
	ChainFantom       int64 = 250
	ChainFraxtal      int64 = 252
	ChainFilecoin     int64 = 314
	ChainZkSync       int64 = 324
	ChainWorldChain   int64 = 480
	ChainPolygonZk    int64 = 1101
	ChainMoonbeam     int64 = 1284
	ChainSei          int64 = 1329
	ChainCentrifuge   int64 = 2031
	ChainKava         int64 = 2222
	ChainMerlin       int64 = 4200
	ChainMantle       int64 = 5000
	ChainBase         int64 = 8453
	ChainMonadTestnet int64 = 10143 // Testnet
	ChainImmutable    int64 = 13371
	ChainMode         int64 = 34443
	ChainArbitrum     int64 = 42161
	ChainCelo         int64 = 42220
	ChainAvalanche    int64 = 43114
	ChainZircuit      int64 = 48900
	ChainLinea        int64 = 59144
	ChainBera         int64 = 80094
	ChainBlast        int64 = 81457
	ChainTaiko        int64 = 167000
	ChainScroll       int64 = 534352
	ChainZkLinkNova   int64 = 810180
	ChainSepolia      int64 = 11155111 // Testnet
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
	ChainSei:       WSEIList[ChainSei],        // 1329
	ChainAvalanche: WAVAXList[ChainAvalanche], // 43114
}

// Map chainID to chain name (self defined).
var ChainNameMap = map[int64]string{
	ChainEthereum:     ChainNameEthereum,
	ChainGoerli:       ChainNameGoerli,
	ChainOptimism:     ChainNameOptimism,
	ChainBSC:          ChainNameBSC,
	ChainOK:           ChainNameOkex,
	ChainGnosis:       ChainNameGnosis,
	ChainHeco:         ChainNameHeco,
	ChainPolygon:      ChainNamePolygon,
	ChainMonad:        ChainNameMonad,
	ChainSonic:        ChainNameSonic,
	ChainTac:          ChainNameTac,
	ChainFantom:       ChainNameFantom,
	ChainFraxtal:      ChainNameFraxtal,
	ChainFilecoin:     ChainNameFilecoin,
	ChainZkSync:       ChainNameZkSync,
	ChainWorldChain:   ChainNameWorldChain,
	ChainPolygonZk:    ChainNamePolygonZkEVM,
	ChainMoonbeam:     ChainNameMoonbeam,
	ChainSei:          ChainNameSei,
	ChainCentrifuge:   ChainNameCentrifuge,
	ChainKava:         ChainNameKava,
	ChainMerlin:       ChainNameMerlin,
	ChainMantle:       ChainNameMantle,
	ChainBase:         ChainNameBase,
	ChainMonadTestnet: ChainNameMonadTestnet,
	ChainImmutable:    ChainNameImmutable,
	ChainMode:         ChainNameMode,
	ChainArbitrum:     ChainNameArbitrum,
	ChainCelo:         ChainNameCelo,
	ChainAvalanche:    ChainNameAvalanche,
	ChainZircuit:      ChainNameZircuit,
	ChainLinea:        ChainNameLinea,
	ChainBera:         ChainNameBera,
	ChainBlast:        ChainNameBlast,
	ChainTaiko:        ChainNameTaiko,
	ChainScroll:       ChainNameScroll,
	ChainZkLinkNova:   ChainNameZkLinkNova,
	ChainSepolia:      ChainNameSepolia,
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
	ChainNameMonad:        ChainMonad,
	ChainNameSonic:        ChainSonic,
	ChainNameTac:          ChainTac,
	ChainNameFantom:       ChainFantom,
	ChainNameFraxtal:      ChainFraxtal,
	ChainNameFilecoin:     ChainFilecoin,
	ChainNameZkSync:       ChainZkSync,
	ChainNameWorldChain:   ChainWorldChain,
	ChainNamePolygonZkEVM: ChainPolygonZk,
	ChainNameMoonbeam:     ChainMoonbeam,
	ChainNameSei:          ChainSei,
	ChainNameCentrifuge:   ChainCentrifuge,
	ChainNameKava:         ChainKava,
	ChainNameMerlin:       ChainMerlin,
	ChainNameMantle:       ChainMantle,
	ChainNameBase:         ChainBase,
	ChainNameMonadTestnet: ChainMonadTestnet,
	ChainNameImmutable:    ChainImmutable,
	ChainNameMode:         ChainMode,
	ChainNameArbitrum:     ChainArbitrum,
	ChainNameCelo:         ChainCelo,
	ChainNameAvalanche:    ChainAvalanche,
	ChainNameZircuit:      ChainZircuit,
	ChainNameLinea:        ChainLinea,
	ChainNameBera:         ChainBera,
	ChainNameBlast:        ChainBlast,
	ChainNameTaiko:        ChainTaiko,
	ChainNameScroll:       ChainScroll,
	ChainNameZkLinkNova:   ChainZkLinkNova,
	ChainNameSepolia:      ChainSepolia,
}
