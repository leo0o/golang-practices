package interface_function

import "fmt"

type Handler interface {
	Handle(k, v string)
}

type HandlerFunc func(k, v string)

func (hf HandlerFunc) Handle(k, v string) {
	hf(k, v)
}

func each(m map[string]string, h Handler) {
	for k, v := range m {
		h.Handle(k, v)
	}
}

func eachFunc(m map[string]string, f func(k, v string)) {
	each(m, HandlerFunc(f))
}

func introduce(k, v string) {
	fmt.Printf("%s => %s \n", k, v)
}
