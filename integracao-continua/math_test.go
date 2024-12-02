package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)
	if total != 30 {
		t.Errorf("Valor recebido = %d, Valor esperado = %d", total, 30)
	}
}
