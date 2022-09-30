package controller

import (
	"net/http"
	"stockpro/models"
	"stockpro/service"

	"github.com/labstack/echo/v4"
)

type InvestmentController struct {
	Service service.ServiceInvestment
}

//function to get top 5 investment opportinties
func (h InvestmentController) GetInvestmentListByType(c echo.Context) error {

	invType := c.QueryParam("invType")

	val := h.Service.GetInvestmentListByType(invType)

	var r models.Response
	r.Code = http.StatusOK
	r.Status = "Success"
	r.Data = val
	return c.JSON(200, r)
}

//func to get portfolio of investments based on required RoI
func (h InvestmentController) GetInvOpportunities(c echo.Context) error {
	var request models.InvRequest
	if err := c.Bind(&request); err != nil {
		return err
	}
	res := h.Service.GetInvOpportunities(request)

	var r models.Response
	r.Code = http.StatusOK
	r.Status = "Success"
	r.Data = res
	return c.JSON(200, r)
}
