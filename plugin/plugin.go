package main

import (
	"context"
	"fmt"
	"github.com/zazin/test-plugin/gateway"
	"net/http"
)

func init() {
	fmt.Println("new plugin loaded!!!")
}

var ClientRegisterer = registerer("new-plugin")

type registerer string

func (r registerer) Run(f func(
	name string,
	handler func(context.Context, map[string]interface{}) (http.Handler, error),
)) {
	f(string(r), func(ctx context.Context, extra map[string]interface{}) (http.Handler, error) {
		return gateway.New(ctx)
	})
}

func main() {}
