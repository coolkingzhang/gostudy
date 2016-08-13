package testing

import "testing"

func Add2(aa int, bb int) int {
	return aa + bb
}
func TestAdd2(t *testing.T) {
	r := Add2(1, 2)
	if r != 3 {
		t.Errorf("error testting")
	}
}
