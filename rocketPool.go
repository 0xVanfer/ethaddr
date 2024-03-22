package ethaddr

// Rocket pool rETH.
//
// map[network] = address.
var RETHList = map[string]string{
	chainArbitrum: "0xEC70Dcb4A1EFa46b8F2D97C310C9c4790ba5ffA8", // rETH, 0xec70dcb4a1efa46b8f2d97c310c9c4790ba5ffa8
	chainEthereum: "0xae78736Cd615f374D3085123A210448E74Fc6393", // rETH, 0xae78736cd615f374d3085123a210448e74fc6393
	chainOptimism: "0x9Bcef72be871e61ED4fBbc7630889beE758eb81D", // rETH, 0x9bcef72be871e61ed4fbbc7630889bee758eb81d
}

// Rocket pool token.
//
// map[network] = address.
var RocketPoolTokenList = map[string]string{
	chainEthereum: "0xD33526068D116cE69F19A9ee46F0bd304F21A51f", // RPL, 0xd33526068d116ce69f19a9ee46f0bd304f21a51f
}

// Rocket pool token.
var RPLList = RocketPoolTokenList
