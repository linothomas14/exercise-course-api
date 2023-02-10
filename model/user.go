package model

type User struct {
	ID       uint32 `json:"id" gorm:"primaryKey;notNull"`
	Email    string `json:"email" gorm:"unique;notNull" validate:"email"`
	Name     string `json:"name" gorm:"notNull"`
	Password string `json:"password" gorm:"notNull" validate:"min=6"`
}

func (User) TableName() string {
	return "user"
}
