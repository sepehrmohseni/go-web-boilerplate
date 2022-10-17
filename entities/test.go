package entities

import "gorm.io/datatypes"

type Test struct {
	ParentEntity

	Name        *string `gorm:"not null;" json:"name"`
	Description *string `json:"description"`

}
