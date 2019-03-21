package main

import (
    "context"
    "log"
    "net"

    pb "github.com/mjkong/selfstudy/grpc/user"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

const (
    port = ":50051"
)

type server struct{}

func (s *server) RegistUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
  log.Printf("Name: %s, Email: %s, Phone: %s", in.Name, in.Email, in.Phone)
  return &pb.UserResponse{Message: "Name " + in.Name}, nil
}


func main(){
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatal("failed to listen: %v", err)
  }

  s := grpc.NewServer()
  pb.RegisterUserServer(s, &server{})

  reflection.Register(s)
  if err := s.Serve(lis); err != nil {
    log.Fatal("failed to serve: %v", err)
  }
}

