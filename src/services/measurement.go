package services

import (
	"context"
	"time"

	pb "./pb"
	"github.com/golang/protobuf/ptypes/empty"
)

type MeasurementRecorderServer struct {
}

// Get はセンサーから値を取得し, クライアントにMeasurementResponseオブジェクト返す関数です.
func (server *MeasurementRecorderServer) Get(_ context.Context, message *pb.MeasurementRequest) (*pb.MeasurementResponse, error) {
	var value float32
	switch message.Type {
	case pb.MeasurementType_TEMPERATURE:
		value = 10.5 // ℃
	case pb.MeasurementType_HUMIDITY:
		value = 92 // %
	case pb.MeasurementType_AIR_PRESSURE:
		value = 1030 // hPa
	case pb.MeasurementType_CO2_CONCENTRATION:
		value = 60 // %
	}
	currentTime := time.Now().Unix()
	measurement := pb.Measurement{
		RoomNumber: pb.RoomNumbers_JUN_ROOM,
		Type:       pb.MeasurementType_TEMPERATURE,
		Value:      value,
		Datetime:   currentTime,
	}
	response := &pb.MeasurementResponse{
		Measurement: &measurement,
		Message:     "Succeeded get temperature value",
		Error:       false,
	}
	return response, nil
}

func (server *MeasurementRecorderServer) GetAll(*empty.Empty, pb.MeasurementRecorder_GetAllServer) error {

}
