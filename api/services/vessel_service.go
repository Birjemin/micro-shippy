package services

import (
    "context"
    vesselPb "github.com/birjemin/micro-shippy/vessel-service/proto/vessel"
    "github.com/micro/go-micro/client"
    "log"
)

type IVesselService interface {
    Create(id, name string, maxWeight, capacity int32) bool
}

type vesselService struct {
    client vesselPb.VesselServiceClient
}

func NewVesselService() IVesselService {
    return &vesselService{
        client: vesselPb.NewVesselServiceClient("go.micro.srv.vessel", client.DefaultClient),
    }
}

func (v vesselService) Create(id, name string, maxWeight, capacity int32) bool {
    resp, err := v.client.Create(context.TODO(), &vesselPb.Vessel{
        Id:        id,
        Name:      name,
        MaxWeight: maxWeight,
        Capacity:  capacity,
    })
    if err != nil {
        log.Fatalf("create vessel failed: %v", err)
        return false
    } else {
        log.Printf("create vessel success: %v", resp)
        return true
    }
}
