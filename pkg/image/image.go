package image

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	ID           string
	Name         string
	Url          string
	SmallUrl     string
	ImageGroupId string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ImageGroup struct {
	gorm.Model
	ID        string
	Name      string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ImagesByGroup(db *gorm.DB, group_id string) ([]Image, error) {
	var images []Image
	err := db.Where("image_group_id = ?", group_id).Find(&images).Error
	return images, err
}

func ImageById(db *gorm.DB, image_id string) (Image, error) {
	var image Image
	err := db.Where("id = ?", image_id).First(&image).Error
	return image, err
}

func ImageGroupsByUser(db *gorm.DB, user_id string) ([]ImageGroup, error) {
	var groups []ImageGroup
	err := db.Where("user_id = ?", user_id).Find(&groups).Error
	return groups, err
}
