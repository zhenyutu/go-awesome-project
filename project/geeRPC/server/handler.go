package server

import "reflect"

type Handler interface {
	Handle(method string, params []interface{}) ([]interface{}, error)
}

type RPCHandler struct {
	//持有代理对象
	obj reflect.Value
}

func (h *RPCHandler) Handle(methodName string, params []interface{}) ([]interface{}, error) {
	argsIn := make([]reflect.Value, len(params))
	for i, p := range params {
		argsIn[i] = reflect.ValueOf(p)
	}

	method := h.obj.MethodByName(methodName)
	argsOut := method.Call(argsIn)

	result := make([]interface{}, len(argsOut))
	for i, r := range argsOut {
		result[i] = r.Interface()
	}

	if _, ok := argsIn[1].Interface().(error); ok {
		return nil, argsIn[1].Interface().(error)
	}
	return result, nil
}
