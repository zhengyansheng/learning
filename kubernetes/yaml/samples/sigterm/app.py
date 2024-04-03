import signal
import time
import sys
from datetime import datetime


def sigterm_handler(signal, frame):
    print('Received SIGTERM, exiting gracefully...')
    # 在这里执行清理操作
    sys.exit(0)


# 注册 SIGTERM 信号处理程序
signal.signal(signal.SIGTERM, sigterm_handler)

# 运行一个无限循环，直到接收到 SIGTERM 信号
print('Waiting for SIGTERM...')
while True:
    print('Still running... ', datetime.now().strftime("%Y-%m-%d %H:%M:%S"))
    time.sleep(1)
