package controllers

import (
    "github.com/birjemin/micro-shippy/api/services"
    "github.com/gin-gonic/gin"
    "github.com/spf13/cast"
)

type VesselController struct {
    Service services.IVesselService
}

var vesselObj = &VesselController{Service: services.NewVesselService()}

func CreateVessel(g *gin.Context) {
    id := g.PostForm("id")
    name := g.PostForm("name")
    maxWeight := g.PostForm("maxWeight")
    capacity := g.PostForm("capacity")
    if id == "" || name == "" || maxWeight == "" || capacity == "" {
        JsonReturn(g, -1, "条件不符合", "")
        return
    }
    res := vesselObj.Service.Create(id, name, cast.ToInt32(maxWeight), cast.ToInt32(capacity))
    if !res {
        JsonReturn(g, -1, "创建失败啦", res)
        return
    }
    JsonReturn(g, 0, "创建成功", "")
    return
}
