package apk

import (
	"file-upload/models"
	"file-upload/serializer"
	"os"
	"path/filepath"
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

	filename := apk.FileName
	filePath, _ := filepath.Abs("/app/apks/" + filename)
	os.Remove(filePath)

	return &serializer.Response{
		Code: 0,
		Msg: "删除成功",
	}
}