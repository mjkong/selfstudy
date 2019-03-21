package main

import (
    "context"
    "log"
    "os"
    "time"

    pb "github.com/mjkong/selfstudy/grpc/user"
    "google.golang.org/grpc"
)

const (
    address = "localhost:50051"
    defaultName = "mjkong"
)

func main(){
  conn, err := grpc.Dial(address, grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }

  defer conn.Close()
  c := pb.NewUserClient(conn)

  name := defaultName

  if len(os.Args) > 1 {
    name = os.Args[1]
  }


  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()

  r, err := c.RegistUser(ctx, &pb.UserRequest{Name: name, Email: "mjkong.kong@gmail.com", Phone: "0000"})
  if err != nil {
    log.Fatalf("could not regist: %v", err)
  }
  log.Printf("RegistUser: %s", r.Message)
}
