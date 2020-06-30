package grpc_services

import (
	"context"
	"fmt"

	pb "./pb"
)

type MeasurementStoreServer struct{}

func (server *MeasurementStoreServer) InsertData(_ context.Context, message *pb.Measurement) (*pb.InsertResponse, error) {
	value := message.Value
	fmt.Printf("Receive value: %f \n", value)

	// TODO: DBにデータを挿入する処理

	res := pb.InsertResponse{
		Message: "SUCCEEDED",
		Error:   false,
	}
	return &res, nil
}
