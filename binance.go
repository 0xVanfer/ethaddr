package ethaddr

// Blockscan: https://bscscan.com/
const ChainNameBSC string = "bsc"

// Website: https://binance.com
const BinanceProtocol string = "binance"

// Binance stable coin: BUSD.
//
// map[chainID] = address.
var BUSDList = map[int64]string{
	ChainAvalanche: "0x9C9e5fD8bbc25984B178FdCE6117Defa39d2db39", // BUSD, 0x9c9e5fd8bbc25984b178fdce6117defa39d2db39
	ChainBSC:       "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56", // BUSD, 0xe9e7cea3dedca5984780bafc599bd69add087d56
	ChainEthereum:  "0x4Fabb145d64652a948d72533023f6E7A623C7C53", // BUSD, 0x4fabb145d64652a948d72533023f6e7a623c7c53
	ChainPolygon:   "0x9C9e5fD8bbc25984B178FdCE6117Defa39d2db39", // BUSD, 0x9c9e5fd8bbc25984b178fdce6117defa39d2db39
}

// Wrapped binance smart chain token: WBNB.
//
// map[chainID] = address.
var WBNBList = map[int64]string{
	ChainBSC:      "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", // WBNB, 0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c
	ChainEthereum: "0x418D75f65a02b3D53B2418FB8E1fe493759c7605", // WBNB, 0x418d75f65a02b3d53b2418fb8e1fe493759c7605
}

// Binance Beacon ETH: BETH.
//
// map[chainID] = address.
var BETHList = map[int64]string{
	ChainBSC: "0x250632378E573c6Be1AC2f97Fcdf00515d0Aa91B", // BETH, 0x250632378e573c6be1ac2f97fcdf00515d0aa91b
}

// Wrapped Binance Beacon ETH: wBETH.
//
// map[chainID] = address.
var WBETHList = map[int64]string{
	ChainEthereum: "0xa2E3356610840701BDf5611a53974510Ae27E2e1", // wBETH, 0xa2e3356610840701bdf5611a53974510ae27e2e1
	ChainBSC:      "0xa2E3356610840701BDf5611a53974510Ae27E2e1", // wBETH, 0xa2e3356610840701bdf5611a53974510ae27e2e1
}

// BTCB, Binance withdrawed BTC on BSC.
var BTCBList = map[int64]string{
	ChainBSC: "0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c", // BTCB, 0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c
}

// BBTC, Binance Wrapped BTC only on Ethereum.
//
// map[chainID] = address.
var BBTCList = map[int64]string{
	ChainEthereum: "0x9BE89D2a4cd102D8Fecc6BF9dA793be995C22541", // BBTC, 0x9be89d2a4cd102d8fecc6bf9da793be995c22541
}

// Binance can cash out BBTC on Ethereum or BTCB on BSC.
//
// map[chainID] = address.
var BinanceBTCList = map[int64]string{
	ChainEthereum: "0x9BE89D2a4cd102D8Fecc6BF9dA793be995C22541", // BBTC, 0x9be89d2a4cd102d8fecc6bf9da793be995c22541
	ChainBSC:      "0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c", // BTCB, 0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c
}
