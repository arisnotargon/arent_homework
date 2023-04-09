package db

type TargetWeight struct {
	ID           uint   `gorm:"primary_key"`
	UserId       uint   `gorm:"index,column:user_id"`
	TargetWeight string `gorm:"column:target_weight"`
	NowWeight    string `gorm:"column:now_weight"`
	Status       uint8  `gorm:"comment:状態、1:有効 2:無効"`
}

type TargetWeightStatus uint8

const (
	TargetWeightStatusValid TargetWeightStatus = iota + 1
	TargetWeightStatusInvalid
)
