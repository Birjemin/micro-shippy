package main

import (
    "context"
    pb "github.com/birjemin/micro-shippy/user-service/proto/user"
    "github.com/micro/go-micro"
    "log"
)

const topic = "user.created"

type Subscriber struct{}

func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
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

    _ = micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

    // Run the server
    if err := srv.Run(); err != nil {
        log.Println(err)
    }
}
