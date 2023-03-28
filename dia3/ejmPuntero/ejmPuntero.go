package main

import (
	"fmt"
)

type Compania struct {
	Nombre string
	Puesto string
}
type Empleado struct {
	Nombre   string
	Apellido string
	Compania Compania
}

func (e Empleado) InformacionEmpleado() {
	fmt.Printf("\nEmpleado: %s %s \n Puesto: %s \n Compania: %s\n", e.Nombre, e.Apellido, e.Compania.Puesto, e.Compania.Nombre)
}

func (c *Compania) CambiarPuesto(newPuesto string) {
	c.Puesto = newPuesto //para actualizar un puesto, asigamos puntero para que este se haga correctamente
}

func main() {
	e := Empleado{
		Nombre:   "Juan",
		Apellido: "Lopez",
		Compania: Compania{
			Nombre: "IT Solutions",
			Puesto: "Programador Junior",
		},
	}

	e.InformacionEmpleado()
	e.Compania.CambiarPuesto("Programador Senior")
	e.InformacionEmpleado()
}