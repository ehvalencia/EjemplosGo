/*
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir 
el objetivo es necesario crear una función que devuelva el impuesto de un salario. 
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 
se le descontará además un 10 % (27% en total).
*/

package main 

func main(){
 impSalary(52500)
}

func impSalary(salary float64) float64{
 switch {
 case salary > 50000 && salary <= 150000:
	return salary * 0.17

 case salary> 150000:
	return salary*0.27
	
 default:
	return salary	
 }

}