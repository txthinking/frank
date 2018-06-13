package function

import (
	"fmt"

	"github.com/robertkrimen/otto"
	"encoding/base64"
)

// Base64Encode registers base64 encode function
// Name: base64encode.
// Arguments: string.
// Return: string.
func Base64Encode(vm *otto.Otto) {
	vm.Set("base64encode", func(call otto.FunctionCall) otto.Value {
		a0 := call.Argument(0)
		if !a0.IsString() {
			fmt.Println("ERROR", "base64encode(string)")
			return otto.Value{}
		}
		s, err := a0.ToString()
		if err != nil {
			fmt.Println("ERROR", err)
			return otto.Value{}
		}
		s = base64.StdEncoding.EncodeToString([]byte(s))
		v, err := vm.ToValue(s)
		if err != nil {
			fmt.Println("ERROR", err)
			return otto.Value{}
		}
		return v
	})
}

// Base64Decode registers base64 decode function
// Name: base64decode.
// Arguments: string.
// Return: string.
func Base64Decode(vm *otto.Otto) {
	vm.Set("base64decode", func(call otto.FunctionCall) otto.Value {
		a0 := call.Argument(0)
		if !a0.IsString() {
			fmt.Println("ERROR", "base64decode(string)")
			return otto.Value{}
		}
		s, err := a0.ToString()
		if err != nil {
			fmt.Println("ERROR", err)
			return otto.Value{}
		}
		sb, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			fmt.Println("ERROR", err)
			return otto.Value{}
		}
		v, err := vm.ToValue(string(sb))
		if err != nil {
			fmt.Println("ERROR", err)
			return otto.Value{}
		}
		return v
	})
}