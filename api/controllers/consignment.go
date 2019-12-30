package controllers

import (
    "github.com/birjemin/micro-shippy/api/services"
    "github.com/gin-gonic/gin"
)

type ConsignmentController struct {
    Service services.IConsignmentService
}

var consignmentObj = &ConsignmentController{Service: services.NewConsignmentService()}

func CreateConsignment(g *gin.Context) {
    token := g.PostForm("token")
    data := g.PostForm("data")
    if data == "" {
        JsonReturn(g, -1, "报错啦", "")
        return
    }
    g.Set("token", token)
    r := consignmentObj.Service.CreateConsignment(token, data)
    if r {
        res := consignmentObj.Service.List(token)
        JsonReturn(g, 0, "", res)
        return
    }
    JsonReturn(g, 0, "", r)
    return
}
