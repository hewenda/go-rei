package cmd

import (
	"fmt"
	"hewenda/go-rei/storage"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var addUrl string
var isList bool
var delId string

var rootCmd = &cobra.Command{
	Use:   "rei",
	Short: "rei spider",
	Run: func(cmd *cobra.Command, args []string) {
		if len(addUrl) > 0 {
			SkuAdd(addUrl)
		} else if isList {
			for _, item := range storage.LoadWish() {
				fmt.Println(item.Id, item.Url, item.Skus)
			}
		} else if len(delId) > 0 {
			storage.DeleteWish(delId)
		} else {
			SkuMonit()
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&addUrl, "add", "a", "", "Add a url to monit")
	rootCmd.Flags().BoolVarP(&isList, "list", "l", false, "List monit")
	rootCmd.Flags().StringVarP(&delId, "del", "d", "", "Del a id to monit")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
