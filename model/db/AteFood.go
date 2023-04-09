package db

import "time"

type AteFood struct {
	ID     uint      `gorm:"primary_key"`
	UserId uint      `gorm:"index,column:user_id"`
	Name   string    `gorm:"column:name"`
	Type   uint8     `gorm:"index,comment:食事の類別区分番号、1:朝食 2:お昼 3:夕食 4:おやつ"`
	AteAt  time.Time `gorm:"index,column:ate_at"`
	Pic    string    `gorm:"column:pic"`
}

type AteFoodType uint8

const (
	AteFoodTypeBreakfast = iota + 1
	AteFoodTypeLunch
	AteFoodTypeDinner
	AteFoodTypeSnack
)
