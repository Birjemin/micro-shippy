package services

import (
    "context"
    "encoding/json"
    pb "github.com/birjemin/micro-shippy/consignment-service/proto/consignment"
    "github.com/micro/go-micro/client"
    "github.com/micro/go-micro/metadata"
    "log"
)

type IConsignmentService interface {
    List(token string) map[string]interface{}
    CreateConsignment(token, data string) bool
}

type consignmentService struct {
    client pb.ShippingServiceClient
}

func NewConsignmentService() IConsignmentService {
    return &consignmentService{
        client: pb.NewShippingServiceClient("go.micro.srv.consignment", client.DefaultClient),
    }
}

func (c consignmentService) CreateConsignment(token, data string) bool {
    var con *pb.Consignment
    err := json.Unmarshal([]byte(data), &con)
    if err != nil {
        return false
    }
    // 鉴权
    ctx := metadata.NewContext(context.Background(), map[string]string{
        "token": token,
    })

    res, err := c.client.CreateConsignment(ctx, con)
    if err != nil {
        log.Printf("Could not create: %v, %v", res, err)
        return false
    }
    return true
}

func (c consignmentService) List(token string) map[string]interface{} {
    ctx := metadata.NewContext(context.Background(), map[string]string{
        "token": token,
    })
    resp, err := c.client.GetConsignments(ctx, &pb.GetRequest{})
    maps := make(map[string]interface{}, 2)
    if err != nil {
        log.Printf("Could not list consignments: %v", err)
    } else {
        maps["Total"] = len(resp.Consignments)
        maps["List"] = resp.Consignments
    }
    return maps
}