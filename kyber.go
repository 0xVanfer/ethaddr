package ethaddr

import "github.com/0xVanfer/chainId"

const KyberProtocol string = "kyber"

// Kyber old token: KNC.
//
// map[network] = address.
var KyberOldTokenList = map[string]string{
	chainId.EthereumChainName: "0xdd974d5c2e2928dea5f71b9825b8b646686bd200", // KNC
}

// Kyber token in use: KNC.
//
// map[network] = address.
var KyberTokenList = map[string]string{
	chainId.EthereumChainName:     "0xdefa4e8a7bcba345f687a2f1456f5edd9ce97202", // KNC
	chainId.PolygonChainName:      "0x1c954e8fe737f99f68fa1ccda3e51ebdb291948c", // KNC
	chainId.BinanceSmartChainName: "0xfe56d5892bdffc7bf58f2e84be1b2c32d21c308b", // KNC
	chainId.AvalancheChainName:    "0x39fC9e94Caeacb435842FADeDeCB783589F50f5f", // KNC
}

// Same as KyberTokenList.
var KNCList = KyberTokenList

// Kyber factory.
//
// map[network] = address.
var KyberFactoryList = map[string]string{
	chainId.EthereumChainName:     "0x833e4083b7ae46cea85695c4f7ed25cdad8886de",
	chainId.ArbitrumChainName:     "0x51e8d106c646ca58caf32a47812e95887c071a62",
	chainId.PolygonChainName:      "0x5f1fe642060b5b9658c15721ea22e982643c095c",
	chainId.BinanceSmartChainName: "0x878dfe971d44e9122048308301f540910bbd934c",
	chainId.FantomChainName:       "0x78df70615ffc8066cc0887917f2cd72092c86409",
	chainId.AvalancheChainName:    "0x10908c875d865c66f271f5d3949848971c9595c9",
}

// Kyber DMM factory.
//
// map[network] = address.
var KyberDMMFactoryList = map[string]string{
	chainId.ArbitrumChainName: "0xd9bfe9979e9ca4b2fe84ba5d4cf963bbcb376974",
}

// Kyber router.
//
// map[network] = address.
var KyberRouterList = map[string]string{
	chainId.EthereumChainName:     "0x1c87257f5e8609940bc751a07bb085bb7f8cdbe6",
	chainId.ArbitrumChainName:     "0xc3e2aed41ecdfb1ad41ed20d45377da98d5489dd",
	chainId.PolygonChainName:      "0x546c79662e028b661dfb4767664d0273184e4dd1",
	chainId.BinanceSmartChainName: "0x78df70615ffc8066cc0887917f2cd72092c86409",
	chainId.FantomChainName:       "0x5d5a5a0a465129848c2549669e12cdc2f8de039a",
	chainId.AvalancheChainName:    "0x8efa5a9ad6d594cf76830267077b78ce0bc5a5f8",
}

// Kyber DMM router.
//
// map[network] = address.
var KyberDMMRouterList = map[string]string{
	chainId.ArbitrumChainName: "0xeae47c5d99f7b31165a7f0c5f7e0d6afa25cfd55",
}

// Kyber reward locker.
//
// map[network] = address.
var KyberRewardLockerList = map[string]string{
	chainId.EthereumChainName:     "0xfab5186a194588f5ad5074bd52659302906b4522",
	chainId.PolygonChainName:      "0x063dd8b5a42aae93a014ce5fabb5b70474667961",
	chainId.BinanceSmartChainName: "0xfab5186a194588f5ad5074bd52659302906b4522",
	chainId.AvalancheChainName:    "0xf530a090ef6481cfb33f98c63532e7745abab58a",
}

// Kyber migrator.
//
// map[network] = address.
var KyberMigratorList = map[string]string{
	chainId.EthereumChainName: "0x6a65e062ce8290007301296f3c6ae446af7bdeec",
}
