package structs

// Academia es una interface de Curso y Carrera
type Academia interface {
	Subscribe(name string)
}

func InterfaceTest() {
	curso := Curso{
		Name:  "Go",
		Slug:  "go",
		Skill: []string{"1", "2"},
	}
	carrera := new(Carrera)
	carrera.Name = "GoCarrera"
	carrera.Slug = "go"
	carrera.Skill = []string{"1", "2"}

	// curso2 := new(Curso)
	// curso2.Name = "Php"
	// curso2.Slug = "php"
	// curso2.Skill = []string{"sdasd", "swqqw"}
	// fmt.Println(curso)
	// fmt.Println(curso2)
	// curso.Subscribe("Daniel")
	callSubscribe(Curso)
	callSubscribe(Carrera)
}

func callSubscribe(p Academia) {
	p.Subscribe("Daniel")
}
