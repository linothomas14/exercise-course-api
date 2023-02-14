package param

type UserCourseCreate struct {
	UserID   uint32 `validate:"required"`
	CourseID uint32 `json:"course_id" validate:"required"`
}
