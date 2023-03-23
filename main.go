package main

import "fmt"

func main(){
	fmt.Println("=================")
	hi();
	fmt.Println("=================")
	prNameAnddr();
	fmt.Println("=================")
	weather();
	fmt.Println("=================")
	correction2()
}


func hi(){
fmt.Println("Hola Edwin...")
}

//Crear una aplicación donde tengas como variable tu nombre y dirección.
//Imprimir en consola el valor de cada una de las variables.
func prNameAnddr(){
	name := "Edwin Valencia"
	Addres:= "Cra 15a 1 sur 54 - ap101"
	
	fmt.Println("Nombre ", name , "Direccion ", Addres)
}

/*Una empresa de meteorología quiere una aplicación donde pueda tener la temperatura, humedad y presión atmosférica de distintos lugares. 
Declará 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
Imprimí los valores de las variables en consola.
¿Qué tipo de dato le asignarías a las variables?*/

func weather(){
	var temperature string
	var damp float64
	var airPressure float64

	temperature = "22 grados centigrados"
	damp        = 47
	airPressure = 563
	
	fmt.Println("temperature ", temperature, " Humedad ", damp, " Air ", airPressure )
}

/*Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia Programación I para poder brindarles las correspondientes devoluciones. Uno de los puntos del examen consiste en declarar distintas variables. 
Necesita ayuda para:
Detectar cuáles de estas variables que declaró el alumno son correctas.
Corregir las incorrectas.*/
func correction(){
 //error	
  //var 1nombre string
  //var apellido string
  //var int edad
  //1apellido := 6
  //var licencia_de_conducir = true
  //var estatura de la persona int

/*
correcion
  var nombre string
  var apellido string
  var edad int 
  apellido := 6
  var licenciaDeConducir = true
  var estatura int
*/
}

/*
Un estudiante de programación intentó realizar declaraciones de variables con sus respectivos tipos en Go, pero tuvo varias dudas mientras lo hacía. A partir de esto, nos brindó su código y pidió la ayuda de un desarrollador experimentado que pueda:
Verificar su código y realizar las correcciones necesarias.

  var apellido string = "Gomez"
  var edad int = "35"
  boolean := "false";
  var sueldo string = 45857.90
  var nombre string = "Julián"

*/
func correction2(){
	var apellido2 string = "Gomez"
	var edad2 int = 35
	boolean := "false";
	var sueldo2 float64 = 45857.90
	var nombre2 string = "Julián"

	fmt.Println( apellido2 , edad2 , boolean , sueldo2 , nombre2 )
}