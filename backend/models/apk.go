package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Apk struct {
	gorm.Model
	FileName       string    `json:"file_name"`
	FileSize       int64     `json:"file_size"`
	ApkVersion     string    `json:"apk_version"`
	ApkDescription string    `json:"apk_description"`
	ApkDownloadUrl string    `json:"apk_download_url"`
	UploadTime     time.Time `json:"upload_time"`
	UploadBy       string    `json:"upload_by"`
	DeleteBy       string    `json:"delete_by"`
	State          string    `json:"state"`
}
