package entities

type Test struct {
	ParentEntity

	Name        *string `gorm:"not null;" json:"name"`
	Description *string `json:"description"`

}
