package testpack

import (
	"testing"
)


func TestMain(m *testing.M){
	m.Run()
}

func TestSomething(t *testing.T){
	initial := 42
	if TestFun(initial)  != initial + 1 {
		t.Fatal("Int incorrect")
	}
}
