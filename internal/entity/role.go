package entity

import "github.com/google/uuid"

type Role struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name        string    `gorm:"type:varchar(50);unique;not null" json:"name"`
	AltName     string    `gorm:"type:varchar(50);not null" json:"alt_name"`
	Description string    `gorm:"type:text" json:"description"`
}
