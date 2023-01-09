package kredit

import "rema/kredit/model"

type ResultChecklistPencairan struct {
	data          []model.RequestKreditData
	total         int
	total_halaman int
}
type UpdateChecklistPencairan struct{
	Data []UpdateDetailChecklistPencairan `json:"data" binding:"dive"`
}
type UpdateDetailChecklistPencairan struct{
	Id string `json:"id" binding:"required"`
}