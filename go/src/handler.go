package main

import sdk "github.com/keti-openfx/openfx/executor/go/pb"

func Handler(req sdk.Request) string {
	return string(req.Input)
}
