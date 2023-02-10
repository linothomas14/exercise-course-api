package response

import "time"

type CourseCategoryRes struct {
	ID   uint32    `json:"id" `
	Name time.Time `json:"name" `
}
