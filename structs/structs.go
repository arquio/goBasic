package structs

import (
	"fmt"
)

// GetArray imprime dos arrays
func GetArray() {
	var arr1 [2]string
	arr1[0] = "Test1 Array"
	arr1[1] = "Test2 Array"
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr1, arr2)
}

// GetSlide imprime dos slides
func GetSlide() {
	var sli1 []string
	sli1 = append(sli1, "qweqw", "ip")
	sli2 := []int{1, 2, 3}
	fmt.Println(sli1, sli2)
}

// Curso impartido
type Curso struct {
	Name  string
	Slug  string
	Skill []string
}

//Subscribe Nombre a curso
func (c Curso) Subscribe(n string) {
	fmt.Printf("%s se subcribe a %s", n, c.Name)
}

// Carrera es una estructura que toma los métodos de Curso
type Carrera struct {
	Curso
}

//Subscribe Registra una persona a una carrera
func (c Carrera) Subscribe(n string) {
	fmt.Printf("%s se subcribe a %s", n, c.Name)
}
