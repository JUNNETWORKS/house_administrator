import time

import grpc

from pb import measurement_pb2, measurement_pb2_grpc, room_pb2

measurement_store_uri = "localhost:1919"

measurement = measurement_pb2.Measurement(
    room_number=room_pb2.JUN_ROOM,
    type=measurement_pb2.CO2_CONCENTRATION,
    value=440,
    datetime=int(time.time())
)


def send_measurement(measurement):
    with grpc.insecure_channel(measurement_store_uri) as channel:
        stub = measurement_pb2_grpc.MeasurementStoreStub(channel)
        res = stub.InsertData(measurement)
        print(res.message)


if __name__ == "__main__":
    send_measurement(measurement)
