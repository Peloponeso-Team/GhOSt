package conv

//Função que converte kilogramas em libras
func Kg_to_lb(kg float64) float64 {
	lb := kg * 2.2046
	return lb
}

//Função que converte libras em kilogramas
func Lb_to_kg(lb float64) float64 {
	kg := lb / 2.2046
	return kg
}

//Função que converte onças em gramas
func On_to_gm(on float64) float64 {
	gm := on / 0.035274
	return gm
}

//Função que converte gramas em onças
func Gm_to_on(gm float64) float64 {
	on := gm * 0.035274
	return on
}

//Função que converte libras em onças
func Lb_to_on(lb float64) float64 {
	on := lb * 16.000
	return on
}

//Função que converte onças em libras
func On_to_lb(on float64) float64 {
	lb := on * 0.062500
	return lb
}
