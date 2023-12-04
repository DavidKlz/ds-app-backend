package database

import (
	"github.com/davidklz/ds-app-backend/types"
	"github.com/google/uuid"
)

func (db *Database) GetVariable(id uuid.UUID) (*types.Variable, error) {
	variable := &types.Variable{}
	result := db.connection.Find(&variable, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return variable, nil
}

func (db *Database) GetAllVariablen() (*[]types.Variable, error) {
	vars := &[]types.Variable{}
	result := db.connection.Find(&vars)
	if result.Error != nil {
		return nil, result.Error
	}
	return vars, nil
}

func (db *Database) SaveVariable(v *types.Variable) error {
	return db.connection.Save(v).Error
}

func (db *Database) DeleteVariable(id uuid.UUID) error {
	return db.connection.Unscoped().Delete(&types.Variable{}, id).Error
}
