package ethaddr

// Website: https://lido.fi/
//
// Docs: https://docs.lido.fi/
//
// Deployed contracts: https://docs.lido.fi/deployed-contracts/
const LidoProtocol string = "lido"

// Lido token: LDO.
//
// map[chainID] = address.
var LidoTokenList = map[int64]string{
	ChainArbitrum: "0x13Ad51ed4F1B7e9Dc168d8a00cB3f4dDD85EfA60", // LDO, 0x13ad51ed4f1b7e9dc168d8a00cb3f4ddd85efa60
	ChainEthereum: "0x5A98FcBEA516Cf06857215779Fd812CA3beF1B32", // LDO, 0x5a98fcbea516cf06857215779fd812ca3bef1b32
	ChainOptimism: "0xFdb794692724153d1488CcdBE0C56c252596735F", // LDO, 0xfdb794692724153d1488ccdbe0c56c252596735f
}

// Same as LidoTokenList.
var LDOList = LidoTokenList

// Lido stake eth token: stETH.
//
// map[chainID] = address.
var STETHList = map[int64]string{
	ChainEthereum: "0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84", // stETH, 0xae7ab96520de3a18e5e111b5eaab095312d7fe84
	ChainOptimism: "0x76A50b8c7349cCDDb7578c6627e79b5d99D24138", // stETH, 0x76a50b8c7349ccddb7578c6627e79b5d99d24138
}

// Lido wrapped stake eth token: wstETH.
//
// map[chainID] = address.
var WSTETHList = map[int64]string{
	ChainArbitrum: "0x5979D7b546E38E414F7E9822514be443A4800529", // wstETH, 0x5979d7b546e38e414f7e9822514be443a4800529
	ChainBase:     "0xc1CBa3fCea344f92D9239c08C0568f6F2F0ee452", // wstETH, 0xc1cba3fcea344f92d9239c08c0568f6f2f0ee452
	ChainBSC:      "0x26c5e01524d2E6280A48F2c50fF6De7e52E9611C", // wstETH, 0x26c5e01524d2e6280a48f2c50ff6de7e52e9611c
	ChainEthereum: "0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0", // wstETH, 0x7f39c581f595b53c5cb19bd0b3f8da6c935e2ca0
	ChainLinea:    "0xB5beDd42000b71FddE22D3eE8a79Bd49A568fC8F", // wstETH, 0xb5bedd42000b71fdde22d3ee8a79bd49a568fc8f
	ChainMantle:   "0x458ed78EB972a369799fb278c0243b25e5242A83", // wstETH, 0x458ed78eb972a369799fb278c0243b25e5242a83
	ChainMode:     "0x98f96A4B34D03a2E6f225B28b8f8Cb1279562d81", // wstETH, 0x98f96a4b34d03a2e6f225b28b8f8cb1279562d81
	ChainZircuit:  "0xf0e673Bc224A8Ca3ff67a61605814666b1234833", // wstETH, 0xf0e673bc224a8ca3ff67a61605814666b1234833
	ChainOptimism: "0x1F32b1c2345538c0c6f582fCB022739c4A194Ebb", // wstETH, 0x1f32b1c2345538c0c6f582fcb022739c4a194ebb
	ChainPolygon:  "0x03b54A6e9a984069379fae1a4fC4dBAE93B3bCCD", // wstETH, 0x03b54a6e9a984069379fae1a4fc4dbae93b3bccd
	ChainScroll:   "0xf610A9dfB7C89644979b4A0f27063E9e7d7Cda32", // wstETH, 0xf610a9dfb7c89644979b4a0f27063e9e7d7cda32
}

// Lido stake matic token: stMATIC.
//
// Has been closed in 2025.
//
// map[chainID] = address.
//
// The ethereum is the `ROOT` chain and polygon is the `CHILD` chain.
// The operations and calculation are done on root chain, and child chain is only a operation contract.
var STMATICList = map[int64]string{
	ChainEthereum: "0x9ee91F9f426fA633d227f7a9b000E28b9dfd8599", // stMATIC, 0x9ee91f9f426fa633d227f7a9b000e28b9dfd8599
	ChainPolygon:  "0x3A58a54C066FdC0f2D55FC9C89F0415C92eBf3C4", // stMATIC, 0x3a58a54c066fdc0f2d55fc9c89f0415c92ebf3c4
}
