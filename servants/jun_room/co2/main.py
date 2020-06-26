import serial
from uart import use_uart


def read(serial_dev):
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
    _ = ser.write(b'\xff\x01\x86\x00\x00\x00\x00\x00\x79')
    s = ser.read(9)

    # 取得した形式が正しければCO2濃度を10進法に変換してretrun
    if len(s) >= 4 and s[0] == int('ff', 16) and s[1] == int('86', 16):
        return s[2] * 256 + s[3]


if __name__ == "__main__":
    with use_uart() as serial_dev:
        print(f"CO2 = {read(serial_dev)}[ppm]")
