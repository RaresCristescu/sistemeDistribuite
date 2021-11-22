package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net" /////////////////
	"strconv"
	"strings" /////////////////
)

//pentru id-ul clientilor
var count = 1

func con(c net.Conn) {
	for {
		//obtinem mesajul de la client
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Clientul ", count, " a facut request cu datele: ", netData) /////////////////////////////////////////////////////////////////////////////////////////////
		c.Write([]byte(string("1")))                                             //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" { //comanda de oprire a conexiuni
			break
		}
		va := strings.Split(netData, ",") //despart cuvintele dupa virgula
		n := len(va)                      //retin numarul de cuvinte
		nr_c := len(va[0])                //retin dimensiunea primului cuvant (numarul de caractere)
		not_ok := false
		for i := 0; i < n; i++ { //verific daca sunt toate cuvintele la fel de lungi
			lungime := 0
			if i == n-1 {
				lungime = len(va[i]) - 2
			} else {
				lungime = len(va[i])
			}
			if nr_c != lungime {
				not_ok = true
			}
		}
		if not_ok == true {
			str := "Nu ati introdus cuvinte la fel de mari\n"
			c.Write([]byte(string(str)))
			return
		}
		if nr_c <= conf.DIMMAXIMASIR { //dimensiunea cuvantului sa fie mai mica ca maximul declarat in configurare
			str := ""
			cuvinte := make([]string, 0) //construiesc un slice de cuvinte noi
			for i := 0; i < nr_c; i++ {
				cuv := "" //construiesc noul cuvant
				for j := 0; j < n; j++ {
					cuv += string(va[j][i])
				}
				cuvinte = append(cuvinte, cuv)
				str = str + cuv + " "
			}
			c.Write([]byte(str + "\n")) //trimit noul cuvant la client
		} else {
			str := "Cuvintele trebuie sa fie mai mici decat " + strconv.Itoa(conf.DIMMAXIMASIR) + " caractere\n"
			c.Write([]byte(str))
			return
		}

	}
	count--
	c.Close()
}

//Generam (declarma) o structura referitoare la fructele noastre
type configurare struct {
	PORT           int
	NUMARMAXCLIENT int
	DIMMAXIMASIR   int
}

//obiect pentru gestionarea fisirului (obiectului) json
var conf configurare

func main() {
	fmt.Println("Start server...")

	date_fisier, err := ioutil.ReadFile("Tema1/config.json")
	if err != nil {
		fmt.Print(err)
	}
	//decodificare - unmarshalling process
	err = json.Unmarshal(date_fisier, &conf)
	if err != nil {
		fmt.Println("eroare: ", err)
	}
	//construire port
	PORT := ":" + strconv.Itoa(conf.PORT)
	//Pasul 1 - ascultam pe portul din config
	asculta, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Error ascultare port: ", err)
	}
	//pasul 2 - rulam serverul pana la infinit sau pana cand este apasata combinatia de taste Ctrl-c
	for {
		//pasul 3 - primirea conexiunilor (acceptarea conexiunilor)
		if count <= conf.NUMARMAXCLIENT { //verific sa nu fie mai multi clienti decat doriti
			conexiune, err := asculta.Accept()
			fmt.Println("Client ", count, " connectat") //////////////////////////////////////////////
			if err != nil {
				fmt.Println("Error la acceptarea cereri:", err)
			}
			go con(conexiune)
			count++
		}

	}
}
