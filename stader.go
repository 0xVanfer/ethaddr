package ethaddr

// Website: https://www.staderlabs.com/
//
// Docs: https://www.staderlabs.com/docs/
//
// Deployed contracts: https://www.staderlabs.com/docs-v1/Ethereum/smart-contracts
const StaderProtocol string = "stader"

// Stader token SD.
//
// map[chainID] = address.
var StaderTokenList = map[int64]string{
	ChainEthereum: "0x30D20208d987713f46DFD34EF128Bb16C404D10f", // SD, 0x30d20208d987713f46dfd34ef128bb16c404d10f
}

// Same as StaderTokenList.
var SDList = StaderTokenList

// MaticX by Stader.
//
// map[chainID] = address.
var MaticXList = map[int64]string{
	ChainEthereum: "0xf03A7Eb46d01d9EcAA104558C732Cf82f6B6B645", // MaticX, 0xf03a7eb46d01d9ecaa104558c732cf82f6b6b645
	ChainPolygon:  "0xfa68FB4628DFF1028CFEc22b4162FCcd0d45efb6", // MaticX, 0xfa68fb4628dff1028cfec22b4162fccd0d45efb6
}

// ETHx by stader.
//
// map[chainID] = address.
var ETHxList = map[int64]string{
	ChainArbitrum: "0xED65C5085a18Fa160Af0313E60dcc7905E944Dc7", // ETHx, 0xed65c5085a18fa160af0313e60dcc7905e944dc7
	ChainEthereum: "0xA35b1B31Ce002FBF2058D22F30f95D405200A15b", // ETHx, 0xa35b1b31ce002fbf2058d22f30f95d405200a15b
}

// Stader oracle for ETHx and SD.
//
// map[chainID] = address.
var StaderETHxOracle = map[int64]string{
	ChainEthereum: "0xF64bAe65f6f2a5277571143A24FaaFDFC0C2a737", // 0xf64bae65f6f2a5277571143a24faafdfc0c2a737
}

// Stader PoS staking contract for Matic.
//
// map[chainID] = address.
var StaderStakeManagerList = map[int64]string{
	ChainEthereum: "0x5e3Ef299fDDf15eAa0432E6e66473ace8c13D908", // 0x5e3ef299fddf15eaa0432e6e66473ace8c13d908
}

// Polygon pool for MaticX.
//
// map[chainID] = address.
var StaderChildPoolList = map[int64]string{
	ChainPolygon: "0xfd225C9e6601C9d38d8F98d8731BF59eFcF8C0E3", // 0xfd225c9e6601c9d38d8f98d8731bf59efcf8c0e3
}

// Stader staked BNB, BNBx.
//
// map[chainID] = address.
var BNBxList = map[int64]string{
	ChainBSC: "0x1bdd3Cf7F79cfB8EdbB955f20ad99211551BA275", // 0x1bdd3cf7f79cfb8edbb955f20ad99211551ba275
}
