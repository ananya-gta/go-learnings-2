package main

import (
    "fmt"
)

// Define an interface
type Animal interface {
    Speak() string
}

// Define a Dog type that implements the Animal interface
type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

// Define a Cat type that implements the Animal interface
type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}

// A function that takes an Animal interface
func MakeItSpeak(a Animal) {
    fmt.Println(a.Speak())
}

func main() {
    dog := Dog{}
    cat := Cat{}

    MakeItSpeak(dog) // Output: Woof!
    MakeItSpeak(cat) // Output: Meow!
}
