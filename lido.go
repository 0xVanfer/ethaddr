package ethaddr

import "github.com/0xVanfer/chainId"

const LidoProtocol string = "lido"

// Lido token: LDO.
//
// map[network] = address.
var LidoTokenList = map[string]string{
	chainId.EthereumChainName: "0x5a98fcbea516cf06857215779fd812ca3bef1b32", // LDO
}

// Same as LidoTokenList.
var LDOList = LidoTokenList

// Lido stake eth token: stETH.
//
// map[network] = address.
var STETHList = map[string]string{
	chainId.EthereumChainName: "0xae7ab96520de3a18e5e111b5eaab095312d7fe84", // stETH
}

// Lido wrapped stake eth token: wstETH.
//
// map[network] = address.
var WSTETHList = map[string]string{
	chainId.EthereumChainName: "0x7f39c581f595b53c5cb19bd0b3f8da6c935e2ca0", // wstETH
}

// Lido stake matic token: stMATIC.
//
// map[network] = address.
var STMATICList = map[string]string{
	chainId.PolygonChainName: "0x3a58a54c066fdc0f2d55fc9c89f0415c92ebf3c4", // stMATIC
}
