package config_test

import (
	"fmt"
	"os"

	"github.com/Minimal-Viable-Software/config-go"
)

func Example_basic() {
	os.Setenv("NUM", "69")

	var n int
	fmt.Printf("%d\n", n)

	config.Int(&n, "num")
	fmt.Printf("%d\n", n)

	// Output:
	// 0
	// 69
}

func ExampleSetPrefix() {
	os.Setenv("FOO_NUM", "69")

	config.SetPrefix("foo_")

	var n int
	fmt.Printf("%d\n", n)

	config.Int(&n, "num")
	fmt.Printf("%d\n", n)

	// Output:
	// 0
	// 69
}
