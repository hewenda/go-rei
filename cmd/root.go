package cmd

import (
	"bytes"
	"fmt"
	"hewenda/go-rei/storage"
	"log"
	"os"

	"github.com/robfig/cron"
	"github.com/spf13/cobra"
)

var addUrl string
var isMonit bool
var delId string

func listProcut() string {
	output := new(bytes.Buffer)

	for _, item := range storage.QueryProduct() {
		output.WriteString(fmt.Sprintf("[%d] %s\n", item.Id, item.Url))
	}

	return output.String()
}

var rootCmd = &cobra.Command{
	Use:   "rei",
	Short: "rei spider",
	Run: func(cmd *cobra.Command, args []string) {
		if len(addUrl) > 0 {
			SkuAdd(addUrl)
			fmt.Println(listProcut())
		} else if isMonit {
			SkuMonit()
		} else if len(delId) > 0 {
			storage.DeleteProduct(delId)
			fmt.Println(listProcut())
		} else {
			fmt.Println(listProcut())
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&addUrl, "add", "a", "", "Add a url to monit")
	rootCmd.Flags().BoolVarP(&isMonit, "monit", "m", false, "Monit")
	rootCmd.Flags().StringVarP(&delId, "del", "d", "", "Del a id to monit")
}

func SkuMonit() {

	c := cron.New()
	c.AddFunc("0 0 10-23 * * *", func() {
		SkuLoadAndNotify()
	})

	c.AddFunc("0 30 16 * * *", func() {
		storage.ClearDaily()

		for _, user := range storage.QueryUser() {
			PostMessage(Message{
				Content:      listProcut(),
				Notification: false,
				Token:        user.Token,
			})
		}
	})

	c.Start()
	select {}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
