package types

import "gorm.io/gorm"

type Exercise struct {
	gorm.Model
	Name             string `json:"name" gorm:"unique"`
	Equipment        string `json:"equipment"`
	Category         string `json:"category"`
	MainMuscles      string `json:"main_muscles"`
	SecondaryMuscles string `json:"secondary_muscles"`
	Description      string `json:"description"`
}
