package model

type CourseCategory struct {
	ID   uint32 `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;notNull"`
}

func (CourseCategory) TableName() string {
	return "course_category"
}
