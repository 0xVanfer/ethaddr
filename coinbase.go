package ethaddr

// Website: https://www.coinbase.com/home
const CoinbaseProtocol string = "coinbase"

// Blockscan: https://basescan.org/
const ChainNameBase string = "base"

// USD Base Coin
//
// map[chainID] = address.
var USDbCList = map[int64]string{
	ChainBase: "0xd9aAEc86B65D86f6A7B5B1b0c42FFA531710b6CA", // USDbC, 0xd9aaec86b65d86f6a7b5b1b0c42ffa531710b6ca
}

// Coinbase cbETH.
//
// map[chainID] = address.
var CbETHList = map[int64]string{
	ChainArbitrum: "0x1DEBd73E752bEaF79865Fd6446b0c970EaE7732f", // cbETH, 0x1debd73e752beaf79865fd6446b0c970eae7732f
	ChainBase:     "0x2Ae3F1Ec7F1F5012CFEab0185bfc7aa3cf0DEc22", // cbETH, 0x2ae3f1ec7f1f5012cfeab0185bfc7aa3cf0dec22
	ChainEthereum: "0xBe9895146f7AF43049ca1c1AE358B0541Ea49704", // cbETH, 0xbe9895146f7af43049ca1c1ae358b0541ea49704
	ChainOptimism: "0xadDb6A0412DE1BA0F936DCaeb8Aaa24578dcF3B2", // cbETH, 0xaddb6a0412de1ba0f936dcaeb8aaa24578dcf3b2
}

// Coinbase cbBTC.
//
// map[chainID] = address.
var CbBTCList = map[int64]string{
	ChainBase:     "0xcbB7C0000aB88B473b1f5aFd9ef808440eed33Bf", // cbBTC, 0xcbb7c0000ab88b473b1f5afd9ef808440eed33bf
	ChainEthereum: "0xcbB7C0000aB88B473b1f5aFd9ef808440eed33Bf", // cbBTC, 0xcbb7c0000ab88b473b1f5afd9ef808440eed33bf
}
