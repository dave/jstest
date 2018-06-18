package ob

import "github.com/dave/jstest/ob/sub"

type Person struct {
	Name string
	Age  int
}

var John interface{} = Person{
	Name: "John",
	Age:  38,
}

var Foo = sub.Bar{
	Baz: "baz",
}
