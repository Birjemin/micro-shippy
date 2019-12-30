package services

import (
    "context"
    userPb "github.com/birjemin/micro-shippy/user-service/proto/user"
    "github.com/micro/go-micro/client"
    "log"
)

type IUserService interface {
    Get(Id string) *userPb.User
    List(m map[string]interface{}) map[string]interface{}
    Create(name, email, password, company string) bool
    Auth(email, password string) (res string)
}

type userService struct {
    client userPb.UserServiceClient
}

func NewUserService() IUserService {
    return &userService{
        client: userPb.NewUserServiceClient("go.micro.srv.user", client.DefaultClient),
    }
}

func (u userService) Create(name, email, password, company string) bool {
    resp, err := u.client.Create(context.TODO(), &userPb.User{
        Name:     name,
        Email:    email,
        Password: password,
        Company:  company,
    })
    if err != nil {
        log.Fatalf("create user failed: %v", err)
        return false
    } else {
        log.Printf("create user success: %v", resp)
        return true
    }
}

func (u userService) Get(Id string) *userPb.User {
    resp, err := u.client.Get(context.TODO(), &userPb.User{
        Id: Id,
    })
    if err != nil {
        log.Printf("get user failed: %v", err)
        return nil
    } else {
        return resp.User
    }
}

func (u userService) List(m map[string]interface{}) map[string]interface{} {
    resp, err := u.client.GetAll(context.TODO(), &userPb.Request{})
    maps := make(map[string]interface{}, 2)
    if err != nil {
        log.Printf("get users failed: %v", err)
    } else {
        maps["Total"] = len(resp.Users)
        maps["List"] = resp.Users
    }
    return maps
}

func (u userService) Auth(email, password string) (res string) {
    authResponse, err := u.client.Auth(context.TODO(), &userPb.User{
        Email:    email,
        Password: password,
    })
    if err != nil {
        log.Printf("Could not authenticate user: %s error: %v\n", email, err)
        return ""
    }
    return authResponse.Token
}
