package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sukenda/go-grpc-api-gateway/pkg/auth"
	"github.com/sukenda/go-grpc-api-gateway/pkg/config"
	"github.com/sukenda/go-grpc-api-gateway/pkg/order"
	"github.com/sukenda/go-grpc-api-gateway/pkg/product"
)

// From https://levelup.gitconnected.com/microservices-with-go-grpc-api-gateway-and-authentication-part-1-2-393ad9fc9d30
func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
