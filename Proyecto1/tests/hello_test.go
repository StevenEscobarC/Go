package tests

import "testing"

func TestHello(t *testing.T) {
	respuesta := hello()
	if respuesta != "hola mundo" {
		t.Errorf("hello() = %q, want %q", respuesta, "hola mundo")
	}
}
