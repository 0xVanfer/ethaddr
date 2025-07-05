package ethaddr

// Docs: https://docs.sushi.com/docs/intro
//
// Deployed contracts: https://docs.sushi.com/docs/Products/Classic%20AMM/Deployment%20Addresses
const SushiProtocol string = "sushi"

// Sushi token: SUSHI.
//
// map[chainID] = address.
var SushiTokenList = map[int64]string{
	ChainArbitrum:  "0xd4d42F0b6DEF4CE0383636770eF773390d85c61A", // SUSHI, 0xd4d42f0b6def4ce0383636770ef773390d85c61a
	ChainAvalanche: "0x37B608519F91f70F2EeB0e5Ed9AF4061722e4F76", // SUSHI.e, 0x37b608519f91f70f2eeb0e5ed9af4061722e4f76
	ChainEthereum:  "0x6B3595068778DD592e39A122f4f5a5cF09C90fE2", // SUSHI, 0x6b3595068778dd592e39a122f4f5a5cf09c90fe2
	ChainPolygon:   "0x0b3F868E0BE5597D5DB7fEB59E1CADBb0fdDa50a", // SUSHI, 0x0b3f868e0be5597d5db7feb59e1cadbb0fdda50a
}

// Same as SushiTokenList.
var SUSHIList = SushiTokenList

// XSUSHI.
//
// map[chainID] = address.
var XSushiList = map[int64]string{
	ChainEthereum: "0x8798249c2E607446EfB7Ad49eC89dD1865Ff4272", // XSUSHI, 0x8798249c2e607446efb7ad49ec89dd1865ff4272
}

// v2 Router02
//
// https://docs.sushi.com/docs/Products/Classic%20AMM/Deployment%20Addresses
//
// map[chainID] = address
var SushiV2RouterList = map[int64]string{
	ChainArbitrum:  "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", // 0x1b02da8cb0d097eb8d57a175b88c7d8b47997506
	ChainAvalanche: "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", // 0x1b02da8cb0d097eb8d57a175b88c7d8b47997506
	ChainBSC:       "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", // 0x1b02da8cb0d097eb8d57a175b88c7d8b47997506
	ChainEthereum:  "0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F", // 0xd9e1ce17f2641f24ae83637ab66a2cca9c378b9f
	ChainFantom:    "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", // 0x1b02da8cb0d097eb8d57a175b88c7d8b47997506
	ChainGoerli:    "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", // 0x1b02da8cb0d097eb8d57a175b88c7d8b47997506
	ChainOK:        "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", // 0x1b02da8cb0d097eb8d57a175b88c7d8b47997506
	ChainPolygon:   "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", // 0x1b02da8cb0d097eb8d57a175b88c7d8b47997506
}

// v2 factory
//
// https://docs.sushi.com/docs/Products/Classic%20AMM/Deployment%20Addresses
//
// map[chainID] = address
var SushiV2FactoryList = map[int64]string{
	ChainArbitrum:  "0xc35DADB65012eC5796536bD9864eD8773aBc74C4", // 0xc35dadb65012ec5796536bd9864ed8773abc74c4
	ChainAvalanche: "0xc35DADB65012eC5796536bD9864eD8773aBc74C4", // 0xc35dadb65012ec5796536bd9864ed8773abc74c4
	ChainBSC:       "0xc35DADB65012eC5796536bD9864eD8773aBc74C4", // 0xc35dadb65012ec5796536bd9864ed8773abc74c4
	ChainEthereum:  "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac", // 0xc0aee478e3658e2610c5f7a4a2e1777ce9e4f2ac
	ChainFantom:    "0xc35DADB65012eC5796536bD9864eD8773aBc74C4", // 0xc35dadb65012ec5796536bd9864ed8773abc74c4
	ChainGoerli:    "0xc35DADB65012eC5796536bD9864eD8773aBc74C4", // 0xc35dadb65012ec5796536bd9864ed8773abc74c4
	ChainOK:        "0xc35DADB65012eC5796536bD9864eD8773aBc74C4", // 0xc35dadb65012ec5796536bd9864ed8773abc74c4
	ChainPolygon:   "0xc35DADB65012eC5796536bD9864eD8773aBc74C4", // 0xc35dadb65012ec5796536bd9864ed8773abc74c4
}

// v3 factory
//
// https://docs.sushi.com/docs/Products/V3%20AMM/Core/Deployment%20Addresses
//
// map[chainID] = address
var SushiV3FactoryList = map[int64]string{
	ChainArbitrum:  "0x1af415a1EbA07a4986a52B6f2e7dE7003D82231e", // 0x1af415a1eba07a4986a52b6f2e7de7003d82231e
	ChainAvalanche: "0x3e603C14aF37EBdaD31709C4f848Fc6aD5BEc715", // 0x3e603c14af37ebdad31709c4f848fc6ad5bec715
	ChainBSC:       "0x126555dd55a39328F69400d6aE4F782Bd4C34ABb", // 0x126555dd55a39328f69400d6ae4f782bd4c34abb
	ChainEthereum:  "0xbACEB8eC6b9355Dfc0269C18bac9d6E2Bdc29C4F", // 0xbaceb8ec6b9355dfc0269c18bac9d6e2bdc29c4f
	ChainFantom:    "0x7770978eED668a3ba661d51a773d3a992Fc9DDCB", // 0x7770978eed668a3ba661d51a773d3a992fc9ddcb
	ChainFilecoin:  "0xc35DADB65012eC5796536bD9864eD8773aBc74C4", // 0xc35dadb65012ec5796536bd9864ed8773abc74c4
	ChainOptimism:  "0x9c6522117e2ed1fE5bdb72bb0eD5E3f2bdE7DBe0", // 0x9c6522117e2ed1fe5bdb72bb0ed5e3f2bde7dbe0
	ChainPolygon:   "0x917933899c6a5F8E37F31E19f92CdBFF7e8FF0e2", // 0x917933899c6a5f8e37f31e19f92cdbff7e8ff0e2
	ChainPolygonZk: "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", // 0x1b02da8cb0d097eb8d57a175b88c7d8b47997506
	ChainScroll:    "0x46B3fDF7b5CDe91Ac049936bF0bDb12c5d22202e", // 0x46b3fdf7b5cde91ac049936bf0bdb12c5d22202e
}
