package voxvoid

import (
	"encoding/binary"
	"math"
)

const (
	ResponseType_Unknown                             = 0
	ResponseType_ErrorControllerNotConnectedToPlayer = 1
	ResponseType_Success                             = 1000
	ResponseType_ErrorIndexOutOfBounds               = 2
	ResponseType_ErrorItemNotFound                   = 3
)

type Response struct {
	// 0..3 this is an uint32_t (little endian)
	Action_t uint32
	// 4..7 this is an uint32_t (little endian)
	Result_t uint32

	// 4..11 uint64_t
	Integer_1 uint64
	// 12..19 uint64_t
	Integer_2 uint64
	// 20..27 uint64_t
	Integer_3 uint64
	// 28..35 uint64_t
	Integer_4 uint64

	// 36..43 double
	Double_1 float64
	// 44..51 double
	Double_2 float64
	// 52..59 double
	Double_3 float64
	// 60..67 double
	Double_4 float64

	// 68..131 64 character string
	String_1 string
	// 132..195 64 character string
	String_2 string
	// 196..259 64 character string
	String_3 string
	// 260..323 64 character string
	String_4 string
}

func DeserializeResponse(responseBuffer []byte) Response {
	var resp Response

	// Read the integer values from the buffer
	resp.Action_t = binary.LittleEndian.Uint32(responseBuffer[0:4])
	resp.Result_t = binary.LittleEndian.Uint32(responseBuffer[4:8])

	resp.Integer_1 = binary.LittleEndian.Uint64(responseBuffer[8:16])
	resp.Integer_2 = binary.LittleEndian.Uint64(responseBuffer[16:24])
	resp.Integer_3 = binary.LittleEndian.Uint64(responseBuffer[24:32])
	resp.Integer_4 = binary.LittleEndian.Uint64(responseBuffer[32:40])

	// Read the double values from the buffer
	resp.Double_1 = math.Float64frombits(binary.LittleEndian.Uint64(responseBuffer[40:48]))
	resp.Double_2 = math.Float64frombits(binary.LittleEndian.Uint64(responseBuffer[48:56]))
	resp.Double_3 = math.Float64frombits(binary.LittleEndian.Uint64(responseBuffer[56:64]))
	resp.Double_4 = math.Float64frombits(binary.LittleEndian.Uint64(responseBuffer[64:72]))

	// Read the string values from the buffer
	resp.String_1 = string(responseBuffer[72:136])
	resp.String_2 = string(responseBuffer[136:200])
	resp.String_3 = string(responseBuffer[200:264])
	resp.String_4 = string(responseBuffer[264:328])

	return resp
}
