package gojs

import (
	"fmt"
	"github.com/dop251/goja"
)

type JSRunner struct {
	prog *goja.Program
	vm   *goja.Runtime
}

type CallBackFunction func(args ...interface{}) (interface{}, error)

func New(script string) (*JSRunner, error) {
	prog, err := goja.Compile("", script, true)
	if err != nil {
		return nil, err
	}
	vm := goja.New()
	vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	j := &JSRunner{prog, vm}
	return j, nil
}

func (j *JSRunner) Run() (goja.Value, error) {
	return j.vm.RunProgram(j.prog)
}

func (j *JSRunner) GetObject(name string) (*JSObject, error) {
	obj := j.vm.Get(name)
	if obj == nil {
		return nil, fmt.Errorf("Got nil value for %s ", name)
	}
	robj := obj.ToObject(j.vm)

	return &JSObject{j, robj}, nil
}
func (j *JSRunner) GetGlobalObject() *JSObject {
	robj := j.vm.GlobalObject()
	return &JSObject{j, robj}
}

func (j *JSRunner) VM() *goja.Runtime {
	return j.vm
}

func (j *JSRunner) InjectFn(name string, fn interface{}) {

	j.vm.Set(name, fn)
}

/*func (j *JSRunner) InjectFn(name string, fn CallBackFunction) {

	j.vm.Set(name, func(call goja.FunctionCall) (goja.Value, error) {
		var iargs []interface{}
		for _, arg := range call.Arguments {
			a := arg.Export()
			iargs = append(iargs, a)
		}
		res, err := fn(iargs...)
		if err != nil {
			return j.vm.ToValue(nil), err
		}
		return j.vm.ToValue(res), nil
	})

}
*/

/*

js.InjectObj(window, mywindow)

js.InjectFn(main, myMain)

*/
