package belajargolanggorm

import "time"

type User struct {
	ID          string    `gorm:"primary_key;column:id;<-:create"`
	Password    string    `gorm:"column:password"`
	Name        Name      `gorm:"embedded"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;<-create"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Information string    `gorm:"-"` //ignored
	Wallet      Wallet    `gorm:"foreignKey:user_id;references:id"`
}

// custom table name if not match with gorm convention
func (u *User) TableName() string {
	return "users"
}

type Name struct {
	FirstName  string `gorm:"first_name"`
	MiddleName string `gorm:"middle_name"`
	LastName   string `gorm:"last_name"`
}

type UserLog struct {
	ID        int    `gorm:"primary_key;column:id;autoIncrement"`
	UserId    string `gorm:"column:user_id"`
	Action    string `gorm:"column:action"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}
