package router

import (
	"file-upload/api"
	"file-upload/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 100 << 20
	r.Use(middleware.Cors())

	// 文件上传接口
	r.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		// Source
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		filename, _ := filepath.Abs(`/home/sftp_root/shared/` + file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		// 修改权限
		if err := os.Chown(filename, 1001, 1002); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email))
	})
	/// 获取文件列表
	r.GET("/files", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "获取文件列表成功",
			"data":    traversingDir(`/home/sftp_root/shared/`),
		})
	})

	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// apk
		v1.POST("/apks", api.AddApk)
		v1.DELETE("/apks", api.DeleteApk)
		v1.PUT("/apks", api.EditApk)
		v1.GET("/apks", api.ListApk)
	}

	return r
}

type FileInfo struct {
	Name    string    `json:"name"`
	Size    int64     `json:"size"`
	ModDate string    `json:"mod_date"`
	ModTime time.Time `json:"mod_time"`
	Owner   string    `json:"owner"`
}

/// 遍历文件夹
func traversingDir(path string) []FileInfo {
	var result []FileInfo
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		result = append(result, FileInfo{
			Name:    f.Name(),
			Size:    f.Size(),
			ModDate: f.ModTime().Format("2006/1/2 15:04:05"),
			ModTime: f.ModTime(),
			Owner:   "Jotyy",
		})
		return nil
	})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return result
}
