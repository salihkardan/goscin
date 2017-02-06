package goscin

import "testing"

func TestSum(t *testing.T) {
	if 8 != Sum(3,5){
		t.Failed();
	}
}