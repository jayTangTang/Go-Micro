/*
@Time : 19-11-13 下午4:18 
@Author : itcast
@File : tools
@Software: GoLand
*/

package utils

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro"
)

func GetMicroClient() client.Client {

	consulReg := consul.NewRegistry()
	microService := micro.NewService(
		micro.Registry(consulReg),
	)
	return microService.Client()
}
