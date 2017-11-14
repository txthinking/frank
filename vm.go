package main

import (
	"github.com/robertkrimen/otto"
	"github.com/txthinking/frank/function"
)

var VM *otto.Otto

func InitVM() {
	VM = otto.New()
	RegisterFunctions()
}

func RegisterFunctions() {
	function.MD5(VM)
	function.Must(VM)
}
