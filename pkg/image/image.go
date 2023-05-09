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
	result := db.Model(Image{ImageGroupId: group_id}).Find(&images)
	return images, result.Error
}

func ImagesById(db *gorm.DB, image_id string) (Image, error) {
	var images Image
	result := db.Model(Image{ID: image_id}).First(&images)
	return images, result.Error
}

func ImagesGroupsByUser(db *gorm.DB, user_id string) ([]ImageGroup, error) {
	var groups []ImageGroup
	result := db.Model(ImageGroup{UserId: user_id}).Find(&groups)
	return groups, result.Error
}
