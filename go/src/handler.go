package main

import sdk "github.com/keti-openfx/openfx/executor/go/pb"

func FunctionCall(functionName string, input []byte) string {
	// Do not Change
	address := fmt.Sprintf("%s.openfx-fn:50051", functionName)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %s\n", "sorry")
	}
	defer conn.Close()

	c := sdk.NewFxWatcherClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Call(ctx, &sdk.Request{Input: input, Info: &sdk.Info{FunctionName: functionName, Trigger: &sdk.Trigger{Name: "grpc", Time: time.Now().UTC().String()}}})
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	return r.Output
}

func Handler(req sdk.Request) string {
	return string(req.Input)
}
