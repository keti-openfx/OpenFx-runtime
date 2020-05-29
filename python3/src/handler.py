import grpc

import fxwatcher_pb2
import fxwatcher_pb2_grpc

def Handler(req):
    return req.input

def call_inter_function(functionName: str, input):

    channel = grpc.insecure_channel("{}.openfx-fn:50051".format(functionName))
    # add error logic
    clientStub = fxwatcher_pb2_grpc.FxWatcherStub(channel)
    serviceRequest = fxwatcher_pb2.Request(input = input, info = fxwatcher_pb2.Info(FunctionName = functionName, Trigger = fxwatcher_pb2.Trigger(Name = "grpc")))
    result = clientStub.Call(serviceRequest)
    return result.Output.encode()
