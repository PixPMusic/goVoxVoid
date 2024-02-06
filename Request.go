package voxvoid

import (
	"encoding/binary"
	"math"
)

type Request struct {
	// 0..3 this is an uint32_t (little endian)
	Action_t uint32

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

func (r *Request) Serialize() []byte {
	requestBuffer := make([]byte, 324)

	// Write the action type to the buffer
	binary.LittleEndian.PutUint32(requestBuffer[0:4], r.Action_t)

	// Write the integer values to the buffer
	binary.LittleEndian.PutUint64(requestBuffer[4:12], r.Integer_1)
	binary.LittleEndian.PutUint64(requestBuffer[12:20], r.Integer_2)
	binary.LittleEndian.PutUint64(requestBuffer[20:28], r.Integer_3)
	binary.LittleEndian.PutUint64(requestBuffer[28:36], r.Integer_4)

	// Write the double values to the buffer
	binary.LittleEndian.PutUint64(requestBuffer[36:44], math.Float64bits(r.Double_1))
	binary.LittleEndian.PutUint64(requestBuffer[44:52], math.Float64bits(r.Double_2))
	binary.LittleEndian.PutUint64(requestBuffer[52:60], math.Float64bits(r.Double_3))
	binary.LittleEndian.PutUint64(requestBuffer[60:68], math.Float64bits(r.Double_4))

	// Write the string values to the buffer
	copy(requestBuffer[68:132], r.String_1)
	copy(requestBuffer[132:196], r.String_2)
	copy(requestBuffer[196:260], r.String_3)
	copy(requestBuffer[260:324], r.String_4)

	return requestBuffer
}
