package model

type UserCourse struct {
	ID       uint32 `json:"id" gorm:"primaryKey;notNull"`
	UserID   uint32 `json:"user_id" gorm:""`
	User     User   `json:"-"`
	CourseID uint32 `json:"course_id" gorm:""`
	Course   Course `json:"-"`
}

func (UserCourse) TableName() string {
	return "user_course"
}
