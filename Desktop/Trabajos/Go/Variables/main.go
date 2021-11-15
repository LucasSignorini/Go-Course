package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const ENTER = "y presiona ENTER cuando estes listo"

func main() {
	//seed the random numer generator
	rand.Seed(time.Now().UnixMicro()) //Esta linea sirve para que cada vez que
	//ejecutemos el code, haga randoms distintos, ya que cambia la semilla

	var firstNumber = rand.Intn(8) + 2 //El random genera numeros de 0 a 8, como no queremos ni 0, ni 1, sumamos 2
	var secondNumber = rand.Intn(8) + 2
	var resta = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - resta

	//llamamos la funcion que ejecuta el juego:
	game(firstNumber, secondNumber, resta, answer)

}

func game(firstNumber, secondNumber, resta, answer int) {

	//Creamos la variable reader que lee nuestro input
	reader := bufio.NewReader(os.Stdin)
	//Display a welcome/instructions
	fmt.Println("Adivina el numero")
	fmt.Println("------------------")
	fmt.Println("")

	fmt.Println("Piensa un numero entre 1 y 10", ENTER)
	reader.ReadString('\n')
	//take them through the games
	fmt.Println("Multiplica tu numero por", firstNumber, ENTER)
	reader.ReadString('\n')
	fmt.Println("Ahora, multiplica el resultado por", secondNumber, ENTER)
	reader.ReadString('\n')
	fmt.Println("Divide el resultado por el numero que pensaste originalmente", ENTER)
	reader.ReadString('\n')
	fmt.Println("Ahora restale", resta, ENTER)
	reader.ReadString('\n')
	//give them the answer

	fmt.Println("Su numero es:", answer)
}
