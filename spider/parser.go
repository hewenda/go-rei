package spider

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"hewenda/go-rei/config"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/queue"
)

type Skus map[string]config.Skus

func GetSkus(modelData config.ModelData) Skus {
	skus := make(Skus)

	for _, sku := range modelData.PageData.Product.Skus {
		skus[sku.SkuID] = sku
	}

	return skus
}

func GetAvailableSkus(modelData config.ModelData) Skus {
	skus := make(Skus)

	for _, sku := range modelData.PageData.Product.Skus {
		if sku.Status == "AVAILABLE" {
			skus[sku.SkuID] = sku
		}
	}

	return skus
}

func GetUrlModel(url string) config.ModelData {
	c := colly.NewCollector(
		colly.Async(true),
	)

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	var modelData config.ModelData
	c.OnHTML("#modelData", func(r *colly.HTMLElement) {
		json.Unmarshal([]byte(r.Text), &modelData)
	})

	c.Visit(url)
	c.Wait()

	return modelData
}

func RequestDom() {
	c := colly.NewCollector()

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	q, _ := queue.New(2, &queue.InMemoryQueueStorage{MaxSize: 10000})

	q.AddURL("https://www.rei.com/product/185629/arcteryx-beta-lt-jacket-mens")
	q.AddURL("https://www.rei.com/product/156565/arcteryx-sentinel-ar-pants-womens")
	q.AddURL("https://www.rei.com/product/192241/arcteryx-norvan-lt-hoodie-womens")
	q.AddURL("https://www.rei.com/product/185157/arcteryx-atom-sl-insulated-hoodie-womens")

	c.OnHTML("#modelData", func(r *colly.HTMLElement) {
		var modelData config.ModelData

		json.Unmarshal([]byte(r.Text), &modelData)
		GetAvailableSkus(modelData)
	})

	q.Run(c)
}

func GetFromDom() (config.ModelData, Skus) {
	content, err := ioutil.ReadFile("./model.json")

	if err != nil {
		log.Fatal("Error read file")
	}

	var modelData config.ModelData
	json.Unmarshal(content, &modelData)

	return modelData, GetAvailableSkus(modelData)
}
