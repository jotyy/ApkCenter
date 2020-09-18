package apk

import (
	"file-upload/models"
	"file-upload/serializer"
	"strconv"
)

type EditApkService struct {
	ApkVersion     string `form:"apk_version" json:"apk_version"`
	ApkDescription string `form:"apk_description" json:"apk_description" binding:"required,min=4,max=30"`
	UploadBy       string `form:"upload_by" json:"upload_by" binding:"required,min=2,max=10"`
}

func (service *EditApkService) Edit(id string) *serializer.Response {
	var apk models.Apk

	err := models.DB.First(&apk, id).Error
	if err != nil {
		return &serializer.Response{
			Code: 40000,
			Msg:  "该文件不存在",
		}
	}

	apk = models.Apk{
		ApkVersion:     service.ApkVersion,
		ApkDescription: service.ApkDescription,
		UploadBy:       service.UploadBy,
	}

	apkId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return &serializer.Response{
			Code: 40000,
			Msg:  "文件id有误" + err.Error(),
		}
	}

	if err := models.DB.Model(&apk).Where("id = ?", apkId).Update(&apk).Error; err != nil {
		return &serializer.Response{
			Code: 40000,
			Msg:  "修改文件失败",
		}
	}

	return &serializer.Response{
		Code: 0,
		Data: serializer.BuildApk(apk),
		Msg:  "修改成功",
	}
}
