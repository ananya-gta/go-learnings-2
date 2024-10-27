package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://www.google.com",
		"Facebook": "https://www.facebook.com",
	}

	fmt.Println(websites)

	websites["LinkedIn"] = "https://www.linkedin.com/in"
	fmt.Println(websites)

	// delete a key-value pair from map
	delete(websites, "Facebook")
	fmt.Println(websites)

	useOfMake()
}

// maps vs structs
/* 	Maps: Key-value pairs, dynamic size, unordered, flexible for quick lookups.
	Structs: Fixed fields, structured data, ordered, suitable for complex data types.
*/

// use of make function
func useOfMake() {
	user := make([]string, 2) // this created an empty array of size 2
	// the below user names will be appended after the 2 empty spaces
	user = append(user, "Ananya")
	user = append(user, "Kartik")
	fmt.Println(len(user))


	// In Go, the make function is used to initialize slices, maps, and channels. 
	// It allocates and initializes the data structure, making it ready for use.
	slice := make([]int, 5) // Creates a slice of integers with length 5 and capacity 5
	fmt.Println(slice)
	slice2 := make([]int, 5, 10) // Length 5, capacity 10
	fmt.Println(slice2)
	myMap := make(map[string]int, 10) // Creates a map with a capacity of 10
	fmt.Println(myMap)
	// Using make is essential in Go when dealing with these types to ensure they are properly allocated before use.
}