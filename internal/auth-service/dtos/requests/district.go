package requests

import "mime/multipart"

type ImportDistrict struct {
	File *multipart.FileHeader `form:"file"`
}
