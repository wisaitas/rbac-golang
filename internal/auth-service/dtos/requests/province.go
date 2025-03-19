package requests

import "mime/multipart"

type ImportProvince struct {
	File *multipart.FileHeader `form:"file"`
}
