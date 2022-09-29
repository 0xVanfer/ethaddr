package ethaddr

import "github.com/0xVanfer/chainId"

const MakerProtocol string = "maker"

// Maker token: MKR.
//
// map[network] = address.
var MakerTokenList = map[string]string{
	chainId.EthereumChainName: "0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2", // mkr
}
