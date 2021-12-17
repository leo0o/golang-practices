package functional_option

import "time"

type Server struct {
	Addr     string
	Port     string
	Timeout  time.Duration
	MaxConns int
}

type Option func(server *Server)

func WithPort(port string) Option {
	return func(server *Server) {
		server.Port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.Timeout = timeout
	}
}

func WithMaxConns(maxConns int) Option {
	return func(server *Server) {
		server.MaxConns = maxConns
	}
}

func NewServer(addr string, options ...Option) *Server {
	s := &Server{
		Addr:     addr,
		Port:     ":8080",
		Timeout:  0,
		MaxConns: 0,
	}

	for _, option := range options {
		option(s)
	}

	return s
}
