package ethaddr

// Website: https://www.ether.fi/
//
// Doc: https://etherfi.gitbook.io/etherfi
//
// Deployed contracts: https://etherfi.gitbook.io/etherfi/contracts-and-integrations/deployed-contracts
const EtherfiProtocol string = "etherfi"

// ether.fi governance token: ETHFI
//
// map[chainID] = address.
var EtherFiGovernanceTokenList = map[int64]string{
	ChainEthereum: "0xFe0c30065B384F05761f15d0CC899D4F9F9Cc0eB", // ETHFI, 0xfe0c30065b384f05761f15d0cc899d4f9f9cc0eb
	ChainArbitrum: "0x7189fb5B6504bbfF6a852B13B7B82a3c118fDc27", // ETHFI, 0x7189fb5b6504bbff6a852b13b7b82a3c118fdc27
}

// Same as EtherFiGovernanceTokenList.
var ETHFIList = EtherFiGovernanceTokenList

// Ether.fi ETH. eETH
//
// map[chainID] = address.
var EETHList = map[int64]string{
	ChainEthereum: "0x35fA164735182de50811E8e2E824cFb9B6118ac2", // eETH, 0x35fa164735182de50811e8e2e824cfb9b6118ac2
}

// Wrapped Ether.fi ETH. weETH
//
// map[chainID] = address.
var WEETHList = map[int64]string{
	ChainArbitrum: "0x35751007a407ca6FEFfE80b3cB397736D2cf4dbe", // weETH, 0x35751007a407ca6feffe80b3cb397736d2cf4dbe
	ChainBase:     "0x04C0599Ae5A44757c0af6F9eC3b93da8976c150A", // weETH.base, 0x04c0599ae5a44757c0af6f9ec3b93da8976c150a
	ChainBlast:    "0x04C0599Ae5A44757c0af6F9eC3b93da8976c150A", // weETH, 0x04c0599ae5a44757c0af6f9ec3b93da8976c150a
	ChainBSC:      "0x04C0599Ae5A44757c0af6F9eC3b93da8976c150A", // weETH.bnb, 0x04c0599ae5a44757c0af6f9ec3b93da8976c150a
	ChainEthereum: "0xCd5fE23C85820F7B72D0926FC9b05b43E359b7ee", // weETH, 0xcd5fe23c85820f7b72d0926fc9b05b43e359b7ee
	ChainOptimism: "0x5A7fACB970D094B6C7FF1df0eA68D99E6e73CBFF", // weETH, 0x5a7facb970d094b6c7ff1df0ea68d99e6e73cbff
	ChainScroll:   "0x01f0a31698C4d065659b9bdC21B3610292a1c506", // weETH, 0x01f0a31698c4d065659b9bdc21b3610292a1c506
}

// Ether.fi BTC. eBTC
//
// map[chainID] = address.
var EBTCList = map[int64]string{
	ChainEthereum: "0x657e8C867D8B37dCC18fA4Caead9C45EB088C642", // eBTC, 0x657e8c867d8b37dcc18fa4caead9c45eb088c642
}
