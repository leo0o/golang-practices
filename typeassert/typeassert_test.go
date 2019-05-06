package typeassert

import (
	"testing"
	"fmt"
)

func TestTypeAssert(t *testing.T) {
	//检测类型是否实现了某接口
	// 编译时就能检测出
	var _ Tester = new(Test)    //Test类型实现了Tester接口，编译通过
	//var _ Tester = new(Test1)   //Test1类型未实现Tester接口，编译不通过
	// 运行时检测
	var _ Tester = (*Test)(nil)
	

	//检测某个接口变量的类型
	//所有 case 语句中列举的类型（nil 除外）都必须实现对应的接口（在上例中即 Shaper），如果被检测类型没有在 case 语句列举的类型中，就会执行 default 语句
	var t1 Tester = &Test{"Test"}
	switch t1.(type) {
	case *Test:
		fmt.Println("t1 is a Test")
	case nil:
		fmt.Println("t1 is nil")
	default:
		fmt.Println("t1 is not a Test")
	}
}
