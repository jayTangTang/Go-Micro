package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"ihome/service/getImg/handler"
	getImg "ihome/service/getImg/proto/getImg"
	"github.com/micro/go-micro/registry/consul"
	"ihome/service/getImg/model"
)

func main() {

	//使用consul做服务发现
	consulReg := consul.NewRegistry()

	model.InitRedis()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.getImg"),
		micro.Version("latest"),
		micro.Registry(consulReg),
		//隐藏bug   注册的服务注销了吗   不一定注销 65535
		micro.Address(":9981"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getImg.RegisterGetImgHandler(service.Server(), new(handler.GetImg))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
