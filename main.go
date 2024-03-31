package main

import (
    "fmt"
    "os"
)

func main() {

	username := os.Getenv("USERNAME")

	fmt.Println("%s says: Hola Mundo!", username)
}