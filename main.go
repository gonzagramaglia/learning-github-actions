package main

import (
    "fmt"
    "os"
)

func main() {

	username := os.Getenv("USERNAME")

	fmt.Printf("%s says: Hola Mundo!\n", username)
}