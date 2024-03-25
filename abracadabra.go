package ethaddr

// Website:https://abracadabra.money/
//
// Docs: https://docs.abracadabra.money/learn/
const AbracadabraProtocol string = "abracadabra"

// Abracadabra token: SPELL.
//
// map[network] = address.
var AbracadabraTokenList = map[string]string{
	ChainArbitrum:  "0x3E6648C5a70A150A88bCE65F4aD4d506Fe15d2AF", // SPELL, 0x3e6648c5a70a150a88bce65f4ad4d506fe15d2af
	ChainAvalanche: "0xCE1bFFBD5374Dac86a2893119683F4911a2F7814", // SPELL, 0xce1bffbd5374dac86a2893119683f4911a2f7814
	ChainEthereum:  "0x090185f2135308BaD17527004364eBcC2D37e5F6", // SPELL, 0x090185f2135308bad17527004364ebcc2d37e5f6
}

// Same as AbracadabraTokenList.
var SPELLList = AbracadabraTokenList

// Abracadabra magic internet money: MIM.
//
// map[network] = address.
var MIMList = map[string]string{
	ChainArbitrum:  "0xFEa7a6a0B346362BF88A9e4A88416B77a57D6c2A", // MIM, 0xfea7a6a0b346362bf88a9e4a88416b77a57d6c2a
	ChainAvalanche: "0x130966628846BFd36ff31a822705796e8cb8C18D", // MIM, 0x130966628846bfd36ff31a822705796e8cb8c18d
	ChainEthereum:  "0x99D8a9C45b2ecA8864373A26D1459e3Dff1e17F3", // MIM, 0x99d8a9c45b2eca8864373a26d1459e3dff1e17f3
}
