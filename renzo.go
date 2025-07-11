package ethaddr

// Website: https://www.renzoprotocol.com/
const RenzoProtocol string = "renzo"

// Renzo token: REZ.
//
// map[chainID] = address
var REZList = map[int64]string{
	ChainEthereum: "0x3B50805453023a91a8bf641e279401a0b23FA6F9", // REZ, 0x3b50805453023a91a8bf641e279401a0b23fa6f9
}

// Renzo symbiotic restaked vault: pzETH.
//
// map[chainID] = address
var RenzoSymbioticRestakedLSTList = map[int64]string{
	ChainEthereum: "0x8c9532a60E0E7C6BbD2B2c1303F63aCE1c3E9811", // pzETH, 0x8c9532a60e0e7c6bbd2b2c1303f63ace1c3e9811
}

// Renzo eigen layer restaked vault: ezETH.
//
// map[chainID] = address
var RenzoEigenlayerRestakedLSTList = map[int64]string{
	ChainArbitrum: "0x2416092f143378750bb29b79eD961ab195CcEea5", // ezETH, 0x2416092f143378750bb29b79ed961ab195cceea5
	ChainBase:     "0x2416092f143378750bb29b79eD961ab195CcEea5", // ezETH, 0x2416092f143378750bb29b79ed961ab195cceea5
	ChainEthereum: "0xbf5495Efe5DB9ce00f80364C8B423567e58d2110", // ezETH, 0xbf5495efe5db9ce00f80364c8b423567e58d2110
}

// Same as RenzoSymbioticRestakedLSTList.
var PzETHList = RenzoSymbioticRestakedLSTList

// Same as RenzoEigenlayerRestakedLSTList.
var EzETHList = RenzoEigenlayerRestakedLSTList

// Renzo ezETH restake manager. Read the tvl of ezETH.
//
// map[chainID] = address
var RenzoEzETHRestakeManagerList = map[int64]string{
	ChainEthereum: "0x74a09653A083691711cF8215a6ab074BB4e99ef5", // 0x74a09653a083691711cf8215a6ab074bb4e99ef5
}
