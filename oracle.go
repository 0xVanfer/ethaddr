package ethaddr

import "github.com/0xVanfer/chainId"

// Deprecated: Use ToUsdOracleMap["avalanche"][wavax address] instead.
var AVAXOracleMap = map[string]string{
	chainId.AvalancheChainName: "0x0a77230d17318075983913bc2145db16c7366156",
}

// Deprecated: Use ToUsdOracleMap["avalanche"][savax address] instead.
var SAVAXOracleMap = map[string]string{
	chainId.AvalancheChainName: "0xc9245871D69BF4c36c6F2D15E0D68Ffa883FE1A7",
}

// map[network][token] = oracle address.
//
// Addresses from: https://docs.chain.link/docs/data-feeds/price-feeds/addresses/
//
// Or can read from https://cl-docs-addresses.web.app/addresses.json
var ToUsdOracleMap = map[string]map[string]string{
	chainId.EthereumChainName: {
		OneInchTokenList[chainId.EthereumChainName]: "0xc929ad75b72593967de83e7f7cda0493458261d9",
		AAVEList[chainId.EthereumChainName]:         "0x547a514d5e3769680ce22b2361c10ea13619e8a9",
		AMPLList[chainId.EthereumChainName]:         "0xe20ca8d7546932360e37e9d72c1a47334af57706",
		WAVAXList[chainId.EthereumChainName]:        "0xff3eeb22b5e3de6e705b44749c2559d704923fd7",
		BALList[chainId.EthereumChainName]:          "0xdf2917806e30300537aeb49a7663062f4d1f2b5f",
		WBNBList[chainId.EthereumChainName]:         "0x14e613ac84a31f709eadbdf89c6cc390fdc9540a",
		WBTCList[chainId.EthereumChainName]:         "0xf4030086522a5beea4988f8ca5b36dbc97bee88c",
		BUSDList[chainId.EthereumChainName]:         "0x833d8eb16d306ed1fbb5d7a2e019e106b960965a",
		COMPList[chainId.EthereumChainName]:         "0xdbd020caef83efd542f4de03e3cf0c28a4428bd5",
		CRVList[chainId.EthereumChainName]:          "0xcd627aa160a6fa45eb793d19ef54f5062f20f33f",
		CVXList[chainId.EthereumChainName]:          "0xd962fc30a72a84ce50161031391756bf2876af5d",
		DAIList[chainId.EthereumChainName]:          "0xaed0c38402a5d19df6e4c03f4e2dced6e29c1ee9",
		ENSList[chainId.EthereumChainName]:          "0x5c00128d4d1c2f4f652c267d7bcdd7ac99c16e16",
		EtherAddress:                                "0x5f4ec3df9cbd43714fe2740f5e3616155c5b8419",
		WETHList[chainId.EthereumChainName]:         "0x5f4ec3df9cbd43714fe2740f5e3616155c5b8419",
		FRAXList[chainId.EthereumChainName]:         "0xb9e1e3a9feff48998e45fa90847ed4d467e8bcfd",
		GUSDList[chainId.EthereumChainName]:         "0xa89f5d2365ce98b3cd68012b6f503ab1416245fc",
		KNCList[chainId.EthereumChainName]:          "0xf8ff43e991a81e6ec886a3d281a2c6cc19ae70fc",
		LDOList[chainId.EthereumChainName]:          "0x4e844125952d32acdf339be976c98e22f6f318db",
		LINKList[chainId.EthereumChainName]:         "0x2c1d072e956affc0d435cb7ac38ef18d24d9127c",
		LUSDList[chainId.EthereumChainName]:         "0x3d7ae7e594f2f2091ad8798313450130d0aba3a0",
		MANAList[chainId.EthereumChainName]:         "0x56a4857acbcfe3a66965c251628b1c9f1c408c19",
		WMATICList[chainId.EthereumChainName]:       "0x7bac85a8a13a4bcd8abb3eb7d6b4d632c5a57676",
		MIMList[chainId.EthereumChainName]:          "0x7a364e8770418566e3eb2001a96116e6138eb32f",
		MKRList[chainId.EthereumChainName]:          "0xec1d1b3b0443256cc3860e24a46f108e699484aa",
		SNXList[chainId.EthereumChainName]:          "0xdc3ea94cd0ac27d9a86c180091e7f78c683d3699",
		SPELLList[chainId.EthereumChainName]:        "0x8c110b94c5f1d347facf5e1e938ab2db60e3c9a8",
		UNIList[chainId.EthereumChainName]:          "0x553303d460ee0afb37edff9be42922d8ff63220e",
		USDCList[chainId.EthereumChainName]:         "0x8fffffd4afb6115b954bd326cbe7b4ba576818f6",
		USDPList[chainId.EthereumChainName]:         "0x09023c0da49aaf8fc3fa3adf34c6a7016d38d5e3",
		USDTList[chainId.EthereumChainName]:         "0x3e7d1eab13ad0104d2750b8863b489d65364e32d",
		YFIList[chainId.EthereumChainName]:          "0xa027702dbb89fbd58938e4324ac03b58d812b0e1",
		ZRXList[chainId.EthereumChainName]:          "0x2885d15b8af22648b98b122b22fdf4d2a56c6023",
		SUSDList[chainId.EthereumChainName]:         "0xad35bd71b9afe6e4bdc266b345c198eadef9ad94",
	},
	chainId.AvalancheChainName: {
		AAVEList[chainId.AvalancheChainName]:  "0x3ca13391e9fb38a75330fb28f8cc2eb3d9ceceed",
		WAVAXList[chainId.AvalancheChainName]: "0x0a77230d17318075983913bc2145db16c7366156",
		SAVAXList[chainId.AvalancheChainName]: "0xc9245871d69bf4c36c6f2d15e0d68ffa883fe1a7",
		ALPHAList[chainId.AvalancheChainName]: "0x7b0ca9a6d03fe0467a31ca850f5bca51e027b3af",
		WBTCList[chainId.AvalancheChainName]:  "0x2779d32d5166baaa2b2b658333ba7e6ec0c65743",
		BUSDList[chainId.AvalancheChainName]:  "0x827f8a0dc5c943f7524dda178e2e7f275aad743f",
		CRVList[chainId.AvalancheChainName]:   "0x7cf8a6090a9053b01f3df4d4e6cfedd8c90d9027",
		DAIeList[chainId.AvalancheChainName]:  "0x51d7180eda2260cc4f6e4eebb82fef5c3c2b8300",
		WETHList[chainId.AvalancheChainName]:  "0x976b3d034e162d8bd72d6b9c989d545b839003b0",
		FRAXList[chainId.AvalancheChainName]:  "0xbba56ef1565354217a3353a466edb82e8f25b08e",
		JOEList[chainId.AvalancheChainName]:   "0x02d35d3a8ac3e1626d3ee09a78dd87286f5e8e3a",
		KNCList[chainId.AvalancheChainName]:   "0x9df2195dc96e6ef983b1aac275649f3f28f82aa1",
		LINKList[chainId.AvalancheChainName]:  "0x49ccd9ca821efeab2b98c60dc60f518e765ede9a",
		MIMList[chainId.AvalancheChainName]:   "0x54edab30a7134a16a54218ae64c73e1daf48a8fb",
		QIList[chainId.AvalancheChainName]:    "0x36e039e6391a5e7a7267650979fdf613f659be5d",
		SPELLList[chainId.AvalancheChainName]: "0x4f3ddf9378a4865cf4f28be51e10aecb83b7daee",
		SUSHIList[chainId.AvalancheChainName]: "0x449a373a090d8a1e5f74c63ef831ceff39e94563",
		TUSDList[chainId.AvalancheChainName]:  "0x9cf3ef104a973b351b2c032aa6793c3a6f76b448",
		UNIList[chainId.AvalancheChainName]:   "0x9a1372f9b1b71b3a5a72e092ae67e172dbd7daaa",
		USDCeList[chainId.AvalancheChainName]: "0xf096872672f44d6eba71458d74fe67f9a77a23b9",
		USDCList[chainId.AvalancheChainName]:  "0xf096872672f44d6eba71458d74fe67f9a77a23b9",
		USDTeList[chainId.AvalancheChainName]: "0xebe676ee90fe1112671f19b6b7459bc678b67e8a",
		USDTList[chainId.AvalancheChainName]:  "0xebe676ee90fe1112671f19b6b7459bc678b67e8a",
		YFIList[chainId.AvalancheChainName]:   "0x28043b1ebd41860b93ec1f1ec19560760b6db556",
		ZRXList[chainId.AvalancheChainName]:   "0xc07cdf938aa113741fb639bf39699926123fb58b",
	},
	chainId.PolygonChainName: {
		OneInchTokenList[chainId.PolygonChainName]: "0x443c5116cdf663eb387e72c688d276e702135c87",
		AAVEList[chainId.PolygonChainName]:         "0x72484b12719e23115761d5da1646945632979bb6",
		AgEURList[chainId.PolygonChainName]:        "0x9b88d07b2354ef5f4579690356818e07371c7bed",
		BALList[chainId.PolygonChainName]:          "0xd106b538f2a868c28ca1ec7e298c3325e0251d66",
		WBNBList[chainId.PolygonChainName]:         "0x82a6c4af830caa6c97bb504425f6a66165c2c26e",
		WBTCList[chainId.PolygonChainName]:         "0xc907e116054ad103354f2d350fd2514433d57f6f",
		BUSDList[chainId.PolygonChainName]:         "0xe0dc07d5ed74741ceeda61284ee56a2a0f7a4cc9",
		CRVList[chainId.PolygonChainName]:          "0x336584c8e6dc19637a5b36206b1c79923111b405",
		DAIList[chainId.PolygonChainName]:          "0x4746dec9e833a82ec7c2c1356372ccf2cfcd2f3d",
		ENJList[chainId.PolygonChainName]:          "0x440a341bbc9fa86aa60a195e2409a547e48d4c0c",
		WETHList[chainId.PolygonChainName]:         "0xf9680d99d6c9589e2a93a78a04a279e509205945",
		FRAXList[chainId.PolygonChainName]:         "0x00dbeb1e45485d53df7c2f0df1aa0b6dc30311d3",
		KNCList[chainId.PolygonChainName]:          "0x10e5f3dfc81b3e5ef4e648c4454d04e79e1e41e2",
		LINKList[chainId.PolygonChainName]:         "0xd9ffdb71ebe7496cc440152d43986aae0ab76665",
		WMATICList[chainId.PolygonChainName]:       "0xab594600376ec9fd91f8e885dadf0ce036862de0",
		MIMList[chainId.PolygonChainName]:          "0xd133f916e04ed5d67b231183d85be12eaa018320",
		MiMATICList[chainId.PolygonChainName]:      "0xd8d483d813547cfb624b8dc33a00f2fcbcd2d428",
		QUICKList[chainId.PolygonChainName]:        "0xa058689f4bca95208bba3f265674ae95ded75b6d",
		RAIList[chainId.PolygonChainName]:          "0x7f45273fd7c644714825345670414ea649b50b16",
		SUSHIList[chainId.PolygonChainName]:        "0x49b0c695039243bbfeb8ecd054eb70061fd54aa0",
		TUSDList[chainId.PolygonChainName]:         "0x7c5d415b64312d38c56b54358449d0a4058339d2",
		UNIList[chainId.PolygonChainName]:          "0xdf0fb4e4f928d2dcb76f438575fdd8682386e13c",
		USDCList[chainId.PolygonChainName]:         "0xfe4a8cc5b5b2366c1b58bea3858e81843581b2f7",
		USDTList[chainId.PolygonChainName]:         "0x0a6513e40db6eb1b165753ad52e80663aea50545",
		WBTCList[chainId.PolygonChainName]:         "0xde31f8bfbd8c84b5360cfacca3539b938dd78ae6",
	},
}
