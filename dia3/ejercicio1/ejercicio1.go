/*
Crear un programa que cumpla los siguiente puntos:
Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products instanciado con valores.
2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products y
añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados
en el slice Products.
Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente
al parámetro pasado.
Ejecutar al menos una vez cada método y función definido desde main().
*/
package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
	Category    string
}

var Products = []Product{
	{1, "Teclado", 600.000, "Lenovo", "Tecnologia"},
}

func (p Product) Save() []Product {
	var newP = append(Products, p)
	return newP
}

func (p Product) GetAll() {
	fmt.Println(Products)
}

func getById(id int) Product {
	for _, p := range Products {
		if p.ID == id {
			return p
		}
	}

	return Product{}
}

func main() {
	p := Product{2, "Mouse", 200.000, "Lenovo", "Tecnologia"}
	Products = p.Save()
	p.GetAll()
	fmt.Println(getById(1))
	fmt.Println(getById(2))
}
