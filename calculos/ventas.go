package calculos

func CalcularSubtotal(precio float64, cantidad float64) float64 {
	return precio * cantidad
}

func CalcularIVA(subtotal float64) float64 {
	return subtotal * 0.15
}

func CalcularTotal(subtotal float64, iva float64) float64 {
	return subtotal + iva
}
