package ethaddr

import "github.com/0xVanfer/chainId"

const CompoundProtocol string = "compound"

// Compound token: COMP.
var CompoundTokenList = map[string]string{
	chainId.EthereumChainName: "0xc00e94cb662c3520282e6f5717214004a7f26888", // COMP
}

// Comptroller.
var CompoundComptrollerList = map[string]string{
	chainId.EthereumChainName: "0x3d9819210a31b4961b30ef54be2aed79b9c9cd3b",
}

// Compound c tokens.
var CompoundCTokenList = map[string]map[string]string{
	chainId.EthereumChainName: {
		BATList[chainId.EthereumChainName]:           "0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e", // cBAT
		DAIList[chainId.EthereumChainName]:           "0x5d3a536e4d6dbd6114cc1ead35777bab948e3643", // cDAI
		WETHList[chainId.EthereumChainName]:          "0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5", // cETH (no underlying function)
		REPList[chainId.EthereumChainName]:           "0x158079ee67fce2f58472a96584a73c7ab9ac95c1", // cREP
		USDCList[chainId.EthereumChainName]:          "0x39aa39c021dfbae8fac545936693ac917d5e7563", // cUSDC
		USDTList[chainId.EthereumChainName]:          "0xf650c3d88d12db855b8bf7d11be6c55a4e07dcc9", // cUSDT
		ZRXList[chainId.EthereumChainName]:           "0xb3319f5d18bc0d84dd1b4825dcde5d5f7266d407", // cZRX
		SAIList[chainId.EthereumChainName]:           "0xf5dce57282a584d2746faf1593d3121fcac444dc", // cDAI (should be cSAI)
		UniswapTokenList[chainId.EthereumChainName]:  "0x35a18000230da775cac24873d00ff85bccded550", // cUNI
		CompoundTokenList[chainId.EthereumChainName]: "0x70e36f6bf80a52b3b46b3af8e106cc0ed743e8e4", // cCOMP
		WBTCList[chainId.EthereumChainName]:          "0xccf4429db6322d5c611ee964527d42e5d685dd6a", // cWBTC
		TUSDList[chainId.EthereumChainName]:          "0x12392f67bdf24fae0af363c24ac620a2f67dad86", // cTUSD
		LINKList[chainId.EthereumChainName]:          "0xface851a4921ce59e912d19329929ce6da6eb0c7", // cLINK
		MakerTokenList[chainId.EthereumChainName]:    "0x95b4ef2869ebd94beb4eee400a99824bf5dc325b", // cMKR
		SushiTokenList[chainId.EthereumChainName]:    "0x4b0181102a0112a2ef11abee5563bb4a3176c9d7", // cSUSHI
		AaveTokenList[chainId.EthereumChainName]:     "0xe65cdb6479bac1e22340e4e755fae7e509ecd06c", // cAAVE
		YearnTokenList[chainId.EthereumChainName]:    "0x80a2ae356fc9ef4305676f7a3e2ed04e12c33946", // cYFI
		USDPList[chainId.EthereumChainName]:          "0x041171993284df560249b57358f931d9eb7b925d", // cUSDP
		FEIList[chainId.EthereumChainName]:           "0x7713dd9ca933848f6819f38b8352d9a15ea73f67", // cFEI
		// WBTCList[chainId.EthereumChainName]:          "0xc11b1268c1a384e55c48c2391d8d480264a3a7f4", // cWBTC
	},
}
