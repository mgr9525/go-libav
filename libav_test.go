package golibav

import (
	"github.com/mgr9525/go-libav/avcodec"
	"github.com/mgr9525/go-libav/avutil"
	"testing"
)

func Test1(t *testing.T) {
	avcodec.RegisterAll()

	v1, v2, v3 := avutil.Version()
	println(v1, v2, v3)
}
