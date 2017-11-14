package function

import (
	"fmt"
	"os"

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
			fmt.Println("Exit by must")
			os.Exit(1)
		}
		b, err := a0.ToBoolean()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if !b {
			fmt.Println("Exit by must")
			os.Exit(1)
		}
		return otto.Value{}
	})
}
