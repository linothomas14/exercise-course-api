package response

import "github.com/linothomas14/exercise-course-api/model"

type GetCoursesRes struct {
	ID               uint32               `json:"id" `
	Title            string               `json:"title" `
	CourseCategoryId uint32               `json:"course_category_id" validate:"required"`
	CourseCategory   model.CourseCategory `json:"course_category"`
}
