package order

import (
	"github.com/gin-gonic/gin"
	"github.com/sukenda/go-grpc-api-gateway/pkg/auth"
	"github.com/sukenda/go-grpc-api-gateway/pkg/config"
	"github.com/sukenda/go-grpc-api-gateway/pkg/order/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routerGroup := r.Group("/order")
	routerGroup.Use(a.AuthRequired)
	routerGroup.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
