package ethaddr

// Synthetix token: SNX.
//
// map[network] = address.
var SynthetixTokenList = map[string]string{
	ChainEthereum: "0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F", // SNX, 0xc011a73ee8576fb46f5e1c5751ca3b9fe0af2a6f
}

// Same as SynthetixTokenList.
var SNXList = SynthetixTokenList

// Synthetix stable coin: sUSD.
//
// map[network] = address.
var SUSDList = map[string]string{
	ChainEthereum: "0x57Ab1ec28D129707052df4dF418D58a2D46d5f51", // sUSD, 0x57ab1ec28d129707052df4df418d58a2d46d5f51
	ChainOptimism: "0x8c6f28f2F1A3C87F0f938b96d27520d9751ec8d9", // sUSD, 0x8c6f28f2f1a3c87f0f938b96d27520d9751ec8d9
}
