package controllers

import (
    "github.com/birjemin/micro-shippy/api/services"
    "github.com/gin-gonic/gin"
)

type UserController struct {
    Service services.IUserService
}

var userObj = &UserController{Service: services.NewUserService()}

func GetUsers(g *gin.Context) {
    m := make(map[string]interface{}, 3)
    page := g.Query("page")
    size := g.Query("size")
    if page == "" {
        JsonReturn(g, -1, "page不能为空", "")
        return
    }
    if size == "" {
        JsonReturn(g, -1, "size不能为空", "")
        return
    }
    m["page"] = page
    m["size"] = size
    m["data"] = userObj.Service.List(m)
    JsonReturn(g, 0, "", m)
    return
}

func GetUser(g *gin.Context) {
    id := g.Param("id")
    password := g.Query("password")
    if id == "" || password == "" {
        JsonReturn(g, -1, "id不能为空", "")
        return
    }
    user := userObj.Service.Get(id)
    token := userObj.Service.Auth(user.Email, password)
    JsonReturn(g, 0, "", map[string]interface{}{
        "user": user,
        "token": token,
    })
    return
}

func CreateUser(g *gin.Context) {
    name := g.PostForm("name")
    email := g.PostForm("email")
    password := g.PostForm("password")
    company := g.PostForm("company")
    if name == "" || email == "" || password == "" || company == "" {
        JsonReturn(g, -1, "条件不符合", "")
        return
    }
    res := userObj.Service.Create(name, email, password, company)

    if !res {
        JsonReturn(g, -1, "创建失败啦", res)
        return
    }
    JsonReturn(g, 0, "创建成功", "")
    return
}
