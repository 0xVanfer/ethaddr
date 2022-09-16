package ethaddr

import "github.com/0xVanfer/chainId"

const LidoProtocol string = "lido"

// Lido token: LDO.
// map[network] = address.
var LidoTokenList = map[string]string{
	chainId.EthereumChainName: "0x5a98fcbea516cf06857215779fd812ca3bef1b32", // LDO
}

// Lido stake eth token: stETH.
// map[network] = address.
var LidoSTETHList = map[string]string{
	chainId.EthereumChainName: "0xae7ab96520de3a18e5e111b5eaab095312d7fe84", // stETH
}
