package ethaddr

import "github.com/0xVanfer/chainId"

const (
	ChainEthereum   string = chainId.EthereumChainName     // 1
	ChainGoerli     string = chainId.GoerliChainName       // 5
	ChainOptimism   string = chainId.OptimismChainName     // 10
	ChainBSC        string = chainId.BNBSmartChainName     // 56
	ChainOK         string = chainId.OkChainName           // 66
	ChainGnosis     string = chainId.GnosisChainName       // 100
	ChainHeco       string = chainId.HecoChainName         // 128
	ChainPolygon    string = chainId.PolygonChainName      // 137
	ChainBlast      string = chainId.BlastChainName        // 238
	ChainFantom     string = chainId.FantomChainName       // 250
	ChainFraxtal    string = chainId.FraxtalChainName      // 252
	ChainFilecoin   string = chainId.FilecoinChainName     // 314
	ChainPolygonZk  string = chainId.PolygonZkEVMChainName // 1101
	ChainMoonbeam   string = chainId.MoonbeamChainName     // 1284
	ChainCentrifuge string = chainId.CentrifugeChainName   // 2031
	ChainKava       string = chainId.KavaChainName         // 2222
	ChainMantle     string = chainId.MantleChainName       // 5000
	ChainBase       string = chainId.BaseChainName         // 8453
	ChainImmutable  string = chainId.ImmutableChainName    // 13371
	ChainArbitrum   string = chainId.ArbitrumChainName     // 42161
	ChainCelo       string = chainId.CeloChainName         // 42220
	ChainAvalanche  string = chainId.AvalancheChainName    // 43114
	ChainLinea      string = chainId.LineaChainName        // 59144
	ChainScroll     string = chainId.ScrollChainName       // 534352
)

// Deprecated: Use ChainGnosis instead.
//
// XDai has renamed as gnosis.
const ChainXDai string = ChainGnosis
