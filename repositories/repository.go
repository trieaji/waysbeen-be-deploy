package repositories

import "gorm.io/gorm"

// Declare repository struct here ...
type repository struct {
	db *gorm.DB
}
