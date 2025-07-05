package ethaddr

// Website: https://venus.io/
const (
	VenusProtocol string = "venus"
	// core
	VenusProtocolCore string = "venuscore"
	// isolated btc
	VenusProtocolIsolatedBTC string = "venusisolatedbtc"
	// isolated defi
	VenusProtocolIsolatedDefi string = "venusisolateddefi"
	// isolated gamefi
	VenusProtocolIsolatedGamefi string = "venusisolatedgamefi"
	// isolated liquid staking bnb
	VenusProtocolIsolatedLiquidStakedBNB string = "venusisolatedliquidstakedbnb"
	// isolated liquid staking eth
	VenusProtocolIsolatedLiquidStakedETH string = "venusisolatedliquidstakedeth"
	// isolated meme
	VenusProtocolIsolatedMeme string = "venusisolatedmeme"
	// isolated stablecoins
	VenusProtocolIsolatedStablecoins string = "venusisolatedstablecoins"
	// isolated tron
	VenusProtocolIsolatedTron string = "venusisolatedtron"
	// isolated curve
	VenusProtocolIsolatedCurve string = "venusisolatedcurve"
	// isolated ethena
	VenusProtocolIsolatedEthena string = "venusisolatedethena"
)

// Venus token XVS.
//
// map[chainID] = address.
var XVSList = map[int64]string{
	ChainBSC: "0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63", // XVS, 0xcf6bb5389c92bdda8a3747ddb454cb7a64626c63
}

// Venus comptrollers by venus protocol names.
// Including core and isolated comptrollers.
//
// map[chainID][protocol] = comptroller.
var VenusComptrollersList = map[int64]map[string]string{
	ChainArbitrum: {
		VenusProtocolCore:                    "0x317c1a5739f39046e20b08ac9beea3f10fd43326", // 0x317c1a5739f39046e20b08ac9beea3f10fd43326
		VenusProtocolIsolatedLiquidStakedETH: "0x52bAB1aF7Ff770551BD05b9FC2329a0Bf5E23F16", // 0x52bab1af7ff770551bd05b9fc2329a0bf5e23f16
	},
	ChainBase: {
		VenusProtocolCore: "0x0C7973F9598AA62f9e03B94E92C967fD5437426C", // 0x0c7973f9598aa62f9e03b94e92c967fd5437426c
	},
	ChainBSC: {
		VenusProtocolCore:                    "0xfD36E2c2a6789Db23113685031d7F16329158384", // 0xfd36e2c2a6789db23113685031d7f16329158384
		VenusProtocolIsolatedBTC:             "0x9DF11376Cf28867E2B0741348044780FbB7cb1d6", // 0x9df11376cf28867e2b0741348044780fbb7cb1d6
		VenusProtocolIsolatedDefi:            "0x3344417c9360b963ca93A4e8305361AEde340Ab9", // 0x3344417c9360b963ca93a4e8305361aede340ab9
		VenusProtocolIsolatedGamefi:          "0x1b43ea8622e76627B81665B1eCeBB4867566B963", // 0x1b43ea8622e76627b81665b1ecebb4867566b963
		VenusProtocolIsolatedLiquidStakedBNB: "0xd933909A4a2b7A4638903028f44D1d38ce27c352", // 0xd933909a4a2b7a4638903028f44d1d38ce27c352
		VenusProtocolIsolatedLiquidStakedETH: "0xBE609449Eb4D76AD8545f957bBE04b596E8fC529", // 0xbe609449eb4d76ad8545f957bbe04b596e8fc529
		VenusProtocolIsolatedMeme:            "0x33B6fa34cd23e5aeeD1B112d5988B026b8A5567d", // 0x33b6fa34cd23e5aeed1b112d5988b026b8a5567d
		VenusProtocolIsolatedStablecoins:     "0x94c1495cD4c557f1560Cbd68EAB0d197e6291571", // 0x94c1495cd4c557f1560cbd68eab0d197e6291571
		VenusProtocolIsolatedTron:            "0x23b4404E4E5eC5FF5a6FFb70B7d14E3FabF237B0", // 0x23b4404e4e5ec5ff5a6ffb70b7d14e3fabf237b0
		VenusProtocolIsolatedCurve:           "0x67aA3eCc5831a65A5Ba7be76BED3B5dc7DB60796", // 0x67aa3ecc5831a65a5ba7be76bed3b5dc7db60796
		VenusProtocolIsolatedEthena:          "0x562d2b6FF1dbf5f63E233662416782318cC081E4", // 0x562d2b6ff1dbf5f63e233662416782318cc081e4
	},
	ChainEthereum: {
		VenusProtocolCore:                    "0x687a01ecF6d3907658f7A7c714749fAC32336D1B", // 0x687a01ecf6d3907658f7a7c714749fac32336d1b
		VenusProtocolIsolatedLiquidStakedETH: "0xF522cd0360EF8c2FF48B648d53EA1717Ec0F3Ac3", // 0xf522cd0360ef8c2ff48b648d53ea1717ec0f3ac3
	},
}

// map[chainID][comptroller][underlying] = address.
var VenusCTokensListByComptroller = map[int64]map[string]map[string]string{
	ChainArbitrum: {
		// Last updated: Mar.5.2025
		VenusComptrollersList[ChainArbitrum][VenusProtocolCore]: {
			WBTCList[ChainArbitrum]: "0xaDa57840B372D4c28623E87FC175dE8490792811", // 0xada57840b372d4c28623e87fc175de8490792811
			WETHList[ChainArbitrum]: "0x68a34332983f4Bf866768DD6D6E638b02eF5e1f0", // 0x68a34332983f4bf866768dd6d6e638b02ef5e1f0
			USDCList[ChainArbitrum]: "0x7D8609f8da70fF9027E9bc5229Af4F6727662707", // 0x7d8609f8da70ff9027e9bc5229af4f6727662707
			USDTList[ChainArbitrum]: "0xB9F9117d4200dC296F9AcD1e8bE1937df834a2fD", // 0xb9f9117d4200dc296f9acd1e8be1937df834a2fd
			ARBList[ChainArbitrum]:  "0xAeB0FEd69354f34831fe1D16475D9A83ddaCaDA6", // 0xaeb0fed69354f34831fe1d16475d9a83ddacada6
			// GM
			"0x70d95587d40A2caf56bd97485aB3Eec10Bee6336": "0x9bb8cEc9C0d46F53b4f2173BB2A0221F66c353cC", // 0x9bb8cec9c0d46f53b4f2173bb2a0221f66c353cc
			// GM
			"0x47c031236e19d024b42f8AE6780E44A573170703": "0x4f3a73f318C5EA67A86eaaCE24309F29f89900dF", // 0x4f3a73f318c5ea67a86eaace24309f29f89900df
		},
	},
	ChainBSC: {
		VenusComptrollersList[ChainBSC][VenusProtocolCore]: {
			SXPList[ChainBSC]:      "0x2fF3d0F6990a40261c66E1ff2017aCBc282EB6d0", // vSXP, 0x2ff3d0f6990a40261c66e1ff2017acbc282eb6d0
			XVSList[ChainBSC]:      "0x151B1e2635A717bcDc836ECd6FbB62B674FE3E1D", // v, 0x151b1e2635a717bcdc836ecd6fbb62b674fe3e1d
			USDCList[ChainBSC]:     "0xecA88125a5ADbe82614ffC12D0DB554E2e2867C8", // v, 0xeca88125a5adbe82614ffc12d0db554e2e2867c8
			USDTList[ChainBSC]:     "0xfD5840Cd36d94D7229439859C0112a4185BC0255", // v, 0xfd5840cd36d94d7229439859c0112a4185bc0255
			BUSDList[ChainBSC]:     "0x95c78222B3D6e262426483D42CfA53685A67Ab9D", // v, 0x95c78222b3d6e262426483d42cfa53685a67ab9d
			WBNBList[ChainBSC]:     "0xA07c5b74C9B40447a954e1466938b865b6BBea36", // v, 0xa07c5b74c9b40447a954e1466938b865b6bbea36
			BTCBList[ChainBSC]:     "0x882C173bC7Ff3b7786CA16dfeD3DFFfb9Ee7847B", // v, 0x882c173bc7ff3b7786ca16dfed3dfffb9ee7847b
			ETHList[ChainBSC]:      "0xf508fCD89b8bd15579dc79A6827cB4686A3592c8", // v, 0xf508fcd89b8bd15579dc79a6827cb4686a3592c8
			LTCList[ChainBSC]:      "0x57A5297F2cB2c0AaC9D554660acd6D385Ab50c6B", // v, 0x57a5297f2cb2c0aac9d554660acd6d385ab50c6b
			XRPList[ChainBSC]:      "0xB248a295732e0225acd3337607cc01068e3b9c10", // v, 0xb248a295732e0225acd3337607cc01068e3b9c10
			BCHList[ChainBSC]:      "0x5F0388EBc2B94FA8E123F404b79cCF5f40b29176", // v, 0x5f0388ebc2b94fa8e123f404b79ccf5f40b29176
			DOTList[ChainBSC]:      "0x1610bc33319e9398de5f57B33a5b184c806aD217", // v, 0x1610bc33319e9398de5f57b33a5b184c806ad217
			LINKList[ChainBSC]:     "0x650b940a1033B8A1b1873f78730FcFC73ec11f1f", // v, 0x650b940a1033b8a1b1873f78730fcfc73ec11f1f
			DAIList[ChainBSC]:      "0x334b3eCB4DCa3593BCCC3c7EBD1A1C1d1780FBF1", // v, 0x334b3ecb4dca3593bccc3c7ebd1a1c1d1780fbf1
			FILList[ChainBSC]:      "0xf91d58b5aE142DAcC749f58A49FCBac340Cb0343", // v, 0xf91d58b5ae142dacc749f58a49fcbac340cb0343
			BETHList[ChainBSC]:     "0x972207A639CC1B374B893cc33Fa251b55CEB7c07", // v, 0x972207a639cc1b374b893cc33fa251b55ceb7c07
			ADAList[ChainBSC]:      "0x9A0AF7FDb2065Ce470D72664DE73cAE409dA28Ec", // v, 0x9a0af7fdb2065ce470d72664de73cae409da28ec
			DOGEList[ChainBSC]:     "0xec3422Ef92B2fb59e84c8B02Ba73F1fE84Ed8D71", // v, 0xec3422ef92b2fb59e84c8b02ba73f1fe84ed8d71
			MATICList[ChainBSC]:    "0x5c9476FcD6a4F9a3654139721c949c2233bBbBc8", // v, 0x5c9476fcd6a4f9a3654139721c949c2233bbbbc8
			CakeList[ChainBSC]:     "0x86aC3974e2BD0d60825230fa6F355fF11409df5c", // v, 0x86ac3974e2bd0d60825230fa6f355ff11409df5c
			AAVEList[ChainBSC]:     "0x26DA28954763B92139ED49283625ceCAf52C6f94", // v, 0x26da28954763b92139ed49283625cecaf52c6f94
			TUSDOldList[ChainBSC]:  "0x08CEB3F4a7ed3500cA0982bcd0FC7816688084c3", // v, 0x08ceb3f4a7ed3500ca0982bcd0fc7816688084c3
			TRXOldList[ChainBSC]:   "0x61eDcFe8Dd6bA3c891CB9bEc2dc7657B3B422E93", // v, 0x61edcfe8dd6ba3c891cb9bec2dc7657b3b422e93
			USTList[ChainBSC]:      "0x78366446547D062f45b4C0f320cDaa6d710D87bb", // v, 0x78366446547d062f45b4c0f320cdaa6d710d87bb
			LUNAList[ChainBSC]:     "0xb91A659E88B51474767CD97EF3196A3e7cEDD2c8", // v, 0xb91a659e88b51474767cd97ef3196a3e7cedd2c8
			TRXList[ChainBSC]:      "0xC5D3466aA484B040eE977073fcF337f2c00071c1", // v, 0xc5d3466aa484b040ee977073fcf337f2c00071c1
			WBETHList[ChainBSC]:    "0x6CFdEc747f37DAf3b87a35a1D9c8AD3063A1A8A0", // v, 0x6cfdec747f37daf3b87a35a1d9c8ad3063a1a8a0
			TUSDList[ChainBSC]:     "0xBf762cd5991cA1DCdDaC9ae5C638F5B5Dc3Bee6E", // v, 0xbf762cd5991ca1dcddac9ae5c638f5b5dc3bee6e
			UNIList[ChainBSC]:      "0x27FF564707786720C71A2e5c1490A63266683612", // 0x27ff564707786720c71a2e5c1490a63266683612
			FDUSDList[ChainBSC]:    "0xC4eF4229FEc74Ccfe17B2bdeF7715fAC740BA0ba", // 0xc4ef4229fec74ccfe17b2bdef7715fac740ba0ba
			TWTList[ChainBSC]:      "0x4d41a36D04D97785bcEA57b057C412b278e6Edcc", // 0x4d41a36d04d97785bcea57b057c412b278e6edcc
			SolvBTCList[ChainBSC]:  "0xf841cb62c19fCd4fF5CD0AaB5939f3140BaaC3Ea", // 0xf841cb62c19fcd4ff5cd0aab5939f3140baac3ea
			THEList[ChainBSC]:      "0x86e06EAfa6A1eA631Eab51DE500E3D474933739f", // 0x86e06eafa6a1ea631eab51de500e3d474933739f
			SOLList[ChainBSC]:      "0xBf515bA4D1b52FFdCeaBF20d31D705Ce789F2cEC", // 0xbf515ba4d1b52ffdceabf20d31d705ce789f2cec
			LisUSDList[ChainBSC]:   "0x689E0daB47Ab16bcae87Ec18491692BF621Dc6Ab", // 0x689e0dab47ab16bcae87ec18491692bf621dc6ab
			SUSDeList[ChainBSC]:    "0x699658323d58eE25c69F1a29d476946ab011bD18", // 0x699658323d58ee25c69f1a29d476946ab011bd18
			USDeList[ChainBSC]:     "0x74ca6930108F775CC667894EEa33843e691680d7", // 0x74ca6930108f775cc667894eea33843e691680d7
			USD1List[ChainBSC]:     "0x0C1DA220D301155b87318B90692Da8dc43B67340", // 0x0c1da220d301155b87318b90692da8dc43b67340
			XSolvBTCList[ChainBSC]: "0xd804dE60aFD05EE6B89aab5D152258fD461B07D5", // 0xd804de60afd05ee6b89aab5d152258fd461b07d5
			AsBNBList[ChainBSC]:    "0xCC1dB43a06d97f736C7B045AedD03C6707c09BDF", // 0xcc1db43a06d97f736c7b045aed03c6707c09bdf

			// CAN, token of cannon, cannon.finance, already dead.
			"0x20bff4bbEDa07536FF00e073bd8359E5D80D733d": "0xeBD0070237a0713E8D94fEf1B728d3d993d290ef", // 0xebd0070237a0713e8d94fef1b728d3d993d290ef
			// PT-sUSDE-26JUN2025
			"0xDD809435ba6c9d6903730f923038801781cA66ce": "0x9e4E5fed5Ac5B9F732d0D850A615206330Bf1866", // 0x9e4e5fed5ac5b9f732d0d850a615206330bf1866
		},
		VenusComptrollersList[ChainBSC][VenusProtocolIsolatedDefi]: {
			AnkrBNBList[ChainBSC]: "0x53728FD51060a85ac41974C6C3Eb1DaE42776723", // 0x53728fd51060a85ac41974c6c3eb1dae42776723
			BSWList[ChainBSC]:     "0x8f657dFD3a1354DEB4545765fE6840cc54AFd379", // 0x8f657dfd3a1354deb4545765fe6840cc54afd379
			ALPACAList[ChainBSC]:  "0x02c5Fb0F26761093D297165e902e96D08576D344", // 0x02c5fb0f26761093d297165e902e96d08576d344
			USDTList[ChainBSC]:    "0x1D8bBDE12B6b34140604E18e9f9c6e14deC16854", // 0x1d8bbde12b6b34140604e18e9f9c6e14dec16854
			USDDList[ChainBSC]:    "0xA615467caE6B9E0bb98BC04B4411d9296fd1dFa0", // 0xa615467cae6b9e0bb98bc04b4411d9296fd1dfa0
			ANKRList[ChainBSC]:    "0x19CE11C8817a1828D1d357DFBF62dCf5b0B2A362", // 0x19ce11c8817a1828d1d357dfbf62dcf5b0b2a362
			TWTList[ChainBSC]:     "0x736bf1D21A28b5DC19A1aC8cA71Fc2856C23c03F", // 0x736bf1d21a28b5dc19a1ac8ca71fc2856c23c03f
			PLANETList[ChainBSC]:  "0xFf1112ba7f88a53D4D23ED4e14A117A2aE17C6be", // 0xff1112ba7f88a53d4d23ed4e14a117a2ae17c6be
		},
		VenusComptrollersList[ChainBSC][VenusProtocolIsolatedLiquidStakedBNB]: {
			AnkrBNBList[ChainBSC]: "0xBfe25459BA784e70E2D7a718Be99a1f3521cA17f", // 0xbfe25459ba784e70e2d7a718be99a1f3521ca17f
			BNBxList[ChainBSC]:    "0x5E21bF67a6af41c74C1773E4b473ca5ce8fd3791", // 0x5e21bf67a6af41c74c1773e4b473ca5ce8fd3791
			StkBNBList[ChainBSC]:  "0xcc5D9e502574cda17215E70bC0B4546663785227", // 0xcc5d9e502574cda17215e70bc0b4546663785227
			WBNBList[ChainBSC]:    "0xe10E80B7FD3a29fE46E16C30CC8F4dd938B742e2", // 0xe10e80b7fd3a29fe46e16c30cc8f4dd938b742e2
			SlisBNBList[ChainBSC]: "0xd3CC9d8f3689B83c91b7B59cAB4946B063EB894A", // 0xd3cc9d8f3689b83c91b7b59cab4946b063eb894a
			AsBNBList[ChainBSC]:   "0x4A50a0a1c832190362e1491D5bB464b1bc2Bd288", // 0x4a50a0a1c832190362e1491d5bb464b1bc2bd288
			// PT-clisBNB-24APR2025
			"0xE8F1C9804770e11Ab73395bE54686Ad656601E9e": "0xA537ACf381b12Bbb91C58398b66D1D220f1C77c8", // 0xa537acf381b12bbb91c58398b66d1d220f1c77c8
		},
		VenusComptrollersList[ChainBSC][VenusProtocolIsolatedLiquidStakedETH]: {
			WSTETHList[ChainBSC]: "0x94180a3948296530024Ef7d60f60B85cfe0422c8", // 0x94180a3948296530024ef7d60f60b85cfe0422c8
			WEETHList[ChainBSC]:  "0xc5b24f347254bD8cF8988913d1fd0F795274900F", // 0xc5b24f347254bd8cf8988913d1fd0f795274900f
			ETHList[ChainBSC]:    "0xeCCACF760FEA7943C5b0285BD09F601505A29c05", // 0xeccacf760fea7943c5b0285bd09f601505a29c05
		},
		VenusComptrollersList[ChainBSC][VenusProtocolIsolatedStablecoins]: {
			LisUSDList[ChainBSC]: "0xCa2D81AA7C09A1a025De797600A7081146dceEd9", // 0xca2d81aa7c09a1a025de797600a7081146dceed9
			USDTList[ChainBSC]:   "0x5e3072305F9caE1c7A82F6Fe9E38811c74922c3B", // 0x5e3072305f9cae1c7a82f6fe9e38811c74922c3b
			USDDList[ChainBSC]:   "0xc3a45ad8812189cAb659aD99E64B1376f6aCD035", // 0xc3a45ad8812189cab659ad99e64b1376f6acd035
			EURAList[ChainBSC]:   "0x795DE779Be00Ea46eA97a28BDD38d9ED570BCF0F", // 0x795de779be00ea46ea97a28bdd38d9ed570bcf0f
		},
		VenusComptrollersList[ChainBSC][VenusProtocolIsolatedBTC]: {
			BTCBList[ChainBSC]: "0x8F2AE20b25c327714248C95dFD3b02815cC82302", // 0x8f2ae20b25c327714248c95dfd3b02815cc82302
			// PT-SolvBTC.BBN-27MAR2025
			"0x541B5eEAC7D4434C8f87e2d32019d67611179606": "0x02243F036897E3bE1cce1E540FA362fd58749149", // 0x02243f036897e3be1cce1e540fa362fd58749149
		},
		VenusComptrollersList[ChainBSC][VenusProtocolIsolatedGamefi]: {
			USDTList[ChainBSC]: "0x4978591f17670A846137d9d613e333C38dc68A37", // 0x4978591f17670a846137d9d613e333c38dc68a37
			USDDList[ChainBSC]: "0x9f2FD23bd0A5E08C5f2b9DD6CF9C96Bfb5fA515C", // 0x9f2fd23bd0a5e08c5f2b9dd6cf9c96bfb5fa515c
			// RACA
			"0x12BB890508c125661E03b09EC06E404bc9289040": "0xE5FE5527A5b76C75eedE77FdFA6B80D52444A465", // 0xe5fe5527a5b76c75eede77fdfa6b80d52444a465
			// FLOKI
			"0xfb5B838b6cfEEdC2873aB27866079AC55363D37E": "0xc353B7a1E13dDba393B5E120D4169Da7185aA2cb", // 0xc353b7a1e13ddba393b5e120d4169da7185aa2cb
		},
		VenusComptrollersList[ChainBSC][VenusProtocolIsolatedMeme]: {
			USDTList[ChainBSC]: "0x4a9613D06a241B76b81d3777FCe3DDd1F61D4Bd0", // 0x4a9613d06a241b76b81d3777fce3ddd1f61d4bd0
			// BabyDoge
			"0xc748673057861a797275CD8A068AbB95A902e8de": "0x52eD99Cd0a56d60451dD4314058854bc0845bbB5", // 0x52ed99cd0a56d60451dd4314058854bc0845bbb5
		},
		VenusComptrollersList[ChainBSC][VenusProtocolIsolatedTron]: {
			TRXList[ChainBSC]:  "0x836beb2cB723C498136e1119248436A645845F4E", // 0x836beb2cb723c498136e1119248436a645845f4e
			USDTList[ChainBSC]: "0x281E5378f99A4bc55b295ABc0A3E7eD32Deba059", // 0x281e5378f99a4bc55b295abc0a3e7ed32deba059
			USDDList[ChainBSC]: "0xf1da185CCe5BeD1BeBbb3007Ef738Ea4224025F7", // 0xf1da185cce5bed1bebbb3007ef738ea4224025f7
			// BTT
			"0x352Cb5E19b12FC216548a2677bD0fce83BaE434B": "0x49c26e12959345472E2Fd95E5f79F8381058d3Ee", // 0x49c26e12959345472e2fd95e5f79f8381058d3ee
			// WIN
			"0xaeF0d72a118ce24feE3cD1d43d383897D05B4e99": "0xb114cfA615c828D88021a41bFc524B800E64a9D5", // 0xb114cfa615c828d88021a41bfc524b800e64a9d5
		},
	},
	ChainEthereum: {
		VenusComptrollersList[ChainEthereum][VenusProtocolCore]: {
			WBTCList[ChainEthereum]:   "0x8716554364f20BCA783cb2BAA744d39361fd1D8d", // 0x8716554364f20bca783cb2baa744d39361fd1d8d
			WETHList[ChainEthereum]:   "0x7c8ff7d2A1372433726f879BD945fFb250B94c65", // 0x7c8ff7d2a1372433726f879bd945ffb250b94c65
			USDCList[ChainEthereum]:   "0x17C07e0c232f2f80DfDbd7a95b942D893A4C5ACb", // 0x17c07e0c232f2f80dfdbd7a95b942d893a4c5acb
			USDTList[ChainEthereum]:   "0x8C3e3821259B82fFb32B2450A95d2dcbf161C24E", // 0x8c3e3821259b82ffb32b2450a95d2dcbf161c24e
			CrvUSDList[ChainEthereum]: "0x672208C10aaAA2F9A6719F449C4C8227bc0BC202", // 0x672208c10aaaa2f9a6719f449c4c8227bc0bc202
			DAIList[ChainEthereum]:    "0xd8AdD9B41D4E1cd64Edad8722AB0bA8D35536657", // 0xd8add9b41d4e1cd64edad8722ab0ba8d35536657
			TUSDList[ChainEthereum]:   "0x13eB80FDBe5C5f4a7039728E258A6f05fb3B912b", // 0x13eb80fdbe5c5f4a7039728e258a6f05fb3b912b
			FRAXList[ChainEthereum]:   "0x4fAfbDc4F2a9876Bd1764827b26fb8dc4FD1dB95", // 0x4fafbdc4f2a9876bd1764827b26fb8dc4fd1db95
			SFRAXList[ChainEthereum]:  "0x17142a05fe678e9584FA1d88EfAC1bF181bF7ABe", // 0x17142a05fe678e9584fa1d88efac1bf181bf7abe
			EIGENList[ChainEthereum]:  "0x256AdDBe0a387c98f487e44b85c29eb983413c5e", // 0x256addbe0a387c98f487e44b85c29eb983413c5e
			EBTCList[ChainEthereum]:   "0x325cEB02fe1C2fF816A83a5770eA0E88e2faEcF2", // 0x325ceb02fe1c2ff816a83a5770ea0e88e2faecf2
			LBTCList[ChainEthereum]:   "0x25C20e6e110A1cE3FEbaCC8b7E48368c7b2F0C91", // 0x25c20e6e110a1ce3febacc8b7e48368c7b2f0c91
			USDSList[ChainEthereum]:   "0x0c6B19287999f1e31a5c0a44393b24B62D2C0468", // 0x0c6b19287999f1e31a5c0a44393b24b62d2c0468
			SUSDSList[ChainEthereum]:  "0xE36Ae842DbbD7aE372ebA02C8239cd431cC063d6", // 0xe36ae842dbbd7ae372eba02c8239cd431cc063d6
			BALList[ChainEthereum]:    "0x0Ec5488e4F8f319213a14cab188E01fB8517Faa8", // 0x0ec5488e4f8f319213a14cab188e01fb8517faa8
			SUSDeList[ChainEthereum]:  "0xa836ce315b7A6Bb19397Ee996551659B1D92298e", // 0xa836ce315b7a6bb19397ee996551659b1d92298e
			USDeList[ChainEthereum]:   "0xa0EE2bAA024cC3AA1BC9395522D07B7970Ca75b3", // 0xa0ee2baa024cc3aa1bc9395522d07b7970ca75b3
			TBTCList[ChainEthereum]:   "0x5e35C312862d53FD566737892aDCf010cb4928F7", // 0x5e35c312862d53fd566737892adcf010cb4928f7
			// yvUSDC-1
			"0xBe53A109B494E5c9f97b9Cd39Fe969BE68BF6204": "0xf87c0a64dc3a8622D6c63265FA29137788163879", // 0xf87c0a64dc3a8622d6c63265fa29137788163879
			// yvUSDT-1
			"0x310B7Ea7475A0B449Cfd73bE81522F1B88eFAFaa": "0x475d0C68a8CD275c15D1F01F4f291804E445F677", // 0x475d0c68a8cd275c15d1f01f4f291804e445f677
			// yvUSDS-1
			"0x182863131F9a4630fF9E27830d945B1413e347E8": "0x520d67226Bc904aC122dcE66ed2f8f61AA1ED764", // 0x520d67226bc904ac122dce66ed2f8f61aa1ed764
			// yvWETH-1
			"0xc56413869c6CDf96496f2b1eF801fEDBdFA7dDB0": "0xba3916302cBA4aBcB51a01e706fC6051AaF272A0", // 0xba3916302cba4abcb51a01e706fc6051aaf272a0
			// weETHs
			"0x917ceE801a67f933F2e6b33fC0cD1ED2d5909D88": "0xc42E4bfb996ED35235bda505430cBE404Eb49F77", // 0xc42e4bfb996ed35235bda505430cbe404eb49f77
		},
	},
}

// DEPRECATED: Use VenusComptrollersList[chainID][protocol] instead.
var VenusComptrollerList = map[int64]string{
	ChainArbitrum: VenusComptrollersList[ChainArbitrum][VenusProtocolCore],
	ChainBase:     VenusComptrollersList[ChainBase][VenusProtocolCore],
	ChainBSC:      VenusComptrollersList[ChainBSC][VenusProtocolCore],
	ChainEthereum: VenusComptrollersList[ChainEthereum][VenusProtocolCore],
}

// DEPRECATED: Use VenusCTokensListByComptroller instead.
var VenusCTokenList = map[int64]map[string]string{
	ChainArbitrum: VenusCTokensListByComptroller[ChainArbitrum][VenusProtocolCore],
	ChainBSC:      VenusCTokensListByComptroller[ChainBSC][VenusProtocolCore],
	ChainEthereum: VenusCTokensListByComptroller[ChainEthereum][VenusProtocolCore],
}
