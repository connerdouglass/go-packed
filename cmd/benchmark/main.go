package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	// Our example object
	rect := Rectangle{
		ID: "rect1",
		Position: Position{
			X: 1.0,
			Y: 2.0,
			Z: 3.0,
		},
		Size: Size{
			Width:  100.0,
			Height: 200.0,
		},
		Color: Color{
			R: 255,
			G: 127,
			B: 0,
			A: 255,
		},
	}

	// Encode the object with packed
	startPacked := time.Now()
	encodedPacked, _ := rect.Packed().Encode()
	elapsedPacked := time.Since(startPacked)

	// Encode the object with JSON
	startJSON := time.Now()
	encodedJSON, _ := json.Marshal(rect)
	elapsedJSON := time.Since(startJSON)

	// Print the results
	fmt.Printf("------------------------------------------------\n")
	fmt.Printf("PACKED\n")
	fmt.Printf("Elapsed: %s\n", elapsedPacked)
	fmt.Printf("Size:    %d bytes\n", len(encodedPacked))
	fmt.Printf("------------------------------------------------\n")
	fmt.Printf("JSON\n")
	fmt.Printf("Elapsed: %s\n", elapsedJSON)
	fmt.Printf("Size:    %d bytes\n", len(encodedJSON))
	fmt.Printf("------------------------------------------------\n")
}
