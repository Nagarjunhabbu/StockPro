package main

import (
	"stockpro/controller"
	"stockpro/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes for Investment Opportunities
	e.GET("/api/v1/investment", controller.InvestmentController{Service: service.ServiceInvestment{}}.GetInvestmentListByType)
	e.POST("/api/v1/investment", controller.InvestmentController{Service: service.ServiceInvestment{}}.GetInvOpportunities)

	// Start server
	e.Logger.Fatal(e.Start(":8001"))
}
