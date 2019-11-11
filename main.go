/*
@Time : 19-11-10 下午2:48 
@Author : itcast
@File : main
@Software: GoLand
*/

package main

import (
	"github.com/gin-gonic/gin"
	"ihome/model"
	"fmt"
	"ihome/controller"
)

func main() {
	//初始化路由
	router := gin.Default()
	//数据库处理
	model.InitRedis()
	err := model.InitDb()
	if err != nil {
		//把错误打印到日志文件中
		fmt.Println(err)
		return
	}

	//路由模块
	//router.Group()
	//展示静态页面
	//静态路由
	router.Static("/home", "view")

	r1 := router.Group("/api/v1.0")
	{
		//路由规范
		r1.GET("/areas", controller.GetArea)
	}

	router.Run(":8099")
}
