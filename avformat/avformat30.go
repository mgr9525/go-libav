// +build ffmpeg30

package avformat

//#include <libavutil/avutil.h>
//#include <libavutil/avstring.h>
//#include <libavcodec/avcodec.h>
//#include <libavformat/avformat.h>
import "C"

import (
	"unsafe"

	"github.com/mgr9525/go-libav/avcodec"
	"github.com/mgr9525/go-libav/avutil"
)

func ApplyBitstreamFilters(codecCtx *avcodec.Context, pkt *avcodec.Packet, filtersCtx *avcodec.BitStreamFilterContext) error {
	cCodecCtx := (*C.AVCodecContext)(unsafe.Pointer(codecCtx.CAVCodecContext))
	cPkt := (*C.AVPacket)(unsafe.Pointer(pkt.CAVPacket))
	cFilters := (*C.AVBitStreamFilterContext)(unsafe.Pointer(filtersCtx.CAVBitStreamFilterContext))
	code := C.av_apply_bitstream_filters(cCodecCtx, cPkt, cFilters)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}
