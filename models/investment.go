package models

type InvType string

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type Opportunities struct {
	Name    string
	RoI     float64
	InvType string
	Units   float64
}

type InvRequest struct {
	ERoI    float64
	InvType string
	Amount  float64
}

type Investment struct {
	Name                   string
	ReturnFor1yr           float64
	ReturnFor3yr           float64
	ReturnFor5yr           float64
	InvType                InvType
	LastTwelveMonthReturns float64
	UnitPrice              float64
}

const (
	Stocks       InvType = "Stocks"
	Gold                 = "Gold"
	MutualFunds          = "MutualFunds"
	FixedDeposit         = "FixedDeposit"
)

var Investments = []Investment{

	{
		Name:                   "Tata motors",
		ReturnFor1yr:           5,
		ReturnFor3yr:           8,
		ReturnFor5yr:           12,
		InvType:                Stocks,
		LastTwelveMonthReturns: 10,
		UnitPrice:              250,
	},

	{
		Name:                   "JSW Steel",
		ReturnFor1yr:           6,
		ReturnFor3yr:           10,
		ReturnFor5yr:           14,
		InvType:                Stocks,
		LastTwelveMonthReturns: 8,
		UnitPrice:              120,
	},
	{
		Name:                   "Tata Power",
		ReturnFor1yr:           6.4,
		ReturnFor3yr:           9,
		ReturnFor5yr:           11,
		InvType:                Stocks,
		LastTwelveMonthReturns: 12,
		UnitPrice:              280,
	},
	{
		Name:                   "Zomato",
		ReturnFor1yr:           8.5,
		ReturnFor3yr:           10,
		ReturnFor5yr:           16,
		InvType:                Stocks,
		LastTwelveMonthReturns: 10,
		UnitPrice:              70,
	},
	{
		Name:                   "LIC",
		ReturnFor1yr:           4.5,
		ReturnFor3yr:           5,
		ReturnFor5yr:           8,
		InvType:                Stocks,
		LastTwelveMonthReturns: 5,
		UnitPrice:              180,
	},

	{
		Name:                   "ICIC bank",
		ReturnFor1yr:           9,
		ReturnFor3yr:           6,
		ReturnFor5yr:           16.5,
		InvType:                FixedDeposit,
		LastTwelveMonthReturns: 5.5,
		UnitPrice:              200,
	},
	{
		Name:                   "Union Bank",
		ReturnFor1yr:           4,
		ReturnFor3yr:           5.5,
		ReturnFor5yr:           6,
		InvType:                FixedDeposit,
		LastTwelveMonthReturns: 6.5,
		UnitPrice:              50,
	},
	{
		Name:                   "Canara",
		ReturnFor1yr:           5,
		ReturnFor3yr:           6,
		ReturnFor5yr:           6.5,
		InvType:                FixedDeposit,
		LastTwelveMonthReturns: 6,
		UnitPrice:              50,
	},
	{
		Name:                   "Muthoot",
		ReturnFor1yr:           6,
		ReturnFor3yr:           6.5,
		ReturnFor5yr:           7,
		InvType:                FixedDeposit,
		LastTwelveMonthReturns: 7,
		UnitPrice:              250,
	},
	{
		Name:                   "TMB",
		ReturnFor1yr:           3.5,
		ReturnFor3yr:           5.5,
		ReturnFor5yr:           6.8,
		InvType:                FixedDeposit,
		LastTwelveMonthReturns: 6,
		UnitPrice:              50,
	},
	{
		Name:                   "SBI financial",
		ReturnFor1yr:           6.5,
		ReturnFor3yr:           12,
		ReturnFor5yr:           20,
		InvType:                MutualFunds,
		LastTwelveMonthReturns: 8,
		UnitPrice:              25,
	},
	{
		Name:                   "ICICI prudential",
		ReturnFor1yr:           15,
		ReturnFor3yr:           18,
		ReturnFor5yr:           20,
		InvType:                MutualFunds,
		LastTwelveMonthReturns: 14,
		UnitPrice:              25,
	},
	{
		Name:                   "Nippon India",
		ReturnFor1yr:           10,
		ReturnFor3yr:           12,
		ReturnFor5yr:           15,
		InvType:                MutualFunds,
		LastTwelveMonthReturns: 8,
		UnitPrice:              35,
	},
	{
		Name:                   "TATA Digital",
		ReturnFor1yr:           5,
		ReturnFor3yr:           12,
		ReturnFor5yr:           15,
		InvType:                MutualFunds,
		LastTwelveMonthReturns: 11,
		UnitPrice:              25,
	},
	{
		Name:                   "Adity Birla",
		ReturnFor1yr:           8.5,
		ReturnFor3yr:           13,
		ReturnFor5yr:           28,
		InvType:                MutualFunds,
		LastTwelveMonthReturns: 15,
		UnitPrice:              25,
	},
	{
		Name:                   "Gold ETFs",
		ReturnFor1yr:           7,
		ReturnFor3yr:           8,
		ReturnFor5yr:           11,
		InvType:                Gold,
		LastTwelveMonthReturns: 7,
		UnitPrice:              1500,
	},
	{
		Name:                   "Gold Petal",
		ReturnFor1yr:           10,
		ReturnFor3yr:           9,
		ReturnFor5yr:           7,
		InvType:                Gold,
		LastTwelveMonthReturns: 7,
		UnitPrice:              1000,
	},
	{
		Name:                   "Gold mini",
		ReturnFor1yr:           4,
		ReturnFor3yr:           8,
		ReturnFor5yr:           21,
		InvType:                Gold,
		LastTwelveMonthReturns: 7,
		UnitPrice:              500,
	},
	{
		Name:                   "Gold Guinea",
		ReturnFor1yr:           14,
		ReturnFor3yr:           12,
		ReturnFor5yr:           18,
		InvType:                Gold,
		LastTwelveMonthReturns: 13,
		UnitPrice:              2500,
	},
	{
		Name:                   "Gold X",
		ReturnFor1yr:           3.5,
		ReturnFor3yr:           6,
		ReturnFor5yr:           9,
		InvType:                Gold,
		LastTwelveMonthReturns: 7,
		UnitPrice:              3000,
	},
}
