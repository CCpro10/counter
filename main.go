package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/counter"
	"main/middleware"
	"net/http"
)

func main() {
	//初始化全局counter
	counter.Counter = counter.Init()
	go counter.Counter.Flush2broker(1000*3600*24, FlushCallBack)

	r := gin.Default()
	userRouter := r.Group("/user").Use(middleware.Counter)
	{
		userRouter.POST("/article/", ArticleHandler)
		userRouter.POST("/article_picture/", PictureHandler)
	}

	err := r.Run(":7777")
	if err != nil {
		panic(err)
	}
}

func FlushCallBack() {
	nodes := counter.Counter.GetAll()
	log.Println("In the past 24 hours")
	for _, node := range nodes {
		log.Printf("The %v path was visited %v times\n", node.Key, node.Count)
	}
}

func ArticleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "发送帖子成功",
	})
}

func PictureHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "上传图片成功",
	})
}
