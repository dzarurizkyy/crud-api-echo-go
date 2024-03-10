package cmd

import (
	"crud_product/config"
	"crud_product/pkg/router"
	database "crud_product/shared/db"

	"github.com/labstack/echo/v4"
)

func RunServer() {
	// Initialize Echo Framework
	e := echo.New();
	g := e.Group("");
	conf := config.GetConfig();

	// Run Web Server
	Apply(e, g, conf);

	// Initialize Port Number
	e.Logger.Fatal(e.Start(":5000"));
};

// Initialize Web Server
func Apply(e *echo.Echo, g *echo.Group, conf config.Configuration) {
	db := database.NewInstanceDB(conf);
	router.NewProductRouter(e, g, db);
};