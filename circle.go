package ethaddr

// USDC.
//
// The original USDC mint by Circle, not the bridged ones. Avalanche, Arbitrum, Polygon, Optimism all have USDC.e, the bridged one. See USDCeList.
//
// BSC use the one mint by binance itself, and peg with Circle USDC.
//
// map[network] = address.
var USDCList = map[string]string{
	chainArbitrum:  "0xaf88d065e77c8cC2239327C5EDb3A432268e5831", // USDC, 0xaf88d065e77c8cc2239327c5edb3a432268e5831
	chainAvalanche: "0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E", // USDC, 0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e
	chainBase:      "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913", // USDC, 0x833589fcd6edb6e08f4c7c32d4f71b54bda02913
	chainBSC:       "0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d", // USDC, 0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d
	chainEthereum:  "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", // USDC, 0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48
	chainOptimism:  "0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85", // USDC, 0x0b2c639c533813f4aa9d7837caf62653d097ff85
	chainPolygon:   "0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359", // USDC, 0x3c499c542cef5e3811e1192ce70d8cc03d5c3359
}

// USDC.e. The bridged USDC. See the original USDC mint by Circle in USDCList.
//
// map[network] = address.
var USDCeList = map[string]string{
	chainArbitrum:  "0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8", // USDC.e, 0xff970a61a04b1ca14834a43f5de4533ebddb5cc8
	chainAvalanche: "0xA7D7079b0FEaD91F3e65f86E8915Cb59c1a4C664", // USDC.e, 0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664
	chainOptimism:  "0x7F5c764cBc14f9669B88837ca1490cCa17c31607", // USDC.e, 0x7f5c764cbc14f9669b88837ca1490cca17c31607
	chainPolygon:   "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174", // USDC.e, 0x2791bca1f2de4661ed88a30c99a7a9449aa84174
}
