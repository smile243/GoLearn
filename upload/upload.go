package upload

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUpload(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

func Upload(c *gin.Context) {
	//从请求中读取文件 multipart forms默认最大为8MB
	//通过修改router.MaxMultipartMemory 可以修改最大值
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//将读取到的文件保存在服务器本地
	c.SaveUploadedFile(file, file.Filename)
	c.JSON(http.StatusOK, gin.H{"message": "文件上传成功"})
}

//TODO 上传多个文件
