package apk

import (
	"file-upload/models"
	"file-upload/serializer"
)

type DeleteApkService struct {
}

func (service *DeleteApkService) Delete(id string) *serializer.Response {
	var apk models.Apk
	err := models.DB.First(&apk, id).Error
	if err != nil {
		return &serializer.Response{
			Code: 40000,
			Msg:  "该文件不存在",
		}
	}

	if err := models.DB.Delete(&apk).Error; err != nil {
		return &serializer.Response{
			Code: 40000,
			Msg:  "删除文件失败",
		}
	}

	return &serializer.Response{
		Code: 0,
		Msg: "删除成功",
	}
}