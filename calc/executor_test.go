package calc

import "testing"


func TestAdd(t *testing.T) {
	var ans = Add(10, 50)
	if ans != 60 {
		t.Errorf("Expected 60")
	}

}

func TestSub(t *testing.T) {
	var ans = Sub(50, 10)
	if ans != 40 {
		t.Errorf("Expected 60")
	}

}