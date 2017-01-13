package main

import (
	"GoAtSchool"
	"GoAtSchool/conv"
	"fmt"
)

func main() {
	fmt.Println(conv.Ml_to_mt(45))
	fmt.Println(conv.S_to_hr(15000))
	fmt.Println(conv.Hr_to_s(2))
	fmt.Println(conv.D_to_s(1))
	fmt.Println(conv.S_to_d(455555))
	fmt.Println(conv.Min_to_d(5000))
	fmt.Println(conv.D_to_min(1))
	fmt.Println(conv.Mlh_to_kmh(10))
	fmt.Println(conv.Kmh_to_mlh(10))
	fmt.Println(conv.Kmh_to_ms(10))
	fmt.Println(conv.Ms_to_kmh(10))
	GoAtSchool.Version()
}
