package routes

import (
	"mishop/app/controllers"

	"github.com/gin-gonic/gin"
)

func InitRourte(r *gin.Engine) {
	r.GET("/product/welcome", func(c *gin.Context) {
		controllers.TestController{}.Welcome(c)
	})

	r.GET("/json", func(c *gin.Context) {
		controllers.TestController{}.Json(c)
	})
	r.GET("/json2", func(c *gin.Context) {
		controllers.TestController{}.Json2(c)
	})

	r.GET("/test", func(c *gin.Context) {
		controllers.TestController{}.Test(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		controllers.TestController{}.Test2(c)
	})
	r.GET("/test3", func(c *gin.Context) {
		controllers.TestController{}.Test3(c)
	})

	r.GET("/redis", func(c *gin.Context) {
		controllers.TestController{}.Redis(c)
	})

	r.POST("/ljson", func(c *gin.Context) {
		controllers.TestController{}.Ljson(c)
	})
	r.POST("/fjson", func(c *gin.Context) {
		controllers.TestController{}.Fjson(c)
	})
}
