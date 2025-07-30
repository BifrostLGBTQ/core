package payloads

type Location struct {
	ID        uint    `gorm:"primaryKey"`
	Latitude  float64 `gorm:"type:decimal(10,7);not null"`
	Longitude float64 `gorm:"type:decimal(10,7);not null"`
	Address   *string `gorm:"size:256"`
}
