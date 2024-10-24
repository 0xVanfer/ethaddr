package ethaddr

// About frxETH and sfrxETH: https://docs.frax.finance/frax-ether/frxeth-and-sfrxeth

const FraxProtocol string = "frax"

// Frax token: FRAX.
//
// map[network] = address.
var FRAXList = map[int64]string{
	ChainArbitrum:  "0x17FC002b466eEc40DaE837Fc4bE5c67993ddBd6F", // FRAX, 0x17fc002b466eec40dae837fc4be5c67993ddbd6f
	ChainAvalanche: "0xD24C2Ad096400B6FBcd2ad8B24E7acBc21A1da64", // FRAX, 0xd24c2ad096400b6fbcd2ad8b24e7acbc21a1da64
	ChainEthereum:  "0x853d955aCEf822Db058eb8505911ED77F175b99e", // FRAX, 0x853d955acef822db058eb8505911ed77f175b99e
	ChainOptimism:  "0x2E3D870790dC77A83DD1d18184Acc7439A53f475", // FRAX, 0x2e3d870790dc77a83dd1d18184acc7439a53f475
	ChainPolygon:   "0x104592a158490a9228070E0A8e5343B499e125D0", // FRAX, 0x104592a158490a9228070e0a8e5343b499e125d0
}

// Frax share token: FXS.
//
// map[network] = address.
var FXSList = map[int64]string{
	ChainEthereum: "0x3432B6A60D23Ca0dFCa7761B7ab56459D9C964D0", // FXS, 0x3432b6a60d23ca0dfca7761b7ab56459d9c964d0
}

// Frax Ether: frxETH.
// Similar to stETH.
//
// map[chainID] = address.
var FrxETHList = map[int64]string{
	ChainEthereum: "0x5E8422345238F34275888049021821E8E08CAa1f", // frxETH, 0x5e8422345238f34275888049021821e8e08caa1f
}

// Staked Frax Ether: sfrxETH.
// Similar to wstETH.
//
// map[chainID] = address.
var StakedFraxEtherList = map[int64]string{
	ChainEthereum: "0xac3E018457B222d93114458476f3E3416Abbe38F", // sfrxETH, 0xac3e018457b222d93114458476f3e3416abbe38f
}

// Same as StakedFraxEtherList.
var SfrxETHList = StakedFraxEtherList
