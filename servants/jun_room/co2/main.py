import subprocess

import serial

# モデルによってUARTに対する挙動が違う
# ラズパイ zero W は以下の設定でOK
# https://qiita.com/UedaTakeyuki/items/c5226960a7328155635f#%E4%BA%8B%E5%89%8D%E6%BA%96%E5%82%99uart%E3%82%92%E7%A9%BA%E3%81%91%E3%82%8B
serial_dev = '/dev/ttyS0'
stop_getty = 'sudo systemctl stop serial-getty@ttyS0.service'
start_getty = 'sudo systemctl start serial-getty@ttyS0.service'


def read():
    """CO2センサーからデータを取得する関数
    https://toranoana-lab.hatenablog.com/entry/2019/12/27/180048
    """
    # UART初期化
    ser = serial.Serial(serial_dev,
                        baudrate=9600,
                        bytesize=serial.EIGHTBITS,
                        parity=serial.PARITY_NONE,
                        stopbits=serial.STOPBITS_ONE,
                        timeout=1.0)

    # CO2濃度を取得するコマンド送信
    result = ser.write(b'\xff\x01\x86\x00\x00\x00\x00\x00\x79')
    s = ser.read(9)

    # 取得した形式が正しければCO2濃度を10進法に変換してretrun
    if len(s) >= 4 and s[0] == int('ff', 16) and s[1] == int('86', 16):
        return s[2] * 256 + s[3]


if __name__ == "__main__":
    subprocess.call(stop_getty, stdout=subprocess.PIPE, shell=True)
    print(f"CO2 = {read()}[ppm]")
    subprocess.call(start_getty, stdout=subprocess.PIPE, shell=True)
