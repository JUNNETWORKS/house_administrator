// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: measurement.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 測定値の種類
type MeasurementType int32

const (
	MeasurementType_TEMPERATURE       MeasurementType = 0 // 温度
	MeasurementType_HUMIDITY          MeasurementType = 1 // 湿度
	MeasurementType_AIR_PRESSURE      MeasurementType = 2 // 気圧
	MeasurementType_CO2_CONCENTRATION MeasurementType = 3 // CO2濃度
)

// Enum value maps for MeasurementType.
var (
	MeasurementType_name = map[int32]string{
		0: "TEMPERATURE",
		1: "HUMIDITY",
		2: "AIR_PRESSURE",
		3: "CO2_CONCENTRATION",
	}
	MeasurementType_value = map[string]int32{
		"TEMPERATURE":       0,
		"HUMIDITY":          1,
		"AIR_PRESSURE":      2,
		"CO2_CONCENTRATION": 3,
	}
)

func (x MeasurementType) Enum() *MeasurementType {
	p := new(MeasurementType)
	*p = x
	return p
}

func (x MeasurementType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MeasurementType) Descriptor() protoreflect.EnumDescriptor {
	return file_measurement_proto_enumTypes[0].Descriptor()
}

func (MeasurementType) Type() protoreflect.EnumType {
	return &file_measurement_proto_enumTypes[0]
}

func (x MeasurementType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MeasurementType.Descriptor instead.
func (MeasurementType) EnumDescriptor() ([]byte, []int) {
	return file_measurement_proto_rawDescGZIP(), []int{0}
}

// 測定値メッセージの定義
type Measurement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomNumber RoomNumbers     `protobuf:"varint,1,opt,name=room_number,json=roomNumber,proto3,enum=room.RoomNumbers" json:"room_number,omitempty"` // 部屋番号
	Type       MeasurementType `protobuf:"varint,2,opt,name=type,proto3,enum=measurement.MeasurementType" json:"type,omitempty"`                    // 測定値の種類
	Value      float32         `protobuf:"fixed32,3,opt,name=value,proto3" json:"value,omitempty"`                                                  // 実際の測定値
	Datetime   int64           `protobuf:"varint,4,opt,name=datetime,proto3" json:"datetime,omitempty"`                                             // 測定した時間(unix_time)
}

func (x *Measurement) Reset() {
	*x = Measurement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measurement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Measurement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Measurement) ProtoMessage() {}

func (x *Measurement) ProtoReflect() protoreflect.Message {
	mi := &file_measurement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Measurement.ProtoReflect.Descriptor instead.
func (*Measurement) Descriptor() ([]byte, []int) {
	return file_measurement_proto_rawDescGZIP(), []int{0}
}

func (x *Measurement) GetRoomNumber() RoomNumbers {
	if x != nil {
		return x.RoomNumber
	}
	return RoomNumbers_JUN_ROOM
}

func (x *Measurement) GetType() MeasurementType {
	if x != nil {
		return x.Type
	}
	return MeasurementType_TEMPERATURE
}

func (x *Measurement) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Measurement) GetDatetime() int64 {
	if x != nil {
		return x.Datetime
	}
	return 0
}

type MeasurementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type MeasurementType `protobuf:"varint,1,opt,name=type,proto3,enum=measurement.MeasurementType" json:"type,omitempty"`
}

func (x *MeasurementRequest) Reset() {
	*x = MeasurementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measurement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeasurementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeasurementRequest) ProtoMessage() {}

func (x *MeasurementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_measurement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeasurementRequest.ProtoReflect.Descriptor instead.
func (*MeasurementRequest) Descriptor() ([]byte, []int) {
	return file_measurement_proto_rawDescGZIP(), []int{1}
}

func (x *MeasurementRequest) GetType() MeasurementType {
	if x != nil {
		return x.Type
	}
	return MeasurementType_TEMPERATURE
}

type MeasurementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Measurement *Measurement `protobuf:"bytes,1,opt,name=measurement,proto3" json:"measurement,omitempty"`
	Message     string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Error       bool         `protobuf:"varint,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *MeasurementResponse) Reset() {
	*x = MeasurementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measurement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeasurementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeasurementResponse) ProtoMessage() {}

func (x *MeasurementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_measurement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeasurementResponse.ProtoReflect.Descriptor instead.
func (*MeasurementResponse) Descriptor() ([]byte, []int) {
	return file_measurement_proto_rawDescGZIP(), []int{2}
}

func (x *MeasurementResponse) GetMeasurement() *Measurement {
	if x != nil {
		return x.Measurement
	}
	return nil
}

func (x *MeasurementResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MeasurementResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

var File_measurement_proto protoreflect.FileDescriptor

var file_measurement_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x72,
	0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x0b, 0x4d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x0b, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x30, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x22, 0x46, 0x0a, 0x12, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x13, 0x4d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x59, 0x0a,
	0x0f, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x45, 0x4d, 0x50, 0x45, 0x52, 0x41, 0x54, 0x55, 0x52, 0x45, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x55, 0x4d, 0x49, 0x44, 0x49, 0x54, 0x59, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x41, 0x49, 0x52, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x53, 0x55, 0x52, 0x45, 0x10,
	0x02, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x32, 0x5f, 0x43, 0x4f, 0x4e, 0x43, 0x45, 0x4e, 0x54,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x32, 0xa9, 0x01, 0x0a, 0x13, 0x4d, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x4a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x06,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20,
	0x2e, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_measurement_proto_rawDescOnce sync.Once
	file_measurement_proto_rawDescData = file_measurement_proto_rawDesc
)

func file_measurement_proto_rawDescGZIP() []byte {
	file_measurement_proto_rawDescOnce.Do(func() {
		file_measurement_proto_rawDescData = protoimpl.X.CompressGZIP(file_measurement_proto_rawDescData)
	})
	return file_measurement_proto_rawDescData
}

var file_measurement_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_measurement_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_measurement_proto_goTypes = []interface{}{
	(MeasurementType)(0),        // 0: measurement.MeasurementType
	(*Measurement)(nil),         // 1: measurement.Measurement
	(*MeasurementRequest)(nil),  // 2: measurement.MeasurementRequest
	(*MeasurementResponse)(nil), // 3: measurement.MeasurementResponse
	(RoomNumbers)(0),            // 4: room.RoomNumbers
	(*empty.Empty)(nil),         // 5: google.protobuf.Empty
}
var file_measurement_proto_depIdxs = []int32{
	4, // 0: measurement.Measurement.room_number:type_name -> room.RoomNumbers
	0, // 1: measurement.Measurement.type:type_name -> measurement.MeasurementType
	0, // 2: measurement.MeasurementRequest.type:type_name -> measurement.MeasurementType
	1, // 3: measurement.MeasurementResponse.measurement:type_name -> measurement.Measurement
	2, // 4: measurement.MeasurementRecorder.Get:input_type -> measurement.MeasurementRequest
	5, // 5: measurement.MeasurementRecorder.GetAll:input_type -> google.protobuf.Empty
	3, // 6: measurement.MeasurementRecorder.Get:output_type -> measurement.MeasurementResponse
	3, // 7: measurement.MeasurementRecorder.GetAll:output_type -> measurement.MeasurementResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_measurement_proto_init() }
func file_measurement_proto_init() {
	if File_measurement_proto != nil {
		return
	}
	file_room_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_measurement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Measurement); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_measurement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeasurementRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_measurement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeasurementResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_measurement_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_measurement_proto_goTypes,
		DependencyIndexes: file_measurement_proto_depIdxs,
		EnumInfos:         file_measurement_proto_enumTypes,
		MessageInfos:      file_measurement_proto_msgTypes,
	}.Build()
	File_measurement_proto = out.File
	file_measurement_proto_rawDesc = nil
	file_measurement_proto_goTypes = nil
	file_measurement_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MeasurementRecorderClient is the client API for MeasurementRecorder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MeasurementRecorderClient interface {
	Get(ctx context.Context, in *MeasurementRequest, opts ...grpc.CallOption) (*MeasurementResponse, error)
	GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (MeasurementRecorder_GetAllClient, error)
}

type measurementRecorderClient struct {
	cc grpc.ClientConnInterface
}

func NewMeasurementRecorderClient(cc grpc.ClientConnInterface) MeasurementRecorderClient {
	return &measurementRecorderClient{cc}
}

func (c *measurementRecorderClient) Get(ctx context.Context, in *MeasurementRequest, opts ...grpc.CallOption) (*MeasurementResponse, error) {
	out := new(MeasurementResponse)
	err := c.cc.Invoke(ctx, "/measurement.MeasurementRecorder/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measurementRecorderClient) GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (MeasurementRecorder_GetAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MeasurementRecorder_serviceDesc.Streams[0], "/measurement.MeasurementRecorder/GetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &measurementRecorderGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MeasurementRecorder_GetAllClient interface {
	Recv() (*MeasurementResponse, error)
	grpc.ClientStream
}

type measurementRecorderGetAllClient struct {
	grpc.ClientStream
}

func (x *measurementRecorderGetAllClient) Recv() (*MeasurementResponse, error) {
	m := new(MeasurementResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MeasurementRecorderServer is the server API for MeasurementRecorder service.
type MeasurementRecorderServer interface {
	Get(context.Context, *MeasurementRequest) (*MeasurementResponse, error)
	GetAll(*empty.Empty, MeasurementRecorder_GetAllServer) error
}

// UnimplementedMeasurementRecorderServer can be embedded to have forward compatible implementations.
type UnimplementedMeasurementRecorderServer struct {
}

func (*UnimplementedMeasurementRecorderServer) Get(context.Context, *MeasurementRequest) (*MeasurementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedMeasurementRecorderServer) GetAll(*empty.Empty, MeasurementRecorder_GetAllServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

func RegisterMeasurementRecorderServer(s *grpc.Server, srv MeasurementRecorderServer) {
	s.RegisterService(&_MeasurementRecorder_serviceDesc, srv)
}

func _MeasurementRecorder_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeasurementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasurementRecorderServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/measurement.MeasurementRecorder/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasurementRecorderServer).Get(ctx, req.(*MeasurementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasurementRecorder_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MeasurementRecorderServer).GetAll(m, &measurementRecorderGetAllServer{stream})
}

type MeasurementRecorder_GetAllServer interface {
	Send(*MeasurementResponse) error
	grpc.ServerStream
}

type measurementRecorderGetAllServer struct {
	grpc.ServerStream
}

func (x *measurementRecorderGetAllServer) Send(m *MeasurementResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MeasurementRecorder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "measurement.MeasurementRecorder",
	HandlerType: (*MeasurementRecorderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _MeasurementRecorder_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _MeasurementRecorder_GetAll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "measurement.proto",
}