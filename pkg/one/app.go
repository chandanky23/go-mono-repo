package one

import "fmt"

func Greetings(caller string) *string {
	fmt.Printf("Hello %s \n", caller)
	return &caller
}
