package tools

import "testing"

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		given int
		want  string
	}{
		{
			given: 5,
			want:  "buzz",
		},
		{
			given: 7,
			want:  "7",
		},
		{
			given: 15,
			want:  "fizz-buzz",
		},
	}
	for _, testCase := range testCases {
		got := FizzBuzz(testCase.given)
		if testCase.want != got {
			t.Errorf("got is not equel to want: %s - %s", got, testCase.want)
		}

	}

}
