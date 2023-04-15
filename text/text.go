package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
)

func main() {
	// 连接到 MySQL 数据库
	db, err := sql.Open("mysql", "root:huangzhitao123@/videocut")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
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
		_, err = db.Exec("INSERT INTO videos (data) VALUES (?)", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
