package main

import (
	"context"
	"flag"
	"fmt"
	"goSocketServer/example/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)
type ChatServer   struct {
	pb.UnimplementedChatServer
}
func (s *ChatServer) Send(ctx context.Context, in *pb.MsgRequest ) (*pb.MsgReply, error) {
	log.Printf("Received:hi   %v", in.GetName())
	return &pb.MsgReply{Message: "Hello " + in.GetName()}, nil
}

func (s *ChatServer) GetHistory(ctx context.Context, in *pb.EmptyRequest) (*pb.MsgListReply, error) {
	log.Println ("GetHistory  :")
	//list :=[]string {"1" ,"2 "}
	list2 :=make([]string ,5 )
	list2= append(list2, "AAAA")
	return &pb.MsgListReply{List:  list2 }, nil
}

var (
	port = flag.Int("port", 50052, "The server port")
)
func main()  {

	log.Println("this is grpc server starting  ")

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServer(s, &ChatServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
