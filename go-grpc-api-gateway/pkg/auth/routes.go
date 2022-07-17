package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sukenda/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/sukenda/go-grpc-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routerGroup := r.Group("/auth")
	routerGroup.POST("/register", svc.Register)
	routerGroup.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
