package model

type Course struct {
	ID               uint32 `json:"id" gorm:"primaryKey;notNull"`
	Title            string `json:"title" gorm:"notNull;unique" validate:"required"`
	CourseCategoryId uint32 `json:"course_category_id" gorm:"notNull;foreignKey:CourseCategoryId" validate:"required"`
	CourseCategory   CourseCategory
}

func (Course) TableName() string {
	return "course"
}
