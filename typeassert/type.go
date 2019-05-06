package typeassert

import "fmt"

type Tester interface {
	doTest()
}

type Test struct {
	name string
}

func (t *Test)doTest() {
	fmt.Println(t.name)
}


type Test1 struct {
	name string
}
