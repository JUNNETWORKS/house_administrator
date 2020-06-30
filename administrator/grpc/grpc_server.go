package grpc_services

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "./pb"
)

func Serve() {
	var port string = ":1919"
	lis, err := net.Listen("tcp", fmt.Sprintf(port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMeasurementStoreServer(grpcServer, &MeasurementStoreServer{})
	grpcServer.Serve(lis)
}
