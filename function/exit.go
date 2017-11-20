package function

import (
	"os"

	"github.com/robertkrimen/otto"
)

// Exit registers exit function
// Name: exit.
// Return: undefined.
func Exit(vm *otto.Otto) {
	vm.Set("exit", func(call otto.FunctionCall) otto.Value {
		os.Exit(0)
		return otto.Value{}
	})
}
