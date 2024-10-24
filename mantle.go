package ethaddr

// Website: https://www.mantle.xyz/meth
//
// Docs: https://docs-v2.mantle.xyz/
const MantleProtocol string = "mantle"

// Blockscan: https://explorer.mantle.xyz/ (non etherscan)
const ChainNameMantle string = "mantle"

// Mantle stake ETH: mETH.
//
// map[chainID] = address
var MantleStakedEtherList = map[int64]string{
	ChainEthereum: "0xd5F7838F5C461fefF7FE49ea5ebaF7728bB0ADfa", // mETH, 0xd5f7838f5c461feff7fe49ea5ebaf7728bb0adfa
}

// Same as MantleStakeETHList.
var METHList = MantleStakedEtherList
