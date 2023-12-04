package types

type Datentyp struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
