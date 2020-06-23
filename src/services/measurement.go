package services

import (
	"context"

	pb "./pb"
	"github.com/golang/protobuf/ptypes/empty"
)

type MeasurementRecorderServer struct {
}

func (server *MeasurementRecorderServer) Get(context.Context, *pb.MeasurementRequest) (*pb.MeasurementResponse, error) {

}
func (server *MeasurementRecorderServer) GetAll(*empty.Empty, pb.MeasurementRecorder_GetAllServer) error {

}
