package db

import "time"

type BodyInfo struct {
	ID        uint      `gorm:"primary_key"`
	UserId    uint      `gorm:"index,column:user_id"`
	RecodedAt time.Time `gorm:"index,column:recorded_at"`
	Weight    string    `gorm:"column:weight"`
	FatRate   string    `gorm:"column:fat_rate"`
}
