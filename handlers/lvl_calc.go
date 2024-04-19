package handler

import "math"

func CalculateLvl(xp int) int {
	// Tutaj możesz zaimplementować dowolny algorytm obliczania poziomu na podstawie punktów doświadczenia
	// Na przykład możesz użyć prostego algorytmu, który mówi, że każdy kolejny poziom jest osiągany po zdobyciu określonej liczby punktów doświadczenia
	// Na przykład, jeśli za każde 100 punktów użytkownik zdobywa nowy poziom, możesz użyć poniższego kodu:
	lvl := int(math.Floor(math.Sqrt(float64(xp) / 100)))
	return lvl
}
