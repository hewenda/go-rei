package cmd

import (
	"bytes"
	"fmt"
	"hewenda/go-rei/spider"
	"hewenda/go-rei/storage"
	"net/http"
	"net/url"
	"time"
)

func Contains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}

func GetAvailableSkus() string {
	wishItems := storage.LoadWish()

	output := new(bytes.Buffer)

	for _, item := range wishItems {
		result := new(bytes.Buffer)
		data := spider.GetUrlModel(item.Url)
		skus := spider.GetAvailableSkus(data)

		for _, sku := range skus {
			if Contains(item.Skus, sku.SkuID) {
				result.WriteString(
					fmt.Sprintf(
						"[%s] %s %s $%.2f=>%.2f",
						sku.SkuID, sku.Color.DisplayLabel, sku.Size.Name, sku.Price.CompareAt.Value, sku.Price.Price.Value,
					),
				)
				if sku.Price.SavingsPercentage != nil {
					result.WriteString(fmt.Sprintf("%v%% off\n", sku.Price.SavingsPercentage))
				} else {
					result.WriteString("\n")

				}
			}
		}
		if result.Len() > 0 {
			result.WriteString("\n")
			output.WriteString(fmt.Sprintf("%s\n%s", data.Title, result.String()))
		}
	}

	return output.String()
}

func PostMessage(token string) {
	currentTime := time.Now()

	title := currentTime.Format("2006-01-02 15:04")
	content := GetAvailableSkus()

	baseUrl, _ := url.Parse("http://www.pushplus.plus/send")

	params := url.Values{}
	params.Add("token", token)
	params.Add("title", title)
	params.Add("content", content)

	baseUrl.RawQuery = params.Encode()

	_, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(baseUrl.String(), err)
		return
	}

}

func SkuMonit() {
	for _, user := range storage.QueryUser() {
		PostMessage(user.Token)
	}
}
