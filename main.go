package main

import (
	"fmt"

	"github.com/shortformikael/Hermes/src/testm"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	m := testm.Greet("James")

	fmt.Println(m)
}
