package types

type Controltyp struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
