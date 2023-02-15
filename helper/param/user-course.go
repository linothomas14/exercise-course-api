package param

type UserCourseCreate struct {
	UserID   uint32 `json:"user_id" validate:"required"`
	CourseID uint32 `json:"course_id" validate:"required"`
}
