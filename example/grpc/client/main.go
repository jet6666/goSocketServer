package main

import (
	"context"
	"flag"
	"goSocketServer/example/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)
var (
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)
const (
	defaultName = "world"
)
func main()  {
	log.Println("this is chat client starting ")

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Send(ctx, &pb.MsgRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not send: %v", err)
	}
	log.Printf("get message reply : %s", r.GetMessage())


	r2, err  := c.GetHistory(ctx, &pb.EmptyRequest{}  )
	if err != nil {
		log.Fatalf("could not GetHistory send: %v", err)
	}
	log.Printf("get meGetHistory  ssage reply : %s", r2.GetList() )

}
