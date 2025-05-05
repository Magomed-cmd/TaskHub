package model

type User struct {
	ID       uint64 `gorm:"primaryKey"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `gorm:"type:text;not null"`
	Name     string `gorm:"type:varchar(20);not null"`
	UUID     string `gorm:"type:varchar(128);not null"`
}
