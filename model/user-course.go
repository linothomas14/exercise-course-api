package model

type UserCourse struct {
	ID       uint32 `json:"id" gorm:"primaryKey;notNull"`
	UserID   uint32 `json:"user_id" gorm:""`
	User     User   `json:"user"`
	CourseID uint32 `json:"course_id" gorm:""`
	Course   Course `json:"course"`
}

func (UserCourse) TableName() string {
	return "user_course"
}
