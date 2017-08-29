package name

import "fmt"

// GetName get/return a name string
func GetName() string {
	var name string
	name = "Sin nombre"
	fmt.Print("Ingresar tu nombre:")
	fmt.Scanf("%s", &name)
	return name
}
