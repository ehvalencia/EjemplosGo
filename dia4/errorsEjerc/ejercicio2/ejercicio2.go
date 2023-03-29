/*
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: el salario es menor
 a 10.000" y lanzalo en caso de que “salary” sea menor o igual a  10.000. La validación debe ser hecha con la función
 “Is()” dentro del “main”.
*/

package main

import (
	"errors"
	"fmt"
)

type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s", e.message)
}

func main() {
	var salary int = 8000

	if salary < 10000 {
		myErr := &myError{message: "el salario ingresado no alcanza el mínimo imponible"}
		var err *myError
		if errors.Is(myErr, err) {
			fmt.Println(myErr.message)
		}
	} else {
		fmt.Println("Debe pagar impuestos")
	}

}
