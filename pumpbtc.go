package ethaddr

// Website: https://mainnet.pumpbtc.xyz/
const PumpBTCProtocol string = "pumpbtc"

// pumpBTC
var PumpBTCList = map[int64]string{
	ChainArbitrum:   "0xF469fBD2abcd6B9de8E169d128226C0Fc90a012e", // pumpBTC, 0xf469fbd2abcd6b9de8e169d128226c0fc90a012e
	ChainAvalanche:  "0x1fCca65fb6Ae3b2758b9b2B394CB227eAE404e1E", // pumpBTC, 0x1fcca65fb6ae3b2758b9b2b394cb227eae404e1e
	ChainBase:       "0xF469fBD2abcd6B9de8E169d128226C0Fc90a012e", // pumpBTC, 0xf469fbd2abcd6b9de8e169d128226c0fc90a012e
	ChainBSC:        "0xf9C4FF105803A77eCB5DAE300871Ad76c2794fa4", // pumpBTC, 0xf9c4ff105803a77ecb5dae300871ad76c2794fa4
	ChainEthereum:   "0xF469fBD2abcd6B9de8E169d128226C0Fc90a012e", // pumpBTC, 0xf469fbd2abcd6b9de8e169d128226c0fc90a012e
	ChainLinea:      "0xF469fBD2abcd6B9de8E169d128226C0Fc90a012e", // pumpBTC, 0xf469fbd2abcd6b9de8e169d128226c0fc90a012e
	ChainMantle:     "0xC75D7767F2EdFbc6a5b18Fc1fA5d51ffB57c2B37", // pumpBTC, 0xc75d7767f2edfbc6a5b18fc1fa5d51ffb57c2b37
	ChainOptimism:   "0xF469fBD2abcd6B9de8E169d128226C0Fc90a012e", // pumpBTC, 0xf469fbd2abcd6b9de8e169d128226c0fc90a012e
	ChainZircuit:    "0xF469fBD2abcd6B9de8E169d128226C0Fc90a012e", // pumpBTC, 0xf469fbd2abcd6b9de8e169d128226c0fc90a012e
	ChainZkLinkNova: "0xDAB5cD46A968aDb6911613896fC61b8B62Cf3B2C", // pumpBTC, 0xdab5cd46a968adb6911613896fc61b8b62cf3b2c
	// ChainTaiko:"0xDBc80A09dE9b075f9380801De2030B3467e3B8FA", // ? NOT Fully Deployed?, 0xdbc80a09de9b075f9380801de2030b3467e3b8fa
}

// pumpBTC.bera
var PumpBTCBeraList = map[int64]string{
	ChainEthereum: "0xADc9c900b05F39f48bB6F402A1BAE60929F4f9A8", // pumpBTC.bera, 0xadc9c900b05f39f48bb6f402a1bae60929f4f9a8
}

// pumpBTC.sei
var PumpBTCSeiList = map[int64]string{
	ChainEthereum: "0xe9ebd666954B7F0B5B044704c86B126651f6235d", // pumpBTC.sei, 0xe9ebd666954b7f0b5b044704c86b126651f6235d
}

// pumpBTC.bera Minter
// To update: internal/test_abis/pumpbtc_test.go
var PumpBTCBeraMinterList = map[int64]string{
	ChainEthereum: "0x9152e9C04e8fE8373EDaa8f5841E25d4015658B7", // 0x9152e9c04e8fe8373edaa8f5841e25d4015658b7
}

// pumpBTC.sei Minter
// To update: internal/test_abis/pumpbtc_test.go
var PumpBTCSeiMinterList = map[int64]string{
	ChainEthereum: "0xB713A6db829F7862516402e605Abb3DDd6F9d7C8", // 0xb713a6db829f7862516402e605abb3ddd6f9d7c8
}

// pumpBTC.sei -> sei = redBTC
var RedBTCList = map[int64]string{
	ChainSei: "0x64310A6176979ac8a752fFC98c0FBbC4CF861ACe", // redBTC, 0x64310a6176979ac8a752ffc98c0fbbc4cf861ace
}
