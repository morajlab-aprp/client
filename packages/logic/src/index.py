from __future__ import print_function
import logging
import grpc
import packages.logic.monitor_pb2 as monitor__pb2
import packages.logic.monitor_pb2_grpc as monitor_pb2__grpc


def main():
    print("Will try to greet world ...")
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = monitor_pb2__grpc.MonitorStub(channel)
        response = stub.SendData(monitor__pb2.MonitorRequest(data="This is my data"))
    print("Greeter client received: " + response.message)


if __name__ == "__main__":
    logging.basicConfig()
    main()
