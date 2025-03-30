package main

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) Speak() {
	fmt.Println("I am an animal")
}

type Dog struct {
	Animal
	Breed string
}

func (d *Dog) Speak() {
	fmt.Println("Woof! I am a", d.Breed)
}

func main() {
	d := Dog{Animal: Animal{Name: "Buddy"}, Breed: "Labrador"}
	d.Speak()
	d.Animal.Speak()
}
