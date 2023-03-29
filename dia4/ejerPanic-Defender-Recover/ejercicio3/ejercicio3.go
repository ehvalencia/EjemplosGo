/*
El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes. Los datos requeridos son:
Legajo
Nombre
DNI
Número de teléfono
Domicilio

Tarea 1: Antes de registrar a un cliente, debés verificar si el mismo ya existe. Para ello, necesitás leer los datos de un array. En caso de que esté
repetido, debes manipular adecuadamente el error como hemos visto hasta aquí. Ese error deberá:
1.- generar un panic;
2.- lanzar por consola el mensaje: “Error: el cliente ya existe”, y continuar con la ejecución del programa normalmente.
Tarea 2: Luego de intentar verificar si el cliente a registrar ya existe, desarrollá una función para validar que todos los datos a registrar de un
cliente contienen un valor distinto de cero. Esta función debe retornar, al menos, dos valores. Uno de ellos tendrá que ser del tipo error para el caso de que se
ingrese por parámetro algún valor cero (recordá los valores cero de cada tipo de dato, ej: 0, “”, nil).
Tarea 3: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes mensajes: “Fin de la ejecución” y “Se
detectaron varios errores en tiempo de ejecución”. Utilizá defer para cumplir con este requerimiento.

Requerimientos generales:
Utilizá recover para recuperar el valor de los panics que puedan surgir
Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error.
Generá algún error, personalizandolo a tu gusto utilizando alguna de las funciones de Go (realiza también la validación pertinente para el caso de error
retornado).
*/
package main

import (
	"fmt"
)

type Customer struct {
	Legajo    int
	Nombre    string
	DNI       int
	Telefono  string
	Domicilio string
}

func isCustomerValid(customer Customer) (bool, error) {
	if customer.Legajo == 0 || customer.Nombre == "" ||
		customer.DNI == 0 || customer.Telefono == "" ||
		customer.Domicilio == "" {
		return false, fmt.Errorf("Data no valid")
	}
	return true, nil
}

func customerExist(dni int) bool {
	for _, customer := range customers {
		if customer.DNI == dni {
			return true
		}
	}
	return false
}

func addCustomer(custumer Customer) {
	if customerExist(custumer.DNI) {
		panic("Error: el cliente ya existe")
	}
	if _, err := isCustomerValid(custumer); err != nil {
		panic(err)
	}
	customers = append(customers, custumer)
}

var customers = []Customer{}

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Fin de la ejecución")
	}()

	customer1 := Customer{
		Legajo:    1234234,
		Nombre:    "Carlos Alberto",
		DNI:       124345,
		Telefono:  "4545456456",
		Domicilio: "Calle privada",
	}

	customer2 := Customer{
		Legajo:    0,
		Nombre:    "",
		DNI:       0,
		Telefono:  "",
		Domicilio: "",
	}

	addCustomer(customer1)
	fmt.Println("Cliente agregado")
	fmt.Println(customers)

	addCustomer(customer2)
	fmt.Println("Cliente agregado")
	fmt.Println(customers)

	return
}
