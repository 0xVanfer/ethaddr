package ethaddr

// Website: https://app.morpho.org/
//
// Docs: https://docs.morpho.org/
//
// Deployed contracts: https://docs.morpho.org/addresses
const MorphoBlueProtocol string = "morphoblue"

// MORPHO token.
//
// map[chainID] = address.
var MorphoTokenList = map[int64]string{
	ChainEthereum: "0x9994E35Db50125E0DF82e4c2dde62496CE330999", // MORPHO, 0x9994e35db50125e0df82e4c2dde62496ce330999
}

var MorphoBlueList = map[int64]string{
	ChainArbitrum: "0x6c247b1F6182318877311737BaC0844bAa518F5e", // 0x6c247b1f6182318877311737bac0844baa518f5e
	ChainBase:     "0xBBBBBbbBBb9cC5e90e3b3Af64bdAF62C37EEFFCb", // 0xbbbbbbbbbb9cc5e90e3b3af64bdaf62c37eeffcb
	ChainEthereum: "0xBBBBBbbBBb9cC5e90e3b3Af64bdAF62C37EEFFCb", // 0xbbbbbbbbbb9cc5e90e3b3af64bdaf62c37eeffcb
	ChainFraxtal:  "0xa6030627d724bA78a59aCf43Be7550b4C5a0653b", // 0xa6030627d724ba78a59acf43be7550b4c5a0653b
	ChainMode:     "0xd85cE6BD68487E0AaFb0858FDE1Cd18c76840564", // 0xd85ce6bd68487e0aafb0858fde1cd18c76840564
	ChainOptimism: "0xce95AfbB8EA029495c66020883F87aaE8864AF92", // 0xce95afbb8ea029495c66020883f87aae8864af92
	ChainPolygon:  "0x1bF0c2541F820E775182832f06c0B7Fc27A25f67", // 0x1bf0c2541f820e775182832f06c0b7fc27a25f67
	ChainScroll:   "0x2d012EdbAdc37eDc2BC62791B666f9193FDF5a55", // 0x2d012edbadc37edc2bc62791b666f9193fdf5a55
	ChainSonic:    "0xd6c916eB7542D0Ad3f18AEd0FCBD50C582cfa95f", // 0xd6c916eb7542d0ad3f18aed0fcbd50c582cfa95f
}

var MORPHOList = MorphoTokenList

// Deprecated: this map is meaningless.
//
// map[network][collateral token][borrow token] = id
var MorphoMarketIDMapList = map[int64]map[string]map[string][]string{
	ChainEthereum: {
		WSTETHList[ChainEthereum]: {
			WETHList[ChainEthereum]: {
				"0xc54d7acf14de29e0e5527cabd7a576506870346a78a11a6762e2cca66322ec41", // wstETH - WETH 94.5% (2024.02.02)
			},
			USDCList[ChainEthereum]: {
				"0xb323495f7e4148be5643a4ea4a8221eef163e4bccfdedc2a6f4696baacbc86cc", // wstETH - USDC 86% (2024.02.02)
			},
			USDTList[ChainEthereum]: {
				"0xe7e9694b754c4d4f7e21faf7223f6fa71abaeb10296a4c43a54a7977149687d2", // wstETH - USDT 86% (2024.02.02)
			},
		},
		SDAIList[ChainEthereum]: {
			USDCList[ChainEthereum]: {
				"0x06f2842602373d247c4934f7656e513955ccc4c377f0febc0d9ca2c3bcc191b1", // sDAI - USDC 96.5% (2024.02.02)
			},
		},
		WETHList[ChainEthereum]: {
			USDCList[ChainEthereum]: {
				"0xf9acc677910cc17f650416a22e2a14d5da7ccb9626db18f1bf94efe64f92b372", // WETH - USDC 91.5% (2024.02.02)
				"0x7dde86a1e94561d9690ec678db673c1a6396365f7d1d65e129c5fff0990ff758", // WETH - USDC 86% (2024.02.02)
			},
			USDTList[ChainEthereum]: {
				"0x608929d6de2a10bacf1046ff157ae38df5b9f466fb89413211efb8f63c63833a", // WETH - USDT 86% (2024.02.02)
				"0xdbffac82c2dc7e8aa781bd05746530b0068d80929f23ac1628580e27810bc0c5", // WETH - USDT 91.5% (2024.02.02)
			},
		},
		WBTCList[ChainEthereum]: {
			USDCList[ChainEthereum]: {
				"0x3a85e619751152991742810df6ec69ce473daef99e28a64ab2340d7b7ccfee49", // WBTC - USDC 86% (2024.02.02)
			},
			USDTList[ChainEthereum]: {
				"0xa921ef34e2fc7a27ccc50ae7e4b154e16c9799d3387076c421423ef52ac4df99", // WBTC - USDT 86% (2024.02.02)
			},
		},
		OSETHList[ChainEthereum]: {
			WETHList[ChainEthereum]: {
				"0xd5211d0e3f4a30d5c98653d988585792bb7812221f04801be73a44ceecb11e89", // osETH - WETH 86% (2024.02.02)
			},
		},
		WEETHList[ChainEthereum]: {
			WETHList[ChainEthereum]: {
				"0x698fe98247a40c5771537b5786b2f3f9d78eb487b4ce4d75533cd0e94d88a115", // weETH - WETH 86% (2024.02.02)
			},
		},
	},
}

// Deprecated: this map is meaningless.
//
// map[network] = [all ids]
var MorphoMarketIDs = map[int64][]string{
	ChainEthereum: {
		"0xc54d7acf14de29e0e5527cabd7a576506870346a78a11a6762e2cca66322ec41", // wstETH - WETH 94.5%
		"0x06f2842602373d247c4934f7656e513955ccc4c377f0febc0d9ca2c3bcc191b1", // sDAI - USDC 96.5% (2024.02.02)
		"0xf9acc677910cc17f650416a22e2a14d5da7ccb9626db18f1bf94efe64f92b372", // WETH - USDC 91.5% (2024.02.02)
		"0x7dde86a1e94561d9690ec678db673c1a6396365f7d1d65e129c5fff0990ff758", // WETH - USDC 86% (2024.02.02)
		"0xb323495f7e4148be5643a4ea4a8221eef163e4bccfdedc2a6f4696baacbc86cc", // wstETH - USDC 86% (2024.02.02)
		"0x495130878b7d2f1391e21589a8bcaef22cbc7e1fbbd6866127193b3cc239d8b1", // wbIB01 - USDC 96.5% (2024.02.02) // Not listed in MorphoMarketIDMapList
		"0x3a85e619751152991742810df6ec69ce473daef99e28a64ab2340d7b7ccfee49", // WBTC - USDC 86% (2024.02.02)
		"0xd5211d0e3f4a30d5c98653d988585792bb7812221f04801be73a44ceecb11e89", // osETH - WETH 86% (2024.02.02)
		"0x698fe98247a40c5771537b5786b2f3f9d78eb487b4ce4d75533cd0e94d88a115", // weETH - WETH 86% (2024.02.02)
		"0x608929d6de2a10bacf1046ff157ae38df5b9f466fb89413211efb8f63c63833a", // WETH - USDT 86% (2024.02.02)
		"0xdbffac82c2dc7e8aa781bd05746530b0068d80929f23ac1628580e27810bc0c5", // WETH - USDT 91.5% (2024.02.02)
		"0xa921ef34e2fc7a27ccc50ae7e4b154e16c9799d3387076c421423ef52ac4df99", // WBTC - USDT 86% (2024.02.02)
		"0xe7e9694b754c4d4f7e21faf7223f6fa71abaeb10296a4c43a54a7977149687d2", // wstETH - USDT 86% (2024.02.02)
	},
}
