package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	for _, data := range []struct {
		input  int
		output string
	}{
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{16, "16"},
	} {
		result := FizzBuzz(data.input)
		if result != data.output {
			t.Errorf("FizzBuzz(%d) == %s, shoud be equal %s", data.input, result, data.output)
		}
	}
}