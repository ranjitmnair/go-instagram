package models

import (
	"time"
)

type Post struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Caption   string    `gorm:"size:255;not null;unique" json:"caption"`
	Image     string    `gorm:"size:255;not null;" json:"image"`
	UserId    uint32    `gorm:"not null" json:"user_id"`
	TimeStamp time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"timestamp"`
}
