package testing

import "testing"
import "fmt"

func Add(aa int, bb int) int {
	return aa + bb
}

func TestAdd(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("error testting")
	}

	fmt.Print("web test")
}
