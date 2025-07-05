package ethaddr

// Website: https://decentraland.org/
const DecentralandProtocol string = "decentraland"

// Decentraland token: MANA.
//
// map[chainID] = address.
var DecentralandTokenList = map[int64]string{
	ChainEthereum: "0x0F5D2fB29fb7d3CFeE444a200298f468908cC942", // MANA, 0x0f5d2fb29fb7d3cfee444a200298f468908cc942
}

// Same as DecentralandTokenList.
var MANAList = DecentralandTokenList
