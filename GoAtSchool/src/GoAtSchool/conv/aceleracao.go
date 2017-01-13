package conv

//Função que converte milhas por hora em km por hora
func Mlh_to_kmh(ml float64) float64 {
	return ml * 1.609344
}

//Função que converte km por hora em milhas por hora
func Kmh_to_mlh(km float64) float64 {
	return km / 1.609344
}

//Função que converte km por hora em metros por segundos
func Kmh_to_ms(km float64) float64 {
	return km / 3.6
}

//Função que converte metros por segundo em km por hora
func Ms_to_kmh(ms float64) float64 {
	return ms * 3.6
}
