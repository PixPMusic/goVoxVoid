package voxvoid

import (
	"encoding/binary"
	"net"
)

var conn net.Conn

func VzxControllerApiCall(req Request) (Response, error) {
	var err error

	// Open a socket connection to 44100
	conn, err := net.Dial("tcp", "localhost:44100")
	if err != nil {
		return Response{}, err
	}
	defer conn.Close()

	// Serialize the request
	requestBuffer := req.Serialize()

	// Send the length of the request to the server
	conn.Write(binary.LittleEndian.AppendUint32([]byte{}, uint32(len(requestBuffer))))

	// Send the request to the server
	conn.Write(requestBuffer)

	// Receive the response from the server
	responseBuffer := make([]byte, 328)
	_, err = conn.Read(responseBuffer)
	if err != nil {
		return Response{}, err
	}

	// Deserialize the response from the buffer
	var resp = DeserializeResponse(responseBuffer)

	return resp, err
}
