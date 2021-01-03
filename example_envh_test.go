package envh_test

import (
	"fmt"
	"os"

	"github.com/felipeneuwald/envh"
)

func ExampleString() {
	os.Setenv("STRING_KEY", "SomeString")

	s, err := envh.String("STRING_KEY")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T %v\n", s, s)

	// Output: string SomeString
}

func ExampleMustString() {
	os.Setenv("STRING_KEY", "SomeString")

	s := envh.MustString("STRING_KEY")

	fmt.Printf("%T %v\n", s, s)

	// Output: string SomeString
}

func ExampleBool() {
	os.Setenv("BOOL_KEY", "true")

	b, err := envh.Bool("BOOL_KEY")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T %v\n", b, b)

	// Output: bool true
}

func ExampleMustBool() {
	os.Setenv("BOOL_KEY", "true")

	b := envh.MustBool("BOOL_KEY")

	fmt.Printf("%T %v\n", b, b)

	// Output: bool true
}

func ExampleInt() {
	os.Setenv("INT_KEY", "9")

	i, err := envh.Int("INT_KEY")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T %v\n", i, i)

	// Output: int 9
}

func ExampleMustInt() {
	os.Setenv("INT_KEY", "9")

	i := envh.MustInt("INT_KEY")

	fmt.Printf("%T %v\n", i, i)

	// Output: int 9
}
