package ethaddr

import "github.com/0xVanfer/chainId"

const BenqiProtocol string = "benqi"

// Benqi token: QI.
//
// map[network] = address.
var BenqiTokenList = map[string]string{
	chainId.AvalancheChainName: "0x8729438eb15e2c8b576fcc6aecda6a148776c0f5", // QI
}

// Same as BenqiTokenList.
var QIList = BenqiTokenList

// Benqi savax token: SAVAX.
//
// map[network] = address.
var SAVAXList = map[string]string{
	chainId.AvalancheChainName: "0x2b2c81e08f1af8835a78bb2a90ae924ace0ea4be", // sAVAX
}

// Benqi comptroller, similar to compound.
//
// map[network] = address.
var BenqiComptrollerList = map[string]string{
	chainId.AvalancheChainName: "0x486af39519b4dc9a7fccd318217352830e8ad9b4",
}

// Benqi c tokens(qi tokens).
//
// map[network][underlying] = address.
var BenqiCTokenList = map[string]map[string]string{
	chainId.AvalancheChainName: {
		DAIeList[chainId.AvalancheChainName]:  "0x835866d37afb8cb8f8334dccdaf66cf01832ff5d", // qiDAI
		USDCeList[chainId.AvalancheChainName]: "0xbeb5d47a3f720ec0a390d04b4d41ed7d9688bc7f", // qiUSDC
		USDTeList[chainId.AvalancheChainName]: "0xc9e5999b8e75c3feb117f6f73e664b9f3c8ca65c", // qiUSDT
		WETHList[chainId.AvalancheChainName]:  "0x334ad834cd4481bb02d09615e7c11a00579a7909", // qiETH
		WBTCList[chainId.AvalancheChainName]:  "0xe194c4c5ac32a3c9ffdb358d9bfd523a0b6d1568", // qiBTC
		WAVAXList[chainId.AvalancheChainName]: "0x5c0401e81bc07ca70fad469b451682c0d747ef1c", // qiAVAX
		LINKList[chainId.AvalancheChainName]:  "0x4e9f683a27a6bdad3fc2764003759277e93696e6", // qiLINK
		QIList[chainId.AvalancheChainName]:    "0x35bd6aeda81a7e5fc7a7832490e71f757b0cd9ce", // qiQI
		USDCList[chainId.AvalancheChainName]:  "0xb715808a78f6041e46d61cb123c9b4a27056ae9c", // qiUSDCn
		USDTList[chainId.AvalancheChainName]:  "0xd8fcda6ec4bdc547c0827b8804e89acd817d56ef", // qiUSDTn
		SAVAXList[chainId.AvalancheChainName]: "0xf362fea9659cf036792c9cb02f8ff8198e21b4cb", // qisAVAX
		BTCbList[chainId.AvalancheChainName]:  "0x89a415b3d20098e6a6c8f7a59001c67bd3129821", // qiBTC.b
		BUSDList[chainId.AvalancheChainName]:  "0x872670ccae8c19557cc9443eff587d7086b8043a", // qiBUSD
	},
}
