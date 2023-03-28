/*
Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. 
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. 
No se pueden introducir notas negativas.
*/

package main

func main(){

}

func calculationAverange(notes ... int )int{
	   var count int
        var result int

        if len(notes) == 0 {
                return 0
        }

        for _, note := range notes {
                if !(note < 0) {
                        result += note
                        count++
                }
        }

        return result / count
}