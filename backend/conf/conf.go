package conf

import (
	"file-upload/models"
	"github.com/joho/godotenv"
	"os"
)

var DownloadUrlPrefix string

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {

	}

	downloadUrlPrefix := os.Getenv("DOWNLOAD_PREFIX")
	DownloadUrlPrefix = downloadUrlPrefix

	// 连接数据库
	models.SetUp(os.Getenv("POSTGRESQL_ADDRESS"))
}
