package main

import (
	"fmt"
	"log"
	"net"

	grpcServices "./grpc_services"
	pb "./grpc_services/pb"
	"google.golang.org/grpc"
)

func serve(portNum int) {
	var port string = fmt.Sprintf(":%d", portNum)
	lis, err := net.Listen("tcp", fmt.Sprintf(port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMeasurementStoreServer(grpcServer, &grpcServices.MeasurementStoreServer{})
	grpcServer.Serve(lis)
}

func main() {
	serve(1919)
}
