package ethaddr

// Website: https://www.fasanara.com/
const FasanaraProtocol string = "fasanara"

// Midas Finance ONE mF-ONE
//
// map[chainID] = address.
var MFONEList = map[int64]string{
	ChainEthereum: "0x238a700eD6165261Cf8b2e544ba797BC11e466Ba", // mF-ONE, 0x238a700ed6165261cf8b2e544ba797bc11e466ba
}

// Midas Finance Oracle for mF-ONE
//
// map[chainID] = address.
var MFONEOracleList = map[int64]string{
	ChainEthereum: "0x8D51DBC85cEef637c97D02bdaAbb5E274850e68C", // 0x8d51dbc85ceef637c97d02bdaabb5e274850e68c
}
