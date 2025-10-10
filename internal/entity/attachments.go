package entity

import "io"

type AttachmentFlat struct {
	Id              *int       `json:"id" db:"id"`
	ParentId        *int       `json:"parent_id" db:"parent_id"`
	Module          *string    `json:"module" db:"module"`
	Title           *string    `json:"title" db:"title"`
	Link      		string     `json:"link" db:"link"`
	ContentType     *string    `json:"content_type" db:"content_type"`
	Description     *string    `json:"description" db:"description"`
	UpdatedAt       *string    `json:"updated_at" db:"updated_at"`
	FileObject      *File      `json:"-" db:"-"`
}

type AttachmentResp struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Link            string    `json:"link"`
	Description     *string   `json:"description,omitempty"`
}

type AttachmentRespArray []AttachmentResp

type File struct {
	Name        string
	Size        int64
	ContentType string
	Content     io.Reader
}