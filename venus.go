package ethaddr

import "github.com/0xVanfer/chainId"

const VenusProtocol string = "venus"

// Venus token XVS.
//
// map[network] = address.
var XVSList = map[string]string{
	chainId.BNBSmartChainName: "0xcf6bb5389c92bdda8a3747ddb454cb7a64626c63", // XVS
}

// Venus c tokens(v tokens).
//
// map[network][underlying] = address.
var VenusCTokenList = map[string]map[string]string{
	chainId.BNBSmartChainName: {
		SXPList[chainId.BNBSmartChainName]:     "0x2ff3d0f6990a40261c66e1ff2017acbc282eb6d0", // v
		XVSList[chainId.BNBSmartChainName]:     "0x151b1e2635a717bcdc836ecd6fbb62b674fe3e1d", // v
		USDCList[chainId.BNBSmartChainName]:    "0xeca88125a5adbe82614ffc12d0db554e2e2867c8", // v
		USDTList[chainId.BNBSmartChainName]:    "0xfd5840cd36d94d7229439859c0112a4185bc0255", // v
		BUSDList[chainId.BNBSmartChainName]:    "0x95c78222b3d6e262426483d42cfa53685a67ab9d", // v
		WBNBList[chainId.BNBSmartChainName]:    "0xa07c5b74c9b40447a954e1466938b865b6bbea36", // v
		BTCbList[chainId.BNBSmartChainName]:    "0x882c173bc7ff3b7786ca16dfed3dfffb9ee7847b", // v
		ETHList[chainId.BNBSmartChainName]:     "0xf508fcd89b8bd15579dc79a6827cb4686a3592c8", // v
		LTCList[chainId.BNBSmartChainName]:     "0x57a5297f2cb2c0aac9d554660acd6d385ab50c6b", // v
		XRPList[chainId.BNBSmartChainName]:     "0xb248a295732e0225acd3337607cc01068e3b9c10", // v
		BCHList[chainId.BNBSmartChainName]:     "0x5f0388ebc2b94fa8e123f404b79ccf5f40b29176", // v
		DOTList[chainId.BNBSmartChainName]:     "0x1610bc33319e9398de5f57b33a5b184c806ad217", // v
		LINKList[chainId.BNBSmartChainName]:    "0x650b940a1033b8a1b1873f78730fcfc73ec11f1f", // v
		DAIList[chainId.BNBSmartChainName]:     "0x334b3ecb4dca3593bccc3c7ebd1a1c1d1780fbf1", // v
		FILList[chainId.BNBSmartChainName]:     "0xf91d58b5ae142dacc749f58a49fcbac340cb0343", // v
		BETHList[chainId.BNBSmartChainName]:    "0x972207a639cc1b374b893cc33fa251b55ceb7c07", // v
		ADAList[chainId.BNBSmartChainName]:     "0x9a0af7fdb2065ce470d72664de73cae409da28ec", // v
		DOGEList[chainId.BNBSmartChainName]:    "0xec3422ef92b2fb59e84c8b02ba73f1fe84ed8d71", // v
		MATICList[chainId.BNBSmartChainName]:   "0x5c9476fcd6a4f9a3654139721c949c2233bbbbc8", // v
		CakeList[chainId.BNBSmartChainName]:    "0x86ac3974e2bd0d60825230fa6f355ff11409df5c", // v
		AAVEList[chainId.BNBSmartChainName]:    "0x26da28954763b92139ed49283625cecaf52c6f94", // v
		TUSDOldList[chainId.BNBSmartChainName]: "0x08ceb3f4a7ed3500ca0982bcd0fc7816688084c3", // v
		TRXOldList[chainId.BNBSmartChainName]:  "0x61edcfe8dd6ba3c891cb9bec2dc7657b3b422e93", // v
		USTList[chainId.BNBSmartChainName]:     "0x78366446547d062f45b4c0f320cdaa6d710d87bb", // v
		LUNAList[chainId.BNBSmartChainName]:    "0xb91a659e88b51474767cd97ef3196a3e7cedd2c8", // v
		TRXList[chainId.BNBSmartChainName]:     "0xc5d3466aa484b040ee977073fcf337f2c00071c1", // v
		WBETHList[chainId.BNBSmartChainName]:   "0x6cfdec747f37daf3b87a35a1d9c8ad3063a1a8a0", // v
		TUSDList[chainId.BNBSmartChainName]:    "0xbf762cd5991ca1dcddac9ae5c638f5b5dc3bee6e", // v
	},
}
