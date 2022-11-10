package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// ! configurando um canal de comunicação entre a função main e as demais go routines
	c := make(chan string)

	// ! para cada link dentro do range de links, rodar uma função checkLink
	for _, link := range links {
		go checkLink(link, c)
	}

	// ! para cada endereço de link, realiza uma função anônima que cadastra o link em outro endereço de memória, evitando loop
	for l := range c {
		go func(link string) {
			// ! tempo em espera na função anônima para evitar um bloqueio de firewall
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// ! função de checagem de http response GET, sinalizando conclusão por um canal para a função principal
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " deve estar offline.")
		c <- link
		return
	}

	fmt.Println(link, " está online!")
	c <- link
}
