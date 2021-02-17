package main

import (
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/xzlinux/category/common"
	"github.com/xzlinux/category/domain/repository"
	service2 "github.com/xzlinux/category/domain/service"
	"github.com/xzlinux/category/handler"
	category "github.com/xzlinux/category/proto/category"
)

func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addr = []string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8082"),
		micro.Registry(consulRegistry),
	)
	//获取mysql 配置，路径中悄带前缀
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@"+mysqlInfo.Host+"/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)

	service.Init()

	//注册handler
	categoryDataService := service2.NewCategoryDataService(repository.NewCategoryRepository(db))
	err = category.RegisterCategoryHandler(service.Server(), &handler.Category{CategoryDataService: categoryDataService})
	if err != nil {
		log.Error(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
