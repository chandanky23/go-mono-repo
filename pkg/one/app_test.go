package mono_repo

import (
	"testing"
)

func TestGreetings(t *testing.T) {
	greeting := "world"
	if received := Greetings("World"); received == nil {
		t.Errorf("Failed as greeting was %s and err location was %v", greeting, received)
	}
}
