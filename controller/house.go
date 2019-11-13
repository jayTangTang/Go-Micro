/*
@Time : 19-11-11 下午1:18 
@Author : itcast
@File : house
@Software: GoLand
*/

package controller

import (
	"github.com/gin-gonic/gin"
	getArea "ihome/proto/getArea"
	getImg "ihome/proto/getImg"
	"context"
	"fmt"
	"net/http"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"

	"ihome/service/getArea/utils"
	"github.com/afocus/captcha"
	"encoding/json"
	"image/png"
)

//获取所有地区信息
func GetArea(ctx *gin.Context) {
	/*
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)
	//获取数据库数据
	areas, err := model.GetArea()
	if err != nil {
		fmt.Println("获取所有地狱信息错误")
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		return
	}
	//把数据返回给前端
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = areas
	*/

	//调用远程服务,获取所有地域信息
	//初始化客户端
	//从consul中获取服务
	consulRegistry := consul.NewRegistry()
	micService := micro.NewService(
		micro.Registry(consulRegistry),
	)

	microClient := getArea.NewGetAreaService("go.micro.srv.getArea", micService.Client())
	//调用远程服务
	resp, err := microClient.MicroGetArea(context.TODO(), &getArea.Request{})
	if err != nil {
		fmt.Println(err)

		/*ctx.JSON(http.StatusOK,resp)
		return */
	}

	ctx.JSON(http.StatusOK, resp)
}

//写一个假的session请求返回
func GetSession(ctx *gin.Context) {
	//构造未登录
	resp := make(map[string]interface{})

	resp["errno"] = utils.RECODE_LOGINERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)

	ctx.JSON(http.StatusOK, resp)
}

//获取验证码图片方法
func GetImageCd(ctx *gin.Context) {
	//获取数据
	uuid := ctx.Param("uuid")
	//校验数据
	if uuid == "" {
		fmt.Println("获取数据错误")
		return
	}
	//处理数据

	//调用远程服务
	//初始化客户端
	consulReg := consul.NewRegistry()
	microService := micro.NewService(
		micro.Registry(consulReg),
	)

	microClient := getImg.NewGetImgService("go.micro.srv.getImg", microService.Client())
	//调用远程服务
	resp, err := microClient.MicroGetImg(context.TODO(), &getImg.Request{Uuid: uuid})

	//获取数据
	if err != nil {
		fmt.Println("获取远端数据失败:", err)
		ctx.JSON(http.StatusOK, resp)
		return
	}
	//返回json数据
	//反序列化拿到img数据
	var img captcha.Image
	json.Unmarshal(resp.Data, &img)

	png.Encode(ctx.Writer, img)
}
