package main

import (
	"fmt"
	"strings"
)

func main() {
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7f50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	for key, Value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, Value)

	}

	fmt.Printf("\033[0;32;41m%s\033[0m\n", strings.Repeat("-", 66))
	removeColor(colors, "Coral")
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}

func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
