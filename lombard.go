package ethaddr

// Website: https://www.lombard.finance/
//
// Deployed addresses: https://docs.lombard.finance/technical-documentation/smart-contracts
const LombardProtocol string = "lombard"

// Lombard LBTC.
//
// map[chainID] = address
var LBTCList = map[int64]string{
	ChainBase:     "0xecAc9C5F704e954931349Da37F60E39f515c11c1", // LBTC, 0xecac9c5f704e954931349da37f60e39f515c11c1
	ChainBera:     "0xecAc9C5F704e954931349Da37F60E39f515c11c1", // LBTC, 0xecac9c5f704e954931349da37f60e39f515c11c1
	ChainBSC:      "0xecAc9C5F704e954931349Da37F60E39f515c11c1", // LBTC, 0xecac9c5f704e954931349da37f60e39f515c11c1
	ChainEthereum: "0x8236a87084f8B84306f72007F36F2618A5634494", // LBTC, 0x8236a87084f8b84306f72007f36f2618a5634494
	ChainSonic:    "0xecAc9C5F704e954931349Da37F60E39f515c11c1", // LBTC, 0xecac9c5f704e954931349da37f60e39f515c11c1
}

// Lombard LBTC vault.
//
// map[chainID] = address
var LBTCvList = map[int64]string{
	ChainBase:     "0x5401b8620e5fb570064ca9114fd1e135fd77d57c", // LBTCv, 0x5401b8620e5fb570064ca9114fd1e135fd77d57c
	ChainBSC:      "0x5401b8620e5fb570064ca9114fd1e135fd77d57c", // LBTCv, 0x5401b8620e5fb570064ca9114fd1e135fd77d57c
	ChainEthereum: "0x5401b8620E5FB570064CA9114fd1e135fd77D57c", // LBTCv, 0x5401b8620e5fb570064ca9114fd1e135fd77d57c
}
