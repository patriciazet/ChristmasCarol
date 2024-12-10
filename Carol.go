package main

import (
	"fmt"
	"time"
)

func main() {
	// Array of carol lines
	carol := []string{
		"Stille Nacht, heilige Nacht,",
		"Alles schläft, einsam wacht",
		"Nur das traute hochheilige Paar.",
		"Holder Knabe im lockigen Haar,",
		"Schlaf in himmlischer Ruh!",
		"Schlaf in himmlischer Ruh!",
	}

	fmt.Println("German Christmas Carol: Stille Nacht\n")

	// Print each line with a delay
	for _, line := range carol {
		fmt.Println(line)
		time.Sleep(2 * time.Second) // Pause for effect
	}

	fmt.Println("\nFrohe Weihnachten!")
}
