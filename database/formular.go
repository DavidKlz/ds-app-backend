package database

import (
	"github.com/davidklz/ds-app-backend/types"
	"github.com/google/uuid"
)

func (db *Database) GetFormular(id uuid.UUID) (*types.Formular, error) {
	form := &types.Formular{}
	result := db.connection.Find(&form, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return form, nil
}

func (db *Database) GetAllFormulare() (*[]types.Formular, error) {
	forms := &[]types.Formular{}
	result := db.connection.Find(&forms)
	if result.Error != nil {
		return nil, result.Error
	}
	return forms, nil
}

func (db *Database) SaveFormular(f *types.Formular) error {
	return db.connection.Save(f).Error
}

func (db *Database) DeleteFormular(id uuid.UUID) error {
	return db.connection.Unscoped().Delete(&types.Formular{}, id).Error
}
