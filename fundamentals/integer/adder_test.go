package integer

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(4, 2)
	expected := 6

	if sum != expected {
		t.Errorf("expected '%d', but got '%d'", expected, sum)
	}

}

func ExampleAdd() {
	sum := Add(5, 6)
	fmt.Println(sum)
	// Output: 6
}
