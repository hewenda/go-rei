package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hewenda/go-rei/spider"
	"hewenda/go-rei/storage"
	"net/http"
)

func Contains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}

func GetAvailableSkus() (string, bool) {
	wishItems := storage.LoadWish()

	output := new(bytes.Buffer)
	oops := false

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
					oops = true
					result.WriteString(fmt.Sprintf("%v%% off\n", sku.Price.SavingsPercentage))
				} else {
					result.WriteString("\n")

				}
			}
		}
		if result.Len() > 0 {
			result.WriteString("\n")
			output.WriteString(fmt.Sprintf("[%s](%s)\n```%s```", data.Title, item.Url, result.String()))
		}
	}

	return output.String(), oops
}

func PostMessage(token string) {
	content, oops := GetAvailableSkus()

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

func SkuMonit() {
	for _, user := range storage.QueryUser() {
		PostMessage(user.Token)
	}
}
