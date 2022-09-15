package ethaddr

import "github.com/0xVanfer/chainId"

// Synthetix token: SNX.
var SynthetixTokenList = map[string]string{
	chainId.EthereumChainName: "0xc011a73ee8576fb46f5e1c5751ca3b9fe0af2a6f", // SNX
}

// Synthetix stable coin: sUSD.
var SUSDList = map[string]string{
	chainId.EthereumChainName: "0x57ab1ec28d129707052df4df418d58a2d46d5f51", // sUSD
	chainId.OptimismChainName: "0x8c6f28f2f1a3c87f0f938b96d27520d9751ec8d9", // sUSD
}
