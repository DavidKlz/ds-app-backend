package database

import (
	"fmt"
	"os"

	"github.com/davidklz/ds-app-backend/types"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Store interface {
	GetFormular(uuid.UUID) (*types.Formular, error)
	GetVariable(uuid.UUID) (*types.Variable, error)
	GetDatentyp(int) (*types.Datentyp, error)
	GetControltyp(int) (*types.Controltyp, error)
	GetAllFormulare() (*[]types.Formular, error)
	GetAllVariablen() (*[]types.Variable, error)
	GetAllDatentypen() (*[]types.Datentyp, error)
	GetAllControltypen() (*[]types.Controltyp, error)
	SaveFormular(*types.Formular) error
	SaveVariable(*types.Variable) error
	SaveDatentyp(*types.Datentyp) error
	SaveControltyp(*types.Controltyp) error
	DeleteFormular(uuid.UUID) error
	DeleteVariable(uuid.UUID) error
	DeleteDatentyp(int) error
	DeleteControltyp(int) error
}

type Database struct {
	connection *gorm.DB
}

func ConnectToPostgres() (*Database, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&types.Controltyp{})
	db.AutoMigrate(&types.Datentyp{})
	db.AutoMigrate(&types.Variable{})
	db.AutoMigrate(&types.Formular{})

	return &Database{
		connection: db,
	}, nil
}
