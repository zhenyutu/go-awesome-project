package server

import "reflect"

type Handler interface {
	Handle(method string, params []interface{}) (interface{}, error)
}

type RPCHandler struct {
	//持有代理对象
	Object reflect.Value
}

func (h *RPCHandler) Handle(methodName string, params []interface{}) (interface{}, error) {
	argsIn := make([]reflect.Value, len(params))
	if params != nil && len(params) > 0 {
		for i, p := range params {
			argsIn[i] = reflect.ValueOf(p)
		}
	}

	method := h.Object.MethodByName(methodName)
	argsOut := method.Call(argsIn)

	result := make([]interface{}, len(argsOut))
	for i, r := range argsOut {
		result[i] = r.Interface()
	}

	if len(result) > 1 {
		if _, ok := argsIn[1].Interface().(error); ok {
			return nil, argsIn[1].Interface().(error)
		}
	}

	return result[0], nil
}
