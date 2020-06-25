# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import control_pb2 as control__pb2


class GreetingStub(object):
    """Missing associated documentation comment in .proto file"""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.World = channel.unary_unary(
                '/control.Greeting/World',
                request_serializer=control__pb2.Hello.SerializeToString,
                response_deserializer=control__pb2.Hello.FromString,
                )


class GreetingServicer(object):
    """Missing associated documentation comment in .proto file"""

    def World(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_GreetingServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'World': grpc.unary_unary_rpc_method_handler(
                    servicer.World,
                    request_deserializer=control__pb2.Hello.FromString,
                    response_serializer=control__pb2.Hello.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'control.Greeting', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Greeting(object):
    """Missing associated documentation comment in .proto file"""

    @staticmethod
    def World(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/control.Greeting/World',
            control__pb2.Hello.SerializeToString,
            control__pb2.Hello.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)
