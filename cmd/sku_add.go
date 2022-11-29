package cmd

import (
	"hewenda/go-rei/storage"
	"regexp"

	log "github.com/sirupsen/logrus"
)

var productUrl string
var productId string

var urlPartten = "^https://www.rei.com/.*product/([0-9]+)[/.+]?"

func SkuAdd(url string) {
	productUrl = url
	matched, err := regexp.MatchString(urlPartten, productUrl)
	if !matched || err != nil {
		log.Fatal("Format url error: ", productUrl)
	}

	urlCompile := regexp.MustCompile(urlPartten)
	params := urlCompile.FindStringSubmatch(productUrl)

	if len(params) >= 2 {
		productId = params[1]
	}

	insertProduct()
}

func insertProduct() {
	storage.InsertProduct(productId, productUrl)
}
