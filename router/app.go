package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "videos.cn/docs"
	"videos.cn/service"
)

func Router() *gin.Engine {
	maxFileSize := int64(1024 * 1024 * 1024)

	r := gin.Default()
	r.MaxMultipartMemory = maxFileSize
	// Swagg 配置
	// Swagg 配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 视频操作基础
	r.GET("/videos-list", service.GetVideosList)

	r.POST("/api/videos/upload", service.UploadVideo)
	r.GET("/api/videos/download/{name}", service.DownloadVideo)

	return r
}
