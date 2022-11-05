package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hewenda/go-rei/spider"
	"hewenda/go-rei/storage"
	"net/http"
)

func GetAvailableSkus() (string, bool) {
	product := storage.QueryProduct()

	output := new(bytes.Buffer)
	oops := false

	for _, item := range product {
		result := new(bytes.Buffer)

		data := spider.GetUrlModel(item.Url)
		skus := spider.GetAvailableSkus(data)

		for _, sku := range skus {
			if sku.Price.SavingsPercentage != nil {
				oops = true

				result.WriteString(
					fmt.Sprintf(
						"%s %s $%.2f=>%.2f",
						sku.Color.DisplayLabel, sku.Size.Name, sku.Price.CompareAt.Value, sku.Price.Price.Value,
					),
				)
				result.WriteString(fmt.Sprintf(" %v%%Off\n", sku.Price.SavingsPercentage))
			}
		}
		if result.Len() > 0 {
			result.WriteString("\n")
			output.WriteString(fmt.Sprintf("[%s](%s)\n``` %s ```", data.Title, item.Url, result.String()))
		}
	}

	return output.String(), oops
}

func PostMessage(token string) {
	content, oops := GetAvailableSkus()

	if !oops {
		return
	}

	baseUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	data := make(map[string]interface{})
	data["chat_id"] = -794133668
	data["text"] = content
	data["parse_mode"] = "Markdown"
	data["disable_notification"] = !oops
	b, _ := json.Marshal(data)

	fmt.Println(content)

	_, err := http.Post(
		baseUrl,
		"application/json",
		bytes.NewBuffer(b),
	)
	if err != nil {
		fmt.Println(err)
	}

}

func SkuLoadAndNotify() {
	for _, user := range storage.QueryUser() {
		PostMessage(user.Token)
	}
}
