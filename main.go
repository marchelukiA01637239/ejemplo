package main

import "fmt"

func hola(name string) {
	fmt.Printf("Hola %s\n", name)
}

func main() {
	fmt.Println("Hola Mundo!!")
	hola("Abel")
}
