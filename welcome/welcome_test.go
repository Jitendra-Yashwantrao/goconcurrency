package welcome

import "testing"

func TestWelcome(t *testing.T) {
	want := "Welcome to learning concurrency with go"

	if got := Welcome(); want != got {

		t.Errorf("Welcome() = %q,  got- %q", got, want)
	}

}
