package conv

//Função para conveter horas em segundos
func Hr_to_s(hrs float64) float64 {
	return hrs / 0.00027778
}

//Função para conveter segundos em horas
func S_to_hr(seg float64) float64 {
	return seg * 0.00027778
}

//Função para conveter dias em segundos
func D_to_s(days float64) float64 {
	return days / 0.000011574
}

//Função para conveter segundos em dias
func S_to_d(seg float64) float64 {
	return seg * 0.000011574
}

//Função para conveter dias em minutos
func D_to_min(days float64) float64 {
	return days * 1440.0
}

//Função para conveter minutos em dias
func Min_to_d(min float64) float64 {
	return min * 0.00069444
}
