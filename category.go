package main

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func addCategory(db *gorm.DB, category *Category) error {
	if category == nil || strings.TrimSpace(category.Name) == "" {
		return errors.New("category name is required")
	}
	category.Name = strings.TrimSpace(category.Name)
	return db.Create(category).Error
}

// removeCategoryByID refuses to delete a category while it still has products.
func removeCategoryByID(db *gorm.DB, id uint) error {
	var category Category
	if err := db.First(&category, id).Error; err != nil {
		return err
	}

	var itemCount int64
	if err := db.Unscoped().Model(&Product{}).Where("category_id = ?", id).Count(&itemCount).Error; err != nil {
		return err
	}
	if itemCount > 0 {
		return fmt.Errorf("cannot delete category %d: it still has %d product(s)", id, itemCount)
	}
	return db.Delete(&category).Error
}
