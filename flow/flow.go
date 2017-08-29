package flow

import (
	"fmt"
	"strings"
)

// Exec Ejecutar todos los ejemplos de flujo
func Exec() {
	ifTest()
	forTest()
	strings2()
	switchTest()
}

func ifTest() {
	var number = 0
	fmt.Print("Ingresa un número entero:")
	fmt.Scanf("%d", &number)
	if number%2 == 0 {
		fmt.Println("Par")
	} else {
		fmt.Println("Impar")
	}
	if number2 := 3; number2 == 3 {
		fmt.Println("Número 3")
	}
}

func forTest() {
	for i := 1; i < 6; i++ {
		fmt.Println(i)
	}
	c := 100
	for c > 0 {
		fmt.Println(c)
		c -= 10
	}
	s := 1000
	for {
		s--
		if s == 0 {
			fmt.Printf("Fin de Bucle infinito %d", s)
			break
		}
	}
}

func strings2() {
	var text = "asdasd, ewrewer, iqwieqwie"
	println(strings.ToUpper(text))
	println(strings.ToLower(text))
	println(strings.Replace(text, "w", "2", 2))
	println(strings.Replace(text, "a", "5", -1))
	split := strings.Split(text, " ")
	println(split[0], split[1], split[2])

}
func switchTest() {
	var number = 0
	fmt.Print("Ingresa un número entero:")
	fmt.Scanf("%d", &number)
	switch number {
	case 1:
		fmt.Println("Es número 1")
	default:
		fmt.Printf("%dNo es número 1", number)
	}
	switch {
	case number%2 == 0:
		println("Par")
	default:
		println("Impar")
	}
}
