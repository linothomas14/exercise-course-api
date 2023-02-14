package model

type UserCourse struct {
	ID       uint32 `json:"id" gorm:"primaryKey;notNull"`
	UserID   uint32 `json:"user_id" gorm:""`
	CourseID uint32 `json:"course_id" gorm:""`
}

func (UserCourse) TableName() string {
	return "user_course"
}
