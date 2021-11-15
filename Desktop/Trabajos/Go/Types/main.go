package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var keyPressChan chan rune

func main() {
	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Toca cualquier tecla, q para salir")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		fmt.Println("hola")
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChan <- char //Sintaxis para enviar algo por el canal
	}

}

type Animal interface {
	Says() string     //Entonces si creamos un tipo de dato cat o perro, ambas
	HowManyLegs() int // tienen que tener estas funciones disponibles
}

func listenForKeyPress() {
	for {
		key := <-keyPressChan //Sintaxis para recibir algo por el canal
		fmt.Println("Tocaste el: ", string(key))
	}
}
