package service

import (
	"errors"
	"fmt"
	"math"
	"stockpro/models"
)

type ServiceInvestment struct {
}

//function to get top 5 investment opportinties
func (t ServiceInvestment) GetInvestmentListByType(invtype string) (map[string][]models.Investment, error) {
	inv := models.Investments

	var res []models.Investment
	if invtype != "" {

		for _, val := range inv {
			if string(val.InvType) == invtype {
				res = append(res, val)
			}
		}
	} else {
		res = inv
	}
	investment := make(map[string][]models.Investment)
	ret1Yr := GetYearReturns(res, 1)
	ret3Yr := GetYearReturns(res, 3)
	ret5Yr := GetYearReturns(res, 5)

	investment["1 Year Return"] = ret1Yr
	investment["3 Year Return"] = ret3Yr
	investment["5 Year Return"] = ret5Yr

	if len(ret1Yr) == 0 && len(ret3Yr) == 0 && len(ret5Yr) == 0 {
		return nil, errors.New("no investment suggestions available")
	}
	return investment, nil

}

//func to get investment opportunities based on No.of years
func GetYearReturns(inv []models.Investment, invYear int) []models.Investment {
	switch {
	case invYear == 1:
		for i := 0; i < len(inv); i++ {
			for j := i + 1; j < len(inv); j++ {
				if inv[i].ReturnFor1yr < inv[j].ReturnFor1yr {
					temp := inv[i]
					inv[i] = inv[j]
					inv[j] = temp
				}
			}
		}

	case invYear == 3:
		for i := 0; i < len(inv); i++ {
			for j := i + 1; j < len(inv); j++ {
				if inv[i].ReturnFor3yr < inv[j].ReturnFor3yr {
					temp := inv[i]
					inv[i] = inv[j]
					inv[j] = temp
				}
			}
		}
	case invYear == 5:
		for i := 0; i < len(inv); i++ {
			for j := i + 1; j < len(inv); j++ {
				if inv[i].ReturnFor5yr < inv[j].ReturnFor5yr {
					temp := inv[i]
					inv[i] = inv[j]
					inv[j] = temp
				}
			}
		}
	default:
		fmt.Println("Invalid")
	}

	res := make([]models.Investment, 5)

	for k := 0; k < 5; k++ {
		res[k] = inv[k]
	}
	return res
}

//func to get portfolio of investments based on required RoI
func (t ServiceInvestment) GetInvOpportunities(request models.InvRequest) (map[string][][]models.Opportunities, error) {
	inv := models.Investments

	var res []models.Investment
	if request.InvType != "" {

		for _, val := range inv {
			if string(val.InvType) == request.InvType {
				res = append(res, val)
			}
		}
	} else {
		res = inv
	}
	investment := make(map[string][][]models.Opportunities)
	ret1Yr := GetInvOpportunitiesByType(res, request, 1)     //func call to get investment opportunities based on 1 year return
	ret3Yr := GetInvOpportunitiesByType(res, request, 3)     //func call to get investment opportunities based on 3 year return
	ret5Yr := GetInvOpportunitiesByType(res, request, 5)     //func call to get investment opportunities based on 5 year return
	lastYrRet := GetInvOpportunitiesByType(res, request, -1) //func call to get investment opportunities based on LastTwele Months return

	if len(ret1Yr) == 0 && len(ret3Yr) == 0 && len(ret5Yr) == 0 && len(lastYrRet) == 0 {
		return nil, errors.New("no investment suggestions available")
	}
	investment["1 Year Return"] = ret1Yr
	investment["3 Year Return"] = ret3Yr
	investment["5 Year Return"] = ret5Yr
	investment["Last Year Return"] = lastYrRet

	if investment == nil {
		return nil, errors.New("no investment suggestions available")
	}

	return investment, nil

}

//func to get portfolio of investments based on required RoI for required No.of years
func GetInvOpportunitiesByType(inv []models.Investment, req models.InvRequest, invYear int) [][]models.Opportunities {

	var units float64
	var units1 float64
	var units2 float64
	var opportunities [][]models.Opportunities
	switch {
	case invYear == 1:

		for i := 0; i < len(inv); i++ {
			for j := i; j < len(inv); j++ {
				if i == j {
					if inv[i].ReturnFor1yr == req.ERoI {
						if req.InvType == "Stocks" {
							units = math.Trunc(req.Amount / inv[i].UnitPrice)
						} else {
							units = req.Amount / inv[i].UnitPrice
							units = TwoDecimalPoint(units)
						}

						newInvOpp := models.Opportunities{
							InvType: string(inv[i].InvType),
							Name:    inv[i].Name,
							RoI:     inv[i].ReturnFor1yr,
							Units:   units,
						}
						opportunities = append(opportunities, []models.Opportunities{newInvOpp})
					}
				} else {
					if (inv[i].ReturnFor1yr+inv[j].ReturnFor1yr)/2 >= req.ERoI-0.5 && (inv[i].ReturnFor1yr+inv[j].ReturnFor1yr)/2 <= req.ERoI+0.5 {
						if req.InvType == "Stocks" {
							units1 = math.Trunc((req.Amount / 2) / inv[i].UnitPrice)
						} else {
							units1 = (req.Amount / 2) / inv[i].UnitPrice
							units1 = TwoDecimalPoint(units1)
						}
						newInvOpp1 := models.Opportunities{
							InvType: string(inv[i].InvType),
							Name:    inv[i].Name,
							RoI:     inv[i].ReturnFor1yr,
							Units:   units1,
						}
						if req.InvType == "Stocks" {
							units2 = math.Trunc((req.Amount / 2) / inv[j].UnitPrice)
						} else {
							units2 = (req.Amount / 2) / inv[j].UnitPrice
							units2 = TwoDecimalPoint(units2)
						}
						newInvOpp2 := models.Opportunities{
							InvType: string(inv[j].InvType),
							Name:    inv[j].Name,
							RoI:     inv[j].ReturnFor1yr,
							Units:   units2,
						}
						opportunities = append(opportunities, []models.Opportunities{newInvOpp1, newInvOpp2})
					}
				}
			}
		}

	case invYear == 3:
		for i := 0; i < len(inv); i++ {
			for j := i; j < len(inv); j++ {
				if i == j {
					if inv[i].ReturnFor3yr == req.ERoI {
						if req.InvType == "Stocks" {
							units = math.Trunc(req.Amount / inv[i].UnitPrice)
						} else {
							units = req.Amount / inv[i].UnitPrice
							units = TwoDecimalPoint(units)
						}
						newInvOpp := models.Opportunities{
							InvType: string(inv[i].InvType),
							Name:    inv[i].Name,
							RoI:     inv[i].ReturnFor3yr,
							Units:   units,
						}
						opportunities = append(opportunities, []models.Opportunities{newInvOpp})
					}
				} else {
					if (inv[i].ReturnFor3yr+inv[j].ReturnFor3yr)/2 >= req.ERoI-0.5 && (inv[i].ReturnFor3yr+inv[j].ReturnFor3yr)/2 <= req.ERoI+0.5 {
						if req.InvType == "Stocks" {
							units1 = math.Trunc((req.Amount / 2) / inv[i].UnitPrice)
						} else {
							units1 = (req.Amount / 2) / inv[i].UnitPrice
							units1 = TwoDecimalPoint(units1)
						}
						newInvOpp1 := models.Opportunities{
							InvType: string(inv[i].InvType),
							Name:    inv[i].Name,
							RoI:     inv[i].ReturnFor3yr,
							Units:   units1,
						}
						if req.InvType == "Stocks" {
							units2 = math.Trunc((req.Amount / 2) / inv[j].UnitPrice)
						} else {
							units2 = (req.Amount / 2) / inv[j].UnitPrice
							units2 = TwoDecimalPoint(units2)
						}
						newInvOpp2 := models.Opportunities{
							InvType: string(inv[j].InvType),
							Name:    inv[j].Name,
							RoI:     inv[j].ReturnFor3yr,
							Units:   units2,
						}
						opportunities = append(opportunities, []models.Opportunities{newInvOpp1, newInvOpp2})
					}
				}
			}
		}

	case invYear == 5:
		for i := 0; i < len(inv); i++ {
			for j := i; j < len(inv); j++ {
				if i == j {
					if inv[i].ReturnFor5yr == req.ERoI {
						if req.InvType == "Stocks" {
							units = math.Trunc(req.Amount / inv[i].UnitPrice)
						} else {
							units = req.Amount / inv[i].UnitPrice
							units = TwoDecimalPoint(units)
						}
						newInvOpp := models.Opportunities{
							InvType: string(inv[i].InvType),
							Name:    inv[i].Name,
							RoI:     inv[i].ReturnFor5yr,
							Units:   units,
						}
						opportunities = append(opportunities, []models.Opportunities{newInvOpp})
					}
				} else {
					if (inv[i].ReturnFor5yr+inv[j].ReturnFor5yr)/2 >= req.ERoI-0.5 && (inv[i].ReturnFor5yr+inv[j].ReturnFor5yr)/2 <= req.ERoI+0.5 {

						if req.InvType == "Stocks" {
							units1 = math.Trunc((req.Amount / 2) / inv[i].UnitPrice)
						} else {
							units1 = (req.Amount / 2) / inv[i].UnitPrice
							units1 = TwoDecimalPoint(units1)
						}
						newInvOpp1 := models.Opportunities{
							InvType: string(inv[i].InvType),
							Name:    inv[i].Name,
							RoI:     inv[i].ReturnFor5yr,
							Units:   units1,
						}
						if req.InvType == "Stocks" {
							units2 = math.Trunc((req.Amount / 2) / inv[j].UnitPrice)
						} else {
							units2 = (req.Amount / 2) / inv[j].UnitPrice
							units2 = TwoDecimalPoint(units2)
						}
						newInvOpp2 := models.Opportunities{
							InvType: string(inv[j].InvType),
							Name:    inv[j].Name,
							RoI:     inv[j].ReturnFor5yr,
							Units:   units2,
						}
						opportunities = append(opportunities, []models.Opportunities{newInvOpp1, newInvOpp2})
					}
				}
			}
		}
	case invYear == -1:
		for i := 0; i < len(inv); i++ {
			for j := i; j < len(inv); j++ {
				if i == j {
					if inv[i].LastTwelveMonthReturns == req.ERoI {
						if req.InvType == "Stocks" {
							units = math.Trunc(req.Amount / inv[i].UnitPrice)
						} else {
							units = req.Amount / inv[i].UnitPrice
							units = TwoDecimalPoint(units)
						}
						newInvOpp := models.Opportunities{
							InvType: string(inv[i].InvType),
							Name:    inv[i].Name,
							RoI:     inv[i].LastTwelveMonthReturns,
							Units:   units,
						}
						opportunities = append(opportunities, []models.Opportunities{newInvOpp})
					}
				} else {
					if (inv[i].LastTwelveMonthReturns+inv[j].LastTwelveMonthReturns)/2 >= req.ERoI-0.5 && (inv[i].LastTwelveMonthReturns+inv[j].LastTwelveMonthReturns)/2 <= req.ERoI+0.5 {

						if req.InvType == "Stocks" {
							units1 = math.Trunc((req.Amount / 2) / inv[i].UnitPrice)
						} else {
							units1 = (req.Amount / 2) / inv[i].UnitPrice
							units1 = TwoDecimalPoint(units1)
						}
						newInvOpp1 := models.Opportunities{
							InvType: string(inv[i].InvType),
							Name:    inv[i].Name,
							RoI:     inv[i].LastTwelveMonthReturns,
							Units:   units1,
						}
						if req.InvType == "Stocks" {
							units2 = math.Trunc((req.Amount / 2) / inv[j].UnitPrice)
						} else {
							units2 = (req.Amount / 2) / inv[j].UnitPrice
							units2 = TwoDecimalPoint(units2)
						}
						newInvOpp2 := models.Opportunities{
							InvType: string(inv[j].InvType),
							Name:    inv[j].Name,
							RoI:     inv[j].LastTwelveMonthReturns,
							Units:   units2,
						}
						opportunities = append(opportunities, []models.Opportunities{newInvOpp1, newInvOpp2})
					}
				}
			}
		}

	default:
		fmt.Println("Invalid Option")
	}

	return opportunities

}

//function to round off decimal value to 2 points
func TwoDecimalPoint(units float64) float64 {
	pow := math.Pow(10, float64(2))
	digit := pow * units
	round := math.Floor(digit)
	units = round / pow
	return units
}
