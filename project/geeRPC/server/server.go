package server

import "reflect"

const header_length = 8

type Server struct {
	services map[string]Handler
}

func (s *Server) NewServer(addr string) *Server {
	server := &Server{}
	server.services = make(map[string]Handler)
	return server
}

func (s *Server) RegisterService(service interface{}) {
	objectName := reflect.Indirect(reflect.ValueOf(service)).Type().Name()
	if _, ok := s.services[objectName]; ok {
		return
	}

	s.services[objectName] = &RPCHandler{
		obj: reflect.ValueOf(service),
	}
}

func (s *Server) ListenAndServe(addr string) error {
	return s.ListenAndHandle(addr)
}
