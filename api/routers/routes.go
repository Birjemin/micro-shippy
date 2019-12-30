package routers

import (
    "github.com/birjemin/micro-shippy/api/controllers"
    "github.com/gin-gonic/gin"
)

func SetRouters(g *gin.Engine) {
    consignment := g.Group("/api/consignment")
    user := g.Group("/api/user")
    vessel := g.Group("/api/vessel")
    {
        consignment.POST("", controllers.CreateConsignment)
        user.GET("", controllers.GetUsers)
        user.GET("/:id", controllers.GetUser)
        user.POST("", controllers.CreateUser)
        vessel.POST("", controllers.CreateVessel)
    }
}
