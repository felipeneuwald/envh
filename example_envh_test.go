package envh_test

import (
	"fmt"
	"os"

	"github.com/felipeneuwald/envh"
)

func ExampleBool() {
	os.Setenv("BKEY1", "true")
	k1, err := envh.Bool("BKEY1")
	if err != nil {
		fmt.Println("BKEY1 err:", err)
	}
	fmt.Println("BKEY1    :", k1)

	os.Setenv("BKEY2", "false")
	k2, err := envh.Bool("BKEY2")
	if err != nil {
		fmt.Println("BKEY2 err:", err)
	}
	fmt.Println("BKEY2    :", k2)

	os.Setenv("BKEY3", "TRUE")
	k3, err := envh.Bool("BKEY3")
	if err != nil {
		fmt.Println("BKEY3 err:", err)
	}
	fmt.Println("BKEY3    :", k3)

	k4, err := envh.Bool("BKEY4")
	if err != nil {
		fmt.Println("BKEY4 err:", err)
	}
	fmt.Println("BKEY4    :", k4)

	// Output:
	// BKEY1    : true
	// BKEY2    : false
	// BKEY3 err: Environment variable BKEY3 is "TRUE"; want "true" or "false"
	// BKEY3    : false
	// BKEY4 err: Environment variable BKEY4 is missing
	// BKEY4    : false
}

func ExampleMustBool() {
	os.Setenv("BKEY1", "true")
	os.Setenv("BKEY2", "false")
	os.Setenv("BKEY3", "TRUE")

	k1 := envh.MustBool("BKEY1")
	k2 := envh.MustBool("BKEY2")
	k3 := envh.MustBool("BKEY3")
	k4 := envh.MustBool("BKEY4")

	fmt.Println(k1, k2, k3, k4)

	// Output: true false false false
}
