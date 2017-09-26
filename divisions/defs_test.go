package divisions

import "testing"

func TestList(t *testing.T) {
	divs := []string{"news", "tv", "radio"}

	if len(List) != len(divs) {
		t.Errorf("Incorrect lenght of list of divisions. Got: %d, expected: %d", len(List), len(divs))
	}
}
