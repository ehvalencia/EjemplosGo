/*
Vamos a hacer que nuestro programa sea un poco más complejo y útil.
Desarrollá las funciones necesarias para permitir a la empresa calcular:
  - Salario mensual de un trabajador según la cantidad de horas trabajadas.
  - La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
  - Dicha función deberá retornar más de un valor (salario calculado y error).
  - En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10 % en concepto de
    impuesto.
  - En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe d
    evolver un error.
    El mismo tendrá que indicar “Error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
*/
package main

import (
	"errors"
	"fmt"
)

var ErrNegativeNumbersNotAllowed = errors.New("error: los números negativos no son permitidos")
var ErrMinHours = errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")

func SalaryPerMonth(hours float32, rate float32) (float32, error) {
	salary := hours * rate

	switch {
	case hours < 80:
		return 0, ErrMinHours
	case hours < 0:
		return 0, ErrNegativeNumbersNotAllowed
	case salary < 150000:
		return salary, nil
	default:
		return salary * (1 - 0.1), nil
	}
}

func main() {
	var salary, err = SalaryPerMonth(120, 20)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("El salario es: %.2f", salary)

}
