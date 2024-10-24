package repository

import (
	"fmt"

	"github.com/RomanshkVolkov/test-api/internal/core/domain"
	"gorm.io/gorm"
)

func PrintSeedAction(nameTable string, action string) {
	fmt.Println("Seeding table: " + nameTable + " " + action + " Success")
}

func AutoMigrateTable(db *gorm.DB, table interface{}) {
	fmt.Println("AutoMigrateTable")
	isInitialized := db.Migrator().HasTable(&table)
	if !isInitialized {
		db.AutoMigrate(table)
		PrintSeedAction("Shifts", "Create")
	}
}

func RunSeeds(db *gorm.DB) {
	// db.Exec("drop table users")
	// db.Exec("drop table user_profiles")
	SeedProfiles(db)
	SeedDevAuthorizedIPAddress(db)
	SeedUsers(db)
	SeedKitchens(db)

	AutoMigrateTable(db, &domain.Shifts{})
}

func SeedUsers(db *gorm.DB) {
	AutoMigrateTable(db, &domain.User{})

	var currentRows int64
	db.Model(&domain.User{}).Count(&currentRows)

	if currentRows > 0 {
		return
	}

	rootProfile := domain.UserProfiles{}
	db.Model(&domain.UserProfiles{}).Where("slug = ?", "root").First(&rootProfile)

	users := []*domain.User{
		{
			UserData: domain.UserData{
				Username:  "jose_dwit",
				Email:     "joseguzmandev@gmail.com",
				Name:      "Jose Guzman",
				ProfileID: rootProfile.ID,
			},
			Password: "password",
		},
		{
			UserData: domain.UserData{
				Username:  "diego_dwit",
				Email:     "diegogutcat@gmail.com",
				Name:      "Diego Gutierrez",
				ProfileID: rootProfile.ID,
			},
			Password: "password",
		},
		{
			UserData: domain.UserData{
				Username:  "itzel_dwit",
				Email:     "itzelramonf@gmail.com",
				Name:      "Itzram",
				ProfileID: rootProfile.ID,
			},
			Password: "password",
		},
	}

	for _, user := range users {
		hashedPassword, _ := HashPassword(user.Password)
		user.Password = hashedPassword

		db.Create(&user)
	}
}

func SeedProfiles(db *gorm.DB) {
	AutoMigrateTable(db, &domain.UserProfiles{})

	var currentRows int64
	db.Model(&domain.UserProfiles{}).Count(&currentRows)

	if currentRows > 0 {
		return
	}

	profiles := []*domain.UserProfiles{
		{
			Name: "Super Admin",
			Slug: "root",
		},
		{
			Name: "Administrador",
			Slug: "admin",
		},
		{
			Name: "Cliente",
			Slug: "customer",
		},
		{
			Name: "Encargado",
			Slug: "manager",
		},
		{
			Name: "Cocinero",
			Slug: "cook",
		},
	}

	for _, profile := range profiles {
		db.Create(&profile)
	}
}

func SeedDevAuthorizedIPAddress(db *gorm.DB) {
	AutoMigrateTable(db, &domain.Dev{})

	var currentRows int64
	db.Model(&domain.Dev{}).Count(&currentRows)

	if currentRows > 0 {
		return
	}

	devs := []*domain.Dev{
		{
			IP:  "172.18.0.1",
			Tag: "docker local",
		},
	}

	for _, dev := range devs {
		db.Create(&dev)
	}
}

func SeedKitchens(db *gorm.DB) {
	AutoMigrateTable(db, &domain.Kitchen{})
	AutoMigrateTable(db, &domain.UsersHasKitchens{})
}

func SeedShifts(db *gorm.DB) {
}

func SeedPermissions(db *gorm.DB) {
	AutoMigrateTable(db, &domain.Permission{})

	var currentRows int64
	db.Model(&domain.Permission{}).Count(&currentRows)

	if currentRows > 0 {
		return
	}

	// permissions := []*domain.Permission{
	// 	{
	// 		Name: "Dashboard",
	// 		Path: "/dashboard",
	// 	},
	// 	{
	// 		Name: "Configuración",
	// 		Path: "/dashboard/settings",
	// 	},
	// 	{
	// 		Name: "Usuarios",
	// 		Path: "/dashboard/settings/users",
	// 	},
	// 	{
	// 		Name: "Platillos",
	// 		Path: "/dashboard/dishes",
	// 	},
	// 	{
	// 		Name: "Administración",
	// 		Path: "/dashboard/managment",
	// 	},
	// }
}
