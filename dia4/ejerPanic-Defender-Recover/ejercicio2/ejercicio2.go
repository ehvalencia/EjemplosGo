/*
A continuación, vamos a crear un archivo “customers.txt” con información de los clientes del estudio.
Ahora que el archivo sí existe, el panic no debe ser lanzado.
Creamos el archivo “customers.txt” y le agregamos la información de los clientes.
Extendemos el código del punto uno para que podamos leer este archivo e imprimir los datos que contenga. En el caso de no poder leerlo, se debe lanzar
un “panic”.
Recordemos que siempre que termina la ejecución, independientemente del resultado, debemos tener un “defer” que nos indique que la ejecución finalizó.
También recordemos cerrar los archivos al finalizar su uso.
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Ejecución finalizada")
	}()

	file, err := os.Open("customers.txt")

	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}

	contenido, errRead := io.ReadAll(file)

	if errRead != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}

	file.Close()
	fmt.Println(string(contenido))

	return
}
