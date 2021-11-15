package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	UserName       string
	Age            int
	NumeroFavorito float64
	tienePerro     bool
}

func main() {
	var user User
	reader = bufio.NewReader(os.Stdin)
	user.UserName = readString("Cual es tu nombre?")
	user.Age = readAge("Cuantos años tenes?")
	user.NumeroFavorito = readNumer("Tu numero favorito es: ")
	user.tienePerro = tienePerro("Tiene perro?")
	//fmt.Println(fmt.Sprintf("Tu nombre es %s. Tu edad es %d", userName, age))

	fmt.Printf("Tu nombre es %s. Tu edad es %d.Tu numero favorio es %.2f. Tiene perro: %t \n", user.UserName, user.Age, user.NumeroFavorito, user.tienePerro)
}

func prompt() {
	fmt.Print("->")
}

func readString(s string) string {
	fmt.Println(s)
	prompt()
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)

	return userInput

}

func readAge(s string) int {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Por favor ingrese un numero")

		} else {
			return num
		}
	}

}

func readNumer(s string) float64 {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Por favor ingrese un numero")

		} else {
			return num
		}
	}

}

func tienePerro(s string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err) //Si hay error, con la libreria de GO log, paramos el programa
	}

	defer func() { //defer hace que se ejecute esta funcion cuando se finaliza lo que se esta ejecutando (en este caso el main)
		_ = keyboard.Close()
	}()
	for {
		fmt.Println(s)
		prompt()
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err) //Si hay error, con la libreria de GO log, paramos el programa
		}
		if char == 'y' || char == 'Y' {
			return true

		}
		if char == 'n' || char == 'N' {
			return false
		}

		fmt.Println("Por favor responda con Y o N")

	}

}
