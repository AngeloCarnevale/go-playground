package main


func soma(x int, y int) (int, bool)  {
	if x > 10 {
		return x + y, true
	}
	return x + y, false
}


type Course struct {
	Name string
	Price int
	Description string
}
func main() {
	// resultado, status := soma(10, 20)

	// println(resultado)
	// println(status)

	course := Course {
		Name: "Curso de Go",
		Description: "Curso de Golang muito legal",
		Price: 100,
	}

	println(course.Name)
	
}