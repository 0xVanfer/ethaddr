package ethaddr

// Compound token: COMP.
//
// map[network] = address.
var CompoundTokenList = map[int64]string{
	ChainArbitrum: "0x354A6dA3fcde098F8389cad84b0182725c6C91dE", // COMP, 0x354a6da3fcde098f8389cad84b0182725c6c91de
	ChainBase:     "0x9e1028F5F1D5eDE59748FFceE5532509976840E0", // COMP, 0x9e1028f5f1d5ede59748ffcee5532509976840e0
	ChainEthereum: "0xc00e94Cb662C3520282E6f5717214004A7f26888", // COMP, 0xc00e94cb662c3520282e6f5717214004a7f26888
	ChainOptimism: "0x7e7d4467112689329f7E06571eD0E8CbAd4910eE", // COMP, 0x7e7d4467112689329f7e06571ed0e8cbad4910ee
	ChainScroll:   "0x643e160a3C3E2B7eae198f0beB1BfD2441450e86", // COMP, 0x643e160a3c3e2b7eae198f0beb1bfd2441450e86
}

// Same as CompoundTokenList.
var COMPList = CompoundTokenList
