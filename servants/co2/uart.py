import subprocess
from contextlib import contextmanager


@contextmanager
def use_uart():
    """UARTを使う時に使うコンテキストマネージャー

    デフォルトではUARTはLinuxコンソールに割り当てられているので,
    なので,UARTを使いたい場合にはLinuxシリアルコンソールを止める必要がある.
    """
    # モデルによってUARTに対する挙動が違う
    # ラズパイ zero W は以下の設定でOK
    # https://qiita.com/UedaTakeyuki/items/c5226960a7328155635f#%E4%BA%8B%E5%89%8D%E6%BA%96%E5%82%99uart%E3%82%92%E7%A9%BA%E3%81%91%E3%82%8B
    # ラズパイのモデルごとにttyが異なる場合があるので, それに対応出来たら良いよね(するとは言っていない)
    stop_getty = 'sudo systemctl stop serial-getty@ttyS0.service'
    start_getty = 'sudo systemctl start serial-getty@ttyS0.service'
    serial_dev = '/dev/ttyS0'

    subprocess.call(stop_getty, stdout=subprocess.PIPE, shell=True)
    yield serial_dev
    subprocess.call(start_getty, stdout=subprocess.PIPE, shell=True)
