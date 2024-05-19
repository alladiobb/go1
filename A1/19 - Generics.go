package main

type MyNumber int

type Number interface{
	~int | ~float64
}

func Soma [T Number] (m map[string]T) T {

	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool{
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"A": 100, "B": 200, "C": 300}
	m2 := map[string]float64{"A": 200.1, "B": 400.1, "C": 600.1}

	m3 := map[string]MyNumber{"A": 200, "B": 400, "C": 600}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10,10))
}
