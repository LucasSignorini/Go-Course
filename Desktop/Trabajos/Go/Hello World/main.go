package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hola Git")
	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		fmt.Print("->")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1) // "\r\n" Se pone para los usuarios de Windows
		userInput = strings.Replace(userInput, "\n", "", -1)   //"\n" Para todos los demas sistemas operativos

		if userInput == "quit" {
			break
		} else {
			Answer := doctor.Response(userInput)
			fmt.Println(Answer)
		}

	}
}
