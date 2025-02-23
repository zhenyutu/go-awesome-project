package server

import "reflect"

const header_length = 8

type Server struct {
	addr     string
	services map[string]Handler
}

func NewServer(addr string) *Server {
	server := &Server{addr: addr, services: make(map[string]Handler)}
	return server
}

func (s *Server) RegisterService(service interface{}) {
	objectName := reflect.Indirect(reflect.ValueOf(service)).Type().Name()
	if _, ok := s.services[objectName]; ok {
		return
	}

	s.services[objectName] = &RPCHandler{
		Object: reflect.ValueOf(service),
	}
}

func (s *Server) Run() error {
	return s.ListenAndHandle(s.addr)
}
