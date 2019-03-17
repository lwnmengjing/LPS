package main

import (
	"github.com/lwnmengjing/LPS/handler"
	"github.com/lwnmengjing/LPS/models"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	domain "github.com/lwnmengjing/LPS/proto/domain"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.LPS"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	_ = domain.RegisterDomainHandler(service.Server(), new(handler.Domain))

	// Register Struct as Subscriber

	// Register Function as Subscriber
	//_ = micro.RegisterSubscriber("go.micro.srv.LPS", service.Server(), subscriber.Handler)

	//init db connection
	models.DBInit()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
