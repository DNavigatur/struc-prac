package main

import "fmt"

/*
type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.age, p1.last, p1.first)

	type obi struct {
		person
		ltk   bool
		first string
	}

	obi1 := obi{
		person: person{
			first: "Jenny",
			last:  "Moneypenny",
			age:   27,
		},
		first: "Henry Eze",
		ltk:   true,
	}

	fmt.Println(obi1)

}

//exercise 1
type person struct {
	first_name string
	last_name  string
	icecream   []string
}

func main() {
	a := person{
		first_name: "Henry",
		last_name:  "Eze",
		icecream:   []string{"Vanilla", "Banana"},
	}

	b := person{
		first_name: "John",
		last_name:  "Doe",
		icecream:   []string{"Choclate", "Mango"},
	}

	for _, v := range a.icecream {
		fmt.Println(a.first_name, "likes", v)
	}

	for _, v := range b.icecream {
		fmt.Println(b.first_name, "like", v)

	}

}

//execrise
type person struct {
	first_name string
	last_name  string
	icecream   []string
}

func main() {

	a := person{
		first_name: "Henry",
		last_name:  "Eze",
		icecream:   []string{"Vanilla", "Banana"},
	}

	b := person{
		first_name: "John",
		last_name:  "Doe",
		icecream:   []string{"Choclate", "Mango"},
	}

	mahp := map[string]person{
		a.last_name: a,
		b.last_name: b,
	}

	for _, v := range mahp {
		fmt.Println(v)
	}

}


//exercise 2

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {

	//this practice focus on embeding one struct in another struct
	a := vehicle{
		//this is a composite literal=========
		engine: engine{
			electric: true,
		},
		//================================
		make:  "Tesla",
		model: "X",
		doors: 2,
		color: "Blue",
	}

	b := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Lamborgini",
		model: "Urus",
		doors: 4,
		color: "Red",
	}

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(a.electric, a.model, a.color)
	fmt.Println(b.color, b.make, b.electric)

}
*/

func main() {

	//anonymous struct. "it help me to do for range for the  stuct and its values"
	artist := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: string("Ololade"),
		friends: map[string]int{
			"Olamide":   49,
			"Yhemo lee": 32,
			"Fireboy":   28,
		},
		favDrinks: []string{"fanta", "coke", "henessy"},
	}

	for _, v := range artist.friends {
		fmt.Println(v)
	}

	//My visualizing the map and slice 'process' ============
	//map
	// m := map[string]int{
	// 	"Olamide":   49,
	// 	"Yhemo lee": 32,
	// 	"Fireboy":   28,
	// }

	// //slice
	// f := []string{"fanta", "coke", "henessy"}
	//=================================================================
}
