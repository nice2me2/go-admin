package router

import (
	"go-admin/app/admin/apis"
	"go-admin/common/middleware"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysCustomerRouter)
}

// 需认证的路由代码
func registerSysCustomerRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysCustomer{}
	r := v1.Group("/customer").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.POST("/upload/file", api.UploadCustomerFile)
	}

	//r1 := v1.Group("/customerKey").Use(authMiddleware.MiddlewareFunc())
	//{
	//	r1.GET("/:customerKey", api.GetSysCustomerByKEYForService)
	//}
	//
	//r2 := v1.Group("/app-customer")
	//{
	//	r2.GET("", api.Get2SysApp)
	//}
	//
	//r3 := v1.Group("/set-customer").Use(authMiddleware.MiddlewareFunc())
	//{
	//	r3.PUT("", api.Update2Set)
	//	r3.GET("", api.Get2Set)
	//}

}
