package model

type DownloadRequest struct {
	Link     string `json:"link"`
	FileType string `json:"type"`
}

type FileType string

const (
	EXCEL FileType = "excel"
	CSV   FileType = "csv"
)
