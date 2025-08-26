package server

import (
	"database/sql"
	"encoding/json"

	"github.com/HanmaDevin/workoutdev/database"
	"github.com/HanmaDevin/workoutdev/types"
	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
)

func GetAllWorkouts(c echo.Context, db *sql.DB) error {
	cookie, err := c.Cookie("user")
	if err != nil {
		c.Logger().Error("Failed to get user cookie:", err)
		return c.JSON(401, map[string]string{"error": "Unauthorized"})
	}
	userID := cookie.Value
	workouts, err := database.GetAllWorkouts(db, userID, "date")
	if err != nil {
		c.Logger().Error("Failed to get workouts:", err)
		return c.JSON(500, map[string]string{"error": err.Error()})
	}
	return c.JSON(200, workouts)
}

func CreateWorkout(c echo.Context, db *sql.DB) error {
	cookie, err := c.Cookie("user")
	if err != nil {
		c.Logger().Error("Failed to get user cookie:", err)
		return c.JSON(401, map[string]string{"error": "Unauthorized"})
	}
	userID := cookie.Value

	var workout types.Workout
	if err := json.NewDecoder(c.Request().Body).Decode(&workout); err != nil {
		c.Logger().Error("Failed to decode workout JSON:", err)
		return c.JSON(400, map[string]string{"error": "Invalid JSON format"})
	}
	workout.UserID = userID

	if err := database.CreateWorkout(db, &workout); err != nil {
		c.Logger().Error("Failed to create workout:", err)
		return c.JSON(500, map[string]string{"error": err.Error()})
	}
	return c.JSON(201, workout)
}
