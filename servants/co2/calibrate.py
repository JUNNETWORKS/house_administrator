# -*- coding: utf-8 -*-
# refer http://eleparts.co.kr/data/design/product_file/SENSOR/gas/MH-Z19_CO2%20Manual%20V2.pdf
#
# <C2><A9> Takeyuki UEDA 2015 -
import serial

from .uart import use_uart


def calibrateZeroPoing(serial_dev):
    ser = serial.Serial(serial_dev,
                        baudrate=9600,
                        bytesize=serial.EIGHTBITS,
                        parity=serial.PARITY_NONE,
                        stopbits=serial.STOPBITS_ONE,
                        timeout=1.0)
    _ = ser.write(b"\xff\x01\x87\x00\x00\x00\x00\x00\x78")


if __name__ == '__main__':
    with use_uart() as serial_dev:
        calibrateZeroPoing(serial_dev)
    print("caribration zero point done.")
