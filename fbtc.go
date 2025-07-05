package ethaddr

// Website: https://fbtc.com/get-fbtc
//
// Deployed contracts: https://docs.fbtc.com/developers/smart-contracts
const FBTCProtocol string = "fbtc"

// FBTC, aka FBTC0
//
// map[chain id] = address
var FBTCList = map[int64]string{
	ChainArbitrum: "0xC96dE26018A54D51c097160568752c4E3BD6C364", // FBTC, 0xc96de26018a54d51c097160568752c4e3bd6c364
	ChainBSC:      "0xC96dE26018A54D51c097160568752c4E3BD6C364", // FBTC, 0xc96de26018a54d51c097160568752c4e3bd6c364
	ChainEthereum: "0xC96dE26018A54D51c097160568752c4E3BD6C364", // FBTC, 0xc96de26018a54d51c097160568752c4e3bd6c364
	ChainMantle:   "0xC96dE26018A54D51c097160568752c4E3BD6C364", // FBTC, 0xc96de26018a54d51c097160568752c4e3bd6c364
	ChainSei:      "0x5fabd1D440a90eE57Dd698ea096B91F994B6DF56", // FBTC, 0x5fabd1d440a90ee57dd698ea096b91f994b6df56
}

// WFBTC, wrapped FBTC
//
// map[chain id] = address
var WFBTCList = map[int64]string{
	ChainBera:     "0xA18717e22dfAbB36f0242259df4e9E563b6Ec70a", // WFBTC, 0xa18717e22dfabb36f0242259df4e9e563b6ec70a
	ChainEthereum: "0x7aAd7A95fCf14B826AC96176590C8e7aad19bbd4", // WFBTC, 0x7aad7a95fcf14b826ac96176590c8e7aad19bbd4
	ChainSei:      "0x5fabd1D440a90eE57Dd698ea096B91F994B6DF56", // FBTC, 0x5fabd1d440a90ee57dd698ea096b91f994b6df56
}
