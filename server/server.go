package server

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func NewServer() *echo.Echo {
	e := echo.New()
	return e
}

func StartServer(e *echo.Echo, db *sql.DB) {
	e.Logger.Fatal(e.Start(":8080"))
}
