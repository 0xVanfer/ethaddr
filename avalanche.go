package ethaddr

import "github.com/0xVanfer/chainId"

// USDC.e for avalanche
//
// map[network] = address.
var USDCeList = map[string]string{
	chainId.AvalancheChainName: "0xA7D7079b0FEaD91F3e65f86E8915Cb59c1a4C664", // USDC.e, 0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664
}

// USDT.e for avalanche
//
// map[network] = address.
var USDTeList = map[string]string{
	chainId.AvalancheChainName: "0xc7198437980c041c805A1EDcbA50c1Ce5db95118", // USDT.e, 0xc7198437980c041c805a1edcba50c1ce5db95118
}

// DAI.e
//
// map[network] = address.
var DAIeList = map[string]string{
	chainId.AvalancheChainName: "0xd586E7F844cEa2F87f50152665BCbc2C279D8d70", // DAI.e, 0xd586e7f844cea2f87f50152665bcbc2c279d8d70
}

// Wrapped avalanche chain token: WAVAX.
//
// map[network] = address.
var WAVAXList = map[string]string{
	chainId.EthereumChainName:  "0x85f138bfEE4ef8e540890CFb48F620571d67Eda3", // WAVAX(wormwhole), 0x85f138bfee4ef8e540890cfb48f620571d67eda3
	chainId.AvalancheChainName: "0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7", // WAVAX, 0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7
	chainId.PolygonChainName:   "0x7Bb11E7f8b10E9e571E5d8Eace04735fDFB2358a", // WAVAX(wormwhole), 0x7bb11e7f8b10e9e571e5d8eace04735fdfb2358a
}

// AVAX on Polygon.
//
// map[network] = address.
var AVAXList = map[string]string{
	chainId.PolygonChainName: "0x2C89bbc92BD86F8075d1DEcc58C7F4E0107f286b", // AVAX, 0x2c89bbc92bd86f8075d1decc58c7f4e0107f286b
}

// XAVA
//
// map[network] = address.
var XAVAList = map[string]string{
	chainId.AvalancheChainName: "0xd1c3f94DE7e5B45fa4eDBBA472491a9f4B166FC4", // XAVA, 0xd1c3f94de7e5b45fa4edbba472491a9f4b166fc4
}

// YUSD
//
// map[network] = address.
var YUSDList = map[string]string{
	chainId.AvalancheChainName: "0x111111111111ed1D73f860F57b2798b683f2d325", // YUSD, 0x111111111111ed1d73f860f57b2798b683f2d325
}
