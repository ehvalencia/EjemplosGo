package main

import (
	"time"
	"fmt"
)

func main(){
	fmt.Println("=================")
	hi();
	fmt.Println("=================")
	prNameAnddr();
	fmt.Println("=================")
	weather();
	fmt.Println("=================")
	correction2()
	fmt.Println("==========Estrucura de control=======")
	lettersWord()
	fmt.Println("=================")
	lettersWord2()
	fmt.Println("=================")
	loan()
	fmt.Println("=================")
	impMonths()

	
	
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

/* Estructura de control */

/*
La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla. Para eso tendrán que:
Crear una aplicación que reciba  una variable con la palabra e imprimir la cantidad de letras que contiene la misma.
Luego, imprimí cada una de las letras.
*/
func lettersWord(){
	var word string 
	word = "Software"
	var countWord = len(word)

	fmt.Println("La palabra", word, "tiene", countWord, "letras.")

	for _, letter := range word {
		fmt.Println(string(letter))
	}
}

func lettersWord2(){
	var word2 string 
	word2 = "Software"
	countWord2 := make ([]string,0,len(word2))
	var count int
	for i :=0; i<len(word2); i++{
		countWord2 = append(countWord2 , string(word2[i]))//primera iteraccion(0) toma countWord2 vacio y word[0] toma la letra S, la cual se asigna a la varaibale countWord2= S , segunda iteracion i=1 toma la letra o y lo guadre en countWord2 = S,o asi sucesivamente
		count++
		fmt.Println(countWord2[i])
	}
	fmt.Println("La palabra ",word2, "contiene ",  count )

}

/*
Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. 
Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes 
cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. 
Dentro de los préstamos que otorga no les cobrará interés a los que posean un sueldo superior a $100.000. 
Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de acuerdo a cada caso.
Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/

func loan(){

var age int
var employee bool

fecha1 := time.Date(2023, 2, 24, 0, 0, 0, 0, time.UTC)
fecha2 := time.Date(2023, 3, 24, 0, 0, 0, 0, time.UTC)

var salary float64

age = 23
employee = true
duracion := fecha2.Sub(fecha1)
salary = 3000000

	if age > 22 && employee && duracion >= (time.Hour *24 * 365){
		if salary > 100000{
			fmt.Println("Congratulations, se ha aprobado el prestamo y no se cobrar intereses")
		}else{
			fmt.Println("Congratulations, se ha aprobado el prestamo ")
		}


	}else{
		fmt.Println("Lo siento, no se le puede otorgar un préstamo.")
	}


}

/*
Realizar una aplicación que reciba  una variable con el número del mes. 
Según el número, imprimir el mes que corresponda en texto. 
¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
Ej: 7, Julio.
Nota: Validar que el número del mes, sea correcto.
*/
func impMonths (){
	months := []string{
		" ",
        "Enero",
        "Febrero",
        "Marzo",
        "Abril",
        "Mayo",
        "Junio",
        "Julio",
        "Agosto",
        "Septiembre",
        "Octubre",
        "Noviembre",
        "Diciembre",
    }
	var numMoths int 
    numMoths = 12
	if numMoths <= 12 {
		for  i := 0; numMoths < len(months); i++  {
			fmt.Printf("El mes número %d es %s\n", numMoths, months[numMoths])
		}
		
	}else {
		fmt.Println("Months not exist")
	}


}

/*
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, 
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario. 
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se l
e descontará además un 10 % (27% en total).

func SalaryTax(salary float64) (float64, error) {

	

	if salary <= 0 {
		return 0, errors.New("Salary must be positive and greater than 0")
	} else if salary > 0 && salary < 150000 {
		return salary * 0.17, nil
	}
	return salary * 0.27, nil

}*/