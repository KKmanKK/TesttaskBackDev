package main

import (
	"restjwtgo/controllers"
	"restjwtgo/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.InitEnv()
	// initializers.ConnectToDB()
	initializers.InitDb()

}
func main() {

	r := gin.Default()
	r.POST("/api/singup", controllers.SingUp)

	r.GET("/api/refresh", controllers.Refresh)
	r.Run()
	defer initializers.CloseDb()
}
