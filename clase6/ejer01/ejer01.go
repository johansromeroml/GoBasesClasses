package main

import "fmt"

type Alumno struct {
	Name     string
	Apellido string
	DNI      string
	Fecha    string
}

type Person interface {
	Detalle()
}

func (a Alumno) Detalle() {
	fmt.Printf("Name: %s \n"+
		"Apellido: %s \n"+
		"DNI: %s \n"+
		"Fecha: %s \n", a.Name, a.Apellido, a.DNI, a.Fecha)
}

func main() {
	alum := Alumno{Name: "Jhon", Apellido: "Doe", DNI: "123123", Fecha: "01/01/2001"}
	alum.Detalle()
}
