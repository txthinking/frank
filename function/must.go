package function

import (
	"os"

	"github.com/fatih/color"
	"github.com/robertkrimen/otto"
)

// Must registers must function
// Name: must.
// Arguments: bool.
// Return: undefined.
func Must(vm *otto.Otto) {
	vm.Set("must", func(call otto.FunctionCall) otto.Value {
		a0 := call.Argument(0)
		if !a0.IsBoolean() {
			color.Red("Must got false")
			os.Exit(2)
		}
		b, err := a0.ToBoolean()
		if err != nil {
			color.Red(err.Error())
			os.Exit(2)
		}
		if !b {
			color.Red("Must got false")
			os.Exit(2)
		}
		return otto.Value{}
	})
}
