/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el
detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que
tenga un método detalle
*/
package main

import (
	"fmt"
)

type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (alumno Alumnos) detalle() {
	fmt.Println(alumno.Nombre)
	fmt.Println(alumno.Apellido)
	fmt.Println(alumno.DNI)
	fmt.Println(alumno.Fecha)
}

func main() {
	alumno := Alumnos{Nombre: "Edwin", Apellido: "Valencia", DNI: 1085344100, Fecha: "2022/06/09"}
	alumno.detalle()
}
