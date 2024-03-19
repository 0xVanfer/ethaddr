package ethaddr

import "github.com/0xVanfer/chainId"

// Docs: https://docs.benqi.fi/
//
// Deployed contracts: https://docs.benqi.fi/resources/contracts/benqi-liquidity-market
const BenqiProtocol string = "benqi"

// Benqi token: QI.
//
// map[network] = address.
var BenqiTokenList = map[string]string{
	chainId.AvalancheChainName: "0x8729438EB15e2C8B576fCc6AeCdA6A148776C0F5", // QI, 0x8729438eb15e2c8b576fcc6aecda6a148776c0f5
}

// Same as BenqiTokenList.
var QIList = BenqiTokenList

// Benqi savax token: SAVAX.
//
// map[network] = address.
var SAVAXList = map[string]string{
	chainId.AvalancheChainName: "0x2b2C81e08f1Af8835a78Bb2A90AE924ACE0eA4bE", // sAVAX, 0x2b2c81e08f1af8835a78bb2a90ae924ace0ea4be
}

// Benqi comptroller, similar to compound.
//
// map[network] = address.
var BenqiComptrollerList = map[string]string{
	chainId.AvalancheChainName: "0x486Af39519B4Dc9a7fCcd318217352830E8AD9b4", // 0x486af39519b4dc9a7fccd318217352830e8ad9b4
}

// Benqi c tokens(qi tokens).
//
// map[network][underlying] = address.
var BenqiCTokenList = map[string]map[string]string{
	chainId.AvalancheChainName: {
		DAIeList[chainId.AvalancheChainName]:  "0x835866d37AFB8CB8F8334dCCdaf66cf01832Ff5D", // qiDAI, 0x835866d37afb8cb8f8334dccdaf66cf01832ff5d
		USDCeList[chainId.AvalancheChainName]: "0xBEb5d47A3f720Ec0a390d04b4d41ED7d9688bC7F", // qiUSDC, 0xbeb5d47a3f720ec0a390d04b4d41ed7d9688bc7f
		USDTeList[chainId.AvalancheChainName]: "0xc9e5999b8e75C3fEB117F6f73E664b9f3C8ca65C", // qiUSDT, 0xc9e5999b8e75c3feb117f6f73e664b9f3c8ca65c
		WETHList[chainId.AvalancheChainName]:  "0x334AD834Cd4481BB02d09615E7c11a00579A7909", // qiETH, 0x334ad834cd4481bb02d09615e7c11a00579a7909
		WBTCList[chainId.AvalancheChainName]:  "0xe194c4c5aC32a3C9ffDb358d9Bfd523a0B6d1568", // qiBTC, 0xe194c4c5ac32a3c9ffdb358d9bfd523a0b6d1568
		WAVAXList[chainId.AvalancheChainName]: "0x5C0401e81Bc07Ca70fAD469b451682c0d747Ef1c", // qiAVAX, 0x5c0401e81bc07ca70fad469b451682c0d747ef1c
		LINKList[chainId.AvalancheChainName]:  "0x4e9f683A27a6BdAD3FC2764003759277e93696e6", // qiLINK, 0x4e9f683a27a6bdad3fc2764003759277e93696e6
		QIList[chainId.AvalancheChainName]:    "0x35Bd6aedA81a7E5FC7A7832490e71F757b0cD9Ce", // qiQI, 0x35bd6aeda81a7e5fc7a7832490e71f757b0cd9ce
		USDCList[chainId.AvalancheChainName]:  "0xB715808a78F6041E46d61Cb123C9B4A27056AE9C", // qiUSDCn, 0xb715808a78f6041e46d61cb123c9b4a27056ae9c
		USDTList[chainId.AvalancheChainName]:  "0xd8fcDa6ec4Bdc547C0827B8804e89aCd817d56EF", // qiUSDTn, 0xd8fcda6ec4bdc547c0827b8804e89acd817d56ef
		SAVAXList[chainId.AvalancheChainName]: "0xF362feA9659cf036792c9cb02f8ff8198E21B4cB", // qisAVAX, 0xf362fea9659cf036792c9cb02f8ff8198e21b4cb
		BTCbList[chainId.AvalancheChainName]:  "0x89a415b3D20098E6A6C8f7a59001C67BD3129821", // qiBTC.b, 0x89a415b3d20098e6a6c8f7a59001c67bd3129821
		BUSDList[chainId.AvalancheChainName]:  "0x872670CcAe8C19557cC9443Eff587D7086b8043A", // qiBUSD, 0x872670ccae8c19557cc9443eff587d7086b8043a
	},
}
