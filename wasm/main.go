package main

import (
	"fmt"

	// The generated bindings register themselves from this package's init.
	_ "github.com/Vilsol/timeless-jewels/wasm/exposition"
)

func main() {
	fmt.Println("Calculator Initialized")
	select {}
}
