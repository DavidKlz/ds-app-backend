package database

import "github.com/davidklz/ds-app-backend/types"

func (db *Database) GetDatentyp(id int) (*types.Datentyp, error) {
	dt := &types.Datentyp{}
	result := db.connection.Find(&dt, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return dt, nil
}

func (db *Database) GetControltyp(id int) (*types.Controltyp, error) {
	ct := &types.Controltyp{}
	result := db.connection.Find(&ct, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return ct, nil
}

func (db *Database) GetAllDatentypen() (*[]types.Datentyp, error) {
	types := &[]types.Datentyp{}
	result := db.connection.Find(&types)
	if result.Error != nil {
		return nil, result.Error
	}
	return types, nil
}

func (db *Database) GetAllControltypen() (*[]types.Controltyp, error) {
	types := &[]types.Controltyp{}
	result := db.connection.Find(&types)
	if result.Error != nil {
		return nil, result.Error
	}
	return types, nil
}

func (db *Database) SaveDatentyp(d *types.Datentyp) error {
	return db.connection.Save(d).Error
}

func (db *Database) SaveControltyp(c *types.Controltyp) error {
	return db.connection.Save(c).Error
}

func (db *Database) DeleteDatentyp(id int) error {
	return db.connection.Delete(&types.Datentyp{}, id).Error
}

func (db *Database) DeleteControltyp(id int) error {
	return db.connection.Delete(&types.Controltyp{}, id).Error
}
