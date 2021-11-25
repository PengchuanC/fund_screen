import time
import grpc
import api_pb2
import api_pb2_grpc


def run():
    channel = grpc.insecure_channel("10.170.139.12:80")
    channel = grpc.insecure_channel("localhost:50053")
    stub = api_pb2_grpc.ScreenRpcServerStub(channel=channel)
    resp = stub.FundCategory(
        api_pb2.classify__pb2.ClassifyReq(secucode=['110011', '000001', '000191', '000147']))
    for r in resp.data:
        print(r.secucode, r.first, r.second)
    # resp = stub.FundBasicInfoHandler(api_pb2.basic__pb2.FundBasicInfoRequest())
    # for r in resp.data:
    #     print(r.secucode, r.launch_date)
    resp = stub.FundStyleNature(api_pb2.style__pb2.StyleReq(funds=['000001', '110011'], many=True))
    for k, v in resp.data.items():
        v = v.styles[0]
        print(v.secucode, v.date, v.style)

    resp = stub.FundRelatedIndex(api_pb2.index__pb2.IndexCorrReq(indexes=['000001']))
    for k, v in resp.data.items():
        print(k, v)


s = time.time()
run()
e = time.time()
print(f'finished in {round(e-s, 4)} sec.')
