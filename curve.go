package ethaddr

import "github.com/0xVanfer/chainId"

// Docs: https://curve.readthedocs.io/
//
// Deployed contracts: https://curve.readthedocs.io/ref-addresses.html
const CurveProtocol string = "curve"

// Curve token: CRV.
//
// map[network] = address.
var CurveTokenlist = map[string]string{
	chainId.EthereumChainName:  "0xD533a949740bb3306d119CC777fa900bA034cd52", // CRV, 0xd533a949740bb3306d119cc777fa900ba034cd52
	chainId.AvalancheChainName: "0x47536f17f4ff30e64a96a7555826b8f9e66ec468", // CRV
	chainId.PolygonChainName:   "0x172370d5cd63279efa6d502dab29171933a610af", // CRV
	chainId.ArbitrumChainName:  "0x11cdb42b0eb46d95f990bedd4695a6e3fa034978", // CRV
}

// Same as CurveTokenlist.
var CRVList = CurveTokenlist

// Curve USD: crvUSD.
//
// map[network] = address.
var CrvUSDList = map[string]string{
	chainId.EthereumChainName: "0xf939E0A03FB07F59A73314E73794Be0E57ac1b4E", // crvUSD, 0xf939e0a03fb07f59a73314e73794be0e57ac1b4e
}

var CurveSTETHETHRouterAddress string = "0xDC24316b9AE028F1497c275EB9192a3Ea0f67022"

// Curve folded pools(a pool folded into other pools).
//
// map[network] = []address.
var CurveFoldedPoolList = map[string][]string{
	chainId.EthereumChainName: {
		"0x6c3f90f043a72fa612cbac8115ee7e52bde6e490", // 3crv
		"0x075b1bb99792c9e1041ba13afef80c91a1e70fb3", // 3btc
	},
	chainId.AvalancheChainName: {
		"0x1337bedc9d22ecbe766df105c9623922a27963ec", // av3CRV
	},
}

// Curve gauage factory.
//
// map[network] = address.
var CurveGaugeFactoryList = map[string]string{
	chainId.AvalancheChainName: "0xabc000d88f23bb45525e447528dbf656a9d55bf5",
}

// Curve pools.
//
// map[network][pool name] = address.
var CurvePoolsList = map[string]map[string]string{
	chainId.EthereumChainName: {
		"ib3CRV":      "0x5282a4ef67d9c33135340fb3289cc1711c13638c", // cyDAI/cyUSDC/cyUSDT
		"eCRV":        "0xa3d87fffce63b53e0d54faa1cc983b7eb0b74a9c", // ETH/sETH
		"steCRV":      "0x06325440d014e39736583c165c2963ba99faf14e", // ETH/stETH
		"crvRenWSBTC": "0x075b1bb99792c9e1041ba13afef80c91a1e70fb3", // renBTC/wBTC/sBTC
		"crvRenWBTC":  "0x49849c98ae39fff122806c06791fa73784fb3675", // renBTC/wBTC
		"oBTC":        "0x2fe94ea3d5d4a175184081439753de15aef9d614", // oBTC/sbtcCRV
		"pBTC":        "0xde5331ac4b3630f94853ff322b66407e0d6331e8", // pBTC/sbtcCRV
		"tBTC":        "0x64eda51d3ad40d56b9dfc5554e06f94e1dd786fd", // tBTC/sbtcCRV
		"bBTC":        "0x410e3e86ef427e30b9235497143881f717d93c2a", // bBTC/sbtcCRV
		"FRAX3CRV-f":  "0xd632f22692fac7611d2aa1c0d552930d43caed3b", // FRAX/3CRV
		"LUSD3CRV-f":  "0xed279fdd11ca84beef15af5d39bb4d4bee23f0ca", // LUSD/3CRV
		"saCRV":       "0x02d341ccb60faaf662bc0554d13778015d1b285c", // aDAI/aSUSD
		"yveCRV-DAO":  "0xc5bddf9843308380375a611c18b50fb9341f502a", // CRV/3CRV
	},
}

// Curve gauges using different contracts.
//
// map[network][gauge type] = []address.
var CurveGauges = map[string]map[string][]string{
	chainId.EthereumChainName: {
		"liquidity gauge": {
			"0x7ca5b0a2910B33e9759DC7dDB0413949071D7575", // compound
			"0xBC89cd85491d81C6AD2954E6d0362Ee29fCa8F53", // usdt
			"0xFA712EE4788C042e2B7BB55E6cb8ec569C4530c1", // y
			"0x69Fb7c45726cfE2baDeE8317005d3F94bE838840", // busd
			"0x64E3C23bfc40722d3B649844055F1D51c1ac041d", // pax
			"0xB1F2cdeC61db658F091671F5f199635aEF202CAC", // ren
			"0x4c18E409Dc8619bFb6a1cB56D114C3f592E0aE79", // hbtc
			"0xbFcF63294aD7105dEa65aA58F8AE5BE2D9d0952A", // 3pool
			"0xC5cfaDA84E902aD92DD40194f0883ad49639b023", // gusd
			"0x2db0E83599a91b508Ac268a6197b8B14F5e72840", // husd
			"0xC2b1DF84112619D190193E48148000e3990Bf627", // usdk
			"0xF98450B5602fa59CC66e1379DFfB6FDDc724CfC4", // usdn
			"0xd8b712d29381748dB89c36BCa0138d7c75866ddF", // mim
			"0x903dA6213a5A12B61c821598154EfAd98C3B20E4", // f-cvxcrv
			"0xeFF437A56A22D7dD86C1202A308536ED8C7da7c1", // f-ibjpy
			"0x63d9f3aB7d0c528797A12a0684E50C397E9e79dC", // f-ibgbp
			"0x05ca5c01629a8E5845f12ea3A03fF7331932233A", // f-ibaud
			"0x99fb76F75501039089AAC8f20f487bf84E51d76F", // f-ibeur
			"0x2fA53e8fa5fAdb81f4332C8EcE39Fe62eA2f919E", // f-ibchf
			"0x1750a3a3d80A3F5333BBe9c4695B0fAd41061ab1", // f-ibkrw
			"0x25f0cE4E2F8dbA112D9b115710AC297F816087CD", // ousd
			"0x12dCD9E8D1577b5E4F066d8e7D404404Ef045342", // f-aleth
			"0xD9277b0D007464eFF133622eC0d42081c93Cef02", // f-eurn
			"0x9AF13a7B1f1Bbf1A2B05c6fBF23ac23A9E573b4E", // f-usdm
			"0xB518f5e3242393d4eC792BD3f44946A3b98d0E48", // f-ust-mim
			"0x16C2beE6f55dAB7F494dBa643fF52ef2D47FBA36", // f-d3pool
			"0xC95bdf13A08A547E4dD9f29B00aB7fF08C5d093d", // f-usdpax
			"0xb0f5d00e5916c8b8981e99191A1458704B587b2b", // f-ustw
			"0x346C7BB1A7a6A30c8e81c14e90FC2f0FBddc54d8", // f-ibbtc
			"0x8Fa728F393588E8D8dD1ca397E9a710E53fA553a", // f-dola
			"0x29284d30bcb70e86a6C3f84CbC4de0Ce16b0f1CA", // f-tbtc2
			"0x1E212e054d74ed136256fc5a5DDdB4867c6E003F", // f-ageur
			"0xdC69D4cB5b86388Fff0b51885677e258883534ae", // f-fei
			"0x009aCD89535DAbC270C93F9b39D3232105Fef453", // f-fxseth
			"0x02246583870b36Be0fEf2819E1d3A771d6C07546", // f-badgerwbtc
			"0x05255C5BD33672b9FEA4129C13274D1E6193312d", // f-yfieth
			"0x38039dD47636154273b287F74C432Cac83Da97e2", // f-ageuribeur
			"0xb07d00e0eE9b1b2eb9f1B483924155Af7AF0c8Fa", // f-pwrd
			"0xE786Df7076AFeECC3faCD841ED4AD20d0F04CF19", // f-csdcusdc
			"0x60355587a8D4aa67c2E64060Ab36e566B9bCC000", // f-sdteth
			"0xB81465Ac19B9a57158a79754bDaa91C60fDA91ff", // f-dydxeth
			"0x82d0aDea8C4CF2fc84A499b568F4C1194d63113d", // f-aavepalstkaave
			"0xAB1927160EC7414C6Fa71763E2a9f3D107c126dd", // f-crvfxs
			"0x8aD7e0e6EDc61bC48ca0DD07f9021c249044eD30", // f-rocketethwsteth
			"0x5AC6886Edd18ED0AD01C0B0910660637c551FBd6", // f-btrflyeth
			"0xB5efA93d5D23642f970aF41a1ea9A26f19CbD2Eb", // f-pbtc
			"0x784342E983E9283A7108F20FcA21995534b3fE65", // f-silofrax
			"0x95d16646311fDe101Eb9F897fE06AC881B7Db802", // f-stgusdc
			"0x6F98dA2D5098604239C07875C6B7Fd583BC520b9", // f-bean
			"0x46521Db0D31A62A2CBF8D1A7Cdc6bBBBC441A1fc", // f-lfteth
			"0xeCb860e54E33FEA8fAb5B076734e2591D1A9ebA4", // f-seth2steth
			"0xFA49B2a5D9E77f6748bf05801aa22356D514137b", // f-jpegeth
			"0xdB7cbbb1d5D5124F86E92001C9dFDC068C05801D", // f-fpifrax
			"0x1779AEB087C5BdBe48749ab03575f5f25D1DEeaF", // f-ibaudusdc
			"0x36C66bC294fEf4e94B3e40A1801d0AB0085Fe96e", // f-ibchfusdc
			"0xE1D520B1263D6Be5678568BD699c84F7f9086023", // f-ibeurusdc
			"0x1Ba86c33509013c937344f6e231DA2E63ea45197", // f-ibgbpusdc
			"0x3A748A2F4765BDFB119Cb7143b884Db7594a68c3", // f-ibjpyusdc
			"0xb6d7C2bda5a907832d4556AE5f7bA800FF084C2a", // f-ibkrwusdc
			"0x6d3328F0333f6FB0B2FaC87cF5a0FFa7e77beB60", // f-kp3reth
			"0x89664D561E79Ca22Fd2eA4076b3e5deF0b219C15", // f-pusd
			"0x8dF6FdAe05C9405853dd4cF2809D5dc2b5E77b0C", // f-ometh
			"0x4fb13b55D6535584841dbBdb14EDC0258F7aC414", // f-paleth
			"0xa9A9BC60fc80478059A83f516D5215185eeC2fc0", // f-sdfxsfxs
			"0x03fFC218C7A9306D21193565CbDc4378952faA8c", // f-sdagag
			"0x663FC22e92f26C377Ddf3C859b560C4732ee639a", // f-sdcrvcrv
			"0xd5d3efC90fFB38987005FdeA303B68306aA5C624", // f-usdd3crv
			"0xf668e6d326945d499e5b35e7cd2e82acfbcfe6f0", // f-stethconcentrated
			"0xa0C08C0Aede65a0306F7dD042D2560dA174c91fC", // f-tokeeth
		},
		"staking liquidity gauge": {
			"0xA90996896660DEcC6E997655E065b23788857849", // susdv2
			"0x705350c4BcD35c9441419DdD5d2f097d7a55410F", // sbtc
			"0x5f626c30EC1215f4EdCc9982265E8b1F411D1352", // musd
			"0x6828bcF74279eE32f2723eC536c22c51Eed383C6", // tbtc
			"0x4dC4A289a8E33600D8bD4cf5F6313E43a37adec7", // rsv
			"0xAEA6c312f4b3E04D752946d329693F7293bC2e6D", // dusd
		},
		"liquidity gauge v2": {
			"0xd7d147c6Bb90A718c3De8C0568F9B560C79fa416", // pbtc
			"0xdFc7AdFa664b08767b735dE28f9E84cd30492aeE", // bbtc
			"0x11137B10C210b579405c21A07489e28F3c040AB1", // obtc
			"0x3B7020743Bc2A4ca9EaF9D0722d42E20d6935855", // ust
			"0x90Bb609649E0451E5aD952683D64BD2d1f245840", // eurs
			"0x3C0FFFF15EA30C35d7A85B85c0782D6c94e1d238", // seth
			"0xd662908ADA2Ea1916B3318327A97eB18aD588b5d", // aave
			"0x182B723a58739a9c974cFDB385ceaDb237453c28", // steth
			"0x462253b8F74B72304c145DB0e4Eebd326B22ca39", // saave
			"0x6d10ed2cF043E6fcf51A0e7b4C2Af3Fa06695707", // ankreth
			"0xF5194c3325202F456c95c1Cf0cA36f8475C1949F", // ib
			"0xFD4D8a17df4C27c1dD245d153ccf4499e806C87D", // link
			"0x055be5DDB7A925BfEF3417FC157f53CA77cA7222", // usdp
			"0x359FD5d6417aE3D8D6497d9B2e7A890798262BA4", // tusd
			"0xd4B22fEdcA85E684919955061fDf353b9d38389b", // busdv2
			"0x72E158d38dbd50A483501c24f792bDAAA3e7D55C", // frax
			"0x9B8519A9a00100720CCdC8a120fBeD319cA47a14", // lusd
		},
		"liquidity gauge v3": {
			"0x824F13f1a2F29cFEEa81154b46C0fc820677A637", // reth
			"0x9582C4ADACB3BCE56Fea3e590F05c3ca2fb9C477", // alusd
			"0x6955a55416a06839309018A8B0cB72c4DDC11f15", // tricrypto
			"0xDeFd8FdD20e0f34115C7018CCfb655796F6B2168", // tricrypto2
			"0xe8060Ad8971450E624d5289A10017dD30F5dA85F", // eurt
		},
		"root-chain gauge": {
			"0xC48f4653dd6a9509De44c92beb0604BEA3AEe714", // polygon-a3crv
			"0xb9C05B8EE41FDCbd9956114B3aF15834FDEDCb54", // fantom-2pool
			"0xfE1A3dD8b169fB5BF0D5dbFe813d956F39fF6310", // fantom-geist
			"0x488E6ef919C2bB9de535C634a80afb0114DA8F62", // polygon-ren
			"0xfDb129ea4b6f557b07BcDCedE54F665b7b6Bc281", // fantom-ren
			"0x060e386eCfBacf42Aa72171Af9EFe17b3993fC4F", // polygon-atricrypto
			"0x6C09F6727113543Fd061a721da512B7eFCDD0267", // xdai-3pool
			"0x9044E12fB1732f88ed0c93cfa5E9bB9bD2990cE5", // arbitrum-crypto
			"0xFf17560d746F85674FE7629cE986E949602EF948", // arbitrum-2pool
			"0x9F86c5142369B1Ffd4223E5A2F2005FC66807894", // arbitrum-ren
			"0x260e4fBb13DD91e187AE992c3435D0cf97172316", // famtom-tricrypto
			"0xB504b6EB06760019801a91B451d3f7BD9f027fC9", // avalanche-a3crv
			"0x75D05190f35567e79012c2F0a02330D3Ed8a1F74", // avalanche-ren
			"0xa05E565cA0a103FcD999c7A7b8de7Bd15D5f6505", // avalanche-atricrypto
			"0xf2Cde8c47C20aCbffC598217Ad5FE6DB9E00b163", // harmony-3pool
			"0x56eda719d82aE45cBB87B7030D3FB485685Bea45", // arbitrum-eursusd
			"0xAF78381216a8eCC7Ad5957f3cD12a431500E0B0D", // polygon-eurtusd
		},
		"liquidity gauge v4": {
			"0x65CA7Dc5CB661fC58De57B1E1aF404649a27AD35", // eursusd
			"0x4Fd86Ce7Ecea88F7E0aA78DC12625996Fb3a04bC", // eurtusd
			"0x1cEBdB0856dd985fAe9b8fEa2262469360B8a3a6", // crveth
			"0x66ec719045bBD62db5eBB11184c18237D3Cc2E62", // rai
			"0x7E1444BA99dcdFfE8fBdb42C02F0005D14f13BE1", // cvxeth
			"0x1B3E14157ED33F60668f2103bCd5Db39a1573E5B", // xautusd
			"0x08380a4999Be1a958E2abbA07968d703C7A3027C", // spelleth
			"0x6070fBD4E608ee5391189E7205d70cc4A274c017", // teth
		},
		"liquidity gauge v5": {
			"0x34883134A39B206A451c2D3B0E7Cac44BE4D9181", // 4pool
			"0x9f330Db38caAAe5B61B410e2f0aaD63fff2109d8", // 2pool
		},
		"Root Liquidity Gauge Implementation": {
			"0x279f11F8E2825dbe0b00F6776376601AC948d868", // fantom-f-FRAX2pool3CRV
			"0x95069889DF0BCdf15bc3182c1A4D6B20631F3B46", // fantom-f-MAI3Pool
			"0x1c77fB5486545810679D53E325d5bCf6C6A45081", // fantom-f-3poolV2
			"0xA6ff75281eACa4cD5fEEb333e8E15558208295e5", // fantom-f-USDL-3CRV
			"0xc1c5B8aAfE653592627B54B9527C7E98326e83Ff", // fantom-f-FTM/FTML
			"0x9562c4D2E06aAf85efC5367Fb4544ECeB788465E", // fantom-f-UST 3pool
			"0xd0698b2E41C42bcE42B51f977F962Fd127cF82eA", // fantom-f-4POOL
			"0xF7b9c402c4D6c2eDbA04a7a515b53D11B1E9b2cc", // fantom-f-g3CRV-gauge
			"0x15bB164F9827De760174d3d3dAD6816eF50dE13c", // fantom-f-DAI+USDC-gauge
			"0xbC38bD19227F91424eD4132F630f51C9A42Fa338", // fantom-f-btcCRV-gauge
			"0x319E268f0A4C85D404734ee7958857F5891506d7", // fantom-f-crv3crypto-gauge
			"0xdA690c2EA49a058a9966C69f46a05Bfc225939f4", // polygon-f-TUSDam3CRV3CRV
			"0x18006c6A7955Bf6Db72DE34089B975f733601660", // polygon-f-crvEURSUSD-gauge
			"0x20759F567BB3EcDB55c817c9a1d13076aB215EdC", // polygon-f-am3CRV-gauge
			"0xBb1B19495B8FE7C402427479B9aC14886cbbaaeE", // polygon-f-crvUSDBTCETH-gauge
			"0x8D9649e50A0d1da8E939f800fB926cdE8f18B47D", // polygon-f-btcCRV-gauge
			"0x8b397084699Cc64E429F610F81Fac13bf061ef55", // polygon-f-crvEURTUSD-gauge
			"0xD1426C391A7Cbe9DeCd302ac9c44e65C3505d1f0", // polygon-f-aMATICb
			"0xF2dDF89C04d702369Ab9eF8399Edb99a76e951Ce", // arbitrum-f-VSTFRAX
			"0xCE5F24B7A95e9cBa7df4B54E911B4A3Dc8CDAf6f", // arbitrum-f-2CRV-gauge
			"0xDB3fd1bfC67b5D4325cb31C04E0Cae52f1787FD6", // arbitrum-f-btcCRV-gauge
			"0x555766f3da968ecBefa690Ffd49A2Ac02f47aa5f", // arbitrum-f-crv3crypto-gauge
			"0x6339eF8Df0C2d3d3E7eE697E241666a916B81587", // arbitrum-f-crvEURSUSD-gauge
			"0x94A5E05D66834c6C6961E199D34dA576679fC187", // arbitrum-f-4pool
			"0x1AEAA1b998307217D62E9eeFb6407B10598eF3b8", // avalanche-f-3pool
			"0xfbb5b8f2f9b7a4d21ff44dC724C1Fb7b531A6612", // avalanche-f-AVAX/AVAXL
			"0x00f7d467ef51e44f11f52a0c0bef2e56c271b264", // avalanche-f-crv amWBTC/renBTC-gauge
			"0x4620d46b4db7fb04a01a75ffed228bc027c9a899", // avalanche-f-av3CRV-gauge
			"0x1879075f1c055564cb968905ac404a5a01a1699a", // avalanche-f-crvUSDBTCETH-gauge
			"0xc5aE4B5F86332e70f3205a8151Ee9eD9F71e0797", // optimism-f-sUSD3CRV
			"0xF4eA7617E7999710244e2eAbfC8730d35482EE76", // optimism-f-4pool
			"0xB721Cc32160Ab0da2614CC6aB16eD822Aeebc101", // xdai-f-x3CRV-gauge
		},
	},
}
