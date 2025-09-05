package entity

import "github.com/google/uuid"

type UserRole struct {
	UserID uuid.UUID `gorm:"type:uuid;not null;primaryKey" json:"user_id"`
	RoleID uuid.UUID `gorm:"type:uuid;not null;primaryKey" json:"role_id"`

	// Foreign key references
	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Role Role `gorm:"foreignKey:RoleID;references:ID;constraint:OnDelete:CASCADE" json:"role,omitempty"`
}
