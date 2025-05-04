package entities

type User struct {
	ID       uint    `gorm:"primaryKey"`
	Username string  `gorm:"uniqueIndex;not null"`
	Email    string  `gorm:"uniqueIndex;not null"`
	Balance  float64 `gorm:"type:decimal(10,2);default:0"`
}
