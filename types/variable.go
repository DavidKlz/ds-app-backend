package types

import "github.com/google/uuid"

type Variable struct {
	ID           uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name         string     `json:"name"`
	Vorgabewert  string     `json:"vorgabewert"`
	ControltypID int        `json:"controltypId"`
	Controltyp   Controltyp `gorm:"foreignKey:ID" json:"controltyp"`
	DatentypID   int        `json:"datentypId"`
	Datentyp     Datentyp   `gorm:"foreignKey:ID" json:"datentyp"`
	Sichtbar     bool       `json:"sichtbar"`
	Editierbar   bool       `json:"editierbar"`
	Pflichtfeld  bool       `json:"pflichtfeld"`
}
