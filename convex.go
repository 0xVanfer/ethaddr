package ethaddr

// Website: https://www.convexfinance.com/
//
// Docs: https://docs.convexfinance.com/convexfinance/
//
// Deployed contracts: https://docs.convexfinance.com/convexfinance/faq/contract-addresses
const ConvexProtocol string = "convex"

// Convex token: CVX.
//
// map[network] = address.
var ConvexTokenList = map[string]string{
	ChainEthereum: "0x4e3FBD56CD56c3e72c1403e103b45Db9da5B9D2B", // CVX, 0x4e3fbd56cd56c3e72c1403e103b45db9da5b9d2b
}

// Same as ConvexTokenList.
var CVXList = ConvexTokenList
