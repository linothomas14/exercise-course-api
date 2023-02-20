package param

type CreateCourse struct {
	Title            string `json:"title" validate:"required"`
	CourseCategoryId uint32 `json:"course_category_id" validate:"required"`
}

type UpdateCourse struct {
	ID               uint32
	Title            string `json:"title"`
	CourseCategoryId uint32 `json:"course_category_id"`
}
