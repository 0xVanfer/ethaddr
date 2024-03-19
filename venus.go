package ethaddr

import "github.com/0xVanfer/chainId"

const VenusProtocol string = "venus"

// Venus token XVS.
//
// map[network] = address.
var XVSList = map[string]string{
	chainId.BNBSmartChainName: "0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63", // XVS, 0xcf6bb5389c92bdda8a3747ddb454cb7a64626c63
}

// Venus c tokens(v tokens).
//
// map[network][underlying] = address.
var VenusCTokenList = map[string]map[string]string{
	chainId.BNBSmartChainName: {
		SXPList[chainId.BNBSmartChainName]:     "0x2fF3d0F6990a40261c66E1ff2017aCBc282EB6d0", // vSXP, 0x2ff3d0f6990a40261c66e1ff2017acbc282eb6d0
		XVSList[chainId.BNBSmartChainName]:     "0x151B1e2635A717bcDc836ECd6FbB62B674FE3E1D", // v, 0x151b1e2635a717bcdc836ecd6fbb62b674fe3e1d
		USDCList[chainId.BNBSmartChainName]:    "0xecA88125a5ADbe82614ffC12D0DB554E2e2867C8", // v, 0xeca88125a5adbe82614ffc12d0db554e2e2867c8
		USDTList[chainId.BNBSmartChainName]:    "0xfD5840Cd36d94D7229439859C0112a4185BC0255", // v, 0xfd5840cd36d94d7229439859c0112a4185bc0255
		BUSDList[chainId.BNBSmartChainName]:    "0x95c78222B3D6e262426483D42CfA53685A67Ab9D", // v, 0x95c78222b3d6e262426483d42cfa53685a67ab9d
		WBNBList[chainId.BNBSmartChainName]:    "0xA07c5b74C9B40447a954e1466938b865b6BBea36", // v, 0xa07c5b74c9b40447a954e1466938b865b6bbea36
		BTCbList[chainId.BNBSmartChainName]:    "0x882C173bC7Ff3b7786CA16dfeD3DFFfb9Ee7847B", // v, 0x882c173bc7ff3b7786ca16dfed3dfffb9ee7847b
		ETHList[chainId.BNBSmartChainName]:     "0xf508fCD89b8bd15579dc79A6827cB4686A3592c8", // v, 0xf508fcd89b8bd15579dc79a6827cb4686a3592c8
		LTCList[chainId.BNBSmartChainName]:     "0x57A5297F2cB2c0AaC9D554660acd6D385Ab50c6B", // v, 0x57a5297f2cb2c0aac9d554660acd6d385ab50c6b
		XRPList[chainId.BNBSmartChainName]:     "0xB248a295732e0225acd3337607cc01068e3b9c10", // v, 0xb248a295732e0225acd3337607cc01068e3b9c10
		BCHList[chainId.BNBSmartChainName]:     "0x5F0388EBc2B94FA8E123F404b79cCF5f40b29176", // v, 0x5f0388ebc2b94fa8e123f404b79ccf5f40b29176
		DOTList[chainId.BNBSmartChainName]:     "0x1610bc33319e9398de5f57B33a5b184c806aD217", // v, 0x1610bc33319e9398de5f57b33a5b184c806ad217
		LINKList[chainId.BNBSmartChainName]:    "0x650b940a1033B8A1b1873f78730FcFC73ec11f1f", // v, 0x650b940a1033b8a1b1873f78730fcfc73ec11f1f
		DAIList[chainId.BNBSmartChainName]:     "0x334b3eCB4DCa3593BCCC3c7EBD1A1C1d1780FBF1", // v, 0x334b3ecb4dca3593bccc3c7ebd1a1c1d1780fbf1
		FILList[chainId.BNBSmartChainName]:     "0xf91d58b5aE142DAcC749f58A49FCBac340Cb0343", // v, 0xf91d58b5ae142dacc749f58a49fcbac340cb0343
		BETHList[chainId.BNBSmartChainName]:    "0x972207A639CC1B374B893cc33Fa251b55CEB7c07", // v, 0x972207a639cc1b374b893cc33fa251b55ceb7c07
		ADAList[chainId.BNBSmartChainName]:     "0x9A0AF7FDb2065Ce470D72664DE73cAE409dA28Ec", // v, 0x9a0af7fdb2065ce470d72664de73cae409da28ec
		DOGEList[chainId.BNBSmartChainName]:    "0xec3422Ef92B2fb59e84c8B02Ba73F1fE84Ed8D71", // v, 0xec3422ef92b2fb59e84c8b02ba73f1fe84ed8d71
		MATICList[chainId.BNBSmartChainName]:   "0x5c9476FcD6a4F9a3654139721c949c2233bBbBc8", // v, 0x5c9476fcd6a4f9a3654139721c949c2233bbbbc8
		CakeList[chainId.BNBSmartChainName]:    "0x86aC3974e2BD0d60825230fa6F355fF11409df5c", // v, 0x86ac3974e2bd0d60825230fa6f355ff11409df5c
		AAVEList[chainId.BNBSmartChainName]:    "0x26DA28954763B92139ED49283625ceCAf52C6f94", // v, 0x26da28954763b92139ed49283625cecaf52c6f94
		TUSDOldList[chainId.BNBSmartChainName]: "0x08CEB3F4a7ed3500cA0982bcd0FC7816688084c3", // v, 0x08ceb3f4a7ed3500ca0982bcd0fc7816688084c3
		TRXOldList[chainId.BNBSmartChainName]:  "0x61eDcFe8Dd6bA3c891CB9bEc2dc7657B3B422E93", // v, 0x61edcfe8dd6ba3c891cb9bec2dc7657b3b422e93
		USTList[chainId.BNBSmartChainName]:     "0x78366446547D062f45b4C0f320cDaa6d710D87bb", // v, 0x78366446547d062f45b4c0f320cdaa6d710d87bb
		LUNAList[chainId.BNBSmartChainName]:    "0xb91A659E88B51474767CD97EF3196A3e7cEDD2c8", // v, 0xb91a659e88b51474767cd97ef3196a3e7cedd2c8
		TRXList[chainId.BNBSmartChainName]:     "0xC5D3466aA484B040eE977073fcF337f2c00071c1", // v, 0xc5d3466aa484b040ee977073fcf337f2c00071c1
		WBETHList[chainId.BNBSmartChainName]:   "0x6CFdEc747f37DAf3b87a35a1D9c8AD3063A1A8A0", // v, 0x6cfdec747f37daf3b87a35a1d9c8ad3063a1a8a0
		TUSDList[chainId.BNBSmartChainName]:    "0xBf762cd5991cA1DCdDaC9ae5C638F5B5Dc3Bee6E", // v, 0xbf762cd5991ca1dcddac9ae5c638f5b5dc3bee6e
	},
}
