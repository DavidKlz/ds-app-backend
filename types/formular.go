package types

import "github.com/google/uuid"

type Formular struct {
	ID        uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name      string     `json:"name"`
	Variablen []Variable `gorm:"foreignKey:ID" json:"variablen"`
}
