package cmd

import (
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

func listProcut() {
	for _, item := range storage.QueryProduct() {
		fmt.Println(item.Id, item.Url)
	}
}

var rootCmd = &cobra.Command{
	Use:   "rei",
	Short: "rei spider",
	Run: func(cmd *cobra.Command, args []string) {
		if len(addUrl) > 0 {
			SkuAdd(addUrl)
			listProcut()
		} else if isMonit {
			SkuMonit()
		} else if len(delId) > 0 {
			storage.DeleteProduct(delId)
			listProcut()
		} else {
			listProcut()
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
	c.AddFunc("0 0,30 * * * *", func() {
		SkuLoadAndNotify()
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
