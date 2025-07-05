package ethaddr

// Blockscan: ?
//
// Website: https://www.monad.xyz/
const ChainNameMonad string = "monad"

// Testnet
//
// Blockscan for testnet: https://testnet.monadexplorer.com/ or https://monad-testnet.socialscan.io/
const ChainNameMonadTestnet string = "monadtestnet"

// Wrapped Monad.
//
// map[chainID] = address.
var WMONList = map[int64]string{
	ChainMonadTestnet: "0x760AfE86e5de5fa0Ee542fc7B7B713e1c5425701",
}
