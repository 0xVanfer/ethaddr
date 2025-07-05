package ethaddr

// Website: https://app.compound.finance/markets
//
// Deployed contracts: https://docs.compound.finance/#protocol-contracts
const CompoundV3Protocol string = "compoundv3"

// Compound V3 markets.
//
// map[chainID] = address.
var CompoundV3MarketsList = map[int64]map[string]string{
	ChainArbitrum: {
		WETHList[ChainArbitrum]:  "0x6f7D514bbD4aFf3BcD1140B7344b32f063dEe486", // 0x6f7d514bbd4aff3bcd1140b7344b32f063dee486
		USDCList[ChainArbitrum]:  "0x9c4ec768c28520B50860ea7a15bd7213a9fF58bf", // 0x9c4ec768c28520b50860ea7a15bd7213a9ff58bf
		USDCeList[ChainArbitrum]: "0xA5EDBDD9646f8dFF606d7448e414884C7d905dCA", // 0xa5edbdd9646f8dff606d7448e414884c7d905dca
		USDTList[ChainArbitrum]:  "0xd98Be00b5D27fc98112BdE293e487f8D4cA57d07", // 0xd98be00b5d27fc98112bde293e487f8d4ca57d07
	},
	ChainBase: {
		USDCList[ChainBase]:  "0xb125E6687d4313864e53df431d5425969c15Eb2F", // 0xb125e6687d4313864e53df431d5425969c15eb2f
		USDbCList[ChainBase]: "0x9c4ec768c28520B50860ea7a15bd7213a9fF58bf", // 0x9c4ec768c28520b50860ea7a15bd7213a9ff58bf
		WETHList[ChainBase]:  "0x46e6b214b524310239732D51387075E0e70970bf", // 0x46e6b214b524310239732d51387075e0e70970bf
		AEROList[ChainBase]:  "0x784efeB622244d2348d4F2522f8860B96fbEcE89", // 0x784efeB622244d2348d4F2522f8860B96fbEcE89
	},
	ChainEthereum: {
		WETHList[ChainEthereum]:   "0xA17581A9E3356d9A858b789D68B4d866e593aE94", // 0xa17581a9e3356d9a858b789d68b4d866e593ae94
		USDCList[ChainEthereum]:   "0xc3d688B66703497DAA19211EEdff47f25384cdc3", // 0xc3d688b66703497daa19211eedff47f25384cdc3
		USDTList[ChainEthereum]:   "0x3Afdc9BCA9213A35503b077a6072F3D0d5AB0840", // 0x3afdc9bca9213a35503b077a6072f3d0d5ab0840
		WSTETHList[ChainEthereum]: "0x3D0bb1ccaB520A66e607822fC55BC921738fAFE3", // 0x3d0bb1ccab520a66e607822fc55bc921738fafe3
		USDSList[ChainEthereum]:   "0x5D409e56D886231aDAf00c8775665AD0f9897b56", // 0x5d409e56d886231adaf00c8775665ad0f9897b56
	},
	ChainOptimism: {
		USDCList[ChainOptimism]: "0x2e44e174f7D53F0212823acC11C01A11d58c5bCB", // 0x2e44e174f7d53f0212823acc11c01a11d58c5bcb
		USDTList[ChainOptimism]: "0x995E394b8B2437aC8Ce61Ee0bC610D617962B214", // 0x995e394b8b2437ac8ce61ee0bc610d617962b214
		WETHList[ChainOptimism]: "0xE36A30D249f7761327fd973001A32010b521b6Fd", // 0xe36a30d249f7761327fd973001a32010b521b6fd
	},
	ChainPolygon: {
		USDCList[ChainPolygon]: "0xF25212E676D1F7F89Cd72fFEe66158f541246445", // 0xf25212e676d1f7f89cd72ffee66158f541246445
		USDTList[ChainPolygon]: "0xaeB318360f27748Acb200CE616E389A6C9409a07", // 0xaeb318360f27748acb200ce616e389a6c9409a07
	},
	ChainScroll: {
		USDCList[ChainScroll]: "0xB2f97c1Bd3bf02f5e74d13f02E3e26F93D77CE44", // 0xb2f97c1bd3bf02f5e74d13f02e3e26f93d77ce44
	},
	ChainMantle: {
		USDeList[ChainMantle]: "0x606174f62cd968d8e684c645080fa694c1D7786E", // 0x606174f62cd968d8e684c645080fa694c1d7786e
	},
}

var CompoundV3CTokenList = CompoundV3MarketsList

// Compound V3 market extentions.
//
// map[chainID] = address.
var CompoundV3MarketsExtentionList = map[int64]map[string]string{
	ChainArbitrum: {
		WETHList[ChainArbitrum]:  "0x5404872d8f2e24b230EC9B9eC64E3855F637FB93", // 0x5404872d8f2e24b230ec9b9ec64e3855f637fb93
		USDCList[ChainArbitrum]:  "0x1B2E88cC7365d90e7E81392432482925BD8437E9", // 0x1b2e88cc7365d90e7e81392432482925bd8437e9
		USDCeList[ChainArbitrum]: "0x1B2E88cC7365d90e7E81392432482925BD8437E9", // 0x1b2e88cc7365d90e7e81392432482925bd8437e9
		USDTList[ChainArbitrum]:  "0x698A949f3b4f7a5DdE236106F25Fa0eAcA0FcEF1", // 0x698a949f3b4f7a5dde236106f25fa0eaca0fcef1
	},
	ChainBase: {
		USDCList[ChainBase]:  "0x3bac64185786922292266AA92a58cf870D694E2a", // 0x3bac64185786922292266aa92a58cf870d694e2a
		USDbCList[ChainBase]: "0x2F9E3953b2Ef89fA265f2a32ed9F80D00229125B", // 0x2f9e3953b2ef89fa265f2a32ed9f80d00229125b
		WETHList[ChainBase]:  "0x88bB8C109640778D3fB1074bB10a66e31F2c9c17", // 0x88bb8c109640778d3fb1074bb10a66e31f2c9c17
		AEROList[ChainBase]:  "0x6f7D514bbD4aFf3BcD1140B7344b32f063dEe486", // 0x6f7d514bbd4aff3bcd1140b7344b32f063dee486
	},
	ChainEthereum: {
		WETHList[ChainEthereum]:   "0xe2C1F54aFF6b38fD9DF7a69F22cB5fd3ba09F030", // 0xe2c1f54aff6b38fd9df7a69f22cb5fd3ba09f030
		USDCList[ChainEthereum]:   "0x285617313887d43256F852cAE0Ee4de4b68D45B0", // 0x285617313887d43256f852cae0ee4de4b68d45b0
		USDTList[ChainEthereum]:   "0x5C58d4479A1E9b2d19EE052143FA73F0ee79A36e", // 0x5c58d4479a1e9b2d19ee052143fa73f0ee79a36e
		WSTETHList[ChainEthereum]: "0x995E394b8B2437aC8Ce61Ee0bC610D617962B214", // 0x995e394b8b2437ac8ce61ee0bc610d617962b214
		USDSList[ChainEthereum]:   "0x95DeDD64b551F05E9f59a101a519B024b6b116E7", // 0x95dedd64b551f05e9f59a101a519b024b6b116e7
	},
	ChainOptimism: {
		USDCList[ChainOptimism]: "0xE802a0b833f6080FEB46DD54c75444c5dba0c873", // 0xe802a0b833f6080feb46dd54c75444c5dba0c873
		USDTList[ChainOptimism]: "0xC49399814452B41dA8a7cd76a159f5515cb3e493", // 0xc49399814452b41da8a7cd76a159f5515cb3e493
		WETHList[ChainOptimism]: "0x82B8d9A06ccABC1e9B0c0A00f38B858E6925CF2f", // 0x82b8d9a06ccabc1e9b0c0a00f38b858e6925cf2f
	},
	ChainPolygon: {
		USDCList[ChainPolygon]: "0xbdE8F31D2DdDA895264e27DD990faB3DC87b372d", // 0xbde8f31d2ddda895264e27dd990fab3dc87b372d
		USDTList[ChainPolygon]: "0x2F4eAF29dfeeF4654bD091F7112926E108eF4Ed0", // 0x2f4eaf29dfeef4654bd091f7112926e108ef4ed0
	},
	ChainScroll: {
		USDCList[ChainScroll]: "0x27E24C49f95DfF7E231eF1C2849F760cDF25a5Ad", // 0x27e24c49f95dff7e231ef1c2849f760cdf25a5ad
	},
	ChainMantle: {
		USDeList[ChainMantle]: "0xf528B4bCAc12dad0bFa114282b219ad706bA7f18", // 0xf528b4bcac12dad0bfa114282b219ad706ba7f18
	},
}
