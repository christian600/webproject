package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("./static/html/index.html", "./static/html/upload.html")
	//加载静态资源，例如网页的css、js
	router.Static("/static", "./static")

	//加载静态资源，一般是上传的资源，例如用户上传的图片
	router.StaticFS("/upload", http.Dir("upload"))

	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/upload", func(context *gin.Context) {
		context.HTML(http.StatusOK, "upload.html", nil)
	})
	router.POST("/upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		// 上传文件至指定目录
		if err := context.SaveUploadedFile(file, "./upload/"+file.Filename); err != nil {
			fmt.Println(err)
		}
		context.HTML(http.StatusOK, "upload.html", gin.H{"file": "/upload/" + file.Filename})
	})
	router.Run(":8080")
}
