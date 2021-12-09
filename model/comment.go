package model

type Comments struct {
	Count  int    `json:"count"`
	LastAt string `json:"last_at"`
}

type CommentType struct {
	ID          int    `json:"id"`
	Content     string `json:"content"`
	Public      bool   `json:"public"`
	PublishedAt string `json:"published_at"`
	Author      User   `json:"author"`
}
