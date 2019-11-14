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
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/sessions"
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

	router.Use()

	//初始化redis容器
	store, err := redis.NewStore(20, "tcp", "127.0.0.1:6379", "", []byte("session"))
	if err != nil {
		fmt.Println("初始化session容器错误")
		return
	}

	store.Options(
		sessions.Options{
			MaxAge: 0,
		},
	)

	//路由使用中间件 gin中的session默认是生效时间是一个月
	router.Use(sessions.Sessions("mysession", store))

	//使用路由的时候就可以使用session中间件了
	router.GET("/session", func(context *gin.Context) {
		//初始化session对象
		se := sessions.Default(context)
		//设置session的时候除了set函数之外,必须调用save
		se.Set("test", "ihome")
		se.Save()

		context.Writer.WriteString("设置session")
	})

	//获取session
	router.GET("/getSession", func(context *gin.Context) {
		//初始化session对象
		se := sessions.Default(context)
		//获取session
		result := se.Get("test")
		fmt.Println("得到的session数据为", result.(string))
	})

	//测试
	router.GET("/test", func(context *gin.Context) {
		//设置cookie  cookie有两种,一种是有时间效应的,一种是临时cookie
		/*context.SetCookie("myTest","bj5q",0,"","",false,true)
		context.Writer.WriteString("测试cookie")*/
	})

	r1 := router.Group("/api/v1.0")
	{
		//路由规范
		r1.GET("/areas", controller.GetArea)
		r1.GET("/session", controller.GetSession)
		//传参方法,url传值,form表单传值,ajax传值,路径传值
		r1.GET("/imagecode/:uuid", controller.GetImageCd)
		r1.GET("/smscode/:mobile", controller.GetSmscd)
		r1.POST("/users", controller.PostRet)
	}

	router.Run(":8099")
}
