package maps

// setMaps devuelve un map con llave tipo string
func setMaps() map[string]int {
	// mapTest := make(map[string]int)
	// mapTest["cadena1"] = 96
	// mapTest["cadena2"] = 69
	mapTest := map[string]int{
		"juan":  3,
		"pedro": 6,
	}
	mapTest["cadena1"] = 96
	mapTest["cadena2"] = 69
	delete(mapTest, "cadena1")
	delete(mapTest, "cadena2")

	return mapTest
}

//GetMap optine una edad de un nombre
func GetMap(name string) int {
	mapTest := setMaps()
	return mapTest[name]
}
