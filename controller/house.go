/*
@Time : 19-11-11 下午1:18 
@Author : itcast
@File : house
@Software: GoLand
*/

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ihome/utils"
	"ihome/model"
	"fmt"
)

//获取所有地区信息
func GetArea(ctx *gin.Context) {

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

}
