package connection

import (
	"fmt"
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_portAvailable(t *testing.T) {
	assert.Equal(t, true, portAvailable(8081))

	serv, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(os.Stderr, "Unable to start test tcp listener:", err)
		t.Fail()
		return
	}
	defer serv.Close()

	go func() {
		for {
			conn, err := serv.Accept()
			if err == nil {
				conn.Close()
			}
		}
	}()

	assert.Equal(t, false, portAvailable(8081))
	assert.Equal(t, true, portAvailable(8082))
}

func Test_getAvailablePort(t *testing.T) {
	port, err := getAvailablePort(8081, 1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 8081, port)

	serv, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(os.Stderr, "Unable to start test tcp listener:", err)
		t.Fail()
		return
	}
	defer serv.Close()

	go func() {
		for {
			conn, err := serv.Accept()
			if err == nil {
				conn.Close()
			}
		}
	}()

	port, err = getAvailablePort(8081, 0)
	assert.EqualError(t, err, "No available port")
	assert.Equal(t, -1, port)

	port, err = getAvailablePort(8081, 1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 8082, port)
}
