# import mesh

def Handler(req):
    # mesh call
    #
    # functionName = "<FUNCTIONNAME>"
    # input = req.input
    # result = mesh.mesh_call(functionName, input)
    # return result

    # single call
    return str.encode("[Python] ") + req.input

