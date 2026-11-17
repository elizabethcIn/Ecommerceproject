package main

import (
	"Ecommerceproject/calculos" // Importa tu carpeta con el nombre exacto de tu proyecto
	"fmt"
)

func main() {
	fmt.Println("- SISTEMA DE GESTIÓN E-COMMERCE -")
	fmt.Println("Producto: Laptop Lenovo")

	var precio float64
	var cantidad float64

	fmt.Print("Ingrese el precio del producto: ")
	fmt.Scan(&precio)

	fmt.Print("Ingrese la cantidad a comprar: ")
	fmt.Scan(&cantidad)

	subtotal := calculos.CalcularSubtotal(precio, cantidad)
	iva := calculos.CalcularIVA(subtotal)
	total := calculos.CalcularTotal(subtotal, iva)

	var descuento float64 = 0
	if total > 1000.0 {
		descuento = total * 0.10
		total = total - descuento
		fmt.Println("¡Descuento del 10% aplicado por compra mayor a $1000!")
	}

	fmt.Println("\n--- FACTURA ---")
	fmt.Printf("Subtotal:  $%.2f\n", subtotal)
	fmt.Printf("IVA (15%%): $%.2f\n", iva)
	if descuento > 0 {
		fmt.Printf("Descuento: -$%.2f\n", descuento)
	}
	fmt.Printf("TOTAL:     $%.2f\n", total)
}