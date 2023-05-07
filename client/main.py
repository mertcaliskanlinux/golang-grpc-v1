
import grpc

import timeservice_pb2
import timeservice_pb2_grpc

def now(stub):
    request = timeservice_pb2.NowRequest()
    response = stub.Now(request)
    print("Current time: %s" % response.time.value)
    


def stream(stub,length):
    for msg in stub.Stream(timeservice_pb2.TimeStreamRequest(length=length)):
        print("Current time: %s" % msg.time.value)


def run():
    channel = grpc.insecure_channel('localhost:8080')
    stub = timeservice_pb2_grpc.TimeServiceStub(channel)
    # now(stub)
    seconds = 10
    stream(stub,seconds) # buraya kaç saniye yazarsak o kadar döner
if __name__ == '__main__':
    run()



print("GRPC - Client V1.0.0 is   running...")
