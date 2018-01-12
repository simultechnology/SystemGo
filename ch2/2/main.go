package main

import "fmt"

type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
	var talker Talker
	talker = &Greeter{"ishiii"}
	talker.Talk()

	var talker2 Talker
	talker2 = Greeter{}
	talker2.Talk()

	var greeter Greeter
	greeter = Greeter{"taka"}
	greeter.Talk()
}
