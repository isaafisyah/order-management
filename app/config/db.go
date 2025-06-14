package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cnf *Config) (*gorm.DB, error) {
	fmt.Println("Welcome to " + cnf.Server.Name)
	//connect postgres
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		cnf.Database.User,
		cnf.Database.Password,
		cnf.Database.Name,
		cnf.Database.Host,
		cnf.Database.Port,
		"disable",
	)

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	return connection, nil
}
