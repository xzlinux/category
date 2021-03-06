package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul "github.com/micro/go-plugins/registry/consul/v2"
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
		options.Addrs = []string{
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
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", mysqlInfo.User, mysqlInfo.Pwd, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.Database)
	db, err := gorm.Open("mysql", connArgs)

	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)
	rp := repository.NewCategoryRepository(db)
	rp.InitTable()
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
