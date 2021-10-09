package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `gorm:"size:255;not null;unique" json:"name"`
	Email    string        `gorm:"size:100;not null;unique" json:"email"`
	Password int           `gorm:"size:100;not null;" json:"password"`
}
