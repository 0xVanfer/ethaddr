package ethaddr

import "github.com/0xVanfer/chainId"

// TODO: checksum

// Docs: https://docs.aave.com/developers/getting-started/readme
//
// Deployed contracts: https://docs.aave.com/developers/deployed-contracts/deployed-contracts
//
// AToken, VToken, underlying: https://github.com/bgd-labs/aave-address-book/blob/main/tokenlist.json
const AaveV3Protocol string = "aavev3"

// Aave incentive controller V3.
//
// map[network] = address.
var AaveIncentivesControllerV3List = map[string]string{
	chainId.ArbitrumChainName:  "0x929EC64c34a17401F460460D4B9390518E5B473e", // 0x929ec64c34a17401f460460d4b9390518e5b473e
	chainId.AvalancheChainName: "0x929EC64c34a17401F460460D4B9390518E5B473e", // 0x929ec64c34a17401f460460d4b9390518e5b473e
	chainId.BaseChainName:      "0x4D0109d509e66dF298257FfdD55F94A3814343Aa", // 0x4d0109d509e66df298257ffdd55f94a3814343aa
	chainId.EthereumChainName:  "0x8164Cc65827dcFe994AB23944CBC90e0aa80bFcb", // 0x8164cc65827dcfe994ab23944cbc90e0aa80bfcb
	chainId.OptimismChainName:  "0x929EC64c34a17401F460460D4B9390518E5B473e", // 0x929ec64c34a17401f460460d4b9390518e5b473e
	chainId.PolygonChainName:   "0x929EC64c34a17401F460460D4B9390518E5B473e", // 0x929ec64c34a17401f460460d4b9390518e5b473e
}

// Aave lending pool v3.
//
// map[network] = address.
var AaveLendingPoolV3List = map[string]string{
	chainId.ArbitrumChainName:  "0x794a61358D6845594F94dc1DB02A252b5b4814aD", // 0x794a61358d6845594f94dc1db02a252b5b4814ad
	chainId.AvalancheChainName: "0x794a61358D6845594F94dc1DB02A252b5b4814aD", // 0x794a61358d6845594f94dc1db02a252b5b4814ad
	chainId.BaseChainName:      "0xA238Dd80C259a72e81d7e4664a9801593F98d1c5", // 0xa238dd80c259a72e81d7e4664a9801593f98d1c5
	chainId.EthereumChainName:  "0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2", // 0x87870bca3f3fd6335c3f4ce8392d69350b4fa4e2
	chainId.FantomChainName:    "0x794a61358D6845594F94dc1DB02A252b5b4814aD", // 0x794a61358d6845594f94dc1db02a252b5b4814ad
	chainId.OptimismChainName:  "0x794a61358D6845594F94dc1DB02A252b5b4814aD", // 0x794a61358d6845594f94dc1db02a252b5b4814ad
	chainId.PolygonChainName:   "0x794a61358D6845594F94dc1DB02A252b5b4814aD", // 0x794a61358d6845594f94dc1db02a252b5b4814ad
}

// Aave pool data provider v3.
//
// map[network] = address.
var AavePoolDataProviderList = map[string]string{
	chainId.ArbitrumChainName:  "0x69FA688f1Dc47d4B5d8029D5a35FB7a548310654", // 0x69fa688f1dc47d4b5d8029d5a35fb7a548310654
	chainId.AvalancheChainName: "0x69FA688f1Dc47d4B5d8029D5a35FB7a548310654", // 0x69fa688f1dc47d4b5d8029d5a35fb7a548310654
	chainId.BaseChainName:      "0x2d8A3C5677189723C4cB8873CfC9C8976FDF38Ac", // 0x2d8a3c5677189723c4cb8873cfc9c8976fdf38ac
	chainId.EthereumChainName:  "0x7B4EB56E7CD4b454BA8ff71E4518426369a138a3", // 0x7b4eb56e7cd4b454ba8ff71e4518426369a138a3
	chainId.FantomChainName:    "0x69FA688f1Dc47d4B5d8029D5a35FB7a548310654", // 0x69fa688f1dc47d4b5d8029d5a35fb7a548310654
	chainId.OptimismChainName:  "0x69FA688f1Dc47d4B5d8029D5a35FB7a548310654", // 0x69fa688f1dc47d4b5d8029d5a35fb7a548310654
	chainId.PolygonChainName:   "0x69FA688f1Dc47d4B5d8029D5a35FB7a548310654", // 0x69fa688f1dc47d4b5d8029d5a35fb7a548310654
}

// Aave ui pool data provider v3.
//
// map[network] = address.
var AaveUiPoolDataProviderV3List = map[string]string{
	chainId.ArbitrumChainName:  "0x145dE30c929a065582da84Cf96F88460dB9745A7", // 0x145de30c929a065582da84cf96f88460db9745a7
	chainId.AvalancheChainName: "0xF71DBe0FAEF1473ffC607d4c555dfF0aEaDb878d", // 0xf71dbe0faef1473ffc607d4c555dff0aeadb878d
	chainId.BaseChainName:      "0x174446a6741300cD2E7C1b1A636Fee99c8F83502", // 0x174446a6741300cd2e7c1b1a636fee99c8f83502
	chainId.EthereumChainName:  "0x91c0eA31b49B69Ea18607702c5d9aC360bf3dE7d", // 0x91c0ea31b49b69ea18607702c5d9ac360bf3de7d
	chainId.OptimismChainName:  "0xbd83DdBE37fc91923d59C8c1E0bDe0CccCa332d5", // 0xbd83ddbe37fc91923d59c8c1e0bde0cccca332d5
	chainId.PolygonChainName:   "0xC69728f11E9E6127733751c8410432913123acf1", // 0xc69728f11e9e6127733751c8410432913123acf1
}

// Aave ui incentive data provider v3.
//
// map[network] = address.
var AaveUiIncentiveDataProviderV3List = map[string]string{
	chainId.ArbitrumChainName:  "0xDA67AF3403555Ce0AE3ffC22fDb7354458277358", // 0xda67af3403555ce0ae3ffc22fdb7354458277358
	chainId.AvalancheChainName: "0x265d414f80b0fca9505710e6F16dB4b67555D365", // 0x265d414f80b0fca9505710e6f16db4b67555d365
	chainId.BaseChainName:      "0xEdD3b4737C1a0011626631a977b91Cf3E944982d", // 0xedd3b4737c1a0011626631a977b91cf3e944982d
	chainId.EthereumChainName:  "0x162A7AC02f547ad796CA549f757e2b8d1D9b10a6", // 0x162a7ac02f547ad796ca549f757e2b8d1d9b10a6
	chainId.OptimismChainName:  "0x6F143FE2F7B02424ad3CaD1593D6f36c0Aab69d7", // 0x6f143fe2f7b02424ad3cad1593d6f36c0aab69d7
	chainId.PolygonChainName:   "0x874313A46e4957D29FAAC43BF5Eb2B144894f557", // 0x874313a46e4957d29faac43bf5eb2b144894f557
}

// Aave pool address provider v3.
//
// map[network] = address.
var AavePoolAddressProviderV3List = map[string]string{
	chainId.ArbitrumChainName:  "0xa97684ead0e402dC232d5A977953DF7ECBaB3CDb", // 0xa97684ead0e402dc232d5a977953df7ecbab3cdb
	chainId.AvalancheChainName: "0xa97684ead0e402dC232d5A977953DF7ECBaB3CDb", // 0xa97684ead0e402dc232d5a977953df7ecbab3cdb
	chainId.BaseChainName:      "0xe20fCBdBfFC4Dd138cE8b2E6FBb6CB49777ad64D", // 0xe20fcbdbffc4dd138ce8b2e6fbb6cb49777ad64d
	chainId.EthereumChainName:  "0x2f39d218133AFaB8F2B819B1066c7E434Ad94E9e", // 0x2f39d218133afab8f2b819b1066c7e434ad94e9e
	chainId.OptimismChainName:  "0xa97684ead0e402dC232d5A977953DF7ECBaB3CDb", // 0xa97684ead0e402dc232d5a977953df7ecbab3cdb
	chainId.PolygonChainName:   "0xa97684ead0e402dC232d5A977953DF7ECBaB3CDb", // 0xa97684ead0e402dc232d5a977953df7ecbab3cdb
}

// Aave a tokens v3.
//
// map[network][underlying] = address.
var AaveATokenV3List = map[string]map[string]string{
	chainId.EthereumChainName: {
		WETHList[chainId.EthereumChainName]:   "0x4d5F47FA6A74757f35C14fD3a6Ef8E3C9BC514E8", // aEthWETH, 0x4d5f47fa6a74757f35c14fd3a6ef8e3c9bc514e8
		WSTETHList[chainId.EthereumChainName]: "0x0B925eD163218f6662a35e0f0371Ac234f9E9371", // aEthwstETH, 0x0b925ed163218f6662a35e0f0371ac234f9e9371
		WBTCList[chainId.EthereumChainName]:   "0x5Ee5bf7ae06D1Be5997A1A72006FE6C607eC6DE8", // aEthWBTC, 0x5ee5bf7ae06d1be5997a1a72006fe6c607ec6de8
		USDCList[chainId.EthereumChainName]:   "0x98C23E9d8f34FEFb1B7BD6a91B7FF122F4e16F5c", // aEthUSDC, 0x98c23e9d8f34fefb1b7bd6a91b7ff122f4e16f5c
		DAIList[chainId.EthereumChainName]:    "0x018008bfb33d285247A21d44E50697654f754e63", // aEthDAI, 0x018008bfb33d285247a21d44e50697654f754e63
		LINKList[chainId.EthereumChainName]:   "0x5E8C8A7243651DB1384C0dDfDbE39761E8e7E51a", // aEthLINK, 0x5e8c8a7243651db1384c0ddfdbe39761e8e7e51a
		AAVEList[chainId.EthereumChainName]:   "0xA700b4eB416Be35b2911fd5Dee80678ff64fF6C9", // aEthAAVE, 0xa700b4eb416be35b2911fd5dee80678ff64ff6c9
		CbETHList[chainId.EthereumChainName]:  "0x977b6fc5dE62598B08C85AC8Cf2b745874E8b78c", // aEthcbETH, 0x977b6fc5de62598b08c85ac8cf2b745874e8b78c
		USDTList[chainId.EthereumChainName]:   "0x23878914EFE38d27C4D67Ab83ed1b93A74D4086a", // aEthUSDT, 0x23878914efe38d27c4d67ab83ed1b93a74d4086a
		RETHList[chainId.EthereumChainName]:   "0xCc9EE9483f662091a1de4795249E24aC0aC2630f", // aEthrETH, 0xcc9ee9483f662091a1de4795249e24ac0ac2630f
		LUSDList[chainId.EthereumChainName]:   "0x3Fe6a295459FAe07DF8A0ceCC36F37160FE86AA9", // aEthLUSD, 0x3fe6a295459fae07df8a0cecc36f37160fe86aa9
		CRVList[chainId.EthereumChainName]:    "0x7B95Ec873268a6BFC6427e7a28e396Db9D0ebc65", // aEthCRV, 0x7b95ec873268a6bfc6427e7a28e396db9d0ebc65
		MKRList[chainId.EthereumChainName]:    "0x8A458A9dc9048e005d22849F470891b840296619", // aEthMKR, 0x8a458a9dc9048e005d22849f470891b840296619
		SNXList[chainId.EthereumChainName]:    "0xC7B4c17861357B8ABB91F25581E7263E08DCB59c", // aEthSNX, 0xc7b4c17861357b8abb91f25581e7263e08dcb59c
		BALList[chainId.EthereumChainName]:    "0x2516E7B3F76294e03C42AA4c5b5b4DCE9C436fB8", // aEthBAL, 0x2516e7b3f76294e03c42aa4c5b5b4dce9c436fb8
		UNIList[chainId.EthereumChainName]:    "0xF6D2224916DDFbbab6e6bd0D1B7034f4Ae0CaB18", // aEthUNI, 0xf6d2224916ddfbbab6e6bd0d1b7034f4ae0cab18
		LDOList[chainId.EthereumChainName]:    "0x9A44fd41566876A39655f74971a3A6eA0a17a454", // aEthLDO, 0x9a44fd41566876a39655f74971a3a6ea0a17a454
		ENSList[chainId.EthereumChainName]:    "0x545bD6c032eFdde65A377A6719DEF2796C8E0f2e", // aEthENS, 0x545bd6c032efdde65a377a6719def2796c8e0f2e
		OINCHLIST[chainId.EthereumChainName]:  "0x71Aef7b30728b9BB371578f36c5A1f1502a5723e", // aEth1INCH, 0x71aef7b30728b9bb371578f36c5a1f1502a5723e
		FRAXList[chainId.EthereumChainName]:   "0xd4e245848d6E1220DBE62e155d89fa327E43CB06", // aEthFRAX, 0xd4e245848d6e1220dbe62e155d89fa327e43cb06
		GHOList[chainId.EthereumChainName]:    "0x00907f9921424583e7ffBfEdf84F92B7B2Be4977", // aEthGHO, 0x00907f9921424583e7ffbfedf84f92b7b2be4977
		RPLList[chainId.EthereumChainName]:    "0xB76CF92076adBF1D9C39294FA8e7A67579FDe357", // aEthRPL, 0xb76cf92076adbf1d9c39294fa8e7a67579fde357
		SDAIList[chainId.EthereumChainName]:   "0x4C612E3B15b96Ff9A6faED838F8d07d479a8dD4c", // aEthsDAI, 0x4c612e3b15b96ff9a6faed838f8d07d479a8dd4c
		STGList[chainId.EthereumChainName]:    "0x1bA9843bD4327c6c77011406dE5fA8749F7E3479", // aEthSTG, 0x1ba9843bd4327c6c77011406de5fa8749f7e3479
		KNCList[chainId.EthereumChainName]:    "0x5b502e3796385E1e9755d7043B9C945C3aCCeC9C", // aEthKNC, 0x5b502e3796385e1e9755d7043b9c945c3accec9c
		FXSList[chainId.EthereumChainName]:    "0x82F9c5ad306BBa1AD0De49bB5FA6F01bf61085ef", // aEthFXS, 0x82f9c5ad306bba1ad0de49bb5fa6f01bf61085ef
		CrvUSDList[chainId.EthereumChainName]: "0xb82fa9f31612989525992FCfBB09AB22Eff5c85A", // aEthcrvUSD, 0xb82fa9f31612989525992fcfbb09ab22eff5c85a
		PyUSDList[chainId.EthereumChainName]:  "0x0C0d01AbF3e6aDfcA0989eBbA9d6e85dD58EaB1E", // aEthPYUSD, 0x0c0d01abf3e6adfca0989ebba9d6e85dd58eab1e
	},
	chainId.AvalancheChainName: {
		DAIeList[chainId.AvalancheChainName]:  "0x82e64f49ed5ec1bc6e43dad4fc8af9bb3a2312ee", // aAvaDAI
		USDCList[chainId.AvalancheChainName]:  "0x625e7708f30ca75bfd92586e17077590c60eb4cd", // aAvaUSDC
		USDTList[chainId.AvalancheChainName]:  "0x6ab707aca953edaefbc4fd23ba73294241490620", // aAvaUSDT
		WETHList[chainId.AvalancheChainName]:  "0xe50fa9b3c56ffb159cb0fca61f5c9d750e8128c8", // aAvaWETH
		WBTCList[chainId.AvalancheChainName]:  "0x078f358208685046a11c85e8ad32895ded33a249", // aAvaWBTC
		WAVAXList[chainId.AvalancheChainName]: "0x6d80113e533a2c0fe82eabd35f1875dcea89ea97", // aAvaWAVAX
		LINKList[chainId.AvalancheChainName]:  "0x191c10aa4af7c30e871e70c95db0e4eb77237530", // aAvaLINK
		AAVEList[chainId.AvalancheChainName]:  "0xf329e36c7bf6e5e86ce2150875a84ce77f477375", // aAvaAAVE
		SAVAXList[chainId.AvalancheChainName]: "0x513c7e3a9c69ca3e22550ef58ac1c0088e918fff", // aAvaSAVAX
		BTCbList[chainId.AvalancheChainName]:  "0x8ffdf2de812095b1d19cb146e4c004587c0a0692", // aAvaBTC.b
		FRAXList[chainId.AvalancheChainName]:  "0xc45a479877e1e9dfe9fcd4056c699575a1045daa",
		MAIList[chainId.AvalancheChainName]:   "0x8eb270e296023e9d92081fdf967ddd7878724424",
	},
	chainId.OptimismChainName: {
		DAIList[chainId.OptimismChainName]:  "0x82e64f49ed5ec1bc6e43dad4fc8af9bb3a2312ee", // aOptDAI
		USDCList[chainId.OptimismChainName]: "0x625e7708f30ca75bfd92586e17077590c60eb4cd", // aOptUSDC
		USDTList[chainId.OptimismChainName]: "0x6ab707aca953edaefbc4fd23ba73294241490620", // aOptUSDT
		WETHList[chainId.OptimismChainName]: "0xe50fa9b3c56ffb159cb0fca61f5c9d750e8128c8", // aOptWETH
		WBTCList[chainId.OptimismChainName]: "0x078f358208685046a11c85e8ad32895ded33a249", // aOptWBTC
		SUSDList[chainId.OptimismChainName]: "0x6d80113e533a2c0fe82eabd35f1875dcea89ea97", // aOptSUSD
		LINKList[chainId.OptimismChainName]: "0x191c10aa4af7c30e871e70c95db0e4eb77237530", // aOptLINK
		AAVEList[chainId.OptimismChainName]: "0xf329e36C7bF6E5E86ce2150875a84Ce77f477375", // aOptAAVE
	},
	chainId.PolygonChainName: {
		AgEURList[chainId.PolygonChainName]:   "0x8437d7c167dfb82ed4cb79cd44b7a32a1dd95c77", // aPolAGEUR
		DAIList[chainId.PolygonChainName]:     "0x82e64f49ed5ec1bc6e43dad4fc8af9bb3a2312ee", // aPolDAI
		EURSList[chainId.PolygonChainName]:    "0x38d693ce1df5aadf7bc62595a37d667ad57922e5",
		JEURList[chainId.PolygonChainName]:    "0x6533afac2e7bccb20dca161449a13a32d391fb00",
		MiMATICList[chainId.PolygonChainName]: "0xebe517846d0f36eced99c735cbf6131e1feb775d",
		USDCList[chainId.PolygonChainName]:    "0x625e7708f30ca75bfd92586e17077590c60eb4cd",
		USDTList[chainId.PolygonChainName]:    "0x6ab707aca953edaefbc4fd23ba73294241490620",
		AAVEList[chainId.PolygonChainName]:    "0xf329e36c7bf6e5e86ce2150875a84ce77f477375",
		BALList[chainId.PolygonChainName]:     "0x8ffdf2de812095b1d19cb146e4c004587c0a0692",
		CRVList[chainId.PolygonChainName]:     "0x513c7e3a9c69ca3e22550ef58ac1c0088e918fff",
		DPIList[chainId.PolygonChainName]:     "0x724dc807b04555b71ed48a6896b6f41593b8c637",
		GHSTList[chainId.PolygonChainName]:    "0x8eb270e296023e9d92081fdf967ddd7878724424",
		LINKList[chainId.PolygonChainName]:    "0x191c10aa4af7c30e871e70c95db0e4eb77237530",
		MaticXList[chainId.PolygonChainName]:  "0x80ca0d8c38d2e2bcbab66aa1648bd1c7160500fe",
		STMATICList[chainId.PolygonChainName]: "0xea1132120ddcdda2f119e99fa7a27a0d036f7ac9",
		SUSHIList[chainId.PolygonChainName]:   "0xc45a479877e1e9dfe9fcd4056c699575a1045daa",
		WBTCList[chainId.PolygonChainName]:    "0x078f358208685046a11c85e8ad32895ded33a249",
		WETHList[chainId.PolygonChainName]:    "0xe50fa9b3c56ffb159cb0fca61f5c9d750e8128c8",
		WMATICList[chainId.PolygonChainName]:  "0x6d80113e533a2c0fe82eabd35f1875dcea89ea97",
	},
	chainId.ArbitrumChainName: {
		DAIList[chainId.ArbitrumChainName]:  "0x82e64f49ed5ec1bc6e43dad4fc8af9bb3a2312ee", // aArbDAI
		EURSList[chainId.ArbitrumChainName]: "0x6d80113e533a2c0fe82eabd35f1875dcea89ea97", // aArbEURS
		USDCList[chainId.ArbitrumChainName]: "0x625e7708f30ca75bfd92586e17077590c60eb4cd",
		USDTList[chainId.ArbitrumChainName]: "0x6ab707aca953edaefbc4fd23ba73294241490620",
		AAVEList[chainId.ArbitrumChainName]: "0xf329e36c7bf6e5e86ce2150875a84ce77f477375",
		LINKList[chainId.ArbitrumChainName]: "0x191c10aa4af7c30e871e70c95db0e4eb77237530",
		WBTCList[chainId.ArbitrumChainName]: "0x078f358208685046a11c85e8ad32895ded33a249",
		WETHList[chainId.ArbitrumChainName]: "0xe50fa9b3c56ffb159cb0fca61f5c9d750e8128c8",
	},
	chainId.BaseChainName: {
		WETHList[chainId.BaseChainName]:   "0xD4a0e0b9149BCee3C920d2E00b5dE09138fd8bb7", // aBasWETH
		CbETHList[chainId.BaseChainName]:  "0xcf3D55c10DB69f28fD1A75Bd73f3D8A2d9c595ad", // aBascbETH
		USDbCList[chainId.BaseChainName]:  "0x0a1d576f3eFeF75b330424287a95A366e8281D54", // aBasUSDbC
		WSTETHList[chainId.BaseChainName]: "0x99CBC45ea5bb7eF3a5BC08FB1B7E56bB2442Ef0D", // aBaswstETH
		USDCList[chainId.BaseChainName]:   "0x4e65fE4DbA92790696d040ac24Aa414708F5c0AB", // aBasUSDC
	},
}

// Aave static a tokens v3. The wrapped, non-rebase version of a token.
//
// map[network][underlying] = address.
var AaveStaticATokenV3List = map[string]map[string]string{
	// wstETH does not have static a token.
	chainId.BaseChainName: {
		WETHList[chainId.BaseChainName]:  "0x468973e3264F2aEba0417A8f2cD0Ec397E738898", // stataBasWETH, 0x468973e3264f2aeba0417a8f2cd0ec397e738898
		CbETHList[chainId.BaseChainName]: "0x16A004065dfb11276DcB29Dc03fb8A85f9A43C6e", // stataBascbETH, 0x16a004065dfb11276dcb29dc03fb8a85f9a43c6e
		USDbCList[chainId.BaseChainName]: "0x6fCe2756794128B1771324caA860965801DCbCdB", // stataBasUSDbC, 0x6fce2756794128b1771324caa860965801dcbcdb
		USDCList[chainId.BaseChainName]:  "0x4EA71A20e655794051D1eE8b6e4A3269B13ccaCc", // stataBasUSDC, 0x4ea71a20e655794051d1ee8b6e4a3269b13ccacc
	},
}

// Aave s tokens v3.
//
// map[network][underlying] = address.
var AaveSTokenV3List = map[string]map[string]string{
	chainId.EthereumChainName: {
		WETHList[chainId.EthereumChainName]:   "0x102633152313C81cD80419b6EcF66d14Ad68949A", // stableDebtEthWETH, 0x102633152313c81cd80419b6ecf66d14ad68949a
		WSTETHList[chainId.EthereumChainName]: "0x39739943199c0fBFe9E5f1B5B160cd73a64CB85D", // stableDebtEthwstETH, 0x39739943199c0fbfe9e5f1b5b160cd73a64cb85d
		WBTCList[chainId.EthereumChainName]:   "0xA1773F1ccF6DB192Ad8FE826D15fe1d328B03284", // stableDebtEthWBTC, 0xa1773f1ccf6db192ad8fe826d15fe1d328b03284
		USDCList[chainId.EthereumChainName]:   "0xB0fe3D292f4bd50De902Ba5bDF120Ad66E9d7a39", // stableDebtEthUSDC, 0xb0fe3d292f4bd50de902ba5bdf120ad66e9d7a39
		DAIList[chainId.EthereumChainName]:    "0x413AdaC9E2Ef8683ADf5DDAEce8f19613d60D1bb", // stableDebtEthDAI, 0x413adac9e2ef8683adf5ddaece8f19613d60d1bb
		LINKList[chainId.EthereumChainName]:   "0x63B1129ca97D2b9F97f45670787Ac12a9dF1110a", // stableDebtEthLINK, 0x63b1129ca97d2b9f97f45670787ac12a9df1110a
		AAVEList[chainId.EthereumChainName]:   "0x268497bF083388B1504270d0E717222d3A87D6F2", // stableDebtEthAAVE, 0x268497bf083388b1504270d0e717222d3a87d6f2
		CbETHList[chainId.EthereumChainName]:  "0x82bE6012cea6D147B968eBAea5ceEcF6A5b4F493", // stableDebtEthcbETH, 0x82be6012cea6d147b968ebaea5ceecf6a5b4f493
		USDTList[chainId.EthereumChainName]:   "0x822Fa72Df1F229C3900f5AD6C3Fa2C424D691622", // stableDebtEthUSDT, 0x822fa72df1f229c3900f5ad6c3fa2c424d691622
		RETHList[chainId.EthereumChainName]:   "0x1d1906f909CAe494c7441604DAfDDDbD0485A925", // stableDebtEthrETH, 0x1d1906f909cae494c7441604dafdddbd0485a925
		LUSDList[chainId.EthereumChainName]:   "0x37A6B708FDB1483C231961b9a7F145261E815fc3", // stableDebtEthLUSD, 0x37a6b708fdb1483c231961b9a7f145261e815fc3
		CRVList[chainId.EthereumChainName]:    "0x90D9CD005E553111EB8C9c31Abe9706a186b6048", // stableDebtEthCRV, 0x90d9cd005e553111eb8c9c31abe9706a186b6048
		MKRList[chainId.EthereumChainName]:    "0x0496372BE7e426D28E89DEBF01f19F014d5938bE", // stableDebtEthMKR, 0x0496372be7e426d28e89debf01f19f014d5938be
		SNXList[chainId.EthereumChainName]:    "0x478E1ec1A2BeEd94c1407c951E4B9e22d53b2501", // stableDebtEthSNX, 0x478e1ec1a2beed94c1407c951e4b9e22d53b2501
		BALList[chainId.EthereumChainName]:    "0xB368d45aaAa07ee2c6275Cb320D140b22dE43CDD", // stableDebtEthBAL, 0xb368d45aaaa07ee2c6275cb320d140b22de43cdd
		UNIList[chainId.EthereumChainName]:    "0x2FEc76324A0463c46f32e74A86D1cf94C02158DC", // stableDebtEthUNI, 0x2fec76324a0463c46f32e74a86d1cf94c02158dc
		LDOList[chainId.EthereumChainName]:    "0xa0a5bF5781Aeb548db9d4226363B9e89287C5FD2", // stableDebtEthLDO, 0xa0a5bf5781aeb548db9d4226363b9e89287c5fd2
		ENSList[chainId.EthereumChainName]:    "0x7617d02E311CdE347A0cb45BB7DF2926BBaf5347", // stableDebtEthENS, 0x7617d02e311cde347a0cb45bb7df2926bbaf5347
		OINCHLIST[chainId.EthereumChainName]:  "0x4b62bFAff61AB3985798e5202D2d167F567D0BCD", // stableDebtEth1INCH, 0x4b62bfaff61ab3985798e5202d2d167f567d0bcd
		FRAXList[chainId.EthereumChainName]:   "0x219640546c0DFDDCb9ab3bcdA89B324e0a376367", // stableDebtEthFRAX, 0x219640546c0dfddcb9ab3bcda89b324e0a376367
		GHOList[chainId.EthereumChainName]:    "0x3f3DF7266dA30102344A813F1a3D07f5F041B5AC", // stableDebtEthGHO, 0x3f3df7266da30102344a813f1a3d07f5f041b5ac
		RPLList[chainId.EthereumChainName]:    "0x41e330fd8F7eA31E2e8F02cC0C9392D1403597B4", // stableDebtEthRPL, 0x41e330fd8f7ea31e2e8f02cc0c9392d1403597b4
		SDAIList[chainId.EthereumChainName]:   "0x48Bc45f084988bC01933EA93EeFfEBC0416534f6", // stableDebtEthsDAI, 0x48bc45f084988bc01933ea93eeffebc0416534f6
		STGList[chainId.EthereumChainName]:    "0xc3115D0660b93AeF10F298886ae22E3Dd477E482", // stableDebtEthSTG, 0xc3115d0660b93aef10f298886ae22e3dd477e482
		KNCList[chainId.EthereumChainName]:    "0xdfEE0C9eA1309cB9611F33972E72a72166fcF548", // stableDebtEthKNC, 0xdfee0c9ea1309cb9611f33972e72a72166fcf548
		FXSList[chainId.EthereumChainName]:    "0x61dFd349140C239d3B61fEe203Efc811b518a317", // stableDebtEthFXS, 0x61dfd349140c239d3b61fee203efc811b518a317
		CrvUSDList[chainId.EthereumChainName]: "0xb55C604075D79486b8A329c396Fc711Be54B5330", // stableDebtEthcrvUSD, 0xb55c604075d79486b8a329c396fc711be54b5330
		PyUSDList[chainId.EthereumChainName]:  "0x5B393DB4c72B1Bd82CE2834F6485d61b137Bc7aC", // stableDebtEthPYUSD, 0x5b393db4c72b1bd82ce2834f6485d61b137bc7ac

	},
	chainId.AvalancheChainName: {
		DAIeList[chainId.AvalancheChainName]:  "0xd94112b5b62d53c9402e7a60289c6810def1dc9b", // stableDebtAvaDAI
		USDCList[chainId.AvalancheChainName]:  "0x307ffe186f84a3bc2613d1ea417a5737d69a7007", // stableDebtAvaUSDC
		USDTList[chainId.AvalancheChainName]:  "0x70effc565db6eef7b927610155602d31b670e802", // stableDebtAvaUSDT
		WETHList[chainId.AvalancheChainName]:  "0xd8ad37849950903571df17049516a5cd4cbe55f6", // stableDebtAvaWETH
		WBTCList[chainId.AvalancheChainName]:  "0x633b207dd676331c413d4c013a6294b0fe47cd0e", // stableDebtAvaWBTC
		WAVAXList[chainId.AvalancheChainName]: "0xf15f26710c827dde8acba678682f3ce24f2fb56e", // stableDebtAvaWAVAX
		LINKList[chainId.AvalancheChainName]:  "0x89d976629b7055ff1ca02b927ba3e020f22a44e4", // stableDebtAvaLINK
		AAVEList[chainId.AvalancheChainName]:  "0xfaef6a702d15428e588d4c0614aefb4348d83d48", // stableDebtAvaAAVE
		SAVAXList[chainId.AvalancheChainName]: "0x08cb71192985e936c7cd166a8b268035e400c3c3", // stableDebtAvaSAVAX
		BTCbList[chainId.AvalancheChainName]:  "0xa5e408678469d23efdb7694b1b0a85bb0669e8bd", //
		FRAXList[chainId.AvalancheChainName]:  "0x78246294a4c6fbf614ed73ccc9f8b875ca8ee841", //
		MAIList[chainId.AvalancheChainName]:   "0x3ef10dff4928279c004308ebadc4db8b7620d6fc", //
	},
	chainId.PolygonChainName: {
		AgEURList[chainId.PolygonChainName]:   "0x40b4baecc69b882e8804f9286b12228c27f8c9bf", // stableDebtPolAGEUR
		DAIList[chainId.PolygonChainName]:     "0xd94112b5b62d53c9402e7a60289c6810def1dc9b", //
		EURSList[chainId.PolygonChainName]:    "0x8a9fde6925a839f6b1932d16b36ac026f8d3fbdb",
		JEURList[chainId.PolygonChainName]:    "0x6b4b37618d85db2a7b469983c888040f7f05ea3d",
		MiMATICList[chainId.PolygonChainName]: "0x687871030477bf974725232f764aa04318a8b9c8",
		USDCList[chainId.PolygonChainName]:    "0x307ffe186f84a3bc2613d1ea417a5737d69a7007",
		USDTList[chainId.PolygonChainName]:    "0x70effc565db6eef7b927610155602d31b670e802",
		AAVEList[chainId.PolygonChainName]:    "0xfaef6a702d15428e588d4c0614aefb4348d83d48",
		BALList[chainId.PolygonChainName]:     "0xa5e408678469d23efdb7694b1b0a85bb0669e8bd",
		CRVList[chainId.PolygonChainName]:     "0x08cb71192985e936c7cd166a8b268035e400c3c3",
		DPIList[chainId.PolygonChainName]:     "0xdc1fad70953bb3918592b6fcc374fe05f5811b6a",
		GHSTList[chainId.PolygonChainName]:    "0x3ef10dff4928279c004308ebadc4db8b7620d6fc",
		LINKList[chainId.PolygonChainName]:    "0x89d976629b7055ff1ca02b927ba3e020f22a44e4",
		MaticXList[chainId.PolygonChainName]:  "0x62fc96b27a510cf4977b59ff952dc32378cc221d",
		STMATICList[chainId.PolygonChainName]: "0x1ffd28689da7d0148ff0fcb669e9f9f0fc13a219",
		SUSHIList[chainId.PolygonChainName]:   "0x78246294a4c6fbf614ed73ccc9f8b875ca8ee841",
		WBTCList[chainId.PolygonChainName]:    "0x633b207dd676331c413d4c013a6294b0fe47cd0e",
		WETHList[chainId.PolygonChainName]:    "0xd8ad37849950903571df17049516a5cd4cbe55f6",
		WMATICList[chainId.PolygonChainName]:  "0xf15f26710c827dde8acba678682f3ce24f2fb56e",
	},
	chainId.ArbitrumChainName: {
		DAIList[chainId.ArbitrumChainName]:  "0xd94112b5b62d53c9402e7a60289c6810def1dc9b", // sArbDAI
		EURSList[chainId.ArbitrumChainName]: "0xf15f26710c827dde8acba678682f3ce24f2fb56e",
		USDCList[chainId.ArbitrumChainName]: "0x307ffe186f84a3bc2613d1ea417a5737d69a7007",
		USDTList[chainId.ArbitrumChainName]: "0x70effc565db6eef7b927610155602d31b670e802",
		AAVEList[chainId.ArbitrumChainName]: "0xfaef6a702d15428e588d4c0614aefb4348d83d48",
		LINKList[chainId.ArbitrumChainName]: "0x89d976629b7055ff1ca02b927ba3e020f22a44e4",
		WBTCList[chainId.ArbitrumChainName]: "0x633b207dd676331c413d4c013a6294b0fe47cd0e",
		WETHList[chainId.ArbitrumChainName]: "0xd8ad37849950903571df17049516a5cd4cbe55f6",
	},
	chainId.BaseChainName: {
		WETHList[chainId.BaseChainName]:   "0xaED3b56FeA82E809665f02AcBcDEc0816c75f4d9",
		CbETHList[chainId.BaseChainName]:  "0xa9dF5c62d16d3f496673F4d736852017b086eCA0",
		USDbCList[chainId.BaseChainName]:  "0xBBaDd47fbaFa9dE717FE203e4707DEB893C64654",
		WSTETHList[chainId.BaseChainName]: "0xfe742Fa2a84294E8316F05b17c05090Fc68B5105",
		USDCList[chainId.BaseChainName]:   "0x03506214379aA86ad1176af71c260278cfa10B38",
	},
}

// Aave v tokens v3.
//
// map[network][underlying] = address.
var AaveVTokenV3List = map[string]map[string]string{
	chainId.EthereumChainName: {
		WETHList[chainId.EthereumChainName]:   "0xeA51d7853EEFb32b6ee06b1C12E6dcCA88Be0fFE", // variableDebtEthWETH, 0xea51d7853eefb32b6ee06b1c12e6dcca88be0ffe
		WSTETHList[chainId.EthereumChainName]: "0xC96113eED8cAB59cD8A66813bCB0cEb29F06D2e4", // variableDebtEthwstETH, 0xc96113eed8cab59cd8a66813bcb0ceb29f06d2e4
		WBTCList[chainId.EthereumChainName]:   "0x40aAbEf1aa8f0eEc637E0E7d92fbfFB2F26A8b7B", // variableDebtEthWBTC, 0x40aabef1aa8f0eec637e0e7d92fbffb2f26a8b7b
		USDCList[chainId.EthereumChainName]:   "0x72E95b8931767C79bA4EeE721354d6E99a61D004", // variableDebtEthUSDC, 0x72e95b8931767c79ba4eee721354d6e99a61d004
		DAIList[chainId.EthereumChainName]:    "0xcF8d0c70c850859266f5C338b38F9D663181C314", // variableDebtEthDAI, 0xcf8d0c70c850859266f5c338b38f9d663181c314
		LINKList[chainId.EthereumChainName]:   "0x4228F8895C7dDA20227F6a5c6751b8Ebf19a6ba8", // variableDebtEthLINK, 0x4228f8895c7dda20227f6a5c6751b8ebf19a6ba8
		AAVEList[chainId.EthereumChainName]:   "0xBae535520Abd9f8C85E58929e0006A2c8B372F74", // variableDebtEthAAVE, 0xbae535520abd9f8c85e58929e0006a2c8b372f74
		CbETHList[chainId.EthereumChainName]:  "0x0c91bcA95b5FE69164cE583A2ec9429A569798Ed", // variableDebtEthcbETH, 0x0c91bca95b5fe69164ce583a2ec9429a569798ed
		USDTList[chainId.EthereumChainName]:   "0x6df1C1E379bC5a00a7b4C6e67A203333772f45A8", // variableDebtEthUSDT, 0x6df1c1e379bc5a00a7b4c6e67a203333772f45a8
		RETHList[chainId.EthereumChainName]:   "0xae8593DD575FE29A9745056aA91C4b746eee62C8", // variableDebtEthrETH, 0xae8593dd575fe29a9745056aa91c4b746eee62c8
		LUSDList[chainId.EthereumChainName]:   "0x33652e48e4B74D18520f11BfE58Edd2ED2cEc5A2", // variableDebtEthLUSD, 0x33652e48e4b74d18520f11bfe58edd2ed2cec5a2
		CRVList[chainId.EthereumChainName]:    "0x1b7D3F4b3c032a5AE656e30eeA4e8E1Ba376068F", // variableDebtEthCRV, 0x1b7d3f4b3c032a5ae656e30eea4e8e1ba376068f
		MKRList[chainId.EthereumChainName]:    "0x6Efc73E54E41b27d2134fF9f98F15550f30DF9B1", // variableDebtEthMKR, 0x6efc73e54e41b27d2134ff9f98f15550f30df9b1
		SNXList[chainId.EthereumChainName]:    "0x8d0de040e8aAd872eC3c33A3776dE9152D3c34ca", // variableDebtEthSNX, 0x8d0de040e8aad872ec3c33a3776de9152d3c34ca
		BALList[chainId.EthereumChainName]:    "0x3D3efceb4Ff0966D34d9545D3A2fa2dcdBf451f2", // variableDebtEthBAL, 0x3d3efceb4ff0966d34d9545d3a2fa2dcdbf451f2
		UNIList[chainId.EthereumChainName]:    "0xF64178Ebd2E2719F2B1233bCb5Ef6DB4bCc4d09a", // variableDebtEthUNI, 0xf64178ebd2e2719f2b1233bcb5ef6db4bcc4d09a
		LDOList[chainId.EthereumChainName]:    "0xc30808705C01289A3D306ca9CAB081Ba9114eC82", // variableDebtEthLDO, 0xc30808705c01289a3d306ca9cab081ba9114ec82
		ENSList[chainId.EthereumChainName]:    "0xd180D7fdD4092f07428eFE801E17BC03576b3192", // variableDebtEthENS, 0xd180d7fdd4092f07428efe801e17bc03576b3192
		OINCHLIST[chainId.EthereumChainName]:  "0xA38fCa8c6Bf9BdA52E76EB78f08CaA3BE7c5A970", // variableDebtEth1INCH, 0xa38fca8c6bf9bda52e76eb78f08caa3be7c5a970
		FRAXList[chainId.EthereumChainName]:   "0x88B8358F5BC87c2D7E116cCA5b65A9eEb2c5EA3F", // variableDebtEthFRAX, 0x88b8358f5bc87c2d7e116cca5b65a9eeb2c5ea3f
		GHOList[chainId.EthereumChainName]:    "0x786dBff3f1292ae8F92ea68Cf93c30b34B1ed04B", // variableDebtEthGHO, 0x786dbff3f1292ae8f92ea68cf93c30b34b1ed04b
		RPLList[chainId.EthereumChainName]:    "0x8988ECA19D502fd8b9CCd03fA3bD20a6f599bc2A", // variableDebtEthRPL, 0x8988eca19d502fd8b9ccd03fa3bd20a6f599bc2a
		SDAIList[chainId.EthereumChainName]:   "0x8Db9D35e117d8b93C6Ca9b644b25BaD5d9908141", // variableDebtEthsDAI, 0x8db9d35e117d8b93c6ca9b644b25bad5d9908141
		STGList[chainId.EthereumChainName]:    "0x655568bDd6168325EC7e58Bf39b21A856F906Dc2", // variableDebtEthSTG, 0x655568bdd6168325ec7e58bf39b21a856f906dc2
		KNCList[chainId.EthereumChainName]:    "0x253127Ffc04981cEA8932F406710661c2f2c3fD2", // variableDebtEthKNC, 0x253127ffc04981cea8932f406710661c2f2c3fd2
		FXSList[chainId.EthereumChainName]:    "0x68e9f0aD4e6f8F5DB70F6923d4d6d5b225B83b16", // variableDebtEthFXS, 0x68e9f0ad4e6f8f5db70f6923d4d6d5b225b83b16
		CrvUSDList[chainId.EthereumChainName]: "0x028f7886F3e937f8479efaD64f31B3fE1119857a", // variableDebtEthcrvUSD, 0x028f7886f3e937f8479efad64f31b3fe1119857a
		PyUSDList[chainId.EthereumChainName]:  "0x57B67e4DE077085Fd0AF2174e9c14871BE664546", // variableDebtEthPYUSD, 0x57b67e4de077085fd0af2174e9c14871be664546
	},
	chainId.AvalancheChainName: {
		DAIeList[chainId.AvalancheChainName]:  "0x8619d80fb0141ba7f184cbf22fd724116d9f7ffc", // variableDebtAvaDAI
		USDCList[chainId.AvalancheChainName]:  "0xfccf3cabbe80101232d343252614b6a3ee81c989", // variableDebtAvaUSDC
		USDTList[chainId.AvalancheChainName]:  "0xfb00ac187a8eb5afae4eace434f493eb62672df7", // variableDebtAvaUSDT
		WETHList[chainId.AvalancheChainName]:  "0x0c84331e39d6658cd6e6b9ba04736cc4c4734351", // variableDebtAvaWETH
		WBTCList[chainId.AvalancheChainName]:  "0x92b42c66840c7ad907b4bf74879ff3ef7c529473", // variableDebtAvaWBTC
		WAVAXList[chainId.AvalancheChainName]: "0x4a1c3ad6ed28a636ee1751c69071f6be75deb8b8", // variableDebtAvaWAVAX
		LINKList[chainId.AvalancheChainName]:  "0x953a573793604af8d41f306feb8274190db4ae0e", // variableDebtAvaLINK
		AAVEList[chainId.AvalancheChainName]:  "0xe80761ea617f66f96274ea5e8c37f03960ecc679", // variableDebtAvaAAVE
		SAVAXList[chainId.AvalancheChainName]: "0x77ca01483f379e58174739308945f044e1a764dc", // variableDebtAvaSAVAX
		BTCbList[chainId.AvalancheChainName]:  "0xa8669021776bc142dfca87c21b4a52595bcbb40a", // variableDebtAvaBTC.b
		FRAXList[chainId.AvalancheChainName]:  "0x34e2ed44ef7466d5f9e0b782b5c08b57475e7907", //
		MAIList[chainId.AvalancheChainName]:   "0xce186f6cccb0c955445bb9d10c59cae488fea559", //
	},
	chainId.OptimismChainName: {
		DAIList[chainId.OptimismChainName]:  "0x8619d80fb0141ba7f184cbf22fd724116d9f7ffc", // variableDebtOpt
		USDCList[chainId.OptimismChainName]: "0xfccf3cabbe80101232d343252614b6a3ee81c989", // variableDebtOpt
		USDTList[chainId.OptimismChainName]: "0xfb00ac187a8eb5afae4eace434f493eb62672df7", // variableDebtOpt
		WETHList[chainId.OptimismChainName]: "0x0c84331e39d6658Cd6e6b9ba04736cC4c4734351", // variableDebtOpt
		WBTCList[chainId.OptimismChainName]: "0x92b42c66840c7ad907b4bf74879ff3ef7c529473", // variableDebtOpt
		SUSDList[chainId.OptimismChainName]: "0x4a1c3ad6ed28a636ee1751c69071f6be75deb8b8", // variableDebtOpt
		LINKList[chainId.OptimismChainName]: "0x953a573793604af8d41f306feb8274190db4ae0e", // variableDebtOpt
		AAVEList[chainId.OptimismChainName]: "0xe80761ea617f66f96274ea5e8c37f03960ecc679", // variableDebtOpt
	},
	chainId.PolygonChainName: {
		AgEURList[chainId.PolygonChainName]:   "0x3ca5fa07689f266e907439afd1fbb59c44fe12f6", // variableDebtPolAGEUR
		DAIList[chainId.PolygonChainName]:     "0x8619d80fb0141ba7f184cbf22fd724116d9f7ffc", //
		EURSList[chainId.PolygonChainName]:    "0x5d557b07776d12967914379c71a1310e917c7555",
		JEURList[chainId.PolygonChainName]:    "0x44705f578135cc5d703b4c9c122528c73eb87145",
		MiMATICList[chainId.PolygonChainName]: "0x18248226c16bf76c032817854e7c83a2113b4f06",
		USDCList[chainId.PolygonChainName]:    "0xfccf3cabbe80101232d343252614b6a3ee81c989",
		USDTList[chainId.PolygonChainName]:    "0xfb00ac187a8eb5afae4eace434f493eb62672df7",
		AAVEList[chainId.PolygonChainName]:    "0xe80761ea617f66f96274ea5e8c37f03960ecc679",
		BALList[chainId.PolygonChainName]:     "0xa8669021776bc142dfca87c21b4a52595bcbb40a",
		CRVList[chainId.PolygonChainName]:     "0x77ca01483f379e58174739308945f044e1a764dc",
		DPIList[chainId.PolygonChainName]:     "0xf611aeb5013fd2c0511c9cd55c7dc5c1140741a6",
		GHSTList[chainId.PolygonChainName]:    "0xce186f6cccb0c955445bb9d10c59cae488fea559",
		LINKList[chainId.PolygonChainName]:    "0x953a573793604af8d41f306feb8274190db4ae0e",
		MaticXList[chainId.PolygonChainName]:  "0xb5b46f918c2923fc7f26db76e8a6a6e9c4347cf9",
		STMATICList[chainId.PolygonChainName]: "0x6b030ff3fb9956b1b69f475b77ae0d3cf2cc5afa",
		SUSHIList[chainId.PolygonChainName]:   "0x34e2ed44ef7466d5f9e0b782b5c08b57475e7907",
		WBTCList[chainId.PolygonChainName]:    "0x92b42c66840c7ad907b4bf74879ff3ef7c529473",
		WETHList[chainId.PolygonChainName]:    "0x0c84331e39d6658cd6e6b9ba04736cc4c4734351",
		WMATICList[chainId.PolygonChainName]:  "0x4a1c3ad6ed28a636ee1751c69071f6be75deb8b8",
	},
	chainId.ArbitrumChainName: {
		DAIList[chainId.ArbitrumChainName]:  "0x8619d80fb0141ba7f184cbf22fd724116d9f7ffc",
		EURSList[chainId.ArbitrumChainName]: "0x4a1c3ad6ed28a636ee1751c69071f6be75deb8b8",
		USDCList[chainId.ArbitrumChainName]: "0xfccf3cabbe80101232d343252614b6a3ee81c989",
		USDTList[chainId.ArbitrumChainName]: "0xfb00ac187a8eb5afae4eace434f493eb62672df7",
		AAVEList[chainId.ArbitrumChainName]: "0xe80761ea617f66f96274ea5e8c37f03960ecc679",
		LINKList[chainId.ArbitrumChainName]: "0x953a573793604af8d41f306feb8274190db4ae0e",
		WBTCList[chainId.ArbitrumChainName]: "0x92b42c66840c7ad907b4bf74879ff3ef7c529473",
		WETHList[chainId.ArbitrumChainName]: "0x0c84331e39d6658cd6e6b9ba04736cc4c4734351",
	},
	chainId.BaseChainName: {
		WETHList[chainId.BaseChainName]:   "0x24e6e0795b3c7c71D965fCc4f371803d1c1DcA1E",
		CbETHList[chainId.BaseChainName]:  "0x1DabC36f19909425f654777249815c073E8Fd79F",
		USDbCList[chainId.BaseChainName]:  "0x7376b2F323dC56fCd4C191B34163ac8a84702DAB",
		WSTETHList[chainId.BaseChainName]: "0x41A7C3f5904ad176dACbb1D99101F59ef0811DC1",
		USDCList[chainId.BaseChainName]:   "0x59dca05b6c26dbd64b5381374aaac5cd05644c28",
	},
}

// Aave v3 isolated tokens.
//
// map[network]= []underlyings.
var AaveV3IsolatedTokens = map[string][]string{
	chainId.EthereumChainName: {},
	chainId.AvalancheChainName: {
		USDTList[chainId.AvalancheChainName],
		FRAXList[chainId.AvalancheChainName],
		MAIList[chainId.AvalancheChainName],
	},
	chainId.PolygonChainName: {
		USDTList[chainId.PolygonChainName],
		EURSList[chainId.PolygonChainName],
		MiMATICList[chainId.PolygonChainName],
	},
}

// Aave v3 tokens that are not allowed to be set as collateral.
//
// map[network]= []underlyings.
var AaveV3NotAllowCollateralTokens = map[string][]string{
	chainId.EthereumChainName:  {},
	chainId.AvalancheChainName: {},
	chainId.PolygonChainName: {
		JEURList[chainId.PolygonChainName],
		AgEURList[chainId.PolygonChainName],
	},
}
