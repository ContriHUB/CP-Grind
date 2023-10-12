package models

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Handle  string
	Rank string
	Rating float64
	MaxRank string
	MaxRating float64
}
