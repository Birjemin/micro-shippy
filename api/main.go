package main

import (
    "github.com/birjemin/micro-shippy/api/middleware"
    "github.com/birjemin/micro-shippy/api/routers"
    "github.com/gin-gonic/gin"
    "github.com/micro/go-micro/web"
    "log"
)

func main() {

    // Create service
    service := web.NewService(
        web.Name("go.micro.api.all"),
        web.Address(":8080"),
    )

    router := gin.Default()
    router.Use(middleware.Cors())

    router.Static("/public", "/public")
    routers.SetRouters(router)

    // Register Handler
    service.Handle("/", router)

    // Run server
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}
