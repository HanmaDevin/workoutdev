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
	e.GET("/workouts", func(c echo.Context) error {
		return GetAllWorkouts(c, db)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
