package main

import (
	"fmt"
	"time"
)

const helloworld = "Hola %s %s.\n"
const test = "Hello"

func main() {
	// firstname := name.GetName()
	// lastname := "<Modificar con mi apellido>"
	// a, b, c := numbers.GetVars()
	// var num = numbers.Sum(40, b)
	// f32, f64 := numbers.GetFloat()
	// stringUTF8 := name.GetUnicode()
	// fmt.Println("Hola", firstname)
	// fmt.Printf("Hola %s, bienvenido.\n", firstname)
	// fmt.Println(num, a, b, c)
	// fmt.Printf(helloworld, firstname, lastname)
	// fmt.Println(f32, f64)
	// fmt.Println("Texto en Unicode: ", stringUTF8)
	// fmt.Println("Hola"[0], string("Hola"[0]))
	// fmt.Println(len("Hola"))
	// structs.GetArray()
	// structs.GetSlide()
	// Exec()
	//fmt.Println(maps.GetMap("pedro"))
	// number, err := SumInterface(3, 3)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(number)
	// pointerTest()
	forGo(500)
	forGo(400)
	time.Sleep(10000 * time.Millisecond)
}

func pointerTest() {
	a := 100
	var b *int
	b = &a
	fmt.Println("Sin modificar")
	fmt.Println(a, *b)
	fmt.Println(&a, b)
	*b = 10
	fmt.Println("Con una modificación")
	fmt.Println(a, *b)
	fmt.Println(&a, b)

}

func helloGo(index int) {
	fmt.Println("Go rutine #", index)
}

func forGo(n int) {
	for i := 0; i < n; i++ {
		go helloGo(i)
	}
}
