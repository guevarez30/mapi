package image

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	ID             int
	UUID           string
	Name           string
	Url            string
	SmallUrl       string
	ImageGroupUUID string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type ImageGroup struct {
	gorm.Model
	ID        int
	UUID      string
	Name      string
	UserUUID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ImagesByGroup(db *gorm.DB, group_id string) ([]Image, error) {
	var images []Image
	err := db.Where("image_group_uuid = ?", group_id).Find(&images).Error
	return images, err
}

func ImageById(db *gorm.DB, image_id string) (Image, error) {
	var image Image
	err := db.Where("uuid = ?", image_id).First(&image).Error
	return image, err
}

func ImageGroupsByUser(db *gorm.DB, user_id string) ([]ImageGroup, error) {
	var groups []ImageGroup
	err := db.Where("user_uuid = ?", user_id).Find(&groups).Error
	return groups, err
}
