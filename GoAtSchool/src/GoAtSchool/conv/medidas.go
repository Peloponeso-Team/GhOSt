package conv

//Função para converter metros em kilometros
func Mt_to_km(mt float64) float64 {
	return mt / 1000
}

//Função para converter metros em kilometros
func Km_to_mt(km float64) float64 {
	return km * 1000
}

//Função para converter milhas em metros
func Ml_to_mt(ml float64) float64 {
	return ml / 0.00062137
}

//Função para converter metros em milhas
func Mt_to_ml(mt float64) float64 {
	return mt * 0.00062137
}
