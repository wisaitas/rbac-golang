package requests

import "mime/multipart"

type ImportSubDistrict struct {
	File *multipart.FileHeader `form:"file"`
}
