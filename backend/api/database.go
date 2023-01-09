package api

import (
	"fmt"
	"os"
	"rema/kredit/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDb() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	dbUrl := "host="+os.Getenv("host")
	dbUrl = dbUrl+" user="+os.Getenv("user")
	dbUrl = dbUrl+ " password="+os.Getenv("password")
	dbUrl = dbUrl + " dbname="+os.Getenv("dbname")
	dbUrl = dbUrl + " port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	if err := db.AutoMigrate(model.BranchTab{}, model.MstCompanyTab{},model.CustomerDataTab{}, model.LoanDataTab{},model.SkalaRentalTab{}, model.StagingCustomer{}, model.VehicleDataTab{}, model.StagingError{}, model.User{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, err
}
