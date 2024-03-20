package ethaddr

// Rocket pool rETH.
//
// map[network] = address.
var RETHList = map[string]string{
	chainEthereum: "0xae78736Cd615f374D3085123A210448E74Fc6393", // rETH, 0xae78736cd615f374d3085123a210448e74fc6393
	chainArbitrum: "0xae78736Cd615f374D3085123A210448E74Fc6393", // rETH, 0xae78736cd615f374d3085123a210448e74fc6393
}

// Rocket pool token.
//
// map[network] = address.
var RocketPoolTokenList = map[string]string{
	chainEthereum: "0xD33526068D116cE69F19A9ee46F0bd304F21A51f", // RPL, 0xd33526068d116ce69f19a9ee46f0bd304f21a51f
}

// Rocket pool token.
var RPLList = RocketPoolTokenList
