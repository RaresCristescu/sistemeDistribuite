package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//Pasul 1 - conectarea la server
	conexiune, _ := net.Dial("tcp", "127.0.0.1:45623")

	for {
		//pasul 2 - ce trimitem
		cititor := bufio.NewReader(os.Stdin)
		fmt.Print("Textul pe care dorim sa-l trimitem catre server: ")
		mesaj, _ := cititor.ReadString('\n')
		//pasul 3 - trimitem mesajul catre server
		fmt.Fprintf(conexiune, mesaj+"\n")
		//pasul 4 - asteptam raspuns de la server
		mesajServer, _ := bufio.NewReader(conexiune).ReadString('\n')
		fmt.Println(mesajServer)
		if mesajServer == "1" { //////////////////////////////////////////////////////////////////
			fmt.Print("Serverul a primit requestul")
		} else {
			fmt.Print("Mesajul de la server este: " + mesajServer)
			fmt.Fprintf(conexiune, "1\n")
		}
	}
}
