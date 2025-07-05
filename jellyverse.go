package ethaddr

// Website: https://jellyverse.org/
//
// A fork of balancer.
const JellyverseProtocol string = "jellyverse"

// Jellyverse Token JLY.
//
// map[chainID] = address
var JellyverseTokenList = map[int64]string{
	ChainSei: "0xDD7d5e4Ea2125d43C16eEd8f1FFeFffa2F4b4aF6", // JLY, 0xdd7d5e4ea2125d43c16eed8f1ffefffa2f4b4af6
}

var JLYList = JellyverseTokenList

// Jellyverse Vault. (balancer fork)
//
// map[chainID] = address
var JellyverseVault = map[int64]string{
	ChainSei: "0xFB43069f6d0473B85686a85F4Ce4Fc1FD8F00875", // 0xfb43069f6d0473b85686a85f4ce4fc1fd8f00875
}

// Can get by Vault function "getPoolTokens"
//
// map[chainID][pool] = tokens
var JellyversePoolTokens = map[int64]map[string][]string{
	// Updated: Feb.18.2025
	ChainSei: {
		// JLY/50-SEI/50
		"0x55D45C15A95ABfbbCE3C88F90adcD62cD873A2DB": {JLYList[ChainSei], WSEIList[ChainSei]},
		// JLY/50-USDT/50
		"0x8894B8381dcA1d322453282E01aD6D29FC8450DD": {JLYList[ChainSei], USDTList[ChainSei]},
		// JLY/50-USDC/50
		"0x7FbAa67A43F7D5662d4933598E03EDf6D97FF261": {JLYList[ChainSei], USDCList[ChainSei]},
		// JLY/50-wBTC/50
		"0xA5dC1F0A466D5a9CfB4d709Da1a564B137b005bB": {JLYList[ChainSei], WBTCList[ChainSei]},
		// JLY/50-WETH/50
		"0xE026011175CbfA6a754f9db16541BbEDCca5864c": {JLYList[ChainSei], WETHList[ChainSei]},
		// GEM/75-USDC/25
		"0xf3837Ff4AFc5aaf7D1DA30AfD1272A0100D527B5": {
			"0x372B2dC06478AA2c8182EeE0f12eA0e9A15E2913", // GEM, a RWA token.
			USDCList[ChainSei],
		},
		// jUSD/(USDC-USDT-FRAX)
		"0xE42CC0395c68f73dD16fEbed82BCC701011864B6": {
			"0x111A28eE3035E06793CbAc31089dEFeeFa24d438", // FRAX/USDC/USDT, seems to be another jellyverse pool. (balancer fork)
			"0x4c6Dd2CA85Ca55C4607Cd66A7EBdD2C9b58112Cf", // jUSDv1
			"0xE42CC0395c68f73dD16fEbed82BCC701011864B6", // jUSD/(USDC-USDT-FRAX) ??? pool itself as a token???
		},
	},
}

// map[chainID][pool] = poolID
var JellyversePoolIDs = map[int64]map[string]string{
	// Updated: Feb.18.2025
	ChainSei: {
		// JLY/50-SEI/50
		"0x55D45C15A95ABfbbCE3C88F90adcD62cD873A2DB": "0x55d45c15a95abfbbce3c88f90adcd62cd873a2db000200000000000000000005",
		// JLY/50-USDT/50
		"0x8894B8381dcA1d322453282E01aD6D29FC8450DD": "0x8894b8381dca1d322453282e01ad6d29fc8450dd000200000000000000000007",
		// JLY/50-USDC/50
		"0x7FbAa67A43F7D5662d4933598E03EDf6D97FF261": "0x7fbaa67a43f7d5662d4933598e03edf6d97ff261000200000000000000000008",
		// JLY/50-wBTC/50
		"0xA5dC1F0A466D5a9CfB4d709Da1a564B137b005bB": "0xa5dc1f0a466d5a9cfb4d709da1a564b137b005bb0002000000000000000000ed",
		// JLY/50-WETH/50
		"0xE026011175CbfA6a754f9db16541BbEDCca5864c": "0xe026011175cbfa6a754f9db16541bbedcca5864c000200000000000000000015",
		// GEM/75-USDC/25
		"0xf3837Ff4AFc5aaf7D1DA30AfD1272A0100D527B5": "0xf3837ff4afc5aaf7d1da30afd1272a0100d527b50002000000000000000000a7",
		// jUSD/(USDC-USDT-FRAX)
		"0xE42CC0395c68f73dD16fEbed82BCC701011864B6": "0xe42cc0395c68f73dd16febed82bcc701011864b6000000000000000000000102",
	},
}
