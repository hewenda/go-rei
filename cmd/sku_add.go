package cmd

import (
	"fmt"
	"hewenda/go-rei/spider"
	"hewenda/go-rei/storage"
	"log"
	"regexp"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

var productUrl string
var productId string

var urlPartten = "^https://www.rei.com/product/([0-9]+)[/.+]?"

func SkuAdd(url string) {
	productUrl = url
	matched, err := regexp.MatchString(urlPartten, productUrl)
	if !matched || err != nil {
		log.Fatal("Format url error")
	}

	urlCompile := regexp.MustCompile(urlPartten)
	params := urlCompile.FindStringSubmatch(productUrl)

	if len(params) >= 2 {
		fmt.Println(params)
		productId = params[1]
	}

	ask()
}

func filter(filterValue string, optValue string, optIndex int) bool {
	return strings.Contains(optValue, filterValue)
}

func ask() {
	data := spider.GetUrlModel(productUrl)
	skus := spider.GetSkus(data)

	answer := []string{}
	options := []string{}

	for skuId, sku := range skus {
		options = append(
			options,
			fmt.Sprintf("[%s] %s %s (%s)", skuId, sku.Color.DisplayLabel, sku.Size.Name, sku.Status),
		)
	}

	if len(options) == 0 {
		log.Fatal(string(fmt.Sprintf("Get %s fail", productUrl)))
	}

	prompt := &survey.MultiSelect{
		Message: data.Title,
		Options: options,
		Filter:  filter,
	}

	err := survey.AskOne(prompt, &answer)
	if err != nil {
		log.Fatal(err)
	}

	var skuIds []string
	for _, checked := range answer {
		end := strings.Index(checked, "]")
		id := checked[1:end]
		skuIds = append(skuIds, id)
	}

	storage.InsertWish(productId, productUrl, skuIds)
}
