from pypresence import Presence
import time 

RPC = Presence("1179227301993521212")
RPC.connect()

while True:
    RPC.update(state="Red Hat Enterprise Linux 8.9")
    time.sleep(20)