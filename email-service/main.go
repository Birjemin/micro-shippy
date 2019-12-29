package main

import (
    "encoding/json"
    pb "github.com/birjemin/micro-shippy/user-service/proto/user"
    "github.com/micro/go-micro"
    "github.com/micro/go-micro/broker"
    "log"
)

const topic = "user.created"

type Subscriber struct{}

func senEmail(user *pb.User) error {
    log.Println("Picked up a new message")
    log.Println("Sending email to:", user.Name)
    return nil
}

func main() {
    srv := micro.NewService(
        micro.Name("go.micro.srv.email"),
        micro.Version("latest"),
    )

    srv.Init()

    pubSub := srv.Server().Options().Broker

    log.Printf("[pubSub]address: %s\n", pubSub.Address())

    if err := pubSub.Connect(); err != nil {
        log.Fatalf("[pubSub]connect broker connect error: %v\n", err)
    }

    log.Printf("[pubSub]channel: %s\n", pubSub.String())

    _, err := pubSub.Subscribe(topic, func(pub broker.Event) error {
        var user *pb.User
        if err := json.Unmarshal(pub.Message().Body, &user); err != nil {
            return err
        }
        log.Printf("[Create User]: %v\n", user)
        go senEmail(user)
        return nil
    })

    if err != nil {
        log.Printf("sub error: %v\n", err)
    }

    // Run the server
    if err := srv.Run(); err != nil {
        log.Println(err)
    }
}
