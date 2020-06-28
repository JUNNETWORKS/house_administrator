package MeasurementStoreServer

import (
	"context"
	"time"

	pb "./pb"
	"github.com/golang/protobuf/ptypes/empty"
)

type MeasurementStoreServer {}

func (*MeasurementStoreServer) InsertData(context.Context, *Measurement) (*InsertResponse, error){
	// TODO: DBにデータを挿入する処理

	res := pb.InsertResponse{
		message: "SUCCEEDED"
		error: false
	}
	return &res, nil
}