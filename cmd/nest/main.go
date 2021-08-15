package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var (
	projectName          = "MiMo"
	rawCurrentPath       string
	rootPath             string
	resourcePath         string
	resourceRelativePath = "resource"
)

func init() {
	rawCurrentPath, _ = os.Getwd()
	rootPath = strings.Split(rawCurrentPath, projectName)[0]
	resourcePath = path.Join(rootPath, projectName, "resource")

	fmt.Println(rawCurrentPath)
	fmt.Println(rootPath)
	fmt.Println(resourcePath)
	//	检查是否有resource文件夹 以及该文件夹下是否有 doc knowledge video image 等目录， 没有则创建
}

func main() {
	r := gin.Default()
	r.GET("/home", func(c *gin.Context) {
		level := c.Query("level")
		list, _ := ioutil.ReadDir(resourcePath)
		fmt.Println(level)
		var listName []string
		for idx, _ := range list {
			listName = append(listName, list[idx].Name())
		}
		c.JSON(200, gin.H{
			"data":   listName,
			"status": "OK",
		})
	})
	r.GET("/dashboard", func(c *gin.Context) {

		dashboard := []map[string]string{
			{"name": "我的知识库", "code": "knowledge", "coverImg": ""},
			{"name": "广场", "code": "square", "coverImg": ""},
			{"name": "最近浏览", "code": "recent", "coverImg": ""},
		}
		c.JSON(200, gin.H{
			"data":   dashboard,
			"status": "OK",
		})
	})
	r.GET("/resource/list", func(c *gin.Context) {
		level := c.Query("level")
		list, _ := ioutil.ReadDir(resourcePath)
		fmt.Println(level)
		var listName []string
		for idx, _ := range list {
			listName = append(listName, list[idx].Name())
		}
		c.JSON(200, gin.H{
			"data":   listName,
			"status": "OK",
		})
	})
	r.GET("/resource/doc", func(c *gin.Context) {
		content, err := ioutil.ReadFile(resourcePath + "\\test.md") // 一次性读取文件中的所有内容，文件较大时候会影响性能
		if err != nil {
			fmt.Println("read file err:", err)
		}
		c.JSON(200, gin.H{
			"data":   string(content),
			"status": "OK",
		})
	})
	r.Run(":8999") // listen and serve on 0.0.0.0:8080
}
