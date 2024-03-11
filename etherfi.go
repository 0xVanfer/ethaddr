package ethaddr

import "github.com/0xVanfer/chainId"

// Doc: https://etherfi.gitbook.io/etherfi
//
// Deployed contracts: https://etherfi.gitbook.io/etherfi/contracts-and-integrations/deployed-contracts
const EtherfiProtocol string = "etherfi"

// Ether.fi ETH.
//
// map[network] = address.
var EtherfiETHList = map[string]string{
	chainId.EthereumChainName: "0x35fA164735182de50811E8e2E824cFb9B6118ac2", // eETH
}

// Wrapped Ether.fi ETH.
//
// map[network] = address.
var WrappedEtherfiETHList = map[string]string{
	chainId.EthereumChainName: "0xCd5fE23C85820F7B72D0926FC9b05b43E359b7ee", // weETH
}

// eETH
//
// map[network] = address.
var EETHList = EtherfiETHList

// weETH
//
// map[network] = address.
var WEETHList = WrappedEtherfiETHList
