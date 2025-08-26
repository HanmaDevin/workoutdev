package database

import "github.com/HanmaDevin/workoutdev/types"

func CreateWorkout(workout *types.Workout) error {
	result := DB.Create(workout)
	return result.Error
}

func GetWorkouts(userID uint) ([]types.Workout, error) {
	var workouts []types.Workout
	result := DB.Preload("Exercises").Preload("Sets").Where("user_id = ?", userID).Find(&workouts)
	return workouts, result.Error
}

func GetWorkoutByID(id uint, userID uint) (types.Workout, error) {
	var workout types.Workout
	result := DB.Preload("Exercises").Preload("Sets").Where("id = ? AND user_id = ?", id, userID).First(&workout)
	return workout, result.Error
}

func AddCommentsToWorkout(workoutID uint, comments []string) error {
	var workout types.Workout
	if err := DB.First(&workout, workoutID).Error; err != nil {
		return err
	}

	workout.Comments = append(workout.Comments, comments...)
	return DB.Save(&workout).Error
}

func AddExercisesToWorkout(workoutID uint, exerciseNames []string) error {
	var workout types.Workout
	if err := DB.First(&workout, workoutID).Error; err != nil {
		return err
	}

	var exercises []types.Exercise
	if err := DB.Where("name IN ?", exerciseNames).Find(&exercises).Error; err != nil {
		return err
	}

	return DB.Model(&workout).Association("Exercises").Append(&exercises)
}

func AddSetToWorkout(set *types.Set) error {
	result := DB.Create(set)
	return result.Error
}

func DeleteWorkout(id uint, userID uint) error {
	var workout types.Workout
	if err := DB.Where("id = ? AND user_id = ?", id, userID).First(&workout).Error; err != nil {
		return err
	}
	return DB.Delete(&workout).Error
}
