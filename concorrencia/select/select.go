package main

import (
	"fmt"
	"time"

	"github.com/LucasFrezarini/go-html-reader"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1, c2, c3 := html.Titulo(url1), html.Titulo(url2), html.Titulo(url3)

	// Estrutura de controle própria para trabalhar com concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
	}
}
func main() {
	campeao := oMaisRapido(
		"https://www.youtube.com",
		"https://www.google.com",
		"https://www.amazon.com",
	)

	fmt.Println(campeao)
}
