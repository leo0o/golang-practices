package interface_function

import (
	"testing"
)

func TestInterfaceFunction(t *testing.T) {
	m := map[string]string{
		"a": "1",
		"b": "2",
	}
	eachFunc(m, introduce)
}
