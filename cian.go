package ethaddr

var CianProtocol string = "cian"

// ---------------------------------------- system ----------------------------------------

// map[network] = address.
var CianErc2612Verifier = map[int64]string{
	ChainAvalanche: "0x25440d9E199974e705a07DF6F2464291D0ba1e2f", // 0x25440d9e199974e705a07df6f2464291d0ba1e2f
	ChainPolygon:   "0xE946Dd7d03F6F5C440F68c84808Ca88d26475FC5", // 0xe946dd7d03f6f5c440f68c84808ca88d26475fc5
}

// map[network] = address.
var CianTokenApprovalVerifier = map[int64]string{
	ChainAvalanche: "0xB5fF7d0Cdd29BFa19CA16cd955E7385BE7E52CCf", // 0xb5ff7d0cdd29bfa19ca16cd955e7385be7e52ccf
	ChainPolygon:   "0x9B2316cfe980515de7430F1c4E831B89a5921137", // 0x9b2316cfe980515de7430f1c4e831b89a5921137
}

// map[network] = address.
var CianWalletFactory = map[int64]string{
	ChainAvalanche: "0x15cbFF12d53e7BdE3f1618844CaaEf99b2836d2A", // 0x15cbff12d53e7bde3f1618844caaef99b2836d2a
	ChainPolygon:   "0x1cB9CF5439dCEd63d8f5B7f1A5bf9834d8076A9a", // 0x1cb9cf5439dced63d8f5b7f1a5bf9834d8076a9a
}

// map[network] = address.
var CianControllerLib = map[int64]string{
	ChainAvalanche: "0x601954e6AfB77Dac21503DbDfA751fbef9eE5374", // 0x601954e6afb77dac21503dbdfa751fbef9ee5374
	ChainPolygon:   "0xff6771a9565F18638faB2972BA7Fc798ad8bCad0", // 0xff6771a9565f18638fab2972ba7fc798ad8bcad0
}

// map[network] = address.
var CianControllerLibSub = map[int64]string{
	ChainAvalanche: "0x9aa8b1998B1882008c407fbB5BF775A5E2d8e544", // 0x9aa8b1998b1882008c407fbb5bf775a5e2d8e544
	ChainPolygon:   "0xEa5f10A0E612316A47123D818E2b597437D19a17", // 0xea5f10a0e612316a47123d818e2b597437d19a17
}

// map[network] = address.
var CianControllerLink = map[int64]string{
	ChainAvalanche: "0x4792e147bCE02E5FF2b1B70416811704B5625446", // 0x4792e147bce02e5ff2b1b70416811704b5625446
	ChainPolygon:   "0x6E3066412B4e67d2933d6023a7c58d63DD8f800a", // 0x6e3066412b4e67d2933d6023a7c58d63dd8f800a
}

// map[network] = address.
var CianAutomation = map[int64]string{
	ChainAvalanche: "0x056c41b8C2A2E7C6454842C9A62050fa1b5ffbAE", // 0x056c41b8c2a2e7c6454842c9a62050fa1b5ffbae
	ChainPolygon:   "0xA79D00C0feA6bAABE8A1fEd0c41C4d36E7B81895", // 0xa79d00c0fea6baabe8a1fed0c41c4d36e7b81895
}

// Same as CianWrappedNativeTokenGateway.
var (
	CianWavaxGateway  = CianWrappedNativeTokenGateway
	CianWmaticGateway = CianWrappedNativeTokenGateway
)

// map[network] = address.
var CianWrappedNativeTokenGateway = map[int64]string{
	ChainAvalanche: "0x28F83cE214462E888787C5cfD0cc08dD439C9920", // 0x28f83ce214462e888787c5cfd0cc08dd439c9920
	ChainPolygon:   "0xdCB3D91555385DaE23e6B966b5626aa7A75Be940", // 0xdcb3d91555385dae23e6b966b5626aa7a75be940
}

// Same as CianFeeBoxNativeToken.
var (
	CianFeeBoxAVAX  = CianFeeBoxNativeToken
	CianFeeBoxMATIC = CianFeeBoxNativeToken
)

// map[network] = address.
var CianFeeBoxNativeToken = map[int64]string{
	ChainAvalanche: "0xec55E7cfebBE4f878E9dD998d3a038458AC3197D", // 0xec55e7cfebbe4f878e9dd998d3a038458ac3197d
	ChainPolygon:   "0x1C8126e02e8A7dAc69FD6444Ef0b8be5430DF776", // 0x1c8126e02e8a7dac69fd6444ef0b8be5430df776
}

// map[network] = address.
var CianFeeBoxSAVAX = map[int64]string{
	ChainAvalanche: "0xb7ead62ca64A98b21C1212BCC82436D7E7d797c3", // 0xb7ead62ca64a98b21c1212bcc82436d7e7d797c3
}

// map[network] = address.
var CianAdapterManager = map[int64]string{
	ChainAvalanche: "0xf8fE4E5Db46D91cC30eae491363dC456e1DaF2fD", // 0xf8fe4e5db46d91cc30eae491363dc456e1daf2fd
	ChainPolygon:   "0x907883da917ca9750ad202ff6395C4C6aB14e60E", // 0x907883da917ca9750ad202ff6395c4c6ab14e60e
}

// map[network] = address.
var CianTimelock = map[int64]string{
	ChainAvalanche: "0xD3812219eb241053F9cf2b43f9B367c0b28E03DA", // 0xd3812219eb241053f9cf2b43f9b367c0b28e03da
	ChainPolygon:   "0xCe672de0D2d38944716c21BCA7DB1164685Af2aC", // 0xce672de0d2d38944716c21bca7db1164685af2ac
}

// ---------------------------------------- multi chain protocols ----------------------------------------

// -------------------- 1 inch --------------------

// map[network] = address.
var CianOneInchAdapter = map[int64]string{
	ChainAvalanche: "0xD52c283b77c1eE742c600875A1d53e7204611A38", // 0xd52c283b77c1ee742c600875a1d53e7204611a38
	ChainPolygon:   "0x9633D6C81E9449B05954B74c257F5964B6864cAA", // 0x9633d6c81e9449b05954b74c257f5964b6864caa
}

// -------------------- paraswap --------------------

// map[network] = address.
var CianParaswapAdapter = map[int64]string{
	ChainAvalanche: "0x330245E9e46b31B33a79fD63C97d2208a4C1233F", // 0x330245e9e46b31b33a79fd63c97d2208a4c1233f
}

// -------------------- aave v3 --------------------

// map[network] = address.
var CianAaveV3Adapter = map[int64]string{
	ChainPolygon: "0x25bbd860d9627b59296E80e95b6Fe2171D45288C", // 0x25bbd860d9627b59296e80e95b6fe2171d45288c
}

// map[network] = address.
var CianAaveV3QueryHelper = map[int64]string{
	ChainPolygon: "0x6B2BA8F249cC1376f2A02A9FaF8BEcA5D7718DCf", // 0x6b2ba8f249cc1376f2a02a9faf8beca5d7718dcf
}

// -------------------- balancer --------------------

// map[network] = address.
var CianBalancerV2Adapter = map[int64]string{
	ChainPolygon: "0x144F69662e8217ae960cEcAb9733a2cc860f779b", // 0x144f69662e8217ae960cecab9733a2cc860f779b
}

// ---------------------------------------- avalanche only protocols ----------------------------------------

// -------------------- benqi --------------------

// map[network] = address.
var CianSavaxAdapter = map[int64]string{
	ChainAvalanche: "0x83B15AB252482E8AfB0E47460B46AaE5F145ec17", // 0x83b15ab252482e8afb0e47460b46aae5f145ec17
}

// map[network] = address.
var CianBenqiAdapter = map[int64]string{
	ChainAvalanche: "0xe7a5b5783bee4C91c2Bdfb00FF5a34426b6b8a02", // 0xe7a5b5783bee4c91c2bdfb00ff5a34426b6b8a02
}

// -------------------- traderjoe --------------------

// map[network] = address.
var CianTraderJoeAdapter = map[int64]string{
	ChainAvalanche: "0xDA7fBbDFf6225e37D349676f7b65684E96dd5C16", // 0xda7fbbdff6225e37d349676f7b65684e96dd5c16
}

// map[network] = address.
var CianBankerJoeAdapter = map[int64]string{
	ChainAvalanche: "0x123d4F3126B0F57B86d15382ec72A444Bb6E77de", // 0x123d4f3126b0f57b86d15382ec72a444bb6e77de
}

// ---------------------------------------- polygon only protocols ----------------------------------------

// -------------------- stader --------------------

// map[network] = address.
var CianStaderAdapter = map[int64]string{
	ChainPolygon: "0x294fE934F47D95A54a436A148963da3FEF4E8313", // 0x294fe934f47d95a54a436a148963da3fef4e8313
}

// map[network] = address.
var CianStaderAirdrop = map[int64]string{
	ChainPolygon: "0x406e1e0e3cb4201B4AEe409Ad2f6Cd56d3242De7", // 0x406e1e0e3cb4201b4aee409ad2f6cd56d3242de7
}

// -------------------- quickswap --------------------

// map[network] = address.
var CianQuickSwapAdapter = map[int64]string{
	ChainPolygon: "0x054Ab57D364730ed9ed665FF97A92E9813c42515", // 0x054ab57d364730ed9ed665ff97a92e9813c42515
}
