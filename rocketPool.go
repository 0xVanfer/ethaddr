package ethaddr

// Website: https://rocketpool.net/
//
// Docs: https://docs.rocketpool.net/
//
// Deployed contracts: https://docs.rocketpool.net/overview/contracts-integrations
const RocketPoolProtocol string = "rocketpool"

// Rocket pool rETH.
//
// map[chainID] = address.
var RETHList = map[int64]string{
	ChainArbitrum:  "0xEC70Dcb4A1EFa46b8F2D97C310C9c4790ba5ffA8", // rETH, 0xec70dcb4a1efa46b8f2d97c310c9c4790ba5ffa8
	ChainBase:      "0xB6fe221Fe9EeF5aBa221c348bA20A1Bf5e73624c", // rETH, 0xb6fe221fe9eef5aba221c348ba20a1bf5e73624c
	ChainEthereum:  "0xae78736Cd615f374D3085123A210448E74Fc6393", // rETH, 0xae78736cd615f374d3085123a210448e74fc6393
	ChainOptimism:  "0x9Bcef72be871e61ED4fBbc7630889beE758eb81D", // rETH, 0x9bcef72be871e61ed4fbbc7630889bee758eb81d
	ChainPolygon:   "0x0266F4F08D82372CF0FcbCCc0Ff74309089c74d1", // rETH, 0x0266f4f08d82372cf0fcbccc0ff74309089c74d1
	ChainPolygonZk: "0xb23C20EFcE6e24Acca0Cef9B7B7aA196b84EC942", // rETH, 0xb23c20efce6e24acca0cef9b7b7aa196b84ec942
	ChainScroll:    "0x53878B874283351D26d206FA512aEcE1Bef6C0dD", // rETH, 0x53878b874283351d26d206fa512aece1bef6c0dd
}

// Rocket pool token.
//
// map[chainID] = address.
var RocketPoolTokenList = map[int64]string{
	ChainArbitrum:  "0xB766039cc6DB368759C1E56B79AFfE831d0Cc507", // RPL, 0xb766039cc6db368759c1e56b79affe831d0cc507
	ChainBase:      "0x1f73EAf55d696BFFA9b0EA16fa987B93b0f4d302", // RPL, 0x1f73eaf55d696bffa9b0ea16fa987b93b0f4d302
	ChainEthereum:  "0xD33526068D116cE69F19A9ee46F0bd304F21A51f", // RPL, 0xd33526068d116ce69f19a9ee46f0bd304f21a51f
	ChainOptimism:  "0xC81D1F0EB955B0c020E5d5b264E1FF72c14d1401", // RPL, 0xc81d1f0eb955b0c020e5d5b264e1ff72c14d1401
	ChainPolygon:   "0x7205705771547cF79201111B4bd8aaF29467b9eC", // RPL, 0x7205705771547cf79201111b4bd8aaf29467b9ec
	ChainPolygonZk: "0x70d35152fBf63FB312709b11a9Bac87519de0019", // RPL, 0x70d35152fbf63fb312709b11a9bac87519de0019
}

// Rocket pool token.
var RPLList = RocketPoolTokenList
