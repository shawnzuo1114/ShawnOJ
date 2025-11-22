package models

import "time"

type User struct {
	UserID    int64  `db:"userid" json:"userid"`
	Username  string `db:"username" json:"username"`
	Password  string `db:"password" json:"password"`
	Role      string `db:"role" json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
