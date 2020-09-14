package test

import "testing"

func TestHolaMundo(t *testing.T) {
	str := "hola mundo"
	if str != "hola mundo" {
		t.Error("No se pudo saludar ", nil)
	}
}
