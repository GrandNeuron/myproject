package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"

	"CalculatorAppFrontendPantela-main/internal/calculationService"
	"CalculatorAppFrontendPantela-main/internal/db"
	"CalculatorAppFrontendPantela-main/internal/handlers"
	"CalculatorAppFrontendPantela-main/internal/web/tasks"
)

func main() {
	dbConn := db.ConnectDB()
	if err := dbConn.AutoMigrate(&calculationService.Calculation{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	repo := calculationService.NewCalculationRepository(dbConn)
	service := calculationService.NewCalculationService(repo)
	handler := handlers.NewTaskHandler(service)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
