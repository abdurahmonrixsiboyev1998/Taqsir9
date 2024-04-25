package main

import (
	"Taqsir9/git"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("storich/username.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	num1, error := git.GetUserName()
	if err != nil {
		log.Fatal(error)
	}

	file.WriteString(num1)

	num2, error := git.GetEmail()
	if err != nil {
		log.Fatal(error)
	}

	file.WriteString(num2)
	file.Close()

}
