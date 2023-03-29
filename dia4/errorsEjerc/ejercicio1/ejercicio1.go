/*
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: el salario ingresado
no alcanza el mínimo imponible" y lanzalo en caso de que “salary” sea menor a 150.000. De lo contrario, tendrás que
imprimir por consola el mensaje “Debe pagar impuesto”.
*/
package main

import "fmt"

type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s", e.message)
}

func main() {
	var salary int = 140000

	if salary < 150000 {
		myErr := &myError{message: "el salario ingresado no alcanza el mínimo imponible"}
		fmt.Println(myErr)

	} else {
		fmt.Println("Debe pagar impuestos")
	}

}
