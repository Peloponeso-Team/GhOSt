package conv

//Função que converte Graus Kevin para Celsius
func Kv_to_cl(kv int32) int32 {
	cl := kv - Kv
	return cl
}

//Função que converte Graus Celsius para Kevin
func Cl_to_kv(cl int32) int32 {
	kv := cl + Kv
	return kv
}

//Função que converte Graus fahrenheit para Celsius
func Fh_to_cl(fh float64) float64 {
	cl := fh - Fh
	return cl / 1.8
}

//Função que converte Graus Celsius para fahrenheit
func Cl_to_fh(cl float64) float64 {
	fh := cl * 1.8
	return fh + 32
}
