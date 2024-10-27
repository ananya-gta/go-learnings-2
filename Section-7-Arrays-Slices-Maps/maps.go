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
}