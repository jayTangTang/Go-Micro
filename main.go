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
)

func main() {
	//初始化路由
	router := gin.Default()

	//请求分配  路由分组   前端  错误码文件
	r1 := router.Group("/v1")
	{
		r1.GET("/abc", func(ctx *gin.Context) {
			ctx.Writer.WriteString("abcd")
		})
	}
	r2 := router.Group("/v2")
	{
		r2.GET("/abc", func(ctx *gin.Context) {
			ctx.Writer.WriteString("dcba")
		})
	}

	model.InitModel()
	//model.InsertData()
	model.SearchData()
	//model.UpdateData()
	//model.DaleteData()
	router.Run(":8099")
}
