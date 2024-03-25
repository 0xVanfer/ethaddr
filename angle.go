package ethaddr

// Website: https://app.angle.money/
//
// Docs: https://docs.angle.money/overview/readme
//
// Deployed contracts: https://developers.angle.money/overview/smart-contracts/mainnet-contracts
const AngleProtocol string = "angle"

// Angle token: ANGLE.
//
// map[network] = address.
var AngleTokenList = map[string]string{
	ChainPolygon: "0x900F717EA076E1E7a484ad9DD2dB81CEEc60eBF1", // ANGLE, 0x900f717ea076e1e7a484ad9dd2db81ceec60ebf1
}

// Same as AngleTokenList.
var ANGLEList = AngleTokenList

// EURA token: EURA (previously agEUR) and still shows agEUR on blockscans.
//
// map[network] = address.
var EURAList = map[string]string{
	ChainEthereum: "0x1a7e4e63778B4f12a199C062f3eFdD288afCBce8", // EURA (previously agEUR), 0x1a7e4e63778b4f12a199c062f3efdd288afcbce8
	ChainPolygon:  "0xE0B52e49357Fd4DAf2c15e02058DCE6BC0057db4", // EURA (previously agEUR), 0xe0b52e49357fd4daf2c15e02058dce6bc0057db4
	ChainArbitrum: "0xFA5Ed56A203466CbBC2430a43c66b9D8723528E7", // EURA (previously agEUR), 0xfa5ed56a203466cbbc2430a43c66b9d8723528e7
}

// EURA, previously known as AgEUR.
var AgEURLust = EURAList
