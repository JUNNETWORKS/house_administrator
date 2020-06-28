from concurrent.futures import ThreadPoolExecutor

import grpc

from measurement_recorder import MeasurementRecorder
from pb import measurement_pb2_grpc


def serve():
    # Serverオブジェクトを作成する
    server = grpc.server(ThreadPoolExecutor(2))

    # Serverオブジェクトに定義したServicerクラスを登録する
    measurement_pb2_grpc.add_MeasurementRecorderServicer_to_server(
        MeasurementRecorder(), server,
    )

    # 1234番ポートで待ち受けするよう指定する
    server.add_insecure_port('[::]:1234')

    # 待ち受けを開始する
    server.start()

    # 待ち受け終了後の後処理を実行する
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
