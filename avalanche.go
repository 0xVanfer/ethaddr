package ethaddr

// Blockscan for avlanche: https://snowscan.xyz/; previously https://snowtrace.io/.
const ChainNameAvalanche string = "avalanche"

// Wrapped avalanche chain token: WAVAX.
//
// map[chainID] = address.
var WAVAXList = map[int64]string{
	ChainEthereum:  "0x85f138bfEE4ef8e540890CFb48F620571d67Eda3", // WAVAX(wormwhole), 0x85f138bfee4ef8e540890cfb48f620571d67eda3
	ChainAvalanche: "0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7", // WAVAX, 0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7
	ChainPolygon:   "0x7Bb11E7f8b10E9e571E5d8Eace04735fDFB2358a", // WAVAX(wormwhole), 0x7bb11e7f8b10e9e571e5d8eace04735fdfb2358a
}

// AVAX on Polygon, optimism...
//
// map[chainID] = address.
var AVAXList = map[int64]string{
	ChainPolygon: "0x2C89bbc92BD86F8075d1DEcc58C7F4E0107f286b", // AVAX, 0x2c89bbc92bd86f8075d1decc58c7f4e0107f286b
}

// XAVA
//
// map[chainID] = address.
var XAVAList = map[int64]string{
	ChainAvalanche: "0xd1c3f94DE7e5B45fa4eDBBA472491a9f4B166FC4", // XAVA, 0xd1c3f94de7e5b45fa4edbba472491a9f4b166fc4
}

// YUSD
//
// map[chainID] = address.
var YUSDList = map[int64]string{
	ChainAvalanche: "0x111111111111ed1D73f860F57b2798b683f2d325", // YUSD, 0x111111111111ed1d73f860f57b2798b683f2d325
}
