package main

import (
    "context"
    "encoding/json"
    pb "github.com/birjemin/micro-shippy/consignment-service/proto/consignment"
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/metadata"
    "io/ioutil"
    "log"
    "os"
)

const (
    defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
    var consignment *pb.Consignment
    data, err := ioutil.ReadFile(file)
    if err != nil {
        return nil, err
    }
    json.Unmarshal(data, &consignment)
    return consignment, err
}

func main() {
    service := micro.NewService(micro.Name("go.micro.cli.consignment"))
    service.Init()

    client := pb.NewShippingServiceClient("go.micro.srv.consignment", service.Client())

    // Contact the server and print out its response.
    file := defaultFilename
    if len(os.Args) > 1 {
        file = os.Args[1]
    }
    token := os.Getenv("TOKEN")

    consignment, err := parseFile(file)

    if err != nil {
        log.Fatalf("Could not parse file: %v", err)
    }

    ctx := metadata.NewContext(context.Background(), map[string]string{
        "token": token,
    })

    r, err := client.CreateConsignment(ctx, consignment)
    if err != nil {
        log.Fatalf("Could not create: %v", err)
    }
    log.Printf("Created: %t", r.Created)

    getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
    if err != nil {
        log.Fatalf("Could not list consignments: %v", err)
    }
    for _, v := range getAll.Consignments {
        log.Println(v)
    }
}
