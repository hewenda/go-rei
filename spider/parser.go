package spider

import (
	"encoding/json"

	"hewenda/go-rei/config"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type Skus map[string]config.Skus

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
