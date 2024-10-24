package ethaddr

// Website: https://www.bedrock.technology/
//
// Docs: https://docs.bedrock.technology/; https://docs.bedrock.technology/security/smart-contracts
const BedrockProtocol string = "bedrock"

// Bedrock uniBTC.
//
// map[chainID] = address
var UniBTCList = map[int64]string{
	ChainEthereum: "0x004E9C3EF86bc1ca1f0bB5C7662861Ee93350568", // uniBTC, 0x004e9c3ef86bc1ca1f0bb5c7662861ee93350568
	// ChainScroll:   "",
}

// Bedrock uniETH.
//
// map[chainID] = address
var UniETHList = map[int64]string{
	ChainScroll: "0x15EEfE5B297136b8712291B632404B66A8eF4D25", // uniETH, 0x15eefe5b297136b8712291b632404b66a8ef4d25
}
