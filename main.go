package main

import "fmt"

func main() {

	type Saiyan struct {
		Name string
		Power int
	}
		
	goku := Saiyan {
		Name: "Goku",
		Power: 9000,
	}

	taiki := Saiyan {
		Name: "Taiki",
		Power: 1000,
	}
	
	fmt.Printf("%v\n %v\n",goku, taiki)

}
