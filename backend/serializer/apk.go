package serializer

import (
	"file-upload/models"
)

type Apk struct {
	ID             uint   `json:"id"`
	FileName       string `json:"file_name"`
	FileSize       int64  `json:"file_size"`
	ApkVersion     string `json:"apk_version"`
	ApkDescription string `json:"apk_description"`
	ApkDownloadUrl string `json:"apk_download_url"`
	UploadTime     string `json:"upload_time"`
	UploadBy       string `json:"upload_by"`
	State          string `json:"state"`
}

type ApkResponse struct {
	Response
	Data Apk `json:"data"`
}

func BuildApk(apk models.Apk) Apk {
	return Apk{
		ID:             apk.ID,
		FileName:       apk.FileName,
		FileSize:       apk.FileSize,
		ApkVersion:     apk.ApkVersion,
		ApkDescription: apk.ApkDescription,
		ApkDownloadUrl: apk.ApkDownloadUrl,
		UploadTime:     apk.UploadTime.Format("2006-01-02 15:04:05"),
		UploadBy:       apk.UploadBy,
		State:          apk.State,
	}
}

func BuildApkResponse(apk models.Apk) ApkResponse {
	return ApkResponse{
		Data: BuildApk(apk),
	}
}

// BuildBlogs 序列化配置列表
func BuildApkList(items []models.Apk) (apkList []Apk) {
	for _, item := range items {
		blog := BuildApk(item)
		apkList = append(apkList, blog)
	}
	return apkList
}
