package ethaddr

// Doc: https://docs.merlinchain.io/merlin-docs/developers/builder-guides/networks/mainnet
//
// Blockscan: https://scan.merlinchain.io/
const ChainNameMerlin string = "merlin"

// Merlin chain BTC, M-BTC.
var MerlinBTCList = map[int64]string{
	ChainEthereum: "0x2F913C820ed3bEb3a67391a6eFF64E70c4B20b19", // M-BTC, 0x2f913c820ed3beb3a67391a6eff64e70c4b20b19
	ChainMerlin:   "0xB880fd278198bd590252621d4CD071b1842E9Bcd", // M-BTC, 0xb880fd278198bd590252621d4cd071b1842e9bcd
	ChainSei:      "0x9BFA177621119e64CecbEabE184ab9993E2ef727", // M-BTC, 0x9bfa177621119e64cecbebae184ab9993e2ef727
}
