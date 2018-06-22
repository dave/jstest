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

var Slice = []Person{
	{
		Name: "Jim",
		Age:  19,
	},
	{
		Name: "Jen",
		Age:  20,
	},
}

var Map = map[string]Person{
	"dave": {
		Name: "Dave",
		Age:  25,
	},
	"pete": {
		Name: "Pete",
		Age:  27,
	},
}
