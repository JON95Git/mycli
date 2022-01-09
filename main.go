package main

import (
	"fmt"
	"log"
	"mycli/app"
	"os"
)

func main() {
	printHeaderPresentation()
	application := app.Generate()
	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func printHeaderPresentation() {
	fmt.Println("==============================")
	fmt.Println("SEARCH FOR IP AND NAME SERVER")
	fmt.Println("==============================")
}
