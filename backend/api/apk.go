package api

import (
	"file-upload/serializer"
	"file-upload/service/apk"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// @Summary 上传apk
// @Produce  json
// @Param upload_by	query	string	true	"上传者"
// @Param file	query	int	false	"查询条数"
// @Param apk_description	query	string	false	"上传信息"
// @Success 200 {object} serializer.Response
// @Failure 500 {object} serializer.Response
// @Router /api/v1/apks	[POST]
func AddApk(c *gin.Context) {
	var service apk.AddApkService

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	filename := file.Filename[0:len(file.Filename)-4] + "_" + time.Now().Format("20060102150405") + ".apk"
	filePath, _ := filepath.Abs("/app/apks/" + filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	// 修改权限
	if err := os.Chown(filePath, 1001, 1002); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code: 40000,
			Msg:  "chown file err: " + err.Error(),
		})
		return
	}

	if err := c.ShouldBind(&service); err == nil {
		service.FileName = getApkInfo(filePath).Name()
		service.FileSize = getApkInfo(filePath).Size()
		service.UploadTime = getApkInfo(filePath).ModTime()
		if apk, err := service.Add(); err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			res := serializer.BuildApkResponse(apk)
			c.JSON(http.StatusOK, res)
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// @Summary 查询所有Apk
// @Produce  json
// @Param limit	query	int	false	"查询条数"
// @Param start	query	int	false	"开始条数"
// @Success 200 {object} serializer.Response
// @Failure 500 {object} serializer.Response
// @Router /api/v1/apks	[GET]
func ListApk(c *gin.Context) {
	var service apk.ListApkService

	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// @Summary 删除Apk
// @Produce  json
// @Param id query int true	"id"
// @Success 200 {object} serializer.Response
// @Failure 500 {object} serializer.Response
// @Router /api/v1/apks	[DELETE]
func DeleteApk(c *gin.Context) {
	service := apk.DeleteApkService{}

	if err := c.ShouldBind(&service); err == nil {
		if err := service.Delete(c.Query("id")); err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			c.JSON(http.StatusOK, serializer.Response{
				Code: 0,
				Msg:  "删除成功",
			})
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// @Summary 修改Apk信息
// @Produce  json
// @Param id query int true "id"
// @Param upload_by	query	string	false	"上传者"
// @Param apk_description	query	string	false	"上传信息"
// @Success 200 {object} serializer.Response
// @Failure 500 {object} serializer.Response
// @Router /api/v1/apks	[PUT]
func EditApk(c *gin.Context) {
	service := apk.EditApkService{}

	if err := c.ShouldBind(&service); err == nil {
		if err := service.Edit(c.PostForm("id")); err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			c.JSON(http.StatusOK, serializer.Response{
				Code: 0,
				Msg:  "修改成功",
			})
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

/// 获取文件信息
func getApkInfo(path string) os.FileInfo {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Printf("fileInfo returned %v\n", err)
	}

	return fileInfo
}
