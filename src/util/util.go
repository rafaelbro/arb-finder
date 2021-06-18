package util

import "math/big"

var Tokens map[string]string = map[string]string{
	"CAKE": "0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82",
	"DAI":  "0x1af3f329e8be154074d8769d1ffa4ee058b1dbc3",
	"ETH":  "0x2170ed0880ac9a755fd29b2688956bd959f933f8",
	"ADA":  "0x3ee2200efb3400fabb9aacf31297cbdd1d435d47",
	"VAI":  "0x4bd17003473389a42daf6a0a729f6fdb328bbbd7",
	"USDT": "0x55d398326f99059ff775485246999027b3197955",
	"BTCB": "0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c",
	"USDC": "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d",
	"WBNB": "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",
	"BUSD": "0xe9e7cea3dedca5984780bafc599bd69add087d56",
}

var ExchangesMap map[string]int64 = map[string]int64{
	"PANCAKESWAP":    0,
	"PANCAKESWAP_V2": 1,
	"APESWAP":        2,
	"WAULTSWAP":      3,
	"BAKERYSWAP":     4,
	"MDEX":           5,
	"ACRYPTOS_CORE":  6,
	"ACRYPTOS_META":  7,
	"ELLIPSIS_CORE":  8,
	"ELLIPSIS_META":  9,
}

var ContractAddress = "0xFA714b0381fDDA77229f0C801AdF27ff2dE6F0fb"

func ConvertToCryptoValue(value int64) *big.Int {
	landingZeros, exponent := big.NewInt(10), big.NewInt(18)
	landingZeros.Exp(landingZeros, exponent, nil)
	valueBig := big.NewInt(value)
	valueBig.Mul(valueBig, landingZeros)
	return valueBig
}

var Pairs map[string]map[string]string = map[string]map[string]string{
	"CAKE": {
		"WBNB": "0x0eD7e52944161450477ee417DE9Cd3a859b14fD0",
		"BUSD": "0x804678fa97d91B974ec2af3c843270886528a9E6",
	},
	"DAI": {
		"USDT": "0xf6f5CE9a91Dd4FAe2d2eD92E25F2A4dc8564F174",
		"USDC": "0xadBba1EF326A33FDB754f14e62A96D5278b942Bd",
		"WBNB": "0xc7c3cCCE4FA25700fD5574DA7E200ae28BBd36A3",
		"BUSD": "0x66FDB2eCCfB58cF098eaa419e5EfDe841368e489",
	},
	"ETH": {
		"USDT": "0x531FEbfeb9a61D948c384ACFBe6dCc51057AEa7e",
		"BTCB": "0xD171B26E4484402de70e3Ea256bE5A2630d7e88D",
		"USDC": "0xEa26B78255Df2bBC31C1eBf60010D78670185bD0",
		"WBNB": "0x74E4716E431f45807DCF19f284c7aA99F18a4fbc",
		"BUSD": "0x7213a321F1855CF1779f42c0CD85d3D95291D34C",
	},
	"VAI": {
		"BUSD": "0x133ee93FE93320e1182923E1a640912eDE17C90C",
	},
	"USDT": {
		"USDC": "0xEc6557348085Aa57C72514D67070dC863C0a5A8c",
		"WBNB": "0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE",
		"BUSD": "0x7EFaEf62fDdCCa950418312c6C91Aef321375A00",
	},
	"BTCB": {
		"WBNB": "0x61EB789d75A95CAa3fF50ed7E47b96c132fEc082",
		"BUSD": "0xF45cd219aEF8618A92BAa7aD848364a158a24F33",
	},
	"USDC": {
		"WBNB": "0xd99c7F6C65857AC913a8f880A4cb84032AB2FC5b",
		"BUSD": "0x2354ef4DF11afacb85a5C7f98B624072ECcddbB1",
	},
	"WBNB": {
		"BUSD": "0x58F876857a02D6762E0101bb5C46A8c1ED44Dc16",
	},
}

var TradeQuantity map[string]int64 = map[string]int64{
	"CAKE": 500,
	"DAI":  10000,
	"ETH":  4,
	"ADA":  0,
	"VAI":  10000,
	"USDT": 10000,
	"BTCB": 1,
	"USDC": 10000,
	"WBNB": 30,
	"BUSD": 10000,
}
