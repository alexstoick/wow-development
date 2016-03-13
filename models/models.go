package models

import (
	"time"
)

type Realm struct {
	Name string
	Slug string
}

type Modifier struct {
	Type  int
	Value int
}

type Bonus struct {
	Id int `json:"bonusListId"`
}

type AHFile struct {
	ID           int    `gorm:"primary_key"`
	URL          string `sql:"not_null"`
	LastModified int    `sql:"type:bigint;unique"`
	CreatedAt    time.Time
}
