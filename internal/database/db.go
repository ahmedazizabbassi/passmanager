package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ahmedazizabbassi/pass/internal/models"
	_ "github.com/go-sql-driver/mysql"
	mysqldriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() error {
	rootDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	tempDB, err := sql.Open("mysql", rootDSN)
	if err != nil {
		return fmt.Errorf("error connecting to MySQL: %v", err)
	}
	defer tempDB.Close()

	dbName := os.Getenv("DB_NAME")
	rows, err := tempDB.Query(fmt.Sprintf("SHOW DATABASES LIKE '%s'", dbName))
	if err != nil {
		return fmt.Errorf("error checking database existence: %v", err)
	}
	defer rows.Close()

	dbExists := rows.Next()
	if !dbExists {
		log.Printf("Database %s does not exist. Creating...", dbName)
		_, err = tempDB.Exec(fmt.Sprintf("CREATE DATABASE %s CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", dbName))
		if err != nil {
			return fmt.Errorf("error creating database: %v", err)
		}
		log.Printf("Database %s created successfully", dbName)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		dbName,
	)

	DB, err = gorm.Open(mysqldriver.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Vault{},
		&models.Secret{},
		&models.Device{},
		&models.AuditLog{},
	)
	if err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	log.Printf("Database connection established and migrations completed")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
