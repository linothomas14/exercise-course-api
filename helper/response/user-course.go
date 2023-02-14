package response

type UserCourseRes struct {
	ID       uint32 `json:"id"`
	UserID   uint32 `json:"user_id" validate:"required"`
	CourseId uint32 `json:"course_id" validate:"required"`
}
