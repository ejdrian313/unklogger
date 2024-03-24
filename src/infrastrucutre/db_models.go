package infrastrucutre

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ActivityLog struct {
	gorm.Model
	UserName string
	LastSeen time.Time
	UserID   uint
}

type User struct {
	gorm.Model
	Name string
}

type UserAuth struct {
	gorm.Model
	UserID    uint
	Username  string
	Password  string
	CreatedAt time.Time
}
