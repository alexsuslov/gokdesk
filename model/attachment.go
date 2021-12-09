package model

type Attachment struct {
	ID                 int    `json:"id"`
	FileName           string `json:"attachment_file_name"`
	Description        string `json:"description"`
	AttachmentFileSize int    `json:"description"`
	IsPublic           bool   `json:"is_public"`
	CreatedAt          string `json:"created_at"`
}
