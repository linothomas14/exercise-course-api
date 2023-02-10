package param

type UserCourse struct {
	UserId   uint32 `json:"user_id" validate:"required"`
	CourseId uint32 `json:"course_id" validate:"required"`
}
