package integer

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(4, 2)
	expected := 6

	if sum != expected {
		t.Errorf("expected '%d', but got '%d'", expected, sum)
	}

}
