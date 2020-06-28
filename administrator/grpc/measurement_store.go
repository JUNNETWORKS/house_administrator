package grpc_services

import (
	"context"

	pb "./pb"
)

type MeasurementStoreServer struct{}

func (server *MeasurementStoreServer) InsertData(_ context.Context, message *pb.Measurement) (*pb.InsertResponse, error) {
	// TODO: DBにデータを挿入する処理

	res := pb.InsertResponse{
		Message: "SUCCEEDED",
		Error:   false,
	}
	return &res, nil
}
