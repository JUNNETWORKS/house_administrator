package main

import (
	"flag"
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
	var (
		portNum int
	)
	/* コマンドラインオプションの定義 */
	flag.IntVar(&portNum, "p", 1919, "gRPCサーバーの起動に使うポート番号")
	flag.Parse()

	serve(portNum)
}
