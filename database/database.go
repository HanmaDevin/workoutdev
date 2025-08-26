package database

import (
	"log"

	"github.com/HanmaDevin/workoutdev/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("workout.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&types.User{}, &types.Workout{}, &types.Exercise{}, &types.Set{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	seedExercises()
}

func seedExercises() {
	var count int64
	DB.Model(&types.Exercise{}).Count(&count)
	if count == 0 {
		for _, exercise := range Exercises {
			DB.Create(&exercise)
		}
	}
}
