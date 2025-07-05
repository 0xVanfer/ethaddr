package ethaddr

// Website: https://www.avalonfinance.xyz/
const AvalonProtocol string = "avalon"

// Avalon Token.
//
// map[chainID] = address
var AVLList = map[int64]string{
	ChainEthereum: "0x5c8d0c48810fd37a0a824d074ee290e64f7a8fa2", // AVL, 0x5c8d0c48810fd37a0a824d074ee290e64f7a8fa2
	ChainTaiko:    "0xE9cA67e5051e1806546d0a06ee465221c5877feE", // AVL, 0xe9ca67e5051e1806546d0a06ee465221c5877fee
}

// Avalon staked USDa: sUSDa.
//
// map[chainID] = address
var SUSDaList = map[int64]string{
	ChainEthereum: "0x2B66AAdE1e9C062FF411bd47C44E0Ad696d43BD9", // sUSDa, 0x8a60e489004ca22d775c5f2c657598278d17d9c2
	ChainTaiko:    "0x5d5c8Aec46661f029A5136a4411C73647a5714a7", // sUSDa, 0x5d5c8aec46661f029a5136a4411c73647a5714a7
}

// Avalon stable coin USDa.
//
// map[chainID] = address
var USDaList = map[int64]string{
	ChainEthereum: "0x8A60E489004Ca22d775C5F2c657598278d17D9c2", // USDa, 0x8a60e489004ca22d775c5f2c657598278d17d9c2
	ChainTaiko:    "0xff12470a969Dd362EB6595FFB44C82c959Fe9ACc", // USDa, 0xff12470a969dd362eb6595ffb44c82c959fe9acc
}

// sUSDa Saving Account. Convert USDa to sUSDa.
//
// map[chainID] = address
var USDaSavingAccountList = map[int64]string{
	ChainEthereum: "0x01e3cc8E17755989ad2CAFE78A822354Eb5DdFA6", // 0x01e3cc8e17755989ad2cafe78a822354eb5ddfa6
}

// Avalon pool manager.
//
// map[chainID] = address
var AvalonPoolManagerList = map[int64]string{
	ChainEthereum: "0x3f390dD6EF69f68f9877aACC086856a200808693", // 0x3f390dd6ef69f68f9877aacc086856a200808693
}
