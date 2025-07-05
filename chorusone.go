package ethaddr

// Website: https://chorus.one/
const ChorusOneProtocol string = "chorusone"

// Chorus one symbiotic restaking vault ETH: coETH.
//
// map[chainID] = address
var ChorusOneRestakingVaultETHList = map[int64]string{
	ChainEthereum: "0xd6E09a5e6D719d1c881579C9C8670a210437931b", // coETH, 0xd6e09a5e6d719d1c881579c9c8670a210437931b
}

// Same as ChorusOneRestakingVaultETHList.
var CoETHList = ChorusOneRestakingVaultETHList
