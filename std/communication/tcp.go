package communication

import (
	"bufio"
	"net"
	"time"
)

const (
	ProtocolTCP              = "tcp"
	defaultTCPRequestTimeout = 5
	RequestTypeUni           = "uni"
	RequestTypeBi            = "bi"
)

type TCPRequest struct {
	Timeout int
	Address string
	Type    string
	Delim   byte
	Body    []byte
}

type TCPResponse struct {
	Err  error
	Body []byte
}

func (req *TCPRequest) Send() TCPResponse {
	var res TCPResponse
	var err error

	conn, err := net.Dial(ProtocolTCP, req.Address)
	if err != nil {
		res.Err = err
		return res
	}

	_, err = conn.Write(req.Body)
	if err != nil {
		res.Err = err
		conn.Close()
		return res
	}

	if req.Type != RequestTypeUni {

		if req.Timeout == 0 {
			req.Timeout = defaultTCPRequestTimeout
		}
		conn.SetReadDeadline(time.Now().Add(time.Duration(req.Timeout) * time.Second))

		buf, err := bufio.NewReader(conn).ReadBytes(req.Delim)
		if err != nil {
			res.Err = err
			conn.Close()
			return res
		}
		res.Body = buf
	}
	conn.Close()
	return res
}
