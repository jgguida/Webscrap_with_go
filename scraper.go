package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"),
	)

	c.OnHTML(".sg-col-inner", func(e *colly.HTMLElement) {
		name := e.ChildText(".a-size-base-plus.a-color-base.a-text-normal")
		price := e.ChildText(".a-price > .a-offscreen")

		if name != "" && price != "" {
			fmt.Printf("Nome: %s\nPreço: %s\n\n", name, price)
		}

	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Erro:", err)
	})

	c.Visit("https://www.amazon.com.br/s?k=notebook+dell+g15+rtx+3050")
	c.Wait()
}
