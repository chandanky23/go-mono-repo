package mono_repo

import "fmt"

func Greetings(caller string) *string {
	fmt.Printf("Hello %s \n", caller)
	return &caller
}
