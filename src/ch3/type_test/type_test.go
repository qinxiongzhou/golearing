package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64

	//b = a //Cannot use 'a' (type int) as type int64
	b = int64(a)

	var c MyInt
	//c = b //Cannot use 'b' (type int64) as type MyInt
	c = MyInt(b)

	t.Log(a,b,c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPrt := &a
	//aPrt = aPrt +1 //Invalid operation: aPrt +1 (mismatched types *int and untyped int)
	t.Log(a,aPrt)
	t.Log("%T %T",a,aPrt)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))

	if s == "" { //注意这里是“”，而不是nil

	}
}