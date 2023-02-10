package model

type UserCourse struct {
	ID       uint32 `json:"id" gorm:"primaryKey;notNull"`
	UserId   uint32 `json:"user_id" gorm:""`
	User     User   `json:"user"`
	CourseId uint32 `json:"course_id" gorm:""`
	Course   Course
}

func (UserCourse) TableName() string {
	return "user_course"
}
