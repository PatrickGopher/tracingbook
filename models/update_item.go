package models

import (
	"time"

	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
)

type UpdateItem struct {
	gorm.Model
	SiteId        uint   `gorm:"not null"`
	BookId        uint   `gorm:"not null"`
	BookName      string `gorm:"not null"`
	Slug          string `gorm:"not null"`
	LatestName    string `gorm:"not null"`
	LatestChapter string `gorm:"not null"`
	BookUrl       string `gorm:"not null"`
	UserId        uint32 `gorm:"default:null"`
}

func (updateItem *UpdateItem) BeforeSave() (err error) {
	updateItem.Slug = slug.Make(updateItem.LatestName)
	updateItem.UpdatedAt = time.Now()
	return
}
