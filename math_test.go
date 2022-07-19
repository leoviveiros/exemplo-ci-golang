package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(10, 20)
	
	if total != 30 {
		t.Errorf("Soma inválida: resultado=%d, esperado=%d", total, 30)
	}
}