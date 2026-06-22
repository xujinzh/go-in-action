package main

import "fmt"

func main() {
	colors := map[string]string{}

	colors["Red"] = "#da1337"

	value, exists := colors["Blue"]
	if exists {
		fmt.Printf("value: %v\n", value)
	} else {
		fmt.Println("not exist")
	}

	value1 := colors["Blue"]
	if value1 != "" {
		fmt.Printf("value1: %v\n", value1)
	} else {
		fmt.Println("not exist1")
	}

	colors1 := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7f50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	for key, value := range colors1 {
		fmt.Printf("Key: %s \t\t Value: %s\n", key, value)
	}

	delete(colors1, "Coral")
	for key, value := range colors1 {
		fmt.Printf("Key: %s \t\t Value: %s\n", key, value)
	}
}
