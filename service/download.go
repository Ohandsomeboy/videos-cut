package service

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"videos.cn/models"
)

// DownloadVideo
// @Summary 下载视频
// @Description 下载视频
// @Tags Videos
// @Produce mpfd
// @Param name query string true "视频文件名"
// @Success 200
// @Failure 404 {string} json "{"code":"404","msg":"视频文件不存在"}"
// @Router /api/videos/download/{name} [get]
func DownloadVideo(c *gin.Context) {
	name := c.Query("name")

	var videos []models.Video
	if err := models.DB.Where("name like ?", name).Order("id asc").Find(&videos).Error; err != nil {
		log.Println("查询视频数据失败：", err)
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "视频不存在",
		})
		return
	}

	var buf bytes.Buffer
	for _, v := range videos {
		_, err := buf.Write(v.Video)
		if err != nil {
			log.Println("写入响应失败：", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "下载视频失败",
			})
			return
		}
	}

	c.Header("Content-Type", "video/mp4")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
	c.Writer.WriteHeader(http.StatusOK)
	//if _, err := io.Copy(c.Writer, buf); err != nil {
	if _, err := io.Copy(c.Writer, &buf); err != nil {
		log.Println("写入响应失败：", err)
	}
}
