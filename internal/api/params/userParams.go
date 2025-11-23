package params

import "time"

type RegisterParam struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

type LoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type InfoParam struct {
	Username  string    `json:"username"`
	UserID    int64     `json:"userid"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PasswordParam struct {
	UserID      int64  `json:"userid"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
