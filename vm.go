package main

import (
	"github.com/robertkrimen/otto"
	"github.com/txthinking/frank/function"
)

// VM
var VM *otto.Otto

// InitVM
func InitVM() {
	VM = otto.New()
	RegisterFunctions()
}

// Register builtin functions
func RegisterFunctions() {
	function.MD5(VM)
	function.Must(VM)
	function.Exit(VM)
	function.Base64Decode(VM)
	function.Base64Encode(VM)
}
