package service

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"videos.cn/models"
)

// GetVideosList
// @Tags 公共方法
// @Accept mpfd
// @Produce json
// @Summary 视频输入
// @Param file formData file true "上传文件"
// @Success 200 {string} json "{"code":"200","msg":"上传成功"}"
// @Failure 400 {string} json "{"code":"400","msg":"上传失败"}"
// @Router /videos-list [get]
func GetVideosList(c *gin.Context) {
	models.GetVideosList()

	http.HandleFunc("/videos-list", func(w http.ResponseWriter, r *http.Request) {
		// 读取上传的文件
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 将文件存储到数据库中
		models.DB.Exec("INSERT INTO video (video) VALUES (?)", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	c.String(http.StatusOK, "Get Problem List")
}
