package model

type User struct {
	ID string `json:"id" gorm:"primaryKey; type:varchar; not null; unique"`

	UserName string `json:"username" gorm:"type:varchar; column:username; not null; unique; index" validate:"required"`
	Password string `json:"password" gorm:"type:varchar; not null" validate:"required"`

	Role string `json:"role" gorm:"type:varchar; not null; check:role IN ('admin', 'user'); default:user"`

	Gorm
}

type Auth struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}
