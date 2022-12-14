package ethaddr

import "github.com/0xVanfer/chainId"

// Indexcoop Defi Pulse Index: DPI.
//
// map[network] = address.
var IndexcoopDefiPulseIndexList = map[string]string{
	chainId.EthereumChainName: "0x1494ca1f11d487c2bbe4543e90080aeba4ba3c2b", // DPI
	chainId.PolygonChainName:  "0x85955046df4668e1dd369d2de9f3aeb98dd2a369", // DPI
}

// Same as IndexcoopDefiPulseIndexList.
var DPIList = IndexcoopDefiPulseIndexList
