package ethaddr

// Website: https://balancer.fi/
//
// Docs: https://docs.balancer.fi/
//
// Deployed contracts: https://docs.balancer.fi/reference/contracts/deployment-addresses/mainnet.html
const BalancerProtocol string = "balancer"

// Balancer token: BAL.
//
// map[network] = address.
var BalancerTokenList = map[string]string{
	ChainArbitrum: "0x040d1EdC9569d4Bab2D15287Dc5A4F10F56a56B8", // BAL, 0x040d1edc9569d4bab2d15287dc5a4f10f56a56b8
	ChainEthereum: "0xba100000625a3754423978a60c9317c58a424e3D", // BAL, 0xba100000625a3754423978a60c9317c58a424e3d
	ChainPolygon:  "0x9a71012B13CA4d3D0Cdc72A177DF3ef03b0E76A3", // BAL, 0x9a71012b13ca4d3d0cdc72a177df3ef03b0e76a3
}

// Same as BalancerTokenList.
var BALList = BalancerTokenList

// Balancer v2 vault.
//
// map[network] = address.
var BalancerV2VaultList = map[string]string{
	ChainPolygon: "0xBA12222222228d8Ba445958a75a0704d566BF2C8", // 0xba12222222228d8ba445958a75a0704d566bf2c8
}

// Balancer helper.
//
// map[network] = address.
var BalancherHelperList = map[string]string{
	ChainPolygon: "0x239e55F427D44C3cc793f49bFB507ebe76638a2b", // 0x239e55f427d44c3cc793f49bfb507ebe76638a2b
}
