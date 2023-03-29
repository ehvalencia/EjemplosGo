/*
Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”,
para que el mensaje de error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo
imponible (el mensaje mostrado por consola deberá decir: “Error: el mínimo imponible es de 150.000 y el salario
ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
*/
package main

import (
	"fmt"
)

func validate(salary int) error {
	if salary < 10000 {
		return fmt.Errorf("Error: el salario ingresado no alcanza el mínimo imponible y el salario ingresado es de: %d", salary)
	}
	return nil
}

func main() {
	var salary int = 8000
	err := validate(salary)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Debe pagar impuestos")
}
