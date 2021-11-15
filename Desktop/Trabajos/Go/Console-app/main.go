package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err) //Si hay error, con la libreria de GO log, paramos el programa
	}

	defer func() { //defer hace que se ejecute esta funcion cuando se finaliza lo que se esta ejecutando (en este caso el main)
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("--------")
	fmt.Println("1. Capuccino")
	fmt.Println("2. Latte")
	fmt.Println("3. Americano")
	fmt.Println("4. Mocha")
	fmt.Println("5. Macchiato")
	fmt.Println("6. Espresso")
	fmt.Println("Q - Quit")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err) //Si hay error, con la libreria de GO log, paramos el programa
		}
		if char == 'q' || char == 'Q' {
			break
		}
		i, _ := strconv.Atoi(string(char))

		fmt.Println(fmt.Sprintf("Elegiste %s", coffees[i])) //Sprintf reemplaza el %q por lo que haya en char

	}
	fmt.Println("Salida exitosa")
}
