package response

import "github.com/linothomas14/exercise-course-api/model"

type UserCourseRes struct {
	ID       uint32 `json:"id"`
	UserID   uint32 `json:"user_id" validate:"required"`
	User     UserResponse
	CourseId uint32 `json:"course_id" validate:"required"`
	Course   model.Course
}
