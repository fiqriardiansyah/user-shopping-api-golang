package seeders

import (
	"fmt"

	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/config"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/entity"
	"gorm.io/gorm/clause"
)

func RoleSeed() {
	roles := []entity.Role{
		{Name: "admin", AltName: "Admin", Description: "Admin user"},
		{Name: "seller", AltName: "Seller", Description: "Seller user"},
		{Name: "buyer", AltName: "Buyer", Description: "Buyer user"},
	}
	db := config.NewDB()

	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&roles).Error; err != nil {
		panic(fmt.Sprintf("Failed seeding Role: %v", err))
	}

	fmt.Println("SEEDING ROLES SUCCESS ✅")
}
