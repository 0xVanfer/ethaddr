package ethaddr

// WARNING: Kyber was hacked and has very little tvl now. Mar.22.2024
//
// Docs: https://docs.kyberswap.com/
//
// Deployed contracts: https://docs.kyberswap.com/liquidity-solutions/kyberswap-classic/contracts/classic-contract-addresses
const KyberProtocol string = "kyber"

// Kyber old token: KNC. Kyber has another new token, also called KNC.
// You can find the new one by "KyberTokenList" in this package.
// However, this old version token is still in use, such as AAVE V2 Ethereum.
//
// map[network] = address.
var KyberOldTokenList = map[string]string{
	ChainEthereum: "0xdd974d5c2e2928dea5f71b9825b8b646686bd200", // KNC
}

// Kyber token in use: KNC.
//
// map[network] = address.
var KyberTokenList = map[string]string{
	ChainArbitrum:  "0xe4DDDfe67E7164b0FE14E218d80dC4C08eDC01cB", // KNC, 0xe4dddfe67e7164b0fe14e218d80dc4c08edc01cb
	ChainAvalanche: "0x39fC9e94Caeacb435842FADeDeCB783589F50f5f", // KNC, 0x39fc9e94caeacb435842fadedecb783589f50f5f
	ChainBSC:       "0xfe56d5892BDffC7BF58f2E84BE1b2C32D21C308b", // KNC, 0xfe56d5892bdffc7bf58f2e84be1b2c32d21c308b
	ChainEthereum:  "0xdeFA4e8a7bcBA345F687a2f1456F5Edd9CE97202", // KNC, 0xdefa4e8a7bcba345f687a2f1456f5edd9ce97202
	ChainPolygon:   "0x1C954E8fe737F99f68Fa1CCda3e51ebDB291948C", // KNC, 0x1c954e8fe737f99f68fa1ccda3e51ebdb291948c
}

// Same as KyberTokenList.
var KNCList = KyberTokenList

// Kyber factory.
//
// map[network] = address.
var KyberFactoryList = map[string]string{
	ChainArbitrum:  "0x51E8D106C646cA58Caf32A47812e95887C071a62", // 0x51e8d106c646ca58caf32a47812e95887c071a62
	ChainAvalanche: "0x10908C875D865C66f271F5d3949848971c9595C9", // 0x10908c875d865c66f271f5d3949848971c9595c9
	ChainBSC:       "0x878dFE971d44e9122048308301F540910Bbd934c", // 0x878dfe971d44e9122048308301f540910bbd934c
	ChainEthereum:  "0x833e4083B7ae46CeA85695c4f7ed25CDAd8886dE", // 0x833e4083b7ae46cea85695c4f7ed25cdad8886de
	ChainFantom:    "0x78df70615ffc8066cc0887917f2Cd72092C86409", // 0x78df70615ffc8066cc0887917f2cd72092c86409
	ChainPolygon:   "0x5F1fe642060B5B9658C15721Ea22E982643c095c", // 0x5f1fe642060b5b9658c15721ea22e982643c095c
}

// Kyber DMM factory.
//
// map[network] = address.
var KyberDMMFactoryList = map[string]string{
	ChainArbitrum: "0xD9bfE9979e9CA4b2fe84bA5d4Cf963bBcB376974", // 0xd9bfe9979e9ca4b2fe84ba5d4cf963bbcb376974
}

// Kyber router.
//
// map[network] = address.
var KyberRouterList = map[string]string{
	ChainArbitrum:  "0xC3E2aED41ECdFB1ad41ED20D45377Da98D5489dD", // 0xc3e2aed41ecdfb1ad41ed20d45377da98d5489dd
	ChainAvalanche: "0x8Efa5A9AD6D594Cf76830267077B78cE0Bc5A5F8", // 0x8efa5a9ad6d594cf76830267077b78ce0bc5a5f8
	ChainBSC:       "0x78df70615ffc8066cc0887917f2Cd72092C86409", // 0x78df70615ffc8066cc0887917f2cd72092c86409
	ChainEthereum:  "0x1c87257F5e8609940Bc751a07BB085Bb7f8cDBE6", // 0x1c87257f5e8609940bc751a07bb085bb7f8cdbe6
	ChainFantom:    "0x5d5A5a0a465129848c2549669e12cDC2f8DE039A", // 0x5d5a5a0a465129848c2549669e12cdc2f8de039a
	ChainPolygon:   "0x546C79662E028B661dFB4767664d0273184E4dD1", // 0x546c79662e028b661dfb4767664d0273184e4dd1
}

// Kyber DMM router.
//
// map[network] = address.
var KyberDMMRouterList = map[string]string{
	ChainArbitrum: "0xEaE47c5D99f7B31165a7f0c5f7E0D6afA25CFd55", // 0xeae47c5d99f7b31165a7f0c5f7e0d6afa25cfd55
}

// Kyber reward locker.
//
// map[network] = address.
var KyberRewardLockerList = map[string]string{
	ChainAvalanche: "0xf530a090EF6481cfB33F98c63532E7745abab58A", // 0xf530a090ef6481cfb33f98c63532e7745abab58a
	ChainBSC:       "0xfab5186A194588F5AD5074Bd52659302906B4522", // 0xfab5186a194588f5ad5074bd52659302906b4522
	ChainEthereum:  "0xfab5186A194588F5AD5074Bd52659302906B4522", // 0xfab5186a194588f5ad5074bd52659302906b4522
	ChainPolygon:   "0x063DD8b5a42AaE93a014ce5FAbB5B70474667961", // 0x063dd8b5a42aae93a014ce5fabb5b70474667961
}

// Kyber migrator.
//
// map[network] = address.
var KyberMigratorList = map[string]string{
	ChainEthereum: "0x6A65e062cE8290007301296F3C6AE446Af7BDEeC", // 0x6a65e062ce8290007301296f3c6ae446af7bdeec
}
