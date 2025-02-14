package services

import (
	"fmt"
	"log"
	"strings"

	"github.com/HTM1000/table-inss/domain"
	"github.com/gocolly/colly"
)

func ScrapeINSS() []domain.InssTabela {
	var tabelas []domain.InssTabela

	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Visitado: %v\n", r.Request.URL)
	})

	c.OnHTML("div.col", func(e *colly.HTMLElement) {
		data := e.ChildText("h2.h5ms-2")
		if data != "" {
			data = strings.TrimSpace(strings.Replace(data, "INSS a partir de ", "", 1))
			fmt.Printf("Encontrou data: %s\n", data)

			e.ForEach("div.table-responsive table.mtabela tbody tr", func(_ int, row *colly.HTMLElement) {
				faixaTd := row.ChildText("td:first-child")
				aliquotaTd := row.ChildText("td[style='text-align:right;']")

				faixa := strings.TrimSpace(faixaTd)
				aliquota := strings.TrimSpace(strings.Replace(aliquotaTd, "%", "", -1))

				fmt.Printf("Encontrou linha - Faixa: %s, Aliquota: %s\n", faixa, aliquota)

				if faixa != "" && aliquota != "" {
					tabelas = append(tabelas, domain.InssTabela{
						Data:     data,
						Faixa:    faixa,
						Aliquota: aliquota + "%",
					})
				}
			})
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Erro ao visitar %s: %s\n", r.Request.URL, err)
	})

	err := c.Visit("https://www.debit.com.br/tabelas/tabelas-inss")
	if err != nil {
		log.Println("Erro ao acessar a página:", err)
	}

	fmt.Printf("Total de registros encontrados: %d\n", len(tabelas))
	return tabelas
}
