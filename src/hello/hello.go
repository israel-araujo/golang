package main

import (
	"fmt"
	"os"
	"reflect"
)

var name string

func main() {
	intro()
	exibeMenu()
	monitoraCase()
	lecomando()

}

func intro() {

	version := 1.5
	fmt.Println("May you digit your name : ")
	fmt.Println("                          ")
	fmt.Println("**************************")
	fmt.Println("						   ")
	fmt.Println("						   ")
	fmt.Scan(&name)
	fmt.Println("						   ")
	fmt.Println("						   ")
	fmt.Println(" Hello Mr: ", name)
	fmt.Println(" Golang version : ", version)
	fmt.Println(" version from Golang is: ",
		reflect.TypeOf(version))
	fmt.Println("						   ")

}

func lecomando() int {

	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("						   ")
	fmt.Println(" The number chosen was: ", comandoLido)
	fmt.Println("						   ")
	return comandoLido
}

func exibeMenu() {
	fmt.Println("						   ")
	fmt.Println("Please could you choose one set ")
	fmt.Println("						   ")
	fmt.Println(" 1 - starting monitoring ")
	fmt.Println(" 2 - showing logs")
	fmt.Println(" 0 - exit to program")
	fmt.Println("						   ")

}

func monitoraCase() {

	comandoLido := lecomando()

	switch comandoLido {

	case 1:
		fmt.Println(" Monitoring ...")
		fmt.Println("						   ")
		os.Exit(1)
		fmt.Println("						   ")

	case 2:
		fmt.Println(" logging ...")
		fmt.Println("						   ")
		os.Exit(2)

	case 0:
		fmt.Println(" exiting ...")
		os.Exit(0)
		fmt.Println("						   ")

	default:
		fmt.Println(" This number is unknow, sorry ")
		fmt.Println("						   ")
		os.Exit(-1)
		fmt.Println("						   ")

	}

}
