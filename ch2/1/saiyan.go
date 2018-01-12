package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	fmt.Println("start!")

	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	fmt.Printf("%p\n", &goku)
	fmt.Println(goku.Power)
	Super(goku)
	fmt.Println(goku.Power)

	goku2 := SuperReturn(goku)
	fmt.Printf("%p\n", &goku2)
	fmt.Println(goku2.Power)

	Super2(&goku)
	fmt.Println(goku.Power)
}

func Super(s Saiyan) {
	fmt.Printf("%p\n", &s)
	s.Power += 10000
}

func SuperReturn(s Saiyan) Saiyan {
	fmt.Printf("%p\n", &s)
	s.Power += 20000
	return s
}

func Super2(s *Saiyan) {
	s.Power += 10000
}
