package response

type GetCoursesRes struct {
	ID               uint32 `json:"id" `
	Title            string `json:"title" `
	CourseCategoryId uint32 `json:"course_category_id" validate:"required"`
}
