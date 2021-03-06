# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2

from . import measurement_pb2 as measurement__pb2


class MeasurementRecorderStub(object):
    """Missing associated documentation comment in .proto file"""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Get = channel.unary_unary(
            '/measurement.MeasurementRecorder/Get',
            request_serializer=measurement__pb2.MeasurementRequest.SerializeToString,
            response_deserializer=measurement__pb2.MeasurementResponse.FromString,
        )
        self.GetAvailableTypes = channel.unary_unary(
            '/measurement.MeasurementRecorder/GetAvailableTypes',
            request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            response_deserializer=measurement__pb2.AvailableTypes.FromString,
        )


class MeasurementRecorderServicer(object):
    """Missing associated documentation comment in .proto file"""

    def Get(self, request, context):
        """Administrator -> Servant
        ある特定の種類のセンサーの値を返す
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAvailableTypes(self, request, context):
        """利用可能なセンサーの種類を返す
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MeasurementRecorderServicer_to_server(servicer, server):
    rpc_method_handlers = {
        'Get': grpc.unary_unary_rpc_method_handler(
            servicer.Get,
            request_deserializer=measurement__pb2.MeasurementRequest.FromString,
            response_serializer=measurement__pb2.MeasurementResponse.SerializeToString,
        ),
        'GetAvailableTypes': grpc.unary_unary_rpc_method_handler(
            servicer.GetAvailableTypes,
            request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            response_serializer=measurement__pb2.AvailableTypes.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        'measurement.MeasurementRecorder', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

 # This class is part of an EXPERIMENTAL API.


class MeasurementRecorder(object):
    """Missing associated documentation comment in .proto file"""

    @staticmethod
    def Get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/measurement.MeasurementRecorder/Get',
            measurement__pb2.MeasurementRequest.SerializeToString,
            measurement__pb2.MeasurementResponse.FromString,
            options,
            channel_credentials,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata)

    @staticmethod
    def GetAvailableTypes(request,
                          target,
                          options=(),
                          channel_credentials=None,
                          call_credentials=None,
                          compression=None,
                          wait_for_ready=None,
                          timeout=None,
                          metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/measurement.MeasurementRecorder/GetAvailableTypes',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            measurement__pb2.AvailableTypes.FromString,
            options,
            channel_credentials,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata)


class MeasurementStoreStub(object):
    """Missing associated documentation comment in .proto file"""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.InsertData = channel.unary_unary(
            '/measurement.MeasurementStore/InsertData',
            request_serializer=measurement__pb2.Measurement.SerializeToString,
            response_deserializer=measurement__pb2.InsertResponse.FromString,
        )


class MeasurementStoreServicer(object):
    """Missing associated documentation comment in .proto file"""

    def InsertData(self, request, context):
        """Administrator <- Servant
        ある特定の種類のセンサーの値をDBに登録する
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MeasurementStoreServicer_to_server(servicer, server):
    rpc_method_handlers = {
        'InsertData': grpc.unary_unary_rpc_method_handler(
            servicer.InsertData,
            request_deserializer=measurement__pb2.Measurement.FromString,
            response_serializer=measurement__pb2.InsertResponse.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        'measurement.MeasurementStore', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

 # This class is part of an EXPERIMENTAL API.


class MeasurementStore(object):
    """Missing associated documentation comment in .proto file"""

    @staticmethod
    def InsertData(request,
                   target,
                   options=(),
                   channel_credentials=None,
                   call_credentials=None,
                   compression=None,
                   wait_for_ready=None,
                   timeout=None,
                   metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/measurement.MeasurementStore/InsertData',
            measurement__pb2.Measurement.SerializeToString,
            measurement__pb2.InsertResponse.FromString,
            options,
            channel_credentials,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata)
