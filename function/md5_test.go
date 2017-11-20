package function

import (
	"testing"

	"github.com/robertkrimen/otto"
)

func TestMD5(t *testing.T) {
	vm := otto.New()
	MD5(vm)
	vm.Run(`console.log(md5("hello"))`)
}
