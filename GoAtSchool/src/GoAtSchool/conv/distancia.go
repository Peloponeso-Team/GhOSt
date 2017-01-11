package conv

//Função para converter milhas em kilometros
func Ml_to_km(ml float64) float64 {
	km := ml * Ml
	return km
}

//Função para converter kilometros milhas em milhas
func Km_to_ml(km float64) float64 {
	ml := km / Ml
	return ml
}

//Função para converter kilometros quadrados em hectares
func Km_to_hc(km float64) float64 {
	hc := km / 0.010000
	return hc
}

//Função para converter hectares em kilometros quadrados
func Hc_to_km(hc float64) float64 {
	km := hc * 0.010000
	return km
}
