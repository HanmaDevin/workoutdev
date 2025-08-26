package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/HanmaDevin/workoutdev/database"
	"github.com/HanmaDevin/workoutdev/server/middleware"
	"github.com/HanmaDevin/workoutdev/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte("your-secret-key")

func NewServer() *echo.Echo {
	e := echo.New()

	// Public routes
	e.POST("/register", registerUser)
	e.POST("/login", loginUser)

	// Authenticated routes
	r := e.Group("/workouts")
	r.Use(middleware.AuthMiddleware)
	r.GET("", getWorkouts)
	r.GET("/:id", getWorkout)
	r.POST("", createWorkout)
	r.POST("/:id/exercises", addExercisesToWorkout)
	r.POST("/sets", addSetToWorkout)
	r.POST("/:id/comments", addCommentsToWorkout)
	r.DELETE("/:id", deleteWorkout)

	return e
}

func StartServer(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":8080"))
}

func registerUser(c echo.Context) error {
	var user types.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := database.RegisterUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	output := fmt.Sprintf("%s registered successfully", user.FirstName)
	return c.JSON(http.StatusCreated, output)
}

func loginUser(c echo.Context) error {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&creds); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := database.LoginUser(creds.Email, creds.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not generate token")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("Welcome, %s!", user.FirstName),
		"token":   tokenString,
	})
}

func getWorkouts(c echo.Context) error {
	userID := c.Get("userID").(uint)
	workouts, err := database.GetWorkouts(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, workouts)
}

func getWorkout(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	userID := c.Get("userID").(uint)
	workout, err := database.GetWorkoutByID(uint(id), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Workout not found")
	}
	return c.JSON(http.StatusOK, workout)
}

func createWorkout(c echo.Context) error {
	var workout types.Workout
	if err := c.Bind(&workout); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	userID := c.Get("userID").(uint)
	workout.UserID = userID
	if err := database.CreateWorkout(&workout); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, workout)
}

func addCommentsToWorkout(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Workout ID")
	}

	var body struct {
		Comments []string `json:"comments"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := database.AddCommentsToWorkout(uint(id), body.Comments); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "Comments added to workout")
}

func addExercisesToWorkout(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Workout ID")
	}

	var body struct {
		ExerciseNames []string `json:"exercise_names"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	for _, name := range body.ExerciseNames {
		if !database.ExerciseNames[name] {
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid exercise name: %s", name))
		}
	}

	if err := database.AddExercisesToWorkout(uint(id), body.ExerciseNames); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "Exercises added to workout")
}

func addSetToWorkout(c echo.Context) error {
	var set types.Set
	if err := c.Bind(&set); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := database.AddSetToWorkout(&set); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, set)
}

func deleteWorkout(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	userID := c.Get("userID").(uint)
	if err := database.DeleteWorkout(uint(id), userID); err != nil {
		return c.JSON(http.StatusNotFound, "Workout not found")
	}
	return c.NoContent(http.StatusNoContent)
}
