package avutil

//
// THIS FILE WAS AUTOMATICALLY GENERATED by hackgenerator.go.
//

//
//#include <libavutil/avutil.h>
//
//
//int GO_AVERROR_BSF_NOT_FOUND()
//{
//  return (int)(AVERROR_BSF_NOT_FOUND);
//}
//
//int GO_AVERROR_BUG()
//{
//  return (int)(AVERROR_BUG);
//}
//
//int GO_AVERROR_BUFFER_TOO_SMALL()
//{
//  return (int)(AVERROR_BUFFER_TOO_SMALL);
//}
//
//int GO_AVERROR_DECODER_NOT_FOUND()
//{
//  return (int)(AVERROR_DECODER_NOT_FOUND);
//}
//
//int GO_AVERROR_DEMUXER_NOT_FOUND()
//{
//  return (int)(AVERROR_DEMUXER_NOT_FOUND);
//}
//
//int GO_AVERROR_ENCODER_NOT_FOUND()
//{
//  return (int)(AVERROR_ENCODER_NOT_FOUND);
//}
//
//int GO_AVERROR_EOF()
//{
//  return (int)(AVERROR_EOF);
//}
//
//int GO_AVERROR_EXIT()
//{
//  return (int)(AVERROR_EXIT);
//}
//
//int GO_AVERROR_EXTERNAL()
//{
//  return (int)(AVERROR_EXTERNAL);
//}
//
//int GO_AVERROR_FILTER_NOT_FOUND()
//{
//  return (int)(AVERROR_FILTER_NOT_FOUND);
//}
//
//int GO_AVERROR_INVALIDDATA()
//{
//  return (int)(AVERROR_INVALIDDATA);
//}
//
//int GO_AVERROR_MUXER_NOT_FOUND()
//{
//  return (int)(AVERROR_MUXER_NOT_FOUND);
//}
//
//int GO_AVERROR_OPTION_NOT_FOUND()
//{
//  return (int)(AVERROR_OPTION_NOT_FOUND);
//}
//
//int GO_AVERROR_PATCHWELCOME()
//{
//  return (int)(AVERROR_PATCHWELCOME);
//}
//
//int GO_AVERROR_PROTOCOL_NOT_FOUND()
//{
//  return (int)(AVERROR_PROTOCOL_NOT_FOUND);
//}
//
//int GO_AVERROR_STREAM_NOT_FOUND()
//{
//  return (int)(AVERROR_STREAM_NOT_FOUND);
//}
//
//int GO_AVERROR_BUG2()
//{
//  return (int)(AVERROR_BUG2);
//}
//
//int GO_AVERROR_UNKNOWN()
//{
//  return (int)(AVERROR_UNKNOWN);
//}
//
//int GO_AVERROR_EXPERIMENTAL()
//{
//  return (int)(AVERROR_EXPERIMENTAL);
//}
//
//int GO_AVERROR_INPUT_CHANGED()
//{
//  return (int)(AVERROR_INPUT_CHANGED);
//}
//
//int GO_AVERROR_OUTPUT_CHANGED()
//{
//  return (int)(AVERROR_OUTPUT_CHANGED);
//}
//
//int GO_AVERROR_HTTP_BAD_REQUEST()
//{
//  return (int)(AVERROR_HTTP_BAD_REQUEST);
//}
//
//int GO_AVERROR_HTTP_UNAUTHORIZED()
//{
//  return (int)(AVERROR_HTTP_UNAUTHORIZED);
//}
//
//int GO_AVERROR_HTTP_FORBIDDEN()
//{
//  return (int)(AVERROR_HTTP_FORBIDDEN);
//}
//
//int GO_AVERROR_HTTP_NOT_FOUND()
//{
//  return (int)(AVERROR_HTTP_NOT_FOUND);
//}
//
//int GO_AVERROR_HTTP_OTHER_4XX()
//{
//  return (int)(AVERROR_HTTP_OTHER_4XX);
//}
//
//int GO_AVERROR_HTTP_SERVER_ERROR()
//{
//  return (int)(AVERROR_HTTP_SERVER_ERROR);
//}
//
//int64_t GO_AV_NOPTS_VALUE()
//{
//  return (int64_t)(AV_NOPTS_VALUE);
//}
import "C"

const (
	ErrorCodeBSFNotFound      ErrorCode = -1179861752
	ErrorCodeBug              ErrorCode = -558323010
	ErrorCodeBufferTooSmall   ErrorCode = -1397118274
	ErrorCodeDecoderNotFound  ErrorCode = -1128613112
	ErrorCodeDemuxerNotFound  ErrorCode = -1296385272
	ErrorCodeEncoderNotFound  ErrorCode = -1129203192
	ErrorCodeEOF              ErrorCode = -541478725
	ErrorCodeExit             ErrorCode = -1414092869
	ErrorCodeExternal         ErrorCode = -542398533
	ErrorCodeFilterNotFound   ErrorCode = -1279870712
	ErrorCodeInvalidData      ErrorCode = -1094995529
	ErrorCodeMuxerNotFound    ErrorCode = -1481985528
	ErrorCodeOptionNotFound   ErrorCode = -1414549496
	ErrorCodePatchWelcome     ErrorCode = -1163346256
	ErrorCodeProtocolNotFound ErrorCode = -1330794744
	ErrorCodeStreamNotFound   ErrorCode = -1381258232
	ErrorCodeBug2             ErrorCode = -541545794
	ErrorCodeUnknown          ErrorCode = -1313558101
	ErrorCodeExperimental     ErrorCode = -733130664
	ErrorCodeInputChanged     ErrorCode = -1668179713
	ErrorCodeOutputChanged    ErrorCode = -1668179714
	ErrorCodeHttpBadRequest   ErrorCode = -808465656
	ErrorCodeHttpUnauthorized ErrorCode = -825242872
	ErrorCodeHttpForbidden    ErrorCode = -858797304
	ErrorCodeHttpNotFound     ErrorCode = -875574520
	ErrorCodeHttpOther4xx     ErrorCode = -1482175736
	ErrorCodeHttpServerError  ErrorCode = -1482175992
	NoPTSValue                int64     = -9223372036854775808
)

func getErrorCodeBSFNotFound() C.int {
	return C.GO_AVERROR_BSF_NOT_FOUND()
}

func getErrorCodeBug() C.int {
	return C.GO_AVERROR_BUG()
}

func getErrorCodeBufferTooSmall() C.int {
	return C.GO_AVERROR_BUFFER_TOO_SMALL()
}

func getErrorCodeDecoderNotFound() C.int {
	return C.GO_AVERROR_DECODER_NOT_FOUND()
}

func getErrorCodeDemuxerNotFound() C.int {
	return C.GO_AVERROR_DEMUXER_NOT_FOUND()
}

func getErrorCodeEncoderNotFound() C.int {
	return C.GO_AVERROR_ENCODER_NOT_FOUND()
}

func getErrorCodeEOF() C.int {
	return C.GO_AVERROR_EOF()
}

func getErrorCodeExit() C.int {
	return C.GO_AVERROR_EXIT()
}

func getErrorCodeExternal() C.int {
	return C.GO_AVERROR_EXTERNAL()
}

func getErrorCodeFilterNotFound() C.int {
	return C.GO_AVERROR_FILTER_NOT_FOUND()
}

func getErrorCodeInvalidData() C.int {
	return C.GO_AVERROR_INVALIDDATA()
}

func getErrorCodeMuxerNotFound() C.int {
	return C.GO_AVERROR_MUXER_NOT_FOUND()
}

func getErrorCodeOptionNotFound() C.int {
	return C.GO_AVERROR_OPTION_NOT_FOUND()
}

func getErrorCodePatchWelcome() C.int {
	return C.GO_AVERROR_PATCHWELCOME()
}

func getErrorCodeProtocolNotFound() C.int {
	return C.GO_AVERROR_PROTOCOL_NOT_FOUND()
}

func getErrorCodeStreamNotFound() C.int {
	return C.GO_AVERROR_STREAM_NOT_FOUND()
}

func getErrorCodeBug2() C.int {
	return C.GO_AVERROR_BUG2()
}

func getErrorCodeUnknown() C.int {
	return C.GO_AVERROR_UNKNOWN()
}

func getErrorCodeExperimental() C.int {
	return C.GO_AVERROR_EXPERIMENTAL()
}

func getErrorCodeInputChanged() C.int {
	return C.GO_AVERROR_INPUT_CHANGED()
}

func getErrorCodeOutputChanged() C.int {
	return C.GO_AVERROR_OUTPUT_CHANGED()
}

func getErrorCodeHttpBadRequest() C.int {
	return C.GO_AVERROR_HTTP_BAD_REQUEST()
}

func getErrorCodeHttpUnauthorized() C.int {
	return C.GO_AVERROR_HTTP_UNAUTHORIZED()
}

func getErrorCodeHttpForbidden() C.int {
	return C.GO_AVERROR_HTTP_FORBIDDEN()
}

func getErrorCodeHttpNotFound() C.int {
	return C.GO_AVERROR_HTTP_NOT_FOUND()
}

func getErrorCodeHttpOther4xx() C.int {
	return C.GO_AVERROR_HTTP_OTHER_4XX()
}

func getErrorCodeHttpServerError() C.int {
	return C.GO_AVERROR_HTTP_SERVER_ERROR()
}

func getNoPTSValue() C.int64_t {
	return C.GO_AV_NOPTS_VALUE()
}
