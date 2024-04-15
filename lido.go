package ethaddr

// Website: https://lido.fi/
//
// Docs: https://docs.lido.fi/
//
// Deployed contracts: https://docs.lido.fi/deployed-contracts/
const LidoProtocol string = "lido"

// Lido token: LDO.
//
// map[network] = address.
var LidoTokenList = map[string]string{
	ChainArbitrum: "0x13Ad51ed4F1B7e9Dc168d8a00cB3f4dDD85EfA60", // LDO, 0x13ad51ed4f1b7e9dc168d8a00cb3f4ddd85efa60
	ChainEthereum: "0x5A98FcBEA516Cf06857215779Fd812CA3beF1B32", // LDO, 0x5a98fcbea516cf06857215779fd812ca3bef1b32
}

// Same as LidoTokenList.
var LDOList = LidoTokenList

// Lido stake eth token: stETH.
//
// map[network] = address.
var STETHList = map[string]string{
	ChainEthereum: "0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84", // stETH, 0xae7ab96520de3a18e5e111b5eaab095312d7fe84
}

// Lido wrapped stake eth token: wstETH.
//
// map[network] = address.
var WSTETHList = map[string]string{
	ChainArbitrum: "0x5979D7b546E38E414F7E9822514be443A4800529", // wstETH, 0x5979d7b546e38e414f7e9822514be443a4800529
	ChainBase:     "0xc1CBa3fCea344f92D9239c08C0568f6F2F0ee452", // wstETH, 0xc1cba3fcea344f92d9239c08c0568f6f2f0ee452
	ChainEthereum: "0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0", // wstETH, 0x7f39c581f595b53c5cb19bd0b3f8da6c935e2ca0
	ChainOptimism: "0x1F32b1c2345538c0c6f582fCB022739c4A194Ebb", // wstETH, 0x1f32b1c2345538c0c6f582fcb022739c4a194ebb
	ChainPolygon:  "0x03b54A6e9a984069379fae1a4fC4dBAE93B3bCCD", // wstETH, 0x03b54a6e9a984069379fae1a4fc4dbae93b3bccd
	ChainScroll:   "0xf610A9dfB7C89644979b4A0f27063E9e7d7Cda32", // wstETH, 0xf610a9dfb7c89644979b4a0f27063e9e7d7cda32
}

// Lido stake matic token: stMATIC.
//
// map[network] = address.
//
// The ethereum is the `ROOT` chain and polygon is the `CHILD` chain.
// The operations and calculation are done on root chain, and child chain is only a operation contract.
var STMATICList = map[string]string{
	ChainEthereum: "0x9ee91F9f426fA633d227f7a9b000E28b9dfd8599", // stMATIC, 0x9ee91f9f426fa633d227f7a9b000e28b9dfd8599
	ChainPolygon:  "0x3A58a54C066FdC0f2D55FC9C89F0415C92eBf3C4", // stMATIC, 0x3a58a54c066fdc0f2d55fc9c89f0415c92ebf3c4
}
