package functional_option

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	s := NewServer("127.0.0.1",
		WithPort(":8081"),
		WithMaxConns(10))
	t.Logf("%+v", s)
}

/**
go test -v .
&{Addr:127.0.0.1 Port::8081 Timeout:0s MaxConns:10}
*/
