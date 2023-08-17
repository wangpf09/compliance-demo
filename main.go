// Code generated by hertz generator.

package main

import (
	"compliance/workflow/biz/dal"
	"github.com/cloudwego/hertz/pkg/app/server"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func init() {
	dal.Init()
}

func main() {
	h := server.Default()

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	register(h)
	h.Spin()
}
