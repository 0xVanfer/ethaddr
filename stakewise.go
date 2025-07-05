package ethaddr

// Website: https://www.stakewise.io/
const StakeWiseProtocol string = "stakewise"

// Stake Wise staked ETH. osETH
//
// map[chainID] = address.
var OSETHList = map[int64]string{
	ChainEthereum: "0xf1C9acDc66974dFB6dEcB12aA385b9cD01190E38", // osETH, 0xf1c9acdc66974dfb6decb12aa385b9cd01190e38
}

// Where user burn osETH and get WETH.
//
// map[chainID] = address.
var StakewiseWithdrawETHList = map[int64]string{
	ChainEthereum: "0xAC0F906E433d58FA868F936E8A43230473652885", // 0xac0f906e433d58fa868f936e8a43230473652885
}
