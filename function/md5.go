package function

import (
	"fmt"

	"github.com/robertkrimen/otto"
	"github.com/txthinking/x"
)

// MD5 registers md5 function
// Name: md5.
// Arguments: string.
// Return: string.
func MD5(vm *otto.Otto) {
	vm.Set("md5", func(call otto.FunctionCall) otto.Value {
		a0 := call.Argument(0)
		if !a0.IsString() {
			fmt.Println("ERROR", "md5(string)")
			return otto.Value{}
		}
		s, err := a0.ToString()
		if err != nil {
			fmt.Println("ERROR", err)
			return otto.Value{}
		}
		s = x.MD5(s)
		v, err := vm.ToValue(s)
		if err != nil {
			fmt.Println("ERROR", err)
			return otto.Value{}
		}
		return v
	})
}
