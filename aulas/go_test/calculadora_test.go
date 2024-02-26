package main

import (
    "testing"
)

func TestAdicao(t *testing.T) {
    resultado := Adicao(2, 3)
    esperado := 5
    if resultado != esperado {
        t.Errorf("Adicao(2, 3) = %d; esperado %d", resultado, esperado)
    }
}