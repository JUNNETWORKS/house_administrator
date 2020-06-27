import time

from co2 import read_ppm, use_uart
from pb import measurement_pb2, measurement_pb2_grpc, room_pb2

ROOM_NAME = "JUN_ROOM"


class MeasurementRecorder(measurement_pb2_grpc.MeasurementRecorderServicer):
    def Get(self, request: measurement_pb2.MeasurementRequest,
            context) -> measurement_pb2.MeasurementResponse:
        # 欲しいセンサーの種類
        request_sensor = request.type
        sensor_value = None
        if request_sensor == measurement_pb2.TEMPERATURE:
            pass
        elif request_sensor == measurement_pb2.HUMIDITY:
            pass
        elif request_sensor == measurement_pb2.AIR_PRESSURE:
            pass
        elif request_sensor == measurement_pb2.CO2_CONCENTRATION:
            pass

        sensor_name = measurement_pb2.MeasurementType.Name(request_sensor)
        room_number = room_pb2.RoomNumbers.Value(ROOM_NAME)
        now_unix = time.time()

        measurement_obj = measurement_pb2.Measurement(room_number=room_number,
                                                      type=request_sensor,
                                                      datetime=now_unix)

        res = measurement_pb2.MeasurementResponse()

        if sensor_value is None:
            measurement_obj["value"] = 0
            res["message"] = f"{sensor_name}センサーは{ROOM_NAME}に実装されていません."
            res["error"] = True
        else:
            measurement_obj["value"] = sensor_value
            res["message"] = f"{sensor_name}センサーからデータを取得しました."
            res["error"] = False
        res["measurement"] = measurement_obj

        return res

    def GetAvailableTypes(self, request, context):
        pass
