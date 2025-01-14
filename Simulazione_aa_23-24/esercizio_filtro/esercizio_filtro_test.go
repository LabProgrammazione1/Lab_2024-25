package main

import "testing"

var prog = "./esercizio_filtro"

func TestMain1(t *testing.T) {
	LanciaGenerica(t, prog, "", "3", "334")
}


func TestMain2(t *testing.T) {
	LanciaGenerica(t, prog, "", "345", "34456")
}


func TestMain3(t *testing.T) {
	LanciaGenerica(t, prog, "", "", "3")
}

func TestMain4(t *testing.T) {
	LanciaGenerica(t, prog, "", "3", "37777")
}

func TestMain5(t *testing.T) {
	LanciaGenerica(t, prog, "", "37", "37778")
}