package main

import (
	"fmt"
	service "kek/internal/services"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: app <inputFile> [outputFile]")
		os.Exit(1)
	}

	inputPath := args[0]
	outputPath := "output.txt"
	if len(args) > 1 {
		outputPath = args[1]
	}

	producer := service.NewFileProducer(inputPath)
	presenter := service.NewFilePresenter(outputPath)

	srv := service.New(producer, presenter)
	if err := srv.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
