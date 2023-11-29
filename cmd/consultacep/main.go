package main

import (
	"github.com/michelpessoa/segundodesafiogoexpert/pkg/apicep"
	"github.com/michelpessoa/segundodesafiogoexpert/pkg/viacep"

	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

func main() {
	fmt.Println("Digite o CEP:")
	reader := bufio.NewReader(os.Stdin)
	cep, _ := reader.ReadString('\n')
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	cep = reg.ReplaceAllString(cep, "")
	if len(cep) != 8 {
		fmt.Println("CEP precisa ter 8 dígitos.")
		return
	}

	ch1 := make(chan string)
	ch2 := make(chan string)

	go apicep.FetchApiCep(cep, ch1)
	go viacep.FetchViaCep(cep, ch2)

	for {
		select {
		case res := <-ch1:
			fmt.Println(res)
			return
		case res := <-ch2:
			fmt.Println(res)
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout")
			return
		}
	}
}
