package main

import (
	"testing"
)

// TestSum
func TestSum(t *testing.T){
	if (Sum(1,2,3,4)!=10){
		t.Error("Sum did not work as expected.")
	}else{
		t.Log("One test passed.")
	}
}
