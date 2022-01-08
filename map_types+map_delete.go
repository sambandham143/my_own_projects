###################. MAIN.GO ###########



package main

import "fmt"

func main() {

	///////  There are different methods to declare a Map ///////

	//Method -1 //

	// var colors map[string]string

	//Method -2 //

	// colors := make(map[string]string)

	// then u can assign the values like below

	// colors["white"] = "#ffffff"

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
	}

	// deleting a key in a map //

	delete(colors, "red")
	fmt.Println(colors)
}
