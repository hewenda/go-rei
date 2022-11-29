package cmd

import (
	"bytes"
	"fmt"
	"hewenda/go-rei/storage"
	"os"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
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
			log.Println("Run rei --add ", addUrl)

			SkuAdd(addUrl)
			fmt.Println(listProcut())
		} else if isMonit {
			log.Println("Run rei --monit")

			SkuMonit()
		} else if len(delId) > 0 {
			log.Println("Run rei --del ", delId)

			storage.DeleteProduct(delId)
			fmt.Println(listProcut())
		} else {
			log.Println("Run rei --list")

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
	c.AddFunc("0 0,30 10-23 * * *", func() {
		log.Info("Run cron query job")
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
