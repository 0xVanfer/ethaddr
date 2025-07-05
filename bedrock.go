package ethaddr

// Website: https://www.bedrock.technology/
//
// Docs: https://docs.bedrock.technology/
//
// Deployed contracts: https://github.com/Bedrock-Technology/uniBTC/tree/main/deployments
const BedrockProtocol string = "bedrock"

// Bedrock uniBTC.
//
// map[chainID] = address
var UniBTCList = map[int64]string{
	ChainArbitrum: "0x6B2a01A5f79dEb4c2f3c0eDa7b01DF456FbD726a", // uniBTC, 0x6b2a01a5f79deb4c2f3c0eda7b01df456fbd726a
	ChainBSC:      "0x6B2a01A5f79dEb4c2f3c0eDa7b01DF456FbD726a", // uniBTC, 0x6b2a01a5f79deb4c2f3c0eda7b01df456fbd726a
	ChainEthereum: "0x004E9C3EF86bc1ca1f0bB5C7662861Ee93350568", // uniBTC, 0x004e9c3ef86bc1ca1f0bb5c7662861ee93350568
	ChainOptimism: "0x93919784C523f39CACaa98Ee0a9d96c3F32b593e", // uniBTC, 0x93919784c523f39cacaa98ee0a9d96c3f32b593e
	ChainMantle:   "0x93919784C523f39CACaa98Ee0a9d96c3F32b593e", // uniBTC, 0x93919784c523f39cacaa98ee0a9d96c3f32b593e
	ChainMode:     "0x6B2a01A5f79dEb4c2f3c0eDa7b01DF456FbD726a", // uniBTC, 0x6b2a01a5f79deb4c2f3c0eda7b01df456fbd726a
	ChainTaiko:    "0x93919784C523f39CACaa98Ee0a9d96c3F32b593e", // uniBTC, 0x93919784c523f39cacaa98ee0a9d96c3f32b593e
}

// Bedrock uniBTC Router.
//
// map[chainID] = address
var UniBTCRouterList = map[int64]string{
	ChainArbitrum: "0x84E5C854A7fF9F49c888d69DECa578D406C26800", // 0x84e5c854a7ff9f49c888d69deca578d406c26800
	ChainBSC:      "0x84E5C854A7fF9F49c888d69DECa578D406C26800", // 0x84e5c854a7ff9f49c888d69deca578d406c26800
	ChainEthereum: "0x047D41F2544B7F63A8e991aF2068a363d210d6Da", // 0x047d41f2544b7f63a8e991af2068a363d210d6da
	ChainOptimism: "0xF9775085d726E782E83585033B58606f7731AB18", // 0xf9775085d726e782e83585033b58606f7731ab18
	ChainMantle:   "0xF9775085d726E782E83585033B58606f7731AB18", // 0xf9775085d726e782e83585033b58606f7731ab18
	ChainMode:     "0x84E5C854A7fF9F49c888d69DECa578D406C26800", // 0x84e5c854a7ff9f49c888d69deca578d406c26800
}

// Bedrock uniETH.
//
// map[chainID] = address
var UniETHList = map[int64]string{
	ChainArbitrum: "0x3d15fD46CE9e551498328B1C83071D9509E2C3a0", // uniETH, 0x3d15fd46ce9e551498328b1c83071d9509e2c3a0
	ChainEthereum: "0xF1376bceF0f78459C0Ed0ba5ddce976F1ddF51F4", // uniETH, 0xf1376bcef0f78459c0ed0ba5ddce976f1ddf51f4
	ChainLinea:    "0x15EEfE5B297136b8712291B632404B66A8eF4D25", // uniETH, 0x15eefe5b297136b8712291b632404b66a8ef4d25
	ChainScroll:   "0x15EEfE5B297136b8712291B632404B66A8eF4D25", // uniETH, 0x15eefe5b297136b8712291b632404b66a8ef4d25
}
