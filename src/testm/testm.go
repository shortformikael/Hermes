// Package testm: testing
package testm

import "fmt"

func Greet(name string) string {
	message := fmt.Sprintf("Hello, %v. Welcome!", name)
	return message
}
