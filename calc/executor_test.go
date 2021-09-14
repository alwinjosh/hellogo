package calc

import "testing"


func TestAdd(t *testing.T) {
	var ans = Add(10, 50)
	if ans != 60 {
		t.Errorf("Expected 60")
	}

}