# -*- coding: utf-8 -*-
# refer http://eleparts.co.kr/data/design/product_file/SENSOR/gas/MH-Z19_CO2%20Manual%20V2.pdf
#
# <C2><A9> Takeyuki UEDA 2015 -

import subprocess
import time

import serial

serial_dev = '/dev/ttyS0'
stop_getty = 'sudo systemctl stop serial-getty@ttyS0.service'
start_getty = 'sudo systemctl start serial-getty@ttyS0.service'


def calibrateZeroPoing():
    ser = serial.Serial(serial_dev,
                        baudrate=9600,
                        bytesize=serial.EIGHTBITS,
                        parity=serial.PARITY_NONE,
                        stopbits=serial.STOPBITS_ONE,
                        timeout=1.0)
    result = ser.write(b"\xff\x01\x87\x00\x00\x00\x00\x00\x78")


if __name__ == '__main__':
    subprocess.call(stop_getty, stdout=subprocess.PIPE, shell=True)
    calibrateZeroPoing()
    subprocess.call(start_getty, stdout=subprocess.PIPE, shell=True)
    print("caribration zero point done.")
