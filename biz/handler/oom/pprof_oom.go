package oom

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"os"
	"runtime/pprof"
	"strconv"
)

var data []string

func Oom(c context.Context, ctx *app.RequestContext) {
	f, _ := os.OpenFile("cpu.profile", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	_ = pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	oom := ctx.DefaultQuery("oom", "this is test for oom")
	size := ctx.DefaultQuery("size", "1000")
	count, _ := strconv.Atoi(size)
	add2Data(count, oom)
}

func add2Data(count int, oom string) {
	for i := 0; i < count; i++ {
		datum := fmt.Sprintf("%s-%d", oom, i)
		data = append(data, datum)
		//logrus.Info(datum, len(data))
	}
}
