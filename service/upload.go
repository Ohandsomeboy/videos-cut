package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/u2takey/ffmpeg-go"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"videos.cn/models"
)

const (
	maxFileSize   = 1 << 30 // 1GB
	maxPacketSize = 1 << 20 // 1MB
)

// UploadVideo
// @Summary 上传视频
// @Description 上传视频
// @Tags Videos
// @Accept mpfd
// @Produce json
// @Param file formData file true "上传文件"
// @Param start_time query string true "00:00:00"
// @Param end_time query string true "00:00:00"
// @Success 200 {string} json "{"code":"200","msg":"上传成功"}"
// @Failure 400 {string} json "{"code":"400","msg":"上传失败"}"
// @Router /api/videos/upload [post]
func UploadVideo(c *gin.Context) {
	file, err := c.FormFile("file")
	//start_time := c.Query("start_time")
	//end_time := c.Query("end_time")
	//
	//// 解析时间戳
	//startTimeParts := strings.Split(start_time, ":")
	//endTimeParts := strings.Split(end_time, ":")
	//
	//// 计算两个时间戳之间的差值（以秒为单位）
	//startSeconds := toSeconds(startTimeParts)
	//endSeconds := toSeconds(endTimeParts)
	//durationSeconds := endSeconds - startSeconds
	//
	//// 将差值转换为字符串类型
	//durationStr := fmt.Sprintf("%02d:%02d:%02d", durationSeconds/3600, (durationSeconds/60)%60, durationSeconds%60)
	//
	//// 创建临时文件-----------------
	//src, _ := file.Open()
	//defer src.Close()
	//tmpFile, _ := os.CreateTemp("", "tmp-*.mp4")
	//defer os.Remove(tmpFile.Name())
	//
	//// 将上传的文件复制到临时文件中
	//io.Copy(tmpFile, src)
	//tmpFile.Close()
	//
	//// 使用ffmpeg-go库进行剪辑
	//ffmpeg_go.Input(tmpFile.Name()).
	//	//Trim(ffmpeg_go.KwArgs{"start_frame": "10", "end_frame": "30"}).
	//	//Trim(ffmpeg_go.KwArgs{"start_time": "00:00:10", "end_time": "00:00:15"}).
	//	Output("output.mp4", ffmpeg_go.KwArgs{"ss": start_time, "t": durationStr}). // 从10秒开始，往后15秒
	//	OverWriteOutput().
	//	Run()
	////---------------------------------

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "上传失败",
		})
		return
	}

	// 获取上传文件的后缀名
	ext := strings.ToLower(filepath.Ext(file.Filename))

	// 判断后缀名是否为视频文件
	if ext != ".mp4" && ext != ".avi" && ext != ".mkv" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "上传的文件不是视频文件",
		})
		return
	}

	// 判断文件大小是否超过限制
	if file.Size > maxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "上传的文件大小超过限制",
		})
		return
	}

	// 保存文件到本地磁盘
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filepath := filepath.Join("upload", filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil { // 保存的原文件
		log.Println("保存文件失败：", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "上传失败",
		})
		return
	}

	// 打开上传的文件
	f, err := os.Open(filepath)

	start_time := c.Query("start_time")
	end_time := c.Query("end_time")

	// 解析时间戳
	startTimeParts := strings.Split(start_time, ":")
	endTimeParts := strings.Split(end_time, ":")

	// 计算两个时间戳之间的差值（以秒为单位）
	startSeconds := toSeconds(startTimeParts)
	endSeconds := toSeconds(endTimeParts)
	durationSeconds := endSeconds - startSeconds

	// 将差值转换为字符串类型
	durationStr := fmt.Sprintf("%02d:%02d:%02d", durationSeconds/3600, (durationSeconds/60)%60, durationSeconds%60)

	// 创建临时文件-----------------
	src, _ := file.Open()
	defer src.Close()
	tmpFile, _ := os.CreateTemp("", "tmp-*.mp4")
	defer os.Remove(tmpFile.Name())

	// 将上传的文件复制到临时文件中
	io.Copy(tmpFile, src)
	tmpFile.Close()

	cfilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	Cfilepath := "clipVideo\\" + cfilename
	outputFile, err := os.Create(cfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()
	// 使用ffmpeg-go库进行剪辑
	ffmpeg_go.Input(tmpFile.Name()).
		//Trim(ffmpeg_go.KwArgs{"start_frame": "10", "end_frame": "30"}).
		//Trim(ffmpeg_go.KwArgs{"start_time": "00:00:10", "end_time": "00:00:15"}).
		Output("clipVideo/"+cfilename, ffmpeg_go.KwArgs{"ss": start_time, "t": durationStr}). // 从10秒开始，往后15秒
		OverWriteOutput().
		//WithOutput(outputFile).
		Run()
	//---------------------------------

	if err != nil {
		log.Println("打开上传的文件失败：", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "上传失败",
		})
		return
	}
	defer f.Close()

	cf, err := os.Open(Cfilepath)
	if err != nil {
		log.Println("打开上传的文件失败：", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "上传失败",
		})
		return
	}
	// 分批次读取文件数据并保存到数据库
	var offset int64 = 0
	for {
		// 读取一段数据
		buf := make([]byte, maxPacketSize)
		n, err := cf.ReadAt(buf, offset)
		if err != nil && err != io.EOF {
			log.Println("读取上传的文件失败：", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "上传失败",
			})
			return
		}

		// 如果没有数据可读，退出循环
		if n == 0 {
			break
		}

		// 将数据存储到数据库中
		video := models.Video{
			Name:  file.Filename,
			Video: buf[:n],
		}
		if err := models.DB.Create(&video).Error; err != nil {
			log.Println("保存视频数据失败：", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "上传失败",
			})
			return
		}

		// 更新偏移量
		offset += int64(n)
	}

	//if err != nil {
	//	log.Println("打开上传的文件失败：", err)
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": 400,
	//		"msg":  "上传失败",
	//	})
	//	return
	//}
	//defer f.Close()
	//
	//// 分批次读取文件数据并保存到数据库
	//var offset int64 = 0
	//for {
	//	// 读取一段数据
	//	buf := make([]byte, maxPacketSize)
	//	n, err := f.ReadAt(buf, offset)
	//	if err != nil && err != io.EOF {
	//		log.Println("读取上传的文件失败：", err)
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"code": 400,
	//			"msg":  "上传失败",
	//		})
	//		return
	//	}
	//
	//	// 如果没有数据可读，退出循环
	//	if n == 0 {
	//		break
	//	}
	//
	//	// 将数据存储到数据库中
	//	video := models.Video{
	//		Name:  file.Filename,
	//		Video: buf[:n],
	//	}
	//	if err := models.DB.Create(&video).Error; err != nil {
	//		log.Println("保存视频数据失败：", err)
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"code": 400,
	//			"msg":  "上传失败",
	//		})
	//		return
	//	}
	//
	//	// 更新偏移量
	//	offset += int64(n)
	//}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "上传成功",
	})
}

// 将字符串类型的时间戳转换为秒数
func toSeconds(parts []string) int {
	hours, _ := strconv.Atoi(parts[0])
	minutes, _ := strconv.Atoi(parts[1])
	seconds, _ := strconv.Atoi(parts[2])
	return hours*3600 + minutes*60 + seconds
}
