package gojs

import (
	"fmt"
	"github.com/dop251/goja"
	"reflect"
)

type JSObject struct {
	j   *JSRunner
	obj *goja.Object
}

func (jo *JSObject) Call(method string, args ...interface{}) (goja.Value, error) {

	met := jo.obj.Get(method)
	if met == nil {
		return nil, fmt.Errorf("Got nil value for %s ", method)
	}
	var fn goja.Callable
	err := jo.j.vm.ExportTo(met, &fn)
	if err != nil {
		return nil, err
	}
	var vargs []goja.Value
	for _, a := range args {
		vargs = append(vargs, jo.j.vm.ToValue(a))
	}
	return fn(jo.obj, vargs...)
}

func (jo *JSObject) CallReturningObj(method string, args ...interface{}) (*JSObject, error) {
	v, err := jo.Call(method, args...)
	if err != nil {
		return nil, err
	}
	robj := v.ToObject(jo.j.vm)

	return &JSObject{jo.j, robj}, nil
}

func (jo *JSObject) CallReturningStr(method string, args ...interface{}) (string, error) {
	v, err := jo.Call(method, args...)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

func (jo *JSObject) GetNumber(name string) (int64, error) {
	v := jo.obj.Get(name)
	if v == nil {
		return 0, fmt.Errorf("Got nil value for %s ", name)
	}
	if v.ExportType() != reflect.TypeOf(int64(0)) {
		return 0, fmt.Errorf("The variable %s is not number type", name)
	}
	return v.ToInteger(), nil
}

func (jo *JSObject) GetString(name string) (string, error) {
	v := jo.obj.Get(name)
	if v == nil {
		return "", fmt.Errorf("Got nil value for %s ", name)
	}

	return v.String(), nil
}

func (jo *JSObject) GetObject(name string) (*JSObject, error) {

	obj := jo.obj.Get(name)
	if obj == nil {
		return nil, fmt.Errorf("Got nil value for %s ", name)
	}
	robj := obj.ToObject(jo.j.vm)

	return &JSObject{jo.j, robj}, nil
}
