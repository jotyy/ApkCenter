package apk

import (
	"file-upload/models"
	"file-upload/serializer"
)

type ListApkService struct {
	Limit     int    `form:"limit"`
	Start     int    `form:"start"`
	Condition string `form:"condition"`
}

// ListApk 列出Apk
func (service *ListApkService) List() serializer.Response {
	var apkList []models.Apk
	total := 0

	if service.Limit == 0 {
		service.Limit = 20
	}

	if err := models.DB.Model(models.Apk{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Code: 40000,
			Msg:    "查询失败",
		}
	}

	if err := models.DB.Limit(service.Limit).Offset(service.Start).Order("upload_time desc").Find(&apkList).Error; err != nil {
		return serializer.Response{
			Code: 50000,
			Data:   nil,
			Msg:    "查询失败",
		}
	}

	return serializer.BuildListResponse(serializer.BuildApkList(apkList), uint(total))
}

// ListByCondition 根据条件排序
func (service *ListApkService) ListByCondition() serializer.Response {
	var apkList []models.Apk
	total := 0

	if service.Limit == 0 {
		service.Limit = 30
	}

	if err := models.DB.Model(models.Apk{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Code: 40000,
			Msg:    "查询失败",
		}
	}

	if err := models.DB.Limit(10).Order(service.Condition).Find(&apkList).Error; err != nil {
		return serializer.Response{
			Code: 40000,
			Data:   nil,
			Msg:    "查询失败",
		}
	}

	return serializer.BuildListResponse(serializer.BuildApkList(apkList), uint(total))
}
