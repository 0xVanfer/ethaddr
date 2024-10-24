package ethaddr

// Website: https://www.rhomarkets.xyz/
const RhoProtocol string = "rho"

// Rho comptroller as a compound v2 fork.
var RhoComptrollerList = map[int64]string{
	ChainScroll: "0x8a67AB98A291d1AEA2E1eB0a79ae4ab7f2D76041", // 0x8a67ab98a291d1aea2e1eb0a79ae4ab7f2d76041
}

// Rho c tokens(r tokens).
//
// map[network][underlying] = address.
var RhoCTokenList = map[int64]map[string]string{
	ChainScroll: {
		EtherAddress:                "0x639355f34Ca9935E0004e30bD77b9cE2ADA0E692", // rETH, 0x639355f34ca9935e0004e30bd77b9ce2ada0e692
		WETHList[ChainScroll]:       "0x639355f34Ca9935E0004e30bD77b9cE2ADA0E692", // rETH, 0x639355f34ca9935e0004e30bd77b9ce2ada0e692
		USDTList[ChainScroll]:       "0x855CEA8626Fa7b42c13e7A688b179bf61e6c1e81", // rUSDT, 0x855cea8626fa7b42c13e7a688b179bf61e6c1e81
		USDCList[ChainScroll]:       "0xAE1846110F72f2DaaBC75B7cEEe96558289EDfc5", // rUSDC, 0xae1846110f72f2daabc75b7ceee96558289edfc5
		STONEList[ChainScroll]:      "0xAD3d07d431B85B525D81372802504Fa18DBd554c", // rSTONE, 0xad3d07d431b85b525d81372802504fa18dbd554c
		WSTETHList[ChainScroll]:     "0xe4FC4C444efFB5ECa80274c021f652980794Eae6", // rwstETH, 0xe4fc4c444effb5eca80274c021f652980794eae6
		SolvBTCbList[ChainScroll]:   "0x8966993138b95b48142f6ecB590427eb7e18a719", // rsolvBTC.b, 0x8966993138b95b48142f6ecb590427eb7e18a719
		WEETHList[ChainScroll]:      "0x65a5dBEf0D1Bff772822E4652Aed2829718DC43F", // rweETH, 0x65a5dbef0d1bff772822e4652aed2829718dc43f
		WrsETHList[ChainScroll]:     "0x52Fef2B9040BA81e40421660335655D70Fe8Cf03", // rwrsETH, 0x52fef2b9040ba81e40421660335655d70fe8cf03
		USDeList[ChainScroll]:       "0x5fF1926507f6e71bFbd5f9897fBaeF021E2F77CA", // rUSDe, 0x5ff1926507f6e71bfbd5f9897fbaef021e2f77ca
		UniETHList[ChainScroll]:     "0xFE707359517f0d5AD0187a237974D3110A734016", // runiETH, 0xfe707359517f0d5ad0187a237974d3110a734016
		YLstETHBFEList[ChainScroll]: "0x7AE3c19De353ce163Fe81AE3ebFC90709d3868BE", // rylstETH, 0x7ae3c19de353ce163fe81ae3ebfc90709d3868be
		// UniBTCList[ChainScroll]:   "0xFE707359517f0d5AD0187a237974D3110A734016", // runiBTC, 0xfe707359517f0d5ad0187a237974d3110a734016
		// WETHList[ChainScroll]: "0x76DC94562c89D2820E88D1274d4Bb32Cee306d4C", // This seems to be deprecated. WETH (???).
		// WBTCList[ChainScroll]: "0x1D73ead2BBEa318344fccD7142F70488BAb08F44", // This seems to be deprecated. WBTC (???).
		// WEETHList[ChainScroll]: "0x8698fB1093b6DBC345e2aAFEb853C602A3582548", // This seems to be deprecated. rweETH.
		// WEETHList[ChainScroll]: "0x00B49129Af3be28D0Cdb1e60B155B234d0E19190", // This seems to be deprecated. rweETH.
	},
}
