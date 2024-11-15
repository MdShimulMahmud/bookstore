package containers

import (
	"fmt"
	"go-bootcamp/pkg/config"
	"go-bootcamp/pkg/connection"
	"go-bootcamp/pkg/controllers"
	"go-bootcamp/pkg/repositories"
	"go-bootcamp/pkg/routes"
	"go-bootcamp/pkg/services"
	"log"

	"github.com/labstack/echo/v4"
)


func Serve(e *echo.Echo) {


    config.SetConfig()

    db := connection.GetDB()

    bookRepo := repositories.BookDBInstance(db)

    bookService := services.BookServiceInstance(bookRepo)

    controllers.SetBookService(bookService)

    routes.BookRoutes(e)

    log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}