package apk

import (
	"file-upload/models"
	"file-upload/serializer"
	"time"
)

type AddApkService struct {
	FileName       string    `form:"file_name" json:"file_name"`
	FileSize       int64     `form:"file_size" json:"file_size"`
	ApkVersion     string    `form:"apk_version" json:"apk_version"`
	ApkDescription string    `form:"apk_description" json:"apk_description"`
	UploadTime     time.Time `form:"upload_time" json:"upload_time"`
	UploadBy       string    `form:"upload_by" json:"upload_by" binding:"required,min=2,max=10"`
}

func (service *AddApkService) Add() (models.Apk, *serializer.Response) {
	apk := models.Apk{
		FileName:       service.FileName,
		FileSize:       service.FileSize,
		ApkVersion:     service.ApkVersion,
		ApkDescription: service.ApkDescription,
		ApkDownloadUrl: "http://192.168.100.132:8002/" + service.FileName,
		UploadTime:     service.UploadTime,
		UploadBy:       service.UploadBy,
	}

	if err := models.DB.Create(&apk).Error; err != nil {
		return apk, &serializer.Response{
			Code: 0,
			Msg:  "上传Apk失败",
		}
	}

	return apk, nil
}
