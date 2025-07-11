package ethaddr

// Website: https://www.circle.com/
const CircleProtocol string = "circle"

// USDC.
//
// The original USDC mint by Circle, not the bridged ones. Avalanche, Arbitrum, Polygon, Optimism all have USDC.e, the bridged one. See USDCeList.
//
// BSC use the one mint by binance itself, and peg with Circle USDC.
//
// map[chainID] = address.
var USDCList = map[int64]string{
	ChainArbitrum:  "0xaf88d065e77c8cC2239327C5EDb3A432268e5831", // USDC, 0xaf88d065e77c8cc2239327c5edb3a432268e5831
	ChainAvalanche: "0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E", // USDC, 0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e
	ChainBase:      "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913", // USDC, 0x833589fcd6edb6e08f4c7c32d4f71b54bda02913
	ChainBera:      "0x549943e04f40284185054145c6E4e9568C1D3241", // USDC, 0x549943e04f40284185054145c6e4e9568c1d3241
	ChainBSC:       "0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d", // USDC, 0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d
	ChainCelo:      "0xcebA9300f2b948710d2653dD7B07f33A8B32118C", // USDC, 0xceba9300f2b948710d2653dd7b07f33a8b32118c
	ChainEthereum:  "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", // USDC, 0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48
	ChainOptimism:  "0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85", // USDC, 0x0b2c639c533813f4aa9d7837caf62653d097ff85
	ChainPolygon:   "0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359", // USDC, 0x3c499c542cef5e3811e1192ce70d8cc03d5c3359
	ChainScroll:    "0x06eFdBFf2a14a7c8E15944D1F4A48F9F95F663A4", // USDC, 0x06efdbff2a14a7c8e15944d1f4a48f9f95f663a4
	ChainSei:       "0x3894085Ef7Ff0f0aeDf52E2A2704928d1Ec074F1", // USDC, 0x3894085ef7ff0f0aedf52e2a2704928d1ec074f1
	ChainTaiko:     "0x07d83526730c7438048D55A4fc0b850e2aaB6f0b", // USDC, 0x07d83526730c7438048d55a4fc0b850e2aab6f0b
	ChainZkSync:    "0x1d17CBcF0D6D143135aE902365D2E5e2A16538D4", // USDC, 0x1d17cbcf0d6d143135ae902365d2e5e2a16538d4
}

// USDC.e. The bridged USDC. See the original USDC mint by Circle in USDCList.
//
// map[chainID] = address.
var USDCeList = map[int64]string{
	ChainArbitrum:  "0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8", // USDC.e, 0xff970a61a04b1ca14834a43f5de4533ebddb5cc8
	ChainAvalanche: "0xA7D7079b0FEaD91F3e65f86E8915Cb59c1a4C664", // USDC.e, 0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664
	ChainOptimism:  "0x7F5c764cBc14f9669B88837ca1490cCa17c31607", // USDC.e, 0x7f5c764cbc14f9669b88837ca1490cca17c31607
	ChainPolygon:   "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174", // USDC.e, 0x2791bca1f2de4661ed88a30c99a7a9449aa84174
	ChainSonic:     "0x29219dd400f2Bf60E5a23d13Be72B486D4038894", // USDC.e, 0x29219dd400f2bf60e5a23d13be72b486d4038894
}

// Circle EURC.
var EURCList = map[int64]string{
	ChainAvalanche: "0xC891EB4cbdEFf6e073e859e987815Ed1505c2ACD", // EURC, 0xc891eb4cbdeff6e073e859e987815ed1505c2acd
	ChainBase:      "0x60a3E35Cc302bFA44Cb288Bc5a4F316Fdb1adb42", // EURC, 0x60a3e35cc302bfa44cb288bc5a4f316fdb1adb42
	ChainEthereum:  "0x1aBaEA1f7C830bD89Acc67eC4af516284b1bC33c", // EURC, 0x1abaea1f7c830bd89acc67ec4af516284b1bc33c
}
