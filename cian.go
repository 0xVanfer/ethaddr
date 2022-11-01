package ethaddr

import "github.com/0xVanfer/chainId"

var CianProtocol string = "cian"

// ---------------------------------------- system ----------------------------------------

// map[network] = address.
var CianErc2612Verifier = map[string]string{
	chainId.AvalancheChainName: "0x25440d9e199974e705a07df6f2464291d0ba1e2f",
	chainId.PolygonChainName:   "0xE946Dd7d03F6F5C440F68c84808Ca88d26475FC5",
}

// map[network] = address.
var CianTokenApprovalVerifier = map[string]string{
	chainId.AvalancheChainName: "0xb5ff7d0cdd29bfa19ca16cd955e7385be7e52ccf",
	chainId.PolygonChainName:   "0x9b2316cfe980515de7430f1c4e831b89a5921137",
}

// map[network] = address.
var CianWalletFactory = map[string]string{
	chainId.AvalancheChainName: "0x15cbff12d53e7bde3f1618844caaef99b2836d2a",
	chainId.PolygonChainName:   "0x1cb9cf5439dced63d8f5b7f1a5bf9834d8076a9a",
}

// map[network] = address.
var CianControllerLib = map[string]string{
	chainId.AvalancheChainName: "0x601954e6afb77dac21503dbdfa751fbef9ee5374",
	chainId.PolygonChainName:   "0xff6771a9565f18638fab2972ba7fc798ad8bcad0",
}

// map[network] = address.
var CianControllerLibSub = map[string]string{
	chainId.AvalancheChainName: "0x9aa8b1998b1882008c407fbb5bf775a5e2d8e544",
	chainId.PolygonChainName:   "0xea5f10a0e612316a47123d818e2b597437d19a17",
}

// map[network] = address.
var CianControllerLink = map[string]string{
	chainId.AvalancheChainName: "0x4792e147bce02e5ff2b1b70416811704b5625446",
	chainId.PolygonChainName:   "0x6e3066412b4e67d2933d6023a7c58d63dd8f800a",
}

// map[network] = address.
var CianAutomation = map[string]string{
	chainId.AvalancheChainName: "0x056c41b8c2a2e7c6454842c9a62050fa1b5ffbae",
	chainId.PolygonChainName:   "0xa79d00c0fea6baabe8a1fed0c41c4d36e7b81895",
}

// Same as CianWrappedChainTokenGateway.
var (
	CianWavaxGateway  = CianWrappedChainTokenGateway
	CianWmaticGateway = CianWrappedChainTokenGateway
)

// map[network] = address.
var CianWrappedChainTokenGateway = map[string]string{
	chainId.AvalancheChainName: "0x28f83ce214462e888787c5cfd0cc08dd439c9920",
	chainId.PolygonChainName:   "0xdcb3d91555385dae23e6b966b5626aa7a75be940",
}

// Same as CianFeeBoxChainToken.
var (
	CianFeeBoxAVAX  = CianFeeBoxChainToken
	CianFeeBoxMATIC = CianFeeBoxChainToken
)

// map[network] = address.
var CianFeeBoxChainToken = map[string]string{
	chainId.AvalancheChainName: "0xec55e7cfebbe4f878e9dd998d3a038458ac3197d",
	chainId.PolygonChainName:   "0x1c8126e02e8a7dac69fd6444ef0b8be5430df776",
}

// map[network] = address.
var CianFeeBoxSAVAX = map[string]string{
	chainId.AvalancheChainName: "0xb7ead62ca64a98b21c1212bcc82436d7e7d797c3",
}

// map[network] = address.
var CianAdapterManager = map[string]string{
	chainId.AvalancheChainName: "0xf8fe4e5db46d91cc30eae491363dc456e1daf2fd",
	chainId.PolygonChainName:   "0x907883da917ca9750ad202ff6395c4c6ab14e60e",
}

// map[network] = address.
var CianTimelock = map[string]string{
	chainId.AvalancheChainName: "0xd3812219eb241053f9cf2b43f9b367c0b28e03da",
	chainId.PolygonChainName:   "0xce672de0d2d38944716c21bca7db1164685af2ac",
}

// ---------------------------------------- multi chain protocols ----------------------------------------

// -------------------- 1 inch --------------------

// map[network] = address.
var CianOneInchAdapter = map[string]string{
	chainId.AvalancheChainName: "0xd52c283b77c1ee742c600875a1d53e7204611a38",
	chainId.PolygonChainName:   "0x9633d6c81e9449b05954b74c257f5964b6864caa",
}

// -------------------- paraswap --------------------

// map[network] = address.
var CianParaswapAdapter = map[string]string{
	chainId.AvalancheChainName: "0x330245e9e46b31b33a79fd63c97d2208a4c1233f",
}

// -------------------- aave v3 --------------------

// map[network] = address.
var CianAaveV3Adapter = map[string]string{
	chainId.PolygonChainName: "0x25bbd860d9627b59296e80e95b6fe2171d45288c",
}

// map[network] = address.
var CianAaveV3QueryHelper = map[string]string{
	chainId.PolygonChainName: "0x6b2ba8f249cc1376f2a02a9faf8beca5d7718dcf",
}

// -------------------- balancer --------------------

// map[network] = address.
var CianBalancerV2Adapter = map[string]string{
	chainId.PolygonChainName: "0x144f69662e8217ae960cecab9733a2cc860f779b",
}

// ---------------------------------------- avalanche only protocols ----------------------------------------

// -------------------- benqi --------------------

// map[network] = address.
var CianSavaxAdapter = map[string]string{
	chainId.AvalancheChainName: "0x83b15ab252482e8afb0e47460b46aae5f145ec17",
}

// map[network] = address.
var CianBenqiAdapter = map[string]string{
	chainId.AvalancheChainName: "0xe7a5b5783bee4c91c2bdfb00ff5a34426b6b8a02",
}

// -------------------- traderjoe --------------------

// map[network] = address.
var CianTraderJoeAdapter = map[string]string{
	chainId.AvalancheChainName: "0xda7fbbdff6225e37d349676f7b65684e96dd5c16",
}

// map[network] = address.
var CianBankerJoeAdapter = map[string]string{
	chainId.AvalancheChainName: "0x123d4f3126b0f57b86d15382ec72a444bb6e77de",
}

// ---------------------------------------- polygon only protocols ----------------------------------------

// -------------------- stader --------------------

// map[network] = address.
var CianStaderAdapter = map[string]string{
	chainId.PolygonChainName: "0x294fe934f47d95a54a436a148963da3fef4e8313",
}

// map[network] = address.
var CianStaderAirdrop = map[string]string{
	chainId.PolygonChainName: "0x406e1e0e3cb4201b4aee409ad2f6cd56d3242de7",
}

// -------------------- quickswap --------------------

// map[network] = address.
var CianQuickSwapAdapter = map[string]string{
	chainId.PolygonChainName: "0x054ab57d364730ed9ed665ff97a92e9813c42515",
}
