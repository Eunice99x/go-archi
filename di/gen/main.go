package main

import (
	"log"
	"os"

	"github.com/eunice99x/dingoSuck/di"
	"github.com/sarulabs/dingo/v4"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run di/gen/main.go <output-path>")
	}

	outputDir := os.Args[1]
	err := dingo.GenerateContainer((*di.Provider)(nil), outputDir)
	if err != nil {
		log.Fatalf("Error generating container: %v", err)
	}
	log.Println("Dingo container generated successfully at:", outputDir)
}