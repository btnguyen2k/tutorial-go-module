package module_test

import (
	"testing"

	mylib "github.com/btnguyen2k/tutorial-go-module"
)

func TestMd5String(t *testing.T) {
	s := "Goc Lap Trinh"
	expected := "696c82bb924cac3bf3006a08704fc0cd"
	if actual := mylib.Md5String(s); actual != expected {
		t.Errorf("expected [%v] but got [%v]", expected, actual)
	}
}

