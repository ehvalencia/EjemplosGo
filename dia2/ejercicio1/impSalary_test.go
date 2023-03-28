/*
La  empresa de chocolates que anteriormente necesitaba calcular el impuesto de sus empleados, 
al momento de depositar el sueldo de los mismos ahora nos solicitó validar que los cálculos de estos impuestos 
están correctos. Para esto nos encargaron el trabajo de realizar los test correspondientes para:
Calcular el impuesto en caso de que el empleado gane por debajo de $50.000.
Calcular el impuesto en caso de que el empleado gane por encima de $50.000.
Calcular el impuesto en caso de que el empleado gane por encima de $150.000.
*/

package main

import (
	"github.com/stretchr/testify/assert"
    "testing"
)

func impSalaryTest(t* testing.T){
	t.Run("Salary less that 50.000", func(t* testing.T){
		//arrange
		salary := 45600.00
		expectedResult := salary
		//act
		result := impSalary(salary)
		//assert
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Salary bigger than $50000", func(t *testing.T) {
		// arrange
		salary := 55600.40
		expectedResult := salary * 17 / 100

		// act
		result := impSalary(salary)

		// assert
		assert.Equal(t, expectedResult, result)
})

t.Run("Salary bigger than $150000", func(t *testing.T) {
		// arrange
		salary := 155600.40
		expectedResult := salary * 27 / 100

		// act
		result := impSalary(salary)

		// assert
		assert.Equal(t, expectedResult, result)
})
}