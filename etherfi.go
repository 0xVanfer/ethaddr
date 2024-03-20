package ethaddr

// Doc: https://etherfi.gitbook.io/etherfi
//
// Deployed contracts: https://etherfi.gitbook.io/etherfi/contracts-and-integrations/deployed-contracts
const EtherfiProtocol string = "etherfi"

// Ether.fi ETH. eETH
//
// map[network] = address.
var EETHList = map[string]string{
	chainEthereum: "0x35fA164735182de50811E8e2E824cFb9B6118ac2", // eETH, 0x35fa164735182de50811e8e2e824cfb9b6118ac2
}

// Wrapped Ether.fi ETH. weETH
//
// map[network] = address.
var WEETHList = map[string]string{
	chainEthereum: "0xCd5fE23C85820F7B72D0926FC9b05b43E359b7ee", // weETH, 0xcd5fe23c85820f7b72d0926fc9b05b43e359b7ee
}
