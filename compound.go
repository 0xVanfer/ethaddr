package ethaddr

// Website: https://compound.finance/
//
// Docs: https://docs.compound.finance/
//
// Deployed contracts: https://docs.compound.finance/#networks
const CompoundProtocol string = "compound"

// Compound token: COMP.
//
// map[network] = address.
var CompoundTokenList = map[string]string{
	ChainEthereum: "0xc00e94Cb662C3520282E6f5717214004A7f26888", // COMP, 0xc00e94cb662c3520282e6f5717214004a7f26888
}

// Same as CompoundTokenList.
var COMPList = CompoundTokenList

// Comptroller.
//
// map[network] = address.
var CompoundComptrollerList = map[string]string{
	ChainEthereum: "0x3d9819210A31b4961b30EF54bE2aeD79B9c9Cd3B", // 0x3d9819210a31b4961b30ef54be2aed79b9c9cd3b
}

// Compound c tokens.
//
// map[network][underlying] = address.
var CompoundCTokenList = map[string]map[string]string{
	ChainEthereum: {
		BATList[ChainEthereum]:   "0x6C8c6b02E7b2BE14d4fA6022Dfd6d75921D90E4E", // cBAT, 0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e
		DAIList[ChainEthereum]:   "0x5d3a536E4D6DbD6114cc1Ead35777bAB948E3643", // cDAI, 0x5d3a536e4d6dbd6114cc1ead35777bab948e3643
		WETHList[ChainEthereum]:  "0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5", // cETH, 0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5, no underlying function
		REPList[ChainEthereum]:   "0x158079Ee67Fce2f58472A96584A73C7Ab9AC95c1", // cREP, 0x158079ee67fce2f58472a96584a73c7ab9ac95c1
		USDCList[ChainEthereum]:  "0x39AA39c021dfbaE8faC545936693aC917d5E7563", // cUSDC, 0x39aa39c021dfbae8fac545936693ac917d5e7563
		USDTList[ChainEthereum]:  "0xf650C3d88D12dB855b8bf7D11Be6C55A4e07dCC9", // cUSDT, 0xf650c3d88d12db855b8bf7d11be6c55a4e07dcc9
		ZRXList[ChainEthereum]:   "0xB3319f5D18Bc0D84dD1b4825Dcde5d5f7266d407", // cZRX, 0xb3319f5d18bc0d84dd1b4825dcde5d5f7266d407
		SAIList[ChainEthereum]:   "0xF5DCe57282A584D2746FaF1593d3121Fcac444dC", // cDAI,  0xf5dce57282a584d2746faf1593d3121fcac444dc
		UNIList[ChainEthereum]:   "0x35A18000230DA775CAc24873d00Ff85BccdeD550", // cUNI, 0x35a18000230da775cac24873d00ff85bccded550
		COMPList[ChainEthereum]:  "0x70e36f6BF80a52b3B46b3aF8e106CC0ed743E8e4", // cCOMP, 0x70e36f6bf80a52b3b46b3af8e106cc0ed743e8e4
		WBTCList[ChainEthereum]:  "0xccF4429DB6322D5C611ee964527D42E5d685DD6a", // cWBTC, 0xccf4429db6322d5c611ee964527d42e5d685dd6a
		TUSDList[ChainEthereum]:  "0x12392F67bdf24faE0AF363c24aC620a2f67DAd86", // cTUSD, 0x12392f67bdf24fae0af363c24ac620a2f67dad86
		LINKList[ChainEthereum]:  "0xFAce851a4921ce59e912d19329929CE6da6EB0c7", // cLINK, 0xface851a4921ce59e912d19329929ce6da6eb0c7
		MKRList[ChainEthereum]:   "0x95b4eF2869eBD94BEb4eEE400a99824BF5DC325b", // cMKR, 0x95b4ef2869ebd94beb4eee400a99824bf5dc325b
		SUSHIList[ChainEthereum]: "0x4B0181102A0112A2ef11AbEE5563bb4a3176c9d7", // cSUSHI, 0x4b0181102a0112a2ef11abee5563bb4a3176c9d7
		AAVEList[ChainEthereum]:  "0xe65cdB6479BaC1e22340E4E755fAE7E509EcD06c", // cAAVE, 0xe65cdb6479bac1e22340e4e755fae7e509ecd06c
		YFIList[ChainEthereum]:   "0x80a2AE356fc9ef4305676f7a3E2Ed04e12C33946", // cYFI, 0x80a2ae356fc9ef4305676f7a3e2ed04e12c33946
		USDPList[ChainEthereum]:  "0x041171993284df560249B57358F931D9eB7b925D", // cUSDP, 0x041171993284df560249b57358f931d9eb7b925d
		FEIList[ChainEthereum]:   "0x7713DD9Ca933848F6819F38B8352D9A15EA73F67", // cFEI, 0x7713dd9ca933848f6819f38b8352d9a15ea73f67
		// WBTCList[ChainEthereum]:          "0xC11b1268C1A384e55C48c2391d8d480264A3A7F4", // cWBTC, 0xc11b1268c1a384e55c48c2391d8d480264a3a7f4
	},
}
