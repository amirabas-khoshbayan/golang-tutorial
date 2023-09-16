package tools

import "testing"

func TestFizzBuzz(t *testing.T) {
	given := 5
	want := "buzz"
	got := FizzBuzz(given)
	if want != got {
		t.Errorf("got is not equel to want: %s - %s", got, want)
	}

}
