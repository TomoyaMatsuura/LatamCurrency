package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

type pageInfo struct {
	Currency   string    `json:"currency"`
	StatusCode int       `json:"statusCode"`
	URL        string    `json:"url"`
	Title      string    `json:"title"`
	Month      []string  `json:"month"`
	Rate       []float64 `json:"rate"`
}

func savePageJson(fName string, p *pageInfo) {
	// JSONファイルの作成
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	// JSONの内容を標準出力
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	err = enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}

	// JSONに出力(プレフィックスとインデントも設定)
	b, _ := json.MarshalIndent(p, "", "  ")
	fmt.Println(string(b))
}

func main() {
	currency := [8]string{"JPY", "BRL", "MXN", "ARS", "CLP", "PEN", "COP", "BOB"}

	// 上の対象8通貨をループ
	for i := 0; i < len(currency); i++ {

		url := "https://finance.yahoo.com/quote/USD" + currency[i] + "%3DX/history?interval=1mo&filter=history&frequency=1mo"

		p := &pageInfo{
			Currency: currency[i],
		}

		// コントローラの作成
		c := colly.NewCollector()

		// タイトル要素を取得
		c.OnHTML("title", func(e *colly.HTMLElement) {
			p.Title = e.Text
			fmt.Println(e.Text)
		})

		c.OnHTML("table[data-test='historical-prices']", func(e *colly.HTMLElement) {

			e.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
				date := el.ChildText("td:nth-child(1)")
				rate := el.ChildText("td:nth-child(5)")

				p.Month = append(p.Month, date)

				// Parsing rate as a float64
				var r float64
				_, err := fmt.Sscanf(rate, "%f", &r)
				if err != nil {
					log.Println("Error parsing rate:", err)
				}
				p.Rate = append(p.Rate, r)
			})
		})

		// Before making a request print "Visiting URL: https://XXX"
		c.OnRequest(func(r *colly.Request) {
			p.URL = r.URL.String()
			fmt.Println("Visiting URL:", r.URL.String())
		})

		// After making a request extract status code
		c.OnResponse(func(r *colly.Response) {
			p.StatusCode = r.StatusCode
			fmt.Println("StatusCode:", r.StatusCode)
		})

		c.OnError(func(r *colly.Response, err error) {
			p.StatusCode = r.StatusCode
			log.Println("error:", r.StatusCode, err)
		})

		// Start scraping on https://XXX
		c.Visit(url)

		// Wait until threads are finished
		c.Wait()

		// Save as JSON format
		savePageJson(fmt.Sprintf("page_%s.json", currency[i]), p)
	}
}
