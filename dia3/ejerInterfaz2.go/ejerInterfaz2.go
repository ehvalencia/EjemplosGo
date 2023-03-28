/*
Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar
el valor del precio total.
La empresa tiene 3 tipos de productos: Pequeño, Mediano y Grande. (Se espera que sean muchos más)

Y los costos adicionales son:
Pequeño: solo tiene el costo del producto
Mediano: el precio del producto + un 3% de mantenerlo en la tienda
Grande: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.

El porcentaje de mantenerlo en la tienda es en base al precio del producto.
El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.

Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Producto que tenga
el método Precio.

Se debe poder ejecutar el método Precio y que el método me devuelva el precio total en base al costo del producto y los
adicionales en caso que los tenga
*/

package main

import "fmt"

const (
	pequenio = "pequenio"
	mediano  = "mediano"
	grande   = "grande"
)

type Producto interface {
	calcularPrecio() float64
}

type ProductoPequenio struct {
	Precio float64
}

func (p ProductoPequenio) calcularPrecio() float64 {
	return p.Precio
}

type ProductoMediano struct {
	Precio float64
}

func (p ProductoMediano) calcularPrecio() float64 {
	return p.Precio + p.Precio*0.03
}

type ProductoGrande struct {
	Precio float64
}

func (p ProductoGrande) calcularPrecio() float64 {
	return p.Precio + p.Precio*0.06 + 2500
}

func productoFactory(tipoProducto string, precio float64) Producto {

	switch tipoProducto {
	case pequenio:
		return ProductoPequenio{Precio: precio}
	case mediano:
		return ProductoMediano{Precio: precio}
	case grande:
		return ProductoGrande{Precio: precio}
	default:
		return nil
	}

}

func main() {
	productoPequenio := productoFactory(pequenio, 10000)
	productoMediano := productoFactory(mediano, 10000)
	productoGrande := productoFactory(grande, 10000)
	fmt.Println(productoPequenio.calcularPrecio())
	fmt.Println(productoMediano.calcularPrecio())
	fmt.Println(productoGrande.calcularPrecio())
}
