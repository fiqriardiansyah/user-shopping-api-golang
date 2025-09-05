package main

import (
	"fmt"

	"github.com/fiqriardiansyah/user-shopping-api-golang/db/seeders"
)

func main() {
	fmt.Println("START SEEDING ...")
	seeders.RoleSeed()
	fmt.Println("SEEDING FINISH 🌱")
}
