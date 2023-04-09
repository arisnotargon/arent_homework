package db

import "time"

type Notification struct {
	ID      uint      `gorm:"primary_key"`
	UserId  uint      `gorm:"index,column:user_id"`
	ReadAt  time.Time `gorm:"index"`
	Title   string
	Content string
}
