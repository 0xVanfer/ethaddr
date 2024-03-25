package ethaddr

import "github.com/0xVanfer/chainId"

const (
	ChainEthereum  string = chainId.EthereumChainName     // 1
	ChainGoerli    string = chainId.GoerliChainName       // 5
	ChainOptimism  string = chainId.OptimismChainName     // 10
	ChainBSC       string = chainId.BNBSmartChainName     // 56
	ChainOK        string = chainId.OkChainName           // 66
	ChainGnosis    string = chainId.GnosisChainName       // 100
	ChainXDai      string = ChainGnosis                   // XDai has renamed as gnosis
	ChainPolygon   string = chainId.PolygonChainName      // 137
	ChainBlast     string = chainId.BlastChainName        // 238
	ChainFantom    string = chainId.FantomChainName       // 250
	ChainFilecoin  string = chainId.FilecoinChainName     // 314
	ChainPolygonZk string = chainId.PolygonZkEVMChainName // 1101
	ChainBase      string = chainId.BaseChainName         // 8453
	ChainArbitrum  string = chainId.ArbitrumChainName     // 42161
	ChainAvalanche string = chainId.AvalancheChainName    // 43114
	ChainScroll    string = chainId.ScrollChainName       // 534352
)
