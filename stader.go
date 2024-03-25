package ethaddr

// Website: https://www.staderlabs.com/
//
// Docs: https://www.staderlabs.com/docs/
//
// Deployed contracts: https://www.staderlabs.com/docs-v1/Ethereum/smart-contracts
const StaderProtocol string = "stader"

// MaticX by Stader.
//
// map[network] = address.
var MaticXList = map[string]string{
	ChainEthereum: "0xf03A7Eb46d01d9EcAA104558C732Cf82f6B6B645", // MaticX, 0xf03a7eb46d01d9ecaa104558c732cf82f6b6b645
	ChainPolygon:  "0xfa68FB4628DFF1028CFEc22b4162FCcd0d45efb6", // MaticX, 0xfa68fb4628dff1028cfec22b4162fccd0d45efb6
}

// ETHx by stader.
//
// map[network] = address.
var ETHxList = map[string]string{
	ChainEthereum: "0xA35b1B31Ce002FBF2058D22F30f95D405200A15b", // ETHx, 0xa35b1b31ce002fbf2058d22f30f95d405200a15b
}

// Stader PoS staking contract.
//
// map[network] = address.
var StaderStakeManagerList = map[string]string{
	ChainEthereum: "0x5e3Ef299fDDf15eAa0432E6e66473ace8c13D908", // 0x5e3ef299fddf15eaa0432e6e66473ace8c13d908
}

var StaderChildPoolList = map[string]string{
	ChainPolygon: "0xfd225C9e6601C9d38d8F98d8731BF59eFcF8C0E3", // 0xfd225c9e6601c9d38d8f98d8731bf59efcf8c0e3
}
