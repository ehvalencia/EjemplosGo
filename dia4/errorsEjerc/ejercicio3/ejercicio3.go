/*
Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,

	se implemente “errors.New()”.
*/
package main

import (
	"errors"
	"fmt"
)

func validate(salary int) error {
	if salary < 10000 {
		return errors.New("Error: el salario ingresado no alcanza el mínimo imponible")
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
