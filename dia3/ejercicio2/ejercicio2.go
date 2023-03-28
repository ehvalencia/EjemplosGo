/*
Una empresa necesita realizar una buena gestión de sus empleados, para esto realizaremos un pequeño programa nos ayudará

	a gestionar correctamente dichos empleados. Los objetivos son:

Crear una estructura Person con los campos ID, Name, DateOfBirth.
Crear una estructura Employee con los campos: ID, Position y una composicion con la estructura Person.
Realizar el método a la estructura Employe que se llame PrintEmployee(), lo que hará es realizar la impresión de los
campos de un empleado.
Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos y por último ejecutar
el método PrintEmployee().
*/
package main

import (
	"fmt"
)

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person
}

func (e Employee) PrintEmployed() {
	fmt.Println(e)
}

func main() {
	p1 := Person{
		ID:          1,
		Name:        "Edwin",
		DateOfBirth: "1999-02-03",
	}

	e1 := Employee{
		ID:       1,
		Position: "Developer",
		Person:   p1,
	}

	e1.PrintEmployed()
}
