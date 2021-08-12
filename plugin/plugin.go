package main

import (
	"context"
	"github.com/zazin/test-plugin/gateway"
	"net/http"
)

type greeting string

func (g greeting) Greet(ctx context.Context) (http.Handler, error) {
	return gateway.New(ctx)
}

// exported
var Greeter greeting
